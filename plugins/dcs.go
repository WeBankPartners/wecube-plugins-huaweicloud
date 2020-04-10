package plugins

import (
	"fmt"
	"github.com/WeBankPartners/wecube-plugins-huaweicloud/plugins/utils"
	"github.com/huaweicloud/golangsdk"
	"github.com/huaweicloud/golangsdk/openstack"
	"github.com/huaweicloud/golangsdk/openstack/dcs/v1/availablezones"
	"github.com/huaweicloud/golangsdk/openstack/dcs/v1/instances"
	"github.com/huaweicloud/golangsdk/openstack/dcs/v1/products"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
	"time"
)

const (
	INSTANCE_TYPE_SINGLE  = "single"
	INSTANCE_TYPE_HA      = "ha"
	INSTANCE_TYPE_CLUSTER = "cluster"
	INSTANCE_TYPE_PROXY   = "proxy"
)

func createDcsServiceClient(params CloudProviderParam) (*golangsdk.ServiceClient, error) {
	client, err := createGolangSdkProviderClient(params)
	if err != nil {
		return nil, err
	}
	cloudMap, _ := GetMapFromString(params.CloudParams)
	sc, err := openstack.NewDCSServiceV1(client, golangsdk.EndpointOpts{
		Region: cloudMap[CLOUD_PARAM_REGION],
	})
	if err != nil {
		logrus.Errorf("createDcsServiceClient meet err=%v", err)
		return nil, err
	}
	return sc, err
}

var dcsActions = make(map[string]Action)

func init() {
	dcsActions["create"] = new(DcsCreateAction)
	dcsActions["delete"] = new(DcsDeleteAction)
}

type DcsPlugin struct {
}

func (plugin *DcsPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := dcsActions[actionName]
	if !found {
		logrus.Errorf("dcs plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("dcs plugin,action = %s not found", actionName)
	}
	return action, nil
}

type DcsCreateInputs struct {
	Inputs []DcsCreateInput `json:"inputs,omitempty"`
}

type DcsCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid          string `json:"guid,omitempty"`
	Seed          string `json:"seed,omitempty"`
	Id            string `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	InstanceType  string `json:"instance_type,omitempty"`
	EngineVersion string `json:"engine_version,omitempty"`
	Capacity      string `json:"capacity,omitempty"`
	Password      string `json:"password,omitempty"`

	VpcId           string `json:"vpc_id,omitempty"`
	SubnetId        string `json:"subnet_id,omitempty"`
	SecurityGroupId string `json:"security_group_id,omitempty"`
	AvailableZones  string `json:"az,omitempty"`
	PrivateIp       string `json:"private_ip,omitempty"`
	//	Labels              string `json:"labels,omitempty"`
	//	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
	ChargeType string `json:"charge_type,omitempty"`

	//包年包月
	PeriodType  string `json:"period_type,omitempty"`   //年或月
	PeriodNum   string `json:"period_num,omitempty"`    //年有效值[1-9],月有效值[1-3]
	IsAutoRenew string `json:"is_auto_renew,omitempty"` //是否自动续费
}

type DcsCreateOutputs struct {
	Outputs []DcsCreateOutput `json:"outputs,omitempty"`
}

type DcsCreateOutput struct {
	CallBackParameter
	Result
	Guid      string `json:"guid,omitempty"`
	Id        string `json:"id,omitempty"`
	PrivateIp string `json:"private_ip,omitempty"`
	Password  string `json:"password,omitempty"`
}

type DcsCreateAction struct {
}

func (action *DcsCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs DcsCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkCreateDcsParams(input DcsCreateInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}

	if input.Name == "" {
		return fmt.Errorf("empty name")
	}

	validInstanceTypes := []string{
		INSTANCE_TYPE_SINGLE, INSTANCE_TYPE_HA,
		INSTANCE_TYPE_CLUSTER, INSTANCE_TYPE_PROXY,
	}
	if err := isValidStringValue("instanceType", input.InstanceType, validInstanceTypes); err != nil {
		return err
	}

	validRedisVersions := []string{
		"3.0", "4.0", "5.0",
	}
	if err := isValidStringValue("redisVersion", input.EngineVersion, validRedisVersions); err != nil {
		return err
	}

	if input.Capacity == "" {
		return fmt.Errorf("empty capacity")
	}
	if input.VpcId == "" {
		return fmt.Errorf("empty vpcId")
	}
	if input.SubnetId == "" {
		return fmt.Errorf("empty subnetId")
	}
	if input.SecurityGroupId == "" {
		return fmt.Errorf("empty SecurityGroupId")
	}
	if input.AvailableZones == "" {
		return fmt.Errorf("empty AvailableZones")
	}

	if err := isValidStringValue("chargeType", input.ChargeType, []string{PRE_PAID, POST_PAID}); err != nil {
		return err
	}

	if _, err := strconv.Atoi(input.Capacity); err != nil {
		return fmt.Errorf("capacity(%v) invalid", input.Capacity)
	}

	if input.ChargeType == PRE_PAID {
		if err := isValidStringValue("periodType", input.PeriodType, []string{PRE_PAID_MONTH, PRE_PAID_YEAR}); err != nil {
			return err
		}

		if _, err := isValidInteger(input.PeriodNum, 1, 12); err != nil {
			return err
		}
	}
	return nil
}

func getDcsInfoById(cloudProviderParam CloudProviderParam, id string) (*instances.Instance, error) {
	sc, err := createDcsServiceClient(cloudProviderParam)
	if err != nil {
		return nil, err
	}

	dcsInfo, err := instances.Get(sc, id).Extract()
	if err != nil {
		logrus.Errorf("getvmInfoById failed err=%v\n", err)
	}
	return dcsInfo, err
}

func waitDcsCreateOk(sc *golangsdk.ServiceClient, id string) (*instances.Instance, error) {
	var finalErr error

	for {
		time.Sleep(time.Duration(5) * time.Second)
		dcsInfo, err := instances.Get(sc, id).Extract()
		if err != nil {
			finalErr = err
			break
		}
		if dcsInfo.Status == "CREATEFAILED" || dcsInfo.Status == "ERROR" {
			finalErr = fmt.Errorf("create dcs status=%v", dcsInfo.Status)
			break
		}
		if dcsInfo.Status == "RUNNING" {
			return dcsInfo, nil
		}
	}
	return nil, finalErr
}

func isDcsExist(cloudProviderParam CloudProviderParam, id string) (*instances.Instance, bool, error) {
	dcsInfo, err := getDcsInfoById(cloudProviderParam, id)
	if err != nil {
		if strings.Contains(err.Error(), "No Nat Gateway exist") {
			return nil, false, nil
		}
		return nil, false, err
	}

	return dcsInfo, true, nil
}

func getPayMode(chargeType string, periodType string) (string, error) {
	if chargeType == POST_PAID {
		return "Hourly", nil
	}

	if chargeType == PRE_PAID {
		if periodType == PRE_PAID_MONTH {
			return "Monthly", nil
		}
		if periodType == PRE_PAID_YEAR {
			return "Yearly", nil
		}
	}
	return "", fmt.Errorf("unkonwn pay mode")
}

func getAvailableAzMap(sc *golangsdk.ServiceClient) (map[string]string, error) {
	azMap := make(map[string]string)

	response, err := availablezones.Get(sc).Extract()
	if err != nil {
		return azMap, err
	}

	for _, az := range response.AvailableZones {
		if az.ResourceAvailability == "true" {
			azMap[az.Code] = az.ID
		}
	}

	return azMap, nil
}

func isAzsAllInAzMaps(azs []string, azMap map[string]string) error {
	for _, az := range azs {
		if _, ok := azMap[az]; !ok {
			return fmt.Errorf("az %v can't be found in map(%++v)", az, azMap)
		}
	}
	return nil
}

func getRedisProductId(input DcsCreateInput, azs []string) (string, []string, error) {
	sc, err := createDcsServiceClient(input.CloudProviderParam)
	if err != nil {
		return "", azs, err
	}

	payMode, err := getPayMode(input.ChargeType, input.PeriodType)
	if err != nil {
		return "", azs, err
	}

	azMap, err := getAvailableAzMap(sc)
	if err != nil {
		return "", azs, err
	}

	//check az have resource
	for _, az := range azs {
		if _, ok := azMap[az]; !ok {
			return "", azs, fmt.Errorf("az(%v) is unavailable,please use %++v", az, azMap)
		}
	}

	response, err := products.Get(sc).Extract()
	if err != nil {
		return "", azs, err
	}
	for _, product := range response.Products {
		if product.Engine != "redis" || product.CpuType != "x86_64" || product.CacheMode != input.InstanceType {
			continue
		}

		//检查付费方式是否匹配
		if false == strings.EqualFold(payMode, product.ChargingType) {
			continue
		}

		//检查版本是否匹配："engine_versions": "4.0;5.0",
		versions := strings.Split(product.EngineVersions, ";")
		if err := isValidStringValue("engineVersion", input.EngineVersion, versions); err != nil {
			continue
		}

		for _, flavor := range product.Flavors {
			//检查容量是否匹配
			if flavor.Capacity != input.Capacity {
				continue
			}
			if err := isAzsAllInAzMaps(azs, azMap); err == nil {
				azCodes := []string{}
				for _, az := range azs {
					azCodes = append(azCodes, azMap[az])
				}

				return product.ProductID, azCodes, nil
			}
		}
	}

	return "", azs, fmt.Errorf("can't find the desire redis productId")
}

func createDcs(input DcsCreateInput) (output DcsCreateOutput, err error) {
	defer func() {
		output.Guid = input.Guid
		output.CallBackParameter.Parameter = input.CallBackParameter.Parameter
		if err == nil {
			output.Result.Code = RESULT_CODE_SUCCESS
		} else {
			output.Result.Code = RESULT_CODE_ERROR
			output.Result.Message = err.Error()
		}
	}()

	if err = checkCreateDcsParams(input); err != nil {
		return
	}

	if input.Id != "" {
		exist := false
		_, exist, err = isDcsExist(input.CloudProviderParam, input.Id)
		if err == nil && exist {
			output.Id = input.Id
			return
		}
	}

	capacity, _ := strconv.Atoi(input.Capacity)
	azs, err := GetArrayFromString(input.AvailableZones, ARRAY_SIZE_REAL, 0)
	if err != nil {
		return
	}

	productId, azCodes, err := getRedisProductId(input, azs)
	if err != nil {
		return
	}

	opts := instances.CreateOps{
		Name:             input.Name,
		Engine:           "Redis",
		EngineVersion:    input.EngineVersion,
		Capacity:         capacity,
		NoPasswordAccess: "false",
		VPCID:            input.VpcId,
		SecurityGroupID:  input.SecurityGroupId,
		SubnetID:         input.SubnetId,
		AvailableZones:   azCodes,
		ProductID:        productId,
	}
	if input.PrivateIp != "" {
		opts.PrivateIp = input.PrivateIp
	}
	if input.Password == "" {
		input.Password = utils.CreateRandomPassword()
	}
	opts.Password = input.Password

	sc, err := createDcsServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	resp, err := instances.Create(sc, opts).Extract()
	fmt.Printf("resp=%++v,err=%v\n", resp, err)
	if err != nil {
		return
	}

	output.Id = resp.InstanceID
	newDcsInstance, err := waitDcsCreateOk(sc, resp.InstanceID)
	if err != nil {
		return
	}
	output.PrivateIp = newDcsInstance.IP

	output.Password, err = utils.AesEnPassword(input.Guid, input.Seed, input.Password, utils.DEFALT_CIPHER)
	if err != nil {
		return
	}
	logrus.Infof("newDcsInstance =%++v", newDcsInstance)

	return
}

func (action *DcsCreateAction) Do(inputs interface{}) (interface{}, error) {
	dcs, _ := inputs.(DcsCreateInputs)
	outputs := DcsCreateOutputs{}
	var finalErr error

	for _, input := range dcs.Inputs {
		output, err := createDcs(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return outputs, finalErr
}

type DcsDeleteInputs struct {
	Inputs []DcsDeleteInput `json:"inputs,omitempty"`
}

type DcsDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type DcsDeleteOutputs struct {
	Outputs []DcsDeleteOutput `json:"outputs,omitempty"`
}

type DcsDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type DcsDeleteAction struct {
}

func (action *DcsDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs DcsDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func deleteDcs(input DcsDeleteInput) (output DcsDeleteOutput, err error) {
	defer func() {
		output.Guid = input.Guid
		output.Id = input.Id
		output.CallBackParameter.Parameter = input.CallBackParameter.Parameter
		if err == nil {
			output.Result.Code = RESULT_CODE_SUCCESS
		} else {
			output.Result.Code = RESULT_CODE_ERROR
			output.Result.Message = err.Error()
		}
	}()

	if err = isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return
	}
	if input.Id == "" {
		err = fmt.Errorf("id is empty")
		return
	}

	_, exist, err := isDcsExist(input.CloudProviderParam, input.Id)
	if err != nil || !exist {
		return
	}

	sc, err := createDcsServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	err = instances.Delete(sc, input.Id).ExtractErr()
	return
}

func (action *DcsDeleteAction) Do(inputs interface{}) (interface{}, error) {
	dcs, _ := inputs.(DcsDeleteInputs)
	outputs := DcsDeleteOutputs{}
	var finalErr error

	for _, input := range dcs.Inputs {
		output, err := deleteDcs(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all dcs= %v are delete", dcs)
	return &outputs, finalErr
}
