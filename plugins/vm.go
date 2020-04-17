package plugins

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/WeBankPartners/wecube-plugins-huaweicloud/plugins/utils"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/images"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/servers"
	v1 "github.com/gophercloud/gophercloud/openstack/ecs/v1/cloudservers"
	flavor "github.com/gophercloud/gophercloud/openstack/ecs/v1/flavor"
	v1_1 "github.com/gophercloud/gophercloud/openstack/ecs/v1_1/cloudservers"
	v2 "github.com/gophercloud/gophercloud/openstack/ecs/v2/cloudservers"
	"github.com/sirupsen/logrus"
)

const (
	PRE_PAID       = "prePaid"  //包年包月
	POST_PAID      = "postPaid" //按量计费
	PRE_PAID_MONTH = "month"
	PRE_PAID_YEAR  = "year"

	CLOUD_SERVER_V1   = "v1"
	CLOUD_SERVER_V1_1 = "v1_1"
	CLOUD_SERVER_V2   = "v2"

	MEMORY_UNIT = "G"
	CPU_UNIT    = "C"
)

var vmActions = make(map[string]Action)

func init() {
	vmActions["create"] = new(VmCreateAction)
	vmActions["terminate"] = new(VmDeleteAction)
	vmActions["start"] = new(VmStartAction)
	vmActions["stop"] = new(VmStopAction)
	vmActions["bind-security-groups"] = new(VmBindSecurityGroupsAction)
}

type VmPlugin struct {
}

func (plugin *VmPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := vmActions[actionName]
	if !found {
		return nil, fmt.Errorf("vmplugin,action[%s] not found", actionName)
	}
	return action, nil
}

func createVmServiceClient(params CloudProviderParam, version string) (*gophercloud.ServiceClient, error) {
	provider, err := createGopherCloudProviderClient(params)
	if err != nil {
		logrus.Errorf("Get gophercloud provider client failed, error=%v", err)
		return nil, err
	}

	switch version {
	case CLOUD_SERVER_V1:
		return openstack.NewECSV1(provider, gophercloud.EndpointOpts{})
	case CLOUD_SERVER_V1_1:
		return openstack.NewECSV1_1(provider, gophercloud.EndpointOpts{})
	case CLOUD_SERVER_V2:
		return openstack.NewECSV2(provider, gophercloud.EndpointOpts{})
	}

	return nil, fmt.Errorf("version(%v) is not support", version)
}

type VmCreateInputs struct {
	Inputs []VmCreateInput `json:"inputs,omitempty"`
}

type VmCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`

	Seed             string `json:"seed,omitempty"`
	ImageId          string `json:"image_id,omitempty"`
	HostType         string `json:"machine_spec,omitempty"` //4c8g
	SystemDiskSize   string `json:"system_disk_size,omitempty"`
	SystemDiskType   string `json:"system_disk_type,omitempty"`
	VpcId            string `json:"vpc_id,omitempty"`
	SubnetId         string `json:"subnet_id,omitempty"`
	PrivateIp        string `json:"private_ip,omitempty"`
	Name             string `json:"name,omitempty"`
	Password         string `json:"password,omitempty"`
	Labels           string `json:"labels,omitempty"`
	AvailabilityZone string `json:"az,omitempty"`
	SecurityGroups   string `json:"security_group,omitempty"`

	ChargeType string `json:"charge_type,omitempty"`

	//包年包月
	PeriodType  string `json:"period_type,omitempty"`   //年或月
	PeriodNum   string `json:"period_num,omitempty"`    //年有效值[1-9],月有效值[1-3]
	IsAutoRenew string `json:"is_auto_renew,omitempty"` //是否自动续费

	EnterpriseProjectId string `json:"enterprise_project_id,omitempty"`
}

type VmCreateOutputs struct {
	Outputs []VmCreateOutput `json:"outputs,omitempty"`
}

type VmCreateOutput struct {
	CallBackParameter
	Result
	Guid      string `json:"guid,omitempty"`
	Id        string `json:"id,omitempty"`
	Cpu       string `json:"cpu,omitempty"`
	Memory    string `json:"memory,omitempty"`
	Password  string `json:"password,omitempty"`
	PrivateIp string `json:"private_ip,omitempty"`
}

type VmCreateAction struct {
}

func (action *VmCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs VmCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func isValidSystemDiskType(systemDiskType string) error {
	validDisks := []string{
		"SATA",  //普通IO磁盘类型。
		"SAS",   //高IO磁盘类型。
		"SSD",   //超高IO磁盘类型。
		"co-p1", //高IO (性能优化Ⅰ型)
		"uh-l1", //超高IO (时延优化)
	}

	return isValidStringValue("systemDiskType", systemDiskType, validDisks)
}

func checkVmCreateParams(input VmCreateInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Seed == "" {
		return fmt.Errorf("seed is empty")
	}
	if input.ImageId == "" {
		return fmt.Errorf("imageId is empty")
	}
	if input.HostType == "" {
		return fmt.Errorf("hostType is empty")
	}
	if input.SystemDiskSize == "" {
		return fmt.Errorf("systemDiskSize is empty")
	}
	if input.VpcId == "" {
		return fmt.Errorf("vpcId is empty")
	}
	if input.SubnetId == "" {
		return fmt.Errorf("subnetId is empty")
	}
	if input.Name == "" {
		return fmt.Errorf("name is empty")
	}
	if input.AvailabilityZone == "" {
		return fmt.Errorf("availabilityZone is empty")
	}
	if err := isValidStringValue("chargeType", input.ChargeType, []string{PRE_PAID, POST_PAID}); err != nil {
		return err
	}

	if input.SystemDiskType != "" {
		if err := isValidSystemDiskType(input.SystemDiskType); err != nil {
			return err
		}
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

func isVmExist(cloudProviderParam CloudProviderParam, id string) (*v1.CloudServer, bool, error) {
	vmInfo, err := getVmInfoById(cloudProviderParam, id)
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "could not be found") {
				return nil, false, nil
			}
		}
		return nil, false, err
	}
	if vmInfo.Status == "DELETED" {
		return vmInfo, false, nil
	}
	return vmInfo, true, nil
}

func getVmInfoById(cloudProviderParam CloudProviderParam, id string) (*v1.CloudServer, error) {
	sc, err := createVmServiceClient(cloudProviderParam, CLOUD_SERVER_V1)
	if err != nil {
		return nil, err
	}

	vmInfo, err := v1.Get(sc, id).Extract()
	if err != nil {
		logrus.Errorf("getvmInfoById failed err=%v\n", err)
	}
	return vmInfo, err
}

func getIpFromVmInfo(vm *v1.CloudServer) (string, error) {
	for _, addresses := range vm.Addresses {
		for _, address := range addresses {
			return address.Addr, nil
		}
	}
	logrus.Errorf("can't get vm(%v) lan ip", vm.ID)
	return "", fmt.Errorf("can't get vm(%v) lan ip", vm.ID)
}

func getVmIpAddress(cloudProviderParam CloudProviderParam, id string) (string, error) {
	vmInfo, err := getVmInfoById(cloudProviderParam, id)
	if err != nil {
		return "", err
	}

	return getIpFromVmInfo(vmInfo)
}

func buildVmNicStruct(input VmCreateInput) []v1_1.Nic {
	nic := v1_1.Nic{
		SubnetId: input.SubnetId,
	}
	if input.PrivateIp != "" {
		nic.IpAddress = input.PrivateIp
	}

	return []v1_1.Nic{nic}
}

func buildServerExtendParam(input VmCreateInput) v1_1.ServerExtendParam {
	param := v1_1.ServerExtendParam{
		ChargingMode: input.ChargeType,
	}
	if input.EnterpriseProjectId != "" {
		param.EnterpriseProjectID = input.EnterpriseProjectId
	}

	if input.ChargeType == PRE_PAID {
		param.PeriodType = input.PeriodType
		param.IsAutoPay = "true"
		if input.IsAutoRenew != "" {
			param.IsAutoRenew = input.IsAutoRenew
		}
		param.PeriodNum, _ = strconv.Atoi(input.PeriodNum)
	}
	return param
}

func buildSecurityGroups(securityGroups string) []v1_1.SecurityGroup {
	scs := []v1_1.SecurityGroup{}
	if securityGroups != "" {
		sc := v1_1.SecurityGroup{
			ID: securityGroups,
		}
		scs = append(scs, sc)
	}
	return scs
}

func buildServerTags(labels string) []v1_1.ServerTags {
	tags := []v1_1.ServerTags{}
	labelMap, _ := GetMapFromString(labels)
	for k, v := range labelMap {
		tag := v1_1.ServerTags{
			Key:   k,
			Value: v,
		}
		tags = append(tags, tag)
	}
	return tags
}

func buildRootVolumeStruct(input VmCreateInput) (v1_1.RootVolume, error) {
	volume := v1_1.RootVolume{
		VolumeType: "SATA",
	}

	if input.SystemDiskType != "" {
		volume.VolumeType = input.SystemDiskType
	}

	rootSize, err := strconv.Atoi(input.SystemDiskSize)
	if err != nil {
		return volume, err
	}
	volume.Size = rootSize
	return volume, nil
}

func getCpuAndMemoryFromHostType(hostType string) (int64, int64, error) {
	//1C2G, 2C4G, 2C8G
	upperCase := strings.ToUpper(hostType)
	index := strings.Index(upperCase, CPU_UNIT)
	if index <= 0 {
		return 0, 0, fmt.Errorf("hostType(%v) invalid", hostType)
	}
	cpu, err := strconv.ParseInt(upperCase[0:index], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("hostType(%v) invalid", hostType)
	}

	memStr := upperCase[index+1:]
	index2 := strings.Index(memStr, MEMORY_UNIT)
	if index2 <= 0 {
		return 0, 0, fmt.Errorf("hostType(%v) invalid", hostType)
	}

	mem, err := strconv.ParseInt(memStr[0:index2], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("hostType(%v) invalid", hostType)
	}
	return cpu, mem, nil
}

func getFlavorByHostType(input VmCreateInput) (string, int64, int64, error) {
	var matchedCpu int64
	var matchedMem int64

	cpu, memory, err := getCpuAndMemoryFromHostType(input.HostType)
	if err != nil {
		return "", matchedCpu, matchedMem, err
	}
	listOpts := flavor.ListOpts{
		AvailabilityZone: input.AvailabilityZone,
	}

	sc, err := createVmServiceClient(input.CloudProviderParam, CLOUD_SERVER_V1)
	if err != nil {
		return "", matchedCpu, matchedMem, err
	}

	allPages, err := flavor.List(sc, listOpts).AllPages()
	if err != nil {
		return "", matchedCpu, matchedMem, err
	}

	flavors, err := flavor.ExtractFlavors(allPages)
	if err != nil {
		return "", matchedCpu, matchedMem, err
	}

	var minScore int64 = 1000000
	matchCpuItems := []flavor.Flavor{}
	for _, item := range flavors {
		if item.FlvDisabled == true || item.AccessIsPublic == false {
			continue
		}
		// status := item.OsExtraSpecs.CondOperationStatus
		// if status != "normal" && status != "promotion" {
		// 	continue
		// }
		vcpus, err := strconv.ParseInt(item.Vcpus, 10, 64)
		if err != nil {
			logrus.Errorf("vpus(%v) is invald", item.Vcpus)
			return "", matchedCpu, matchedMem, err
		}

		score := vcpus - cpu
		if score < 0 {
			continue
		}
		if score <= minScore {
			minScore = score
			matchedCpu = vcpus
			matchCpuItems = append(matchCpuItems, item)
		}
	}

	instanceType := ""
	minScore = 1000000
	for _, item := range matchCpuItems {
		score := int64(item.Ram)/1024 - memory
		if score < 0 {
			continue
		}
		if score < minScore {
			minScore = score
			matchedMem = item.Ram / 1024
			instanceType = item.ID
		}
	}
	if instanceType == "" {
		return "", matchedCpu, matchedMem, fmt.Errorf("could not get suitable instancetype")
	}

	logrus.Infof("get instancetype=%v", instanceType)
	return instanceType, matchedCpu, matchedMem, nil
}

func waitVmJobOk(sc *gophercloud.ServiceClient, jobId string) (string, error) {
	var jobRst v1_1.JobResult
	count := 0

	for {
		time.Sleep(time.Duration(6) * time.Second)
		job, getJobErr := v1_1.GetJobResult(sc, jobId)
		if getJobErr != nil {
			logrus.Errorf("getJobResult failed err =%v", getJobErr)
			return "", getJobErr
		}

		if strings.Compare("SUCCESS", job.Status) == 0 {
			jobRst = job
			break
		} else if strings.Compare("FAIL", job.Status) == 0 {
			jobRst = job
			break
		}
		count++
		//5 minutes
		if count > 50 {
			return "", fmt.Errorf("vm job still in %v statue", job.Status)
		}
	}
	subJobs := jobRst.Entities.SubJobs
	for _, value := range subJobs {
		if strings.Compare("SUCCESS", value.Status) == 0 {
			return value.Entities.ServerId, nil
		} else {
			return "", fmt.Errorf("Vm job failed")
		}
	}
	return "", fmt.Errorf("can't go to here")
}

func createVm(input VmCreateInput) (output VmCreateOutput, err error) {
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

	if err = checkVmCreateParams(input); err != nil {
		return
	}

	if input.Id != "" {
		exist := false
		_, exist, err = isVmExist(input.CloudProviderParam, input.Id)
		if err == nil && exist {
			output.Id = input.Id
			output.PrivateIp, _ = getVmIpAddress(input.CloudProviderParam, input.Id)
			return
		}
	}

	//now create vm
	nics := buildVmNicStruct(input)
	tags := buildServerTags(input.Labels)
	securityGroups := buildSecurityGroups(input.SecurityGroups)
	serverExtendParam := buildServerExtendParam(input)
	rootVolume, err := buildRootVolumeStruct(input)
	if err != nil {
		return
	}

	flavor, matchedCpu, matchedMem, err := getFlavorByHostType(input)
	if err != nil {
		return
	}
	output.Cpu = fmt.Sprintf("%v", matchedCpu)
	output.Memory = fmt.Sprintf("%v", matchedMem)

	opts := v1_1.CreateOpts{
		Name:             input.Name,
		FlavorRef:        flavor,
		ImageRef:         input.ImageId,
		VpcId:            input.VpcId,
		Nics:             nics,
		RootVolume:       rootVolume,
		AvailabilityZone: input.AvailabilityZone,
		Count:            1,
		ExtendParam:      &serverExtendParam,
	}
	if input.Password == "" {
		input.Password = utils.CreateRandomPassword()
	}
	opts.AdminPass = input.Password
	if len(tags) > 0 {
		opts.ServerTags = tags
	}
	if len(securityGroups) > 0 {
		opts.SecurityGroups = securityGroups
	}

	sc, err := createVmServiceClient(input.CloudProviderParam, CLOUD_SERVER_V1_1)
	if err != nil {
		return
	}

	jobId, _, err := v1_1.Create(sc, opts)
	if err != nil {
		return
	}

	output.Id, err = waitVmJobOk(sc, jobId)
	if err != nil {
		return
	}

	output.Password, err = utils.AesEnPassword(input.Guid, input.Seed, input.Password, utils.DEFALT_CIPHER)
	if err != nil {
		return
	}
	output.PrivateIp, err = getVmIpAddress(input.CloudProviderParam, output.Id)

	return
}

func (action *VmCreateAction) Do(inputs interface{}) (interface{}, error) {
	vms, _ := inputs.(VmCreateInputs)
	outputs := VmCreateOutputs{}
	var finalErr error

	for _, input := range vms.Inputs {
		output, err := createVm(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all vms= %v are created", vms)
	return &outputs, finalErr
}

type VmDeleteInputs struct {
	Inputs []VmDeleteInput `json:"inputs,omitempty"`
}

type VmDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type VmDeleteOutputs struct {
	Outputs []VmDeleteOutput `json:"outputs,omitempty"`
}

type VmDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type VmDeleteAction struct {
}

func (action *VmDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs VmDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func waitVmDeleteOk(cloudProviderParam CloudProviderParam, id string) {
	count := 0
	for {
		time.Sleep(time.Second * 5)
		_, exist, err := isVmExist(cloudProviderParam, id)
		if err != nil || !exist {
			break
		}

		count++
		if count > 10 {
			break
		}
	}
}

func deleteVm(input VmDeleteInput) (output VmDeleteOutput, err error) {
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

	vmInfo, exist, err := isVmExist(input.CloudProviderParam, input.Id)
	if err != nil || !exist {
		return
	}

	provider, err := createGopherCloudProviderClient(input.CloudProviderParam)
	if err != nil {
		logrus.Errorf("Get gophercloud provider client failed, error=%v", err)
		return
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		return
	}

	// check whether the type of vm is prePaid
	// TODO: the prePaid need to do it specially
	if vmInfo.Metadata.ChargingMode == PRE_PAID {
		err = fmt.Errorf("can not support to delete prePaid vm now")
		return
	}

	if err = servers.Delete(client, input.Id).ExtractErr(); err != nil {
		logrus.Errorf("delete vm(%v) failed ,err=%v", input.Id, err)
	}

	waitVmDeleteOk(input.CloudProviderParam, input.Id)

	return
}

func (action *VmDeleteAction) Do(inputs interface{}) (interface{}, error) {
	vms, _ := inputs.(VmDeleteInputs)
	outputs := VmDeleteOutputs{}
	var finalErr error

	for _, input := range vms.Inputs {
		output, err := deleteVm(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all vms= %v are delete", vms)
	return &outputs, finalErr
}

type VmStartInput VmDeleteInput
type VmStartInputs struct {
	Inputs []VmStartInput `json:"inputs,omitempty"`
}

type VmStartOutput VmDeleteOutput
type VmStartOutputs struct {
	Outputs []VmStartOutput `json:"outputs,omitempty"`
}

type VmStartAction struct {
}

func (action *VmStartAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs VmStartInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func startVm(input VmStartInput) (output VmStartOutput, err error) {
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

	if err = isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return
	}
	if input.Id == "" {
		err = fmt.Errorf("id is empty")
		return
	}

	sc, err := createVmServiceClient(input.CloudProviderParam, CLOUD_SERVER_V1)
	if err != nil {
		return
	}

	opts := v1.BatchStartOpts{
		Servers: []v1.Server{
			{ID: input.Id},
		},
	}

	resp, err := v1.BatchStart(sc, opts).ExtractJob()
	if err != nil {
		return
	}

	if _, err = waitVmJobOk(sc, resp.ID); err != nil {
		logrus.Errorf("wait start job failed,err=%v", err)
	}

	return
}

func (action *VmStartAction) Do(inputs interface{}) (interface{}, error) {
	vms, _ := inputs.(VmStartInputs)
	outputs := VmStartOutputs{}
	var finalErr error

	for _, input := range vms.Inputs {
		output, err := startVm(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all vms= %v are start", vms)
	return &outputs, finalErr
}

type VmStopInput VmDeleteInput
type VmStopInputs struct {
	Inputs []VmStopInput `json:"inputs,omitempty"`
}

type VmStopOutput VmDeleteOutput
type VmStopOutputs struct {
	Outputs []VmStopOutput `json:"outputs,omitempty"`
}

type VmStopAction struct {
}

func (action *VmStopAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs VmStopInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func stopVm(input VmStopInput) (output VmStopOutput, err error) {
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

	if err = isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return
	}
	if input.Id == "" {
		err = fmt.Errorf("id is empty")
		return
	}

	sc, err := createVmServiceClient(input.CloudProviderParam, CLOUD_SERVER_V1)
	if err != nil {
		return
	}

	opts := v1.BatchStopOpts{
		Type: v1.Type(v1.Hard),
		Servers: []v1.Server{
			{ID: input.Id},
		},
	}

	resp, err := v1.BatchStop(sc, opts).ExtractJob()
	if err != nil {
		return
	}

	if _, err = waitVmJobOk(sc, resp.ID); err != nil {
		logrus.Errorf("wait stop job failed,err=%v", err)
	}
	return
}

func (action *VmStopAction) Do(inputs interface{}) (interface{}, error) {
	vms, _ := inputs.(VmStopInputs)
	outputs := VmStopOutputs{}
	var finalErr error

	for _, input := range vms.Inputs {
		output, err := stopVm(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all vms= %v are stop", vms)
	return &outputs, finalErr
}

func PrintImages(params CloudProviderParam) {
	provider, err := createGopherCloudProviderClient(params)
	if err != nil {
		fmt.Printf("print Image create provider failed,err=%v\n", err)
		return
	}

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		fmt.Printf("print image new computeV2 failed,err=%v\n", err)
		return
	}

	listOpts := images.ListOpts{
		Status: "active",
	}
	// Query all images list information
	allPages, allPagesErr := images.ListDetail(client, listOpts).AllPages()
	if allPagesErr != nil {
		fmt.Println("allPagesErr:", allPagesErr)
		if ue, ok := allPagesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}
	// Transform images structure
	allImages, allImagesErr := images.ExtractImages(allPages)
	if allImagesErr != nil {
		fmt.Println("allImagesErr:", allImagesErr)
		if ue, ok := allImagesErr.(*gophercloud.UnifiedError); ok {
			fmt.Println("ErrCode:", ue.ErrorCode())
			fmt.Println("Message:", ue.Message())
		}
		return
	}

	for _, image := range allImages {
		if strings.Contains(strings.ToUpper(image.Name), "CENTOS") {
			fmt.Printf("imageName=%v,imageId=%v\n", image.Name, image.ID)
		}
	}
}

func GetVmSecurityGroups(params CloudProviderParam, serverId string) ([]string, error) {
	securityGroups := []string{}
	sc, err := createVmServiceClient(params, CLOUD_SERVER_V2)
	if err != nil {
		return nil, err
	}

	result, err := v2.GetSecurityGroups(sc, serverId).Extract()
	if err != nil {
		return securityGroups, err
	}

	for _, securityGroup := range result.SecurityGroups {
		securityGroups = append(securityGroups, securityGroup.ID)
	}

	return securityGroups, nil
}

func DeleteSecurityGroup(params CloudProviderParam, serverId string, securityGroupId string) error {
	sc, err := createVmServiceClient(params, CLOUD_SERVER_V2)
	if err != nil {
		return err
	}

	result := v2.RemoveSecurityGroup(sc, serverId, securityGroupId)
	return result.Err
}

func AddSecurityGroup(params CloudProviderParam, serverId string, securityGroupId string) error {
	sc, err := createVmServiceClient(params, CLOUD_SERVER_V2)
	if err != nil {
		return err
	}

	result := v2.AddSecurityGroup(sc, serverId, securityGroupId)
	return result.Err
}

type VmBindSecurityGroupsAction struct {
}

type VmBindSecurityGroupsInputs struct {
	Inputs []VmBindSecurityGroupsInput `json:"inputs,omitempty"`
}

type VmBindSecurityGroupsInput struct {
	CallBackParameter
	CloudProviderParam
	Guid           string `json:"guid,omtempty"`
	Id             string `json:"id,omtempty"`
	SecurityGroups string `json:"security_groups,omitemoty"`
}

type VmBindSecurityGroupsOutputs struct {
	Outputs []VmBindSecurityGroupsOutput `json:"outputs,omitempty"`
}

type VmBindSecurityGroupsOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omtempty"`
}

func (action *VmBindSecurityGroupsAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs VmBindSecurityGroupsInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkVmBindSecurityGoupsParam(input VmBindSecurityGroupsInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Id == "" {
		return fmt.Errorf("id is empty")
	}
	if input.SecurityGroups == "" {
		return fmt.Errorf("security_groups is empty")
	}
	return nil
}

func getSecurityGroupsByVm(sc *gophercloud.ServiceClient, id string) ([]v2.SecurityGroup, error) {
	sgs, err := v2.GetSecurityGroups(sc, id).Extract()
	if err != nil {
		return nil, err
	}

	return sgs.SecurityGroups, nil
}

func deleteVmSecurityGoups(sc *gophercloud.ServiceClient, id string, securityGroups []v2.SecurityGroup) error {
	for _, sg := range securityGroups {
		err := v2.RemoveSecurityGroup(sc, id, sg.ID).Err
		if err != nil {
			return err
		}
	}

	return nil
}

func addVmSecurityGoups(sc *gophercloud.ServiceClient, id string, securityGroupIds []string) error {
	for _, sg := range securityGroupIds {
		err := v2.AddSecurityGroup(sc, id, sg).Err
		if err != nil {
			return err
		}
	}
	return nil
}

func vmBindSecurityGoups(input *VmBindSecurityGroupsInput) (output VmBindSecurityGroupsOutput, err error) {
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

	if err = checkVmBindSecurityGoupsParam(*input); err != nil {
		return
	}

	sc, err := createVmServiceClient(input.CloudProviderParam, CLOUD_SERVER_V2)
	if err != nil {
		return
	}

	// do input.SecurityGoups to []string
	sgIds, err := GetArrayFromString(input.SecurityGroups, ARRAY_SIZE_REAL, 0)
	if err != nil {
		return
	}

	// check wether input.SecurityGoups exist
	vpcSc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}
	for _, sgId := range sgIds {
		var exist bool
		_, exist, err = isSecurityGroupExist(vpcSc, sgId)
		if err != nil {
			return
		}

		if !exist {
			err = fmt.Errorf("securityGroup[%v] is not exist", sgId)
			return
		}
	}

	// get all security groups of the vm
	sgs, err := getSecurityGroupsByVm(sc, input.Id)
	if err != nil {
		return
	}

	// remove all security groups of the vm
	if err = deleteVmSecurityGoups(sc, input.Id, sgs); err != nil {
		return
	}

	// add input.SecurityGoups to vm
	if err = addVmSecurityGoups(sc, input.Id, sgIds); err != nil {
		return
	}

	return
}

func (action *VmBindSecurityGroupsAction) Do(inputs interface{}) (interface{}, error) {
	vms, _ := inputs.(VmBindSecurityGroupsInputs)
	outputs := VmBindSecurityGroupsOutputs{}
	var finalErr error

	for _, input := range vms.Inputs {
		output, err := vmBindSecurityGoups(&input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all securityGoups had been bind, input = %++v", vms)
	return &outputs, finalErr
}
