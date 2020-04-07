package plugins

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/WeBankPartners/wecube-plugins-huaweicloud/plugins/utils"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack"
	"github.com/gophercloud/gophercloud/openstack/blockstorage/v2/volumes"
	"github.com/gophercloud/gophercloud/openstack/compute/v2/extensions/volumeattach"
	"github.com/sirupsen/logrus"
)

const (
	DISK_TYPE_SSD  = "SSD"
	DISK_TYPE_SAS  = "SAS"
	DISK_TYPE_SATA = "SATA"
)

var blockStorageActions = make(map[string]Action)

func createBlockStorageServiceClient(params CloudProviderParam) (*gophercloud.ServiceClient, error) {
	provider, err := createGopherCloudProviderClient(params)
	if err != nil {
		logrus.Errorf("Get gophercloud provider client failed, error=%v", err)
		return nil, err
	}

	sc, err := openstack.NewBlockStorageV2(provider, gophercloud.EndpointOpts{})
	if err != nil {
		logrus.Errorf("newBlockStorageV2 failed, error=%v", err)
		return nil, err
	}

	return sc, nil
}

func init() {
	blockStorageActions["create-mount"] = new(CreateAndMountDiskAction)
	blockStorageActions["umount-delete"] = new(UmountAndTerminateDiskAction)
}

type BlockStoragePlugin struct {
}

func (plugin *BlockStoragePlugin) GetActionByName(actionName string) (Action, error) {
	action, found := blockStorageActions[actionName]
	if !found {
		return nil, fmt.Errorf("blockStorage plugin,action = %s not found", actionName)
	}
	return action, nil
}

type CreateAndMountDiskAction struct {
}

type CreateAndMountDiskInputs struct {
	Inputs []CreateAndMountDiskInput `json:"inputs,omitempty"`
}

type CreateAndMountDiskInput struct {
	CallBackParameter
	CloudProviderParam
	Guid             string `json:"guid,omitempty"`
	AvailabilityZone string `json:"az,omitempty"`
	DiskType         string `json:"disk_type,omitempty"`
	DiskSize         string `json:"disk_size,omitempty"`
	Id               string `json:"id,omitempty"`
	Name             string `json:"name,omitempty"`

	//use to attach and format
	InstanceId       string `json:"instance_id,omitempty"`
	InstanceGuid     string `json:"instance_guid,omitempty"`
	InstanceSeed     string `json:"seed,omitempty"`
	InstancePassword string `json:"password,omitempty"`

	FileSystemType string `json:"file_system_type,omitempty"`
	MountDir       string `json:"mount_dir,omitempty"`

	SkipMount string
}

type CreateAndMountDiskOutputs struct {
	Outputs []CreateAndMountDiskOutput `json:"outputs,omitempty"`
}

type CreateAndMountDiskOutput struct {
	CallBackParameter
	Result
	Guid       string `json:"guid,omitempty"`
	VolumeName string `json:"volume_name,omitempty"`
	Id         string `json:"id,omitempty"`
	AttachId   string `json:"attach_id,omitempty"`
}

func (action *CreateAndMountDiskAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs CreateAndMountDiskInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkCreateAndMountParam(input CreateAndMountDiskInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}

	//instance related
	if input.InstanceId == "" {
		return fmt.Errorf("empty instanceId")
	}
	if input.InstanceGuid == "" {
		return fmt.Errorf("empty instanceGuid")
	}
	if input.InstanceSeed == "" {
		return fmt.Errorf("empty InstanceSeed")
	}
	if input.InstancePassword == "" {
		return fmt.Errorf("empty instancePassword")
	}

	//buy disk related
	if input.AvailabilityZone == "" {
		return fmt.Errorf("empty availabilityZone")
	}
	if _, err := strconv.Atoi(input.DiskSize); err != nil {
		return err
	}

	validDiskTypes := []string{
		DISK_TYPE_SSD,
		DISK_TYPE_SAS,
		DISK_TYPE_SATA,
	}
	if err := isValidStringValue("diskType", input.DiskType, validDiskTypes); err != nil {
		return err
	}

	if input.MountDir == "" {
		return fmt.Errorf(" mountDir is empty")
	}

	if err := isValidStringValue("fileSystemType", input.FileSystemType, []string{"ext3", "ext4", "xfs"}); err != nil {
		return fmt.Errorf("%s is not valid file system type", input.FileSystemType)
	}
	return nil
}

func waitVolumeInDesireState(sc *gophercloud.ServiceClient, id string, desireState string) error {
	for {
		time.Sleep(time.Duration(5) * time.Second)
		volume, err := volumes.Get(sc, id).Extract()
		if err != nil {
			return err
		}

		logrus.Infof("waitVolumeInDesireState,now status =%v", volume.Status)

		if volume.Status == "error" {
			return fmt.Errorf("waitVolumeInDesireState,meet status ==ERROR")
		}

		if volume.Status == desireState {
			break
		}
	}
	return nil
}

func waitVolumeCreateOk(sc *gophercloud.ServiceClient, id string) error {
	return waitVolumeInDesireState(sc, id, "available")
}

func waitVolumeInAvailableState(sc *gophercloud.ServiceClient, id string) error {
	return waitVolumeInDesireState(sc, id, "available")
}

func waitVolumeAttachOk(sc *gophercloud.ServiceClient, id string) error {
	return waitVolumeInDesireState(sc, id, "in-use")
}

func attachVolumeToVm(input CreateAndMountDiskInput, volumeId string, instanceId string) (string, string, error) {
	sc, err := createComputeV2Client(input.CloudProviderParam)
	if err != nil {
		return "", "", err
	}

	volumeAttachOptions := volumeattach.CreateOpts{
		VolumeID: volumeId,
	}
	resp, err := volumeattach.Create(sc, instanceId, volumeAttachOptions).Extract()
	if err != nil {
		return "", "", err
	}
	return resp.ID, resp.Device, nil
}

func buyDiskAndAttachToVm(input CreateAndMountDiskInput) (diskId string, attachId string, volumeName string, err error) {
	sc, err := createBlockStorageServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	//create new volume
	diskSize, _ := strconv.Atoi(input.DiskSize)
	createOpts := volumes.CreateOpts{
		AvailabilityZone: input.AvailabilityZone,
		Size:             diskSize,
		VolumeType:       input.DiskType,
	}
	if input.Name != "" {
		createOpts.Name = input.Name
	}

	volume, err := volumes.Create(sc, createOpts).Extract()
	if err != nil {
		return
	}
	diskId = volume.ID

	//wait volume status become ok
	if err = waitVolumeCreateOk(sc, volume.ID); err != nil {
		return
	}

	//attach to vm
	if attachId, volumeName, err = attachVolumeToVm(input, volume.ID, input.InstanceId); err != nil {
		return
	}
	logrus.Infof("attachVolumeToVm return ,attachId=%v,volumeName=%v,err=%v", attachId, volumeName, err)

	err = waitVolumeAttachOk(sc, volume.ID)

	return
}

type UnformatedDisks struct {
	Volumes []string `json:"unformatedDisks,omitempty"`
}

func getUnformatDisks(privateIp string, password string) ([]string, error) {
	if err := utils.CopyFileToRemoteHost(privateIp, password, "./scripts/getUnformatedDisk.py", "/tmp/getUnformatedDisk.py"); err != nil {
		return []string{}, err
	}
	output, err := utils.RunRemoteHostScript(privateIp, password, "python /tmp/getUnformatedDisk.py")
	if err != nil {
		return []string{}, err
	}

	unformatedDisks := UnformatedDisks{}
	if err := json.Unmarshal([]byte(output), &unformatedDisks); err != nil {
		return []string{}, err
	}

	return unformatedDisks.Volumes, nil
}

func getNewCreateDiskVolumeName(ip, password string, lastUnformatedDisks []string) (string, error) {
	lastUnformatedDiskNum := len(lastUnformatedDisks)

	for i := 0; i < 20; i++ {
		newDisks, err := getUnformatDisks(ip, password)
		if err != nil {
			return "", err
		}
		if len(newDisks) == lastUnformatedDiskNum {
			time.Sleep(5 * time.Second)
			continue
		}
		for _, volumeName := range newDisks {
			bFind := false
			for _, oldDisk := range lastUnformatedDisks {
				if volumeName == oldDisk {
					bFind = true
					break
				}
			}
			if bFind == false {
				return volumeName, nil
			}
		}
	}

	return "", errors.New("getNewCreateDiskVolumeName timeout")
}

func isBlockStorageExist(param CloudProviderParam, id string) (bool, error) {
	sc, err := createBlockStorageServiceClient(param)
	if err != nil {
		return false, err
	}

	_, err = volumes.Get(sc, id).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "could not be found") {
				return false, nil
			}
		}
		return false, err
	}
	return true, nil
}

func formatAndMountDisk(ip, password, volumeName, fileSystemType, mountDir string) error {
	if err := utils.CopyFileToRemoteHost(ip, password, "./scripts/formatAndMountDisk.py", "/tmp/formatAndMountDisk.py"); err != nil {
		return err
	}

	execArgs := " -d " + volumeName + " -f " + fileSystemType + " -m " + mountDir
	_, err := utils.RunRemoteHostScript(ip, password, "python /tmp/formatAndMountDisk.py"+execArgs)
	return err
}

func createAndMountDisk(input CreateAndMountDiskInput) (output CreateAndMountDiskOutput, err error) {
	output.Guid = input.Guid
	output.CallBackParameter.Parameter = input.CallBackParameter.Parameter
	defer func() {
		if err == nil {
			output.Result.Code = RESULT_CODE_SUCCESS
		} else {
			output.Result.Code = RESULT_CODE_ERROR
			output.Result.Message = err.Error()
		}
	}()

	if err = checkCreateAndMountParam(input); err != nil {
		return
	}

	privateIp, err := getVmIpAddress(input.CloudProviderParam, input.InstanceId)
	if err != nil {
		return
	}
	password, err := utils.AesDePassword(input.InstanceGuid, input.InstanceSeed, input.InstancePassword)
	if err != nil {
		logrus.Errorf("AesDePassword meet error(%v)", err)
		return
	}

	//check if disk already exsit
	if input.Id != "" {
		exist := false
		exist, err = isBlockStorageExist(input.CloudProviderParam, input.Id)
		if err == nil && exist {
			output.Id = input.Id
			return
		}
	}

	oldUnformatDisks := []string{}
	if input.SkipMount != "TRUE" {
		oldUnformatDisks, err = getUnformatDisks(privateIp, password)
		if err != nil {
			return
		}
	}

	output.Id, output.AttachId, _, err = buyDiskAndAttachToVm(input)
	if err != nil {
		return
	}

	if input.SkipMount == "TRUE" {
		return
	}

	output.VolumeName, err = getNewCreateDiskVolumeName(privateIp, password, oldUnformatDisks)
	if err != nil {
		return
	}

	//format and mount
	err = formatAndMountDisk(privateIp, password, output.VolumeName, input.FileSystemType, input.MountDir)
	return
}

func (action *CreateAndMountDiskAction) Do(input interface{}) (interface{}, error) {
	inputs, _ := input.(CreateAndMountDiskInputs)
	outputs := CreateAndMountDiskOutputs{}
	var finalErr error

	for _, input := range inputs.Inputs {
		output, err := createAndMountDisk(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return outputs, finalErr
}

//-----------umount action ------------//
type UmountAndTerminateDiskAction struct {
}

type UmountAndTerminateDiskInputs struct {
	Inputs []UmountAndTerminateDiskInput `json:"inputs,omitempty"`
}

type UmountAndTerminateDiskInput struct {
	CallBackParameter
	CloudProviderParam
	Guid             string `json:"guid,omitempty"`
	AvailabilityZone string `json:"az,omitempty"`
	Id               string `json:"id,omitempty"`
	AttachId         string `json:"attach_id,omitempty"`

	//use to attach and format
	InstanceId       string `json:"instance_id,omitempty"`
	InstanceGuid     string `json:"instance_guid,omitempty"`
	InstanceSeed     string `json:"seed,omitempty"`
	InstancePassword string `json:"password,omitempty"`

	MountDir   string `json:"mount_dir,omitempty"`
	VolumeName string `json:"volume_name,omitempty"`

	SkipUnmount string
}

type UmountAndTerminateDiskOutputs struct {
	Outputs []UmountAndTerminateDiskOutput `json:"outputs,omitempty"`
}

type UmountAndTerminateDiskOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
}

func (action *UmountAndTerminateDiskAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs UmountAndTerminateDiskInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func checkUmountAndTerminateDiskParam(input UmountAndTerminateDiskInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}

	if input.InstanceId == "" {
		return fmt.Errorf("empty instanceId")
	}
	if input.InstanceGuid == "" {
		return fmt.Errorf("empty instanceGuid")
	}
	if input.InstanceSeed == "" {
		return fmt.Errorf("empty InstanceSeed")
	}
	if input.InstancePassword == "" {
		return fmt.Errorf("empty instancePassword")
	}

	if input.Id == "" {
		return fmt.Errorf("id is empty")
	}

	if input.MountDir == "" {
		return fmt.Errorf(" mountDir is empty")
	}

	if input.VolumeName == "" {
		return fmt.Errorf("volumeName is empty")
	}

	return nil
}

func umountDisk(ip, password, volumeName, mountDir string) error {
	if err := utils.CopyFileToRemoteHost(ip, password, "./scripts/umountDisk.py", "/tmp/umountDisk.py"); err != nil {
		return err
	}

	execArgs := " -d " + volumeName + " -m " + mountDir
	_, err := utils.RunRemoteHostScript(ip, password, "python /tmp/umountDisk.py"+execArgs)
	return err
}

func detachVolumeFromVm(input UmountAndTerminateDiskInput) error {
	sc, err := createComputeV2Client(input.CloudProviderParam)
	if err != nil {
		return err
	}

	err = volumeattach.Delete(sc, input.InstanceId, input.AttachId).ExtractErr()
	if err != nil {
		logrus.Errorf("volumeattach delete meet err=%v", err)
		return err
	}

	blockStorageSc, err := createBlockStorageServiceClient(input.CloudProviderParam)
	if err != nil {
		return err
	}
	err = waitVolumeInAvailableState(blockStorageSc, input.Id)

	return err
}

func deleteVolume(input UmountAndTerminateDiskInput) error {
	sc, err := createBlockStorageServiceClient(input.CloudProviderParam)
	if err != nil {
		return err
	}

	if err = volumes.Delete(sc, input.Id).ExtractErr(); err != nil {
		logrus.Errorf("delete volume(%v) meet err=%v", input.Id, err)
	}

	return err
}

func umountAndTerminateDisk(input UmountAndTerminateDiskInput) (output UmountAndTerminateDiskOutput, err error) {
	output.Guid = input.Guid
	output.CallBackParameter.Parameter = input.CallBackParameter.Parameter
	defer func() {
		if err == nil {
			output.Result.Code = RESULT_CODE_SUCCESS
		} else {
			output.Result.Code = RESULT_CODE_ERROR
			output.Result.Message = err.Error()
		}
	}()

	if err = checkUmountAndTerminateDiskParam(input); err != nil {
		return
	}
	if input.AttachId == "" {
		input.AttachId = input.Id
	}

	exist, err := isBlockStorageExist(input.CloudProviderParam, input.Id)
	if err != nil || !exist {
		return
	}
	privateIp, err := getVmIpAddress(input.CloudProviderParam, input.InstanceId)
	if err != nil {
		return
	}

	password, err := utils.AesDePassword(input.InstanceGuid, input.InstanceSeed, input.InstancePassword)
	if err != nil {
		logrus.Errorf("AesDePassword meet error(%v)", err)
		return
	}

	//umount
	if input.SkipUnmount != "TRUE" {
		if err = umountDisk(privateIp, password, input.VolumeName, input.MountDir); err != nil {
			return
		}
	}

	//detach
	if err = detachVolumeFromVm(input); err != nil {
		return
	}

	//delete disk
	err = deleteVolume(input)

	return
}

func (action *UmountAndTerminateDiskAction) Do(input interface{}) (interface{}, error) {
	inputs, _ := input.(UmountAndTerminateDiskInputs)
	outputs := UmountAndTerminateDiskOutputs{}
	var finalErr error

	for _, input := range inputs.Inputs {
		output, err := umountAndTerminateDisk(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return outputs, finalErr
}
