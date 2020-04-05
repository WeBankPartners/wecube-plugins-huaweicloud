package plugins

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/WeBankPartners/wecube-plugins-huaweicloud/plugins/utils"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/backups"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/datastores"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/flavors"
	"github.com/gophercloud/gophercloud/openstack/rds/v3/instances"
	"github.com/sirupsen/logrus"
)

const (
	RDS_PRE_PAID_MONTH_MIN_VALUE = 1
	RDS_PRE_PAID_MONTH_MAX_VALUE = 9
	RDS_PRE_PAID_YEAR_MIN_VALUE  = 1
	RDS_PRE_PAID_YEAR_MAX_VALUE  = 3

	RDS_VOLUME_TYPE_ULTRAHIGH    = "ULTRAHIGH"
	RDS_VOLUME_TYPE_ULTRAHIGHPRO = "ULTRAHIGHPRO"

	RDS_INSTANCE_STATUS_OK  = "ACTIVE"
	RDS_INSTANCE_STATUS_BAD = "FAILED"

	RDS_BACKUP_STATUS_OK       = "COMPLETED"
	RDS_BACKUP_STATUS_BAD      = "FAILED"
	RDS_FLAVORREF_AZ_STATUS_OK = "normal"
)

var rdsActions = make(map[string]Action)
var engineTypeMap = map[string]string{
	"MYSQL":      "MySQL",
	"POSTGRESQL": "PostgreSQL",
	"SQLSERVER":  "SQLServer",
}

func init() {
	rdsActions["create"] = new(RdsCreateAction)
	rdsActions["delete"] = new(RdsDeleteAction)
	rdsActions["create-backup"] = new(RdsCreateBackupAction)
	rdsActions["delete-backup"] = new(RdsDeleteBackupAction)
}

func createRdsServiceClientV3(params CloudProviderParam) (*gophercloud.ServiceClient, error) {
	provider, err := createGopherCloudProviderClient(params)
	if err != nil {
		logrus.Errorf("get gophercloud provider client failed, error=%v", err)
		return nil, err
	}

	sc, err := openstack.NewRDSV3(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Errorf("NewRDSV3 failed, error=%v", err)
		return nil, err
	}

	return sc, nil
}

type RdsPlugin struct {
}

func (plugin *RdsPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := rdsActions[actionName]
	if !found {
		logrus.Errorf("Rds plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("Rds plugin,action = %s not found", actionName)
	}

	return action, nil
}

type RdsCreateInputs struct {
	Inputs []RdsCreateInput `json:"inputs,omitempty"`
}

type RdsCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid     string `json:"guid,omitempty"`
	Id       string `json:"id,omitempty"`
	Seed     string `json:"seed,omitempty"`
	Name     string `json:"name,omitempty"`
	Password string `json:"password,omitempty"`
	Port     string `json:"port,omitempty"`
	HostType string `json:"machine_spec,omitempty"` //4c8g
	// EngineType       string `json:"engine_type,omitempty"`
	EngineVersion     string `json:"engine_version,omitempty"`
	SecurityGroupId   string `json:"security_group_id,omitempty"`
	VpcId             string `json:"vpc_id,omitempty"`
	SubnetId          string `json:"subnet_id,omitempty"`
	AvailabilityZone  string `json:"az,omitempty"`
	SupportHa         string `json:"support_ha,omitempty"`
	HaReplicationMode string `json:"ha_replication_mode,omitempty"`
	VolumeType        string `json:"volume_type,omitempty"`
	VolumeSize        string `json:"volume_size,omitempty"`
	ChargeType        string `json:"charge_type,omitempty"`

	//包年包月
	PeriodType  string `json:"period_type,omitempty"`   //年或月
	PeriodNum   string `json:"period_num,omitempty"`    //年有效值[1-9],月有效值[1-3]
	IsAutoRenew string `json:"is_auto_renew,omitempty"` //是否自动续费

	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

type RdsCreateOutputs struct {
	Outputs []RdsCreateOutput `json:"outputs,omitempty"`
}

type RdsCreateOutput struct {
	CallBackParameter
	Result
	Guid      string `json:"guid,omitempty"`
	Id        string `json:"id,omitempty"`
	PrivateIp string `json:"private_ip,omitempty"`

	//用户名和密码
	Port     string `json:"private_port,omitempty"`
	UserName string `json:"user_name,omitempty"`
	Password string `json:"password,omitempty"`
	Cpu      string `json:"cpu,omitementy"`
	Memory   string `json:"memory,omitempty"`
}

type RdsCreateAction struct {
}

func (action *RdsCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs RdsCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *RdsCreateAction) checkCreateRdsParams(input RdsCreateInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Seed == "" {
		return fmt.Errorf("seed is empty")
	}
	if input.Name == "" {
		return fmt.Errorf("name is empty")
	}
	if input.Password == "" {
		return fmt.Errorf("password is empty")
	}
	if input.HostType == "" {
		return fmt.Errorf("hostType is empty")
	}
	if input.SecurityGroupId == "" {
		return fmt.Errorf("securityGroupId is empty")
	}
	if input.VpcId == "" {
		return fmt.Errorf("vpcId is empty")
	}
	if input.SubnetId == "" {
		return fmt.Errorf("subnetId is empty")
	}
	if input.VolumeSize == "" {
		return fmt.Errorf("volumeSize is empty")
	}

	// if input.VolumeType == "" {
	// 	return fmt.Errorf("volumeType is empty")
	// }
	if input.VolumeType != RDS_VOLUME_TYPE_ULTRAHIGH && input.VolumeType != RDS_VOLUME_TYPE_ULTRAHIGHPRO {
		return fmt.Errorf("volumeType is wrong")
	}

	if err := checkEngineParams(input.CloudProviderParam, engineTypeMap["MYSQL"], input.EngineVersion); err != nil {
		return err
	}
	if input.AvailabilityZone == "" {
		return fmt.Errorf("availabilityZone is empty")
	}
	if err := isValidStringValue("chargeType", input.ChargeType, []string{PRE_PAID, POST_PAID}); err != nil {
		return err
	}

	if input.ChargeType == PRE_PAID {
		if err := isValidStringValue("periodType", input.PeriodType, []string{PRE_PAID_MONTH, PRE_PAID_YEAR}); err != nil {
			return err
		}

		if input.PeriodType == PRE_PAID_MONTH {
			if _, err := isValidInteger(input.PeriodNum, RDS_PRE_PAID_MONTH_MIN_VALUE, RDS_PRE_PAID_MONTH_MAX_VALUE); err != nil {
				return err
			}
		} else {
			if _, err := isValidInteger(input.PeriodNum, RDS_PRE_PAID_YEAR_MIN_VALUE, RDS_PRE_PAID_YEAR_MAX_VALUE); err != nil {
				return err
			}
		}
		if input.IsAutoRenew != "" && strings.ToLower(input.IsAutoRenew) != "true" && strings.ToLower(input.IsAutoRenew) != "false" {
			return fmt.Errorf("isAutoRenew is wrong")
		}
	}

	if input.SupportHa != "" {
		if strings.ToLower(input.SupportHa) != "true" && strings.ToLower(input.SupportHa) != "false" {
			return fmt.Errorf("supportHa is wrong")
		}
		if strings.ToLower(input.SupportHa) == "true" && input.HaReplicationMode == "" {
			return fmt.Errorf("haReplicationMode is empty")
		}
	}

	return nil
}

func checkEngineParams(params CloudProviderParam, engineType, engineVersion string) error {
	if engineType == "" {
		return fmt.Errorf("engineType is empty")
	}
	if engineVersion == "" {
		return fmt.Errorf("engineVersion is empty")
	}
	if _, ok := engineTypeMap[strings.ToUpper(engineType)]; !ok {
		return fmt.Errorf("engineType is wrong")
	}
	versionMap, err := queryEngineVersionInfo(params, engineType)
	if err != nil {
		return err
	}
	if _, ok := versionMap[engineVersion]; !ok {
		return fmt.Errorf("engineVersion is wrong")
	}

	return nil
}

func queryEngineVersionInfo(params CloudProviderParam, engineType string) (map[string]string, error) {
	versions := map[string]string{}
	sc, err := createRdsServiceClientV3(params)
	if err != nil {
		return versions, err
	}
	allPages, err := datastores.List(sc, engineType).AllPages()
	if err != nil {
		return versions, err
	}
	allDatastores, err := datastores.ExtractDataStores(allPages)
	if err != nil {
		return versions, err
	}
	for _, datastore := range allDatastores.DataStores {
		versions[datastore.Name] = datastore.Id
	}

	return versions, nil
}

// return flavor name, cpu, ram
func getRdsFlavorByHostType(input *RdsCreateInput) (string, string, string, error) {
	cpu, memory, err := getCpuAndMemoryFromHostType(input.HostType)
	if err != nil {
		return "", "", "", err
	}

	sc, err := createRdsServiceClientV3(input.CloudProviderParam)
	if err != nil {
		return "", "", "", err
	}
	allPages, err := flavors.List(sc, flavors.DbFlavorsOpts{
		Versionname: input.EngineVersion,
	}, engineTypeMap["MYSQL"]).AllPages()
	if err != nil {
		return "", "", "", err
	}

	allFlavorsResp, err := flavors.ExtractDbFlavors(allPages)
	if err != nil {
		return "", "", "", err
	}

	var minScore int64 = 1000000
	matchCpuItems := []flavors.Flavors{}
	for _, item := range allFlavorsResp.Flavorslist {
		if _, ok := item.Azstatus[input.AvailabilityZone]; !ok {
			continue
		}
		if item.Azstatus[input.AvailabilityZone] != RDS_FLAVORREF_AZ_STATUS_OK {
			continue
		}

		if strings.ToLower(input.SupportHa) == "true" && item.Instancemode != "ha" {
			continue
		}

		if (strings.ToLower(input.SupportHa) == "false" || input.SupportHa == "") && item.Instancemode != "single" {
			continue
		}

		vcpus, err := strconv.ParseInt(item.Vcpus, 10, 64)
		if err != nil {
			logrus.Errorf("vpus(%v) is invald", item.Vcpus)
			return "", "", "", err
		}

		score := vcpus - cpu
		if score < 0 {
			continue
		}
		if score <= minScore {
			minScore = score
			matchCpuItems = append(matchCpuItems, item)
		}
	}

	var flavorRef, newCpu, ram string
	minScore = 1000000
	for _, item := range matchCpuItems {
		score := int64(item.Ram) - memory
		if score < 0 {
			continue
		}
		if score < minScore {
			minScore = score
			flavorRef = item.Speccode
			newCpu = item.Vcpus
			ram = strconv.Itoa(item.Ram)
		}
	}
	if flavorRef == "" {
		return "", "", "", fmt.Errorf("could not get suitable rds flavor")
	}

	logrus.Infof("get rds flavorRef=%v", flavorRef)
	return flavorRef, newCpu, ram, nil
}

func buildRdsVolumeStruct(input *RdsCreateInput) (*instances.Volume, error) {
	if input.VolumeType == "" {
		input.VolumeType = "ULTRAHIGH"
	}
	volume := instances.Volume{
		Type: input.VolumeType,
	}
	size, err := strconv.Atoi(input.VolumeSize)
	if err != nil {
		err = fmt.Errorf("strconv.Atoi(input.VolumeSize) meet err=%v", err)
		return nil, err
	}
	volume.Size = size
	return &volume, nil
}

func buildRdsHaStruct(input *RdsCreateInput) (*instances.Ha, error) {
	ha := instances.Ha{}
	if input.SupportHa == "" {
		return nil, nil
	}
	if strings.ToLower(input.SupportHa) == "true" {
		ha.Mode = "Ha"
		ha.ReplicationMode = input.HaReplicationMode
	}

	return &ha, nil
}

func buildChargeInfoStruct(input *RdsCreateInput) *instances.ChargeInfo {
	chargeInfo := instances.ChargeInfo{
		ChargeMode: input.ChargeType,
		IsAutoPay:  "true",
	}
	if input.ChargeType == PRE_PAID {
		chargeInfo.PeriodType = input.PeriodType
		if input.IsAutoRenew != "" {
			chargeInfo.IsAutoRenew = input.IsAutoRenew
		}
		chargeInfo.PeriodNum, _ = strconv.Atoi(input.PeriodNum)
	}
	return &chargeInfo
}

func (action *RdsCreateAction) createRds(input *RdsCreateInput) (output RdsCreateOutput, err error) {
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

	err = action.checkCreateRdsParams(*input)
	if err != nil {
		logrus.Errorf("RdsCreateAction checkCreateRdsParams meet error=%v", err)
		return
	}

	// create rds service client
	sc, err := createRdsServiceClientV3(input.CloudProviderParam)
	if err != nil {
		return
	}

	// check whether rds is exist.
	if input.Id != "" {
		var rdsInfo *instances.RdsInstanceResponse
		if rdsInfo, _, err = isRdsExist(sc, input.Id); err != nil {
			logrus.Errorf("check whether rds[Id=%v] is exist, meet error=%v", input.Id, err)
			return
		}
		if rdsInfo != nil {
			output.Id = rdsInfo.Id
			output.PrivateIp = rdsInfo.PrivateIps[0]
			output.Port = strconv.Itoa(rdsInfo.Port)
			output.UserName = rdsInfo.DbUserName
			return
		}
	}

	cloudMap, _ := GetMapFromString(input.CloudProviderParam.CloudParams)

	if input.Password == "" {
		input.Password = utils.CreateRandomPassword()
	}
	flavor, cpu, memory, err := getRdsFlavorByHostType(input)
	if err != nil {
		return
	}
	volume, err := buildRdsVolumeStruct(input)
	if err != nil {
		return
	}

	datastore := instances.Datastore{
		Type:    engineTypeMap["MYSQL"],
		Version: input.EngineVersion,
	}

	ha, err := buildRdsHaStruct(input)
	if err != nil {
		return
	}

	request := instances.CreateRdsOpts{
		Name:             input.Name,
		Datastore:        datastore,
		VpcId:            input.VpcId,
		SubnetId:         input.SubnetId,
		SecurityGroupId:  input.SecurityGroupId,
		Password:         input.Password,
		FlavorRef:        flavor,
		AvailabilityZone: input.AvailabilityZone,
		Region:           cloudMap[CLOUD_PARAM_REGION],
		Volume:           volume,
		ChargeInfo:       buildChargeInfoStruct(input),
		Ha:               ha,
	}
	if input.Port != "" {
		request.Port = input.Port
	}
	if input.EnterpriseProjectId != "" {
		request.EnterpriseProjectId = input.EnterpriseProjectId
	}

	response, err := instances.Create(sc, request).Extract()
	if err != nil {
		return
	}
	output.Id = response.Instance.Id
	instance, err := waitRdsInstanceJobOk(sc, response.Instance.Id, "create", 20)
	if err != nil {
		return
	}
	output.Port = strconv.Itoa(instance.Port)
	output.PrivateIp = instance.PrivateIps[0]
	output.UserName = instance.DbUserName
	output.Memory = memory
	output.Cpu = cpu

	password, err := utils.AesEnPassword(input.Guid, input.Seed, input.Password, utils.DEFALT_CIPHER)
	if err != nil {
		return
	}
	output.Password = password
	return

}

func (action *RdsCreateAction) Do(inputs interface{}) (interface{}, error) {
	rdss, _ := inputs.(RdsCreateInputs)
	outputs := RdsCreateOutputs{}
	var finalErr error
	for _, rds := range rdss.Inputs {
		output, err := action.createRds(&rds)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all rds instances = %v are created", rdss)
	return &outputs, finalErr
}

func waitRdsInstanceJobOk(sc *gophercloud.ServiceClient, id string, action string, times int) (*instances.RdsInstanceResponse, error) {
	var instance *instances.RdsInstanceResponse
	var err error
	count := 1
	for {
		instance, _, err = isRdsExist(sc, id)
		if err != nil {
			return nil, err
		}
		if instance == nil {
			if action == "delete" {
				return nil, nil
			}
			return nil, fmt.Errorf("get rds instance[id=%v] meet error=instance not be found", id)
		}

		if action == "create" {
			if instance.Status == RDS_INSTANCE_STATUS_BAD {
				break
			}
			if instance.Status == RDS_INSTANCE_STATUS_OK {
				return instance, nil
			}
		}

		if count > times {
			break
		}
		time.Sleep(30 * time.Second)
		count++
	}
	return nil, fmt.Errorf("after %vs, %v the rds instance[id=%v] status is %v", count*30, action, id, instance.Status)
}

func isRdsExist(sc *gophercloud.ServiceClient, rdsId string) (*instances.RdsInstanceResponse, bool, error) {
	allPages, err := instances.List(sc, instances.ListRdsInstanceOpts{
		Id: rdsId,
	}).AllPages()
	if err != nil {
		// if ue, ok := err.(*gophercloud.UnifiedError); ok {
		// 	if strings.Contains(ue.Message(), "could not be found") {
		// 		return nil, false, nil
		// 	}
		// }
		return nil, false, err
	}
	allRdsInstances, err := instances.ExtractRdsInstances(allPages)
	if err != nil {
		return nil, false, err
	}
	if len(allRdsInstances.Instances) == 0 {
		return nil, false, nil
	}
	if len(allRdsInstances.Instances) > 1 {
		return nil, false, fmt.Errorf("more than one rds instances have been returned")
	}
	return &allRdsInstances.Instances[0], true, nil
}

type RdsDeleteInputs struct {
	Inputs []RdsDeleteInput `json:"inputs,omitempty"`
}

type RdsDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type RdsDeleteOutputs struct {
	Outputs []RdsDeleteOutput `json:"outputs,omitempty"`
}

type RdsDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type RdsDeleteAction struct {
}

func (action *RdsDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs RdsDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *RdsDeleteAction) checkDeleteRdsParams(input RdsDeleteInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Id == "" {
		return fmt.Errorf("Rds id is empty")
	}

	return nil
}

func (action *RdsDeleteAction) deleteRds(input *RdsDeleteInput) (output RdsDeleteOutput, err error) {
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

	if err = action.checkDeleteRdsParams(*input); err != nil {
		logrus.Errorf("RdsDeleteAction checkDeleteRdsParams meet error=%v", err)
		return
	}

	// create rds service client
	sc, err := createRdsServiceClientV3(input.CloudProviderParam)
	if err != nil {
		return
	}

	// check whether rds is exist.
	rdsinfo, ok, err := isRdsExist(sc, input.Id)
	if err != nil {
		logrus.Errorf("check whether rds[Id=%v] is exist, meet error=%v", input.Id, err)
		return
	}
	if !ok {
		logrus.Infof("rds[Id=%v] is not exist", input.Id)
		return
	}

	// TODO: the prePaid need to do it specially
	if rdsinfo.ChargeInfo.ChargeMode == PRE_PAID {
		err = fmt.Errorf("can not support to delete rds instacne now")
		return
	}

	// delete rds instance
	_, err = instances.Delete(sc, input.Id).Extract()
	if err != nil {
		logrus.Errorf("delete rds[I=%v] meet error=%v", input.Id, err)
		return
	}
	_, err = waitRdsInstanceJobOk(sc, input.Id, "delete", 10)
	if err != nil {
		logrus.Errorf("waitRdsInstanceJobOk meet error=%v", err)
		return
	}
	return
}

func (action *RdsDeleteAction) Do(inputs interface{}) (interface{}, error) {
	rdss, _ := inputs.(RdsDeleteInputs)

	outputs := RdsDeleteOutputs{}
	var finalErr error
	for _, rds := range rdss.Inputs {
		output, err := action.deleteRds(&rds)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all rds instances = %v are deleted", rdss)
	return &outputs, finalErr
}

type RdsCreateBackupInputs struct {
	Inputs []RdsCreateBackupInput `json:"inputs,omitempty"`
}

type RdsCreateBackupInput struct {
	CallBackParameter
	CloudProviderParam
	Guid       string `json:"guid,omitempty"`
	Id         string `json:"id,omitempty"`
	InstanceId string `json:"instance_id,omitempty"`
	Name       string `json:"name,omitempty"`
}

type RdsCreateBackupOutputs struct {
	Outputs []RdsCreateBackupOutput `json:"outputs,omitempty"`
}

type RdsCreateBackupOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitepty"`
}

type RdsCreateBackupAction struct {
}

func (action *RdsCreateBackupAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs RdsCreateBackupInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *RdsCreateBackupAction) checkCreateBackupParams(input RdsCreateBackupInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.InstanceId == "" {
		return fmt.Errorf("instanceId is empty")
	}
	if input.Name == "" {
		return fmt.Errorf("name is empty")
	}

	return nil
}

func (action *RdsCreateBackupAction) createRdsBackup(input *RdsCreateBackupInput) (output RdsCreateBackupOutput, err error) {
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

	if err = action.checkCreateBackupParams(*input); err != nil {
		logrus.Errorf("RdsCreateBackupAction checkCreateBackupParams meet error=%v", err)
		return
	}

	// create rds service client
	sc, err := createRdsServiceClientV3(input.CloudProviderParam)
	if err != nil {
		return
	}

	// check whether backup is exist.
	if input.Id != "" {
		var backup *backups.BackupsResp
		backup, _, err = isRdsBackupExist(sc, input.Id, input.InstanceId)
		if err != nil {
			logrus.Errorf("check whether backup is exist meet error=%v", err)
			return
		}
		if backup != nil {
			logrus.Infof("backup[%v] is exist", input.Id)
			output.Id = input.Id
			return
		}
	}

	// create backup
	request := backups.CreateBackupsOpts{
		InstanceId: input.InstanceId,
		Name:       input.Name,
	}
	response, err := backups.Create(sc, request).Extract()
	if err != nil {
		logrus.Errorf("create backup meet error=%v", err)
		return
	}
	output.Id = response.Id

	_, err = waitRdsBackupJobOk(sc, response.Id, response.Instanceid, "create", 20)
	if err != nil {
		logrus.Errorf("waitRdsBackupJobOk meet error=%v", err)
		return
	}

	return
}

func (action *RdsCreateBackupAction) Do(inputs interface{}) (interface{}, error) {
	backups, _ := inputs.(RdsCreateBackupInputs)
	outputs := RdsCreateBackupOutputs{}
	var finalErr error
	for _, backup := range backups.Inputs {
		output, err := action.createRdsBackup(&backup)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all backups = %v are created", backups)
	return &outputs, finalErr
}

func isRdsBackupExist(sc *gophercloud.ServiceClient, backupId, instanceId string) (*backups.BackupsResp, bool, error) {
	request := backups.ListBackupsOpts{
		InstanceId: instanceId,
		BackupId:   backupId,
	}
	allPages, err := backups.List(sc, request).AllPages()
	if err != nil {
		return nil, false, err
	}
	backupResp, err := backups.ExtractBackups(allPages)
	if err != nil {
		return nil, false, err
	}
	if len(backupResp.Backups) == 0 {
		return nil, false, nil
	}
	if len(backupResp.Backups) > 1 {
		return nil, false, fmt.Errorf("more than one backup had been returned")
	}

	return &backupResp.Backups[0], true, nil
}

func waitRdsBackupJobOk(sc *gophercloud.ServiceClient, backupId, instanceId, action string, times int) (*backups.BackupsResp, error) {
	var backup *backups.BackupsResp
	var err error
	count := 1
	for {
		backup, _, err = isRdsBackupExist(sc, backupId, instanceId)
		if err != nil {
			return nil, err
		}

		if backup == nil {
			if action == "delete" {
				return nil, nil
			}
			// if action == "create", prove	that could not found the backup
			return nil, fmt.Errorf("get rds backup[id=%v] meet error=instance not be found", backupId)
		}

		if action == "create" {
			if backup.Status == RDS_BACKUP_STATUS_BAD {
				break
			}
			if backup.Status == RDS_BACKUP_STATUS_OK {
				return backup, nil
			}
		}
		if count > times {
			break
		}
		time.Sleep(30 * time.Second)
		count++
	}
	return nil, fmt.Errorf("after %vs, %v the rds backup[id=%v] status is %v", count*30, action, backupId, backup.Status)
}

type RdsDeleteBackupInputs struct {
	Inputs []RdsDeleteBackupInput `json:"inputs,omitempty"`
}

type RdsDeleteBackupInput struct {
	CallBackParameter
	CloudProviderParam
	Guid       string `json:"guid,omitempty"`
	Id         string `json:"id,omitempty"`
	InstanceId string `json:"instance_id,omitempty"`
}

type RdsDeleteBackupOutputs struct {
	Outputs []RdsDeleteBackupOutput `json:"outputs,omitempty"`
}

type RdsDeleteBackupOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type RdsDeleteBackupAction struct {
}

func (action *RdsDeleteBackupAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs RdsDeleteBackupInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *RdsDeleteBackupAction) checkDeleteBackupParams(input RdsDeleteBackupInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Id == "" {
		return fmt.Errorf("Rds id is empty")
	}
	if input.InstanceId == "" {
		return fmt.Errorf("Rds instanceId is empty")
	}

	return nil
}

func (action *RdsDeleteBackupAction) deleteBackup(input *RdsDeleteBackupInput) (output RdsDeleteBackupOutput, err error) {
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

	if err = action.checkDeleteBackupParams(*input); err != nil {
		logrus.Errorf("RdsDeleteBackupAction checkDeleteBackupParams meet error=%v", err)
		return
	}

	// create rds service client
	sc, err := createRdsServiceClientV3(input.CloudProviderParam)
	if err != nil {
		return
	}

	// check whether backup is exist.
	_, ok, err := isRdsBackupExist(sc, input.Id, input.InstanceId)
	if err != nil {
		logrus.Errorf("check whether backup is exist meet error=%v", err)
		return
	}
	if !ok {
		logrus.Infof("backup[%v] is not exist", input.Id)
		output.Id = input.Id
		return
	}

	// delete backup
	err = backups.Delete(sc, input.Id).ExtractErr()
	if err != nil {
		logrus.Errorf("delete backup meet error=%v", err)
		return
	}

	_, err = waitRdsBackupJobOk(sc, input.Id, input.InstanceId, "delete", 10)
	if err != nil {
		logrus.Errorf("waitRdsBackupJobOk meet error=%v", err)
		return
	}

	logrus.Infof("delete backup[%v] done", input.Id)
	return
}

func (action *RdsDeleteBackupAction) Do(inputs interface{}) (interface{}, error) {
	backups, _ := inputs.(RdsDeleteBackupInputs)
	outputs := RdsDeleteBackupOutputs{}
	var finalErr error
	for _, backup := range backups.Inputs {
		output, err := action.deleteBackup(&backup)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all backups = %v are deleted", backups)
	return &outputs, finalErr
}
