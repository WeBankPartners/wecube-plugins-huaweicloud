package plugins

import (
	"errors"
	"fmt"
	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/listeners"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/loadbalancers"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/monitors"
	"github.com/gophercloud/gophercloud/openstack/networking/v2/extensions/lbaas_v2/pools"

	"github.com/gophercloud/gophercloud/openstack/vpc/v1/subnets"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

var lbTargetActions = make(map[string]Action)

func init() {
	lbTargetActions["add-backtarget"] = new(AddLbHostAction)
	lbTargetActions["del-backtarget"] = new(DelLbHostAction)
}

type LbTargetPlugin struct {
}

func (plugin *LbTargetPlugin) GetActionByName(actionName string) (Action, error) {
	action, found := lbTargetActions[actionName]
	if !found {
		return nil, fmt.Errorf("lbTarget plugin,action = %s not found", actionName)
	}

	return action, nil
}

//-----------add lb host ----------//
type AddLbHostAction struct {
}

type LbHostInputs struct {
	Inputs []LbHostInput `json:"inputs,omitempty"`
}

type LbHostInput struct {
	CallBackParameter
	CloudProviderParam
	Guid      string `json:"guid,omitempty"`
	LbId      string `json:"lb_id,omitempty"`
	Port      string `json:"lb_port"`
	Protocol  string `json:"protocol"`
	HostIds   string `json:"host_ids"`
	HostPorts string `json:"host_ports"`
}

type LbHostOutputs struct {
	Outputs []LbHostOutput `json:"outputs,omitempty"`
}

type LbHostOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
}

func (action *AddLbHostAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs LbHostInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func isValidPort(port string) error {
	if port == "" {
		return errors.New("port is empty")
	}

	portInt, err := strconv.Atoi(port)
	if err != nil || portInt >= 65535 || portInt <= 0 {
		return fmt.Errorf("port(%s) is invalid", port)
	}
	return nil
}

func isValidProtocol(protocol string) error {
	if protocol == "" {
		return errors.New("protocol is empty")
	}

	if !strings.EqualFold(protocol, "TCP") && !strings.EqualFold(protocol, "UDP") {
		return fmt.Errorf("protocol(%s) is invalid", protocol)
	}
	return nil
}

func checkLbHostInputParam(input LbHostInput) error {
	if input.HostIds == "" {
		return errors.New("empty host id")
	}
	if err := isValidPort(input.Port); err != nil {
		return fmt.Errorf("port(%v) is invalid", input.Port)
	}
	if err := isValidProtocol(input.Protocol); err != nil {
		return fmt.Errorf("protocol(%v) is invalid", input.Protocol)
	}
	if input.LbId == "" {
		return errors.New("empty LbId")
	}
	if input.HostPorts == "" {
		return errors.New("empty hostPorts")
	}
	if _, err := getLbInfoById(input.CloudProviderParam, input.LbId); err != nil {
		return err
	}

	return nil
}

func getLbListener(sc *gophercloud.ServiceClient, lbId string, protocol string, port string) (*listeners.Listener, error) {
	lbInfo, err := loadbalancers.Get(sc, lbId).Extract()
	if err != nil {
		logrus.Errorf("getLbListener get lbinfo meet err=%v", err)
		return nil, err
	}

	for _, listerner := range lbInfo.Listeners {
		listenerInfo, err := listeners.Get(sc, listerner.ID).Extract()
		if err != nil {
			logrus.Errorf("getLbListener,getlistenerDetail(%v) failed,err=%v", listerner.ID, err)
			return nil, err
		}
		listernerPort, _ := strconv.Atoi(port)
		if strings.EqualFold(listenerInfo.Protocol, protocol) && listenerInfo.ProtocolPort == listernerPort {
			return listenerInfo, nil
		}
	}

	logrus.Infof("getLbListener not found match listener,lbId(%v) protocol(%v) port(%v)", lbId, protocol, port)
	return nil, nil
}

func createLbPool(sc *gophercloud.ServiceClient, lbId string, protocol string) (*pools.Pool, error) {
	opts := pools.CreateOpts{
		Name:           "wecube_created",
		LBMethod:       "ROUND_ROBIN",
		Protocol:       pools.Protocol(strings.ToUpper(protocol)),
		LoadbalancerID: lbId,
	}

	pool, err := pools.Create(sc, opts).Extract()
	if err != nil {
		logrus.Errorf("createLbPool meet err=%v", err)
	}
	return pool, err
}

func createLbListener(sc *gophercloud.ServiceClient, lbId, protocol, port, poolId string) (*listeners.Listener, error) {
	portInt, err := strconv.Atoi(port)
	opts := listeners.CreateOpts{
		Name:           "wecubeCreated",
		Protocol:       listeners.Protocol(strings.ToUpper(protocol)),
		ProtocolPort:   portInt,
		DefaultPoolID:  poolId,
		LoadbalancerID: lbId,
	}

	listener, err := listeners.Create(sc, opts).Extract()
	if err != nil {
		logrus.Errorf("createLbListener meet err=%v", err)
	}
	return listener, err
}

func ensureLbListenerAndPoolCreated(input LbHostInput) (string, string, error) {
	sc, err := createLbServiceClient(input.CloudProviderParam)
	if err != nil {
		return "", "", err
	}

	listener, err := getLbListener(sc, input.LbId, input.Protocol, input.Port)
	if err != nil {
		return "", "", err
	}

	if listener == nil {
		pool, err := createLbPool(sc, input.LbId, input.Protocol)
		if err != nil {
			return "", "", err
		}

		listener, err := createLbListener(sc, input.LbId, input.Protocol, input.Port, pool.ID)
		if err != nil {
			return "", "", err
		}
		return listener.ID, pool.ID, nil
	}
	if listener.DefaultPoolID == nil {
		return "", "", fmt.Errorf("listener's default pool is nil")
	}

	return listener.ID, *listener.DefaultPoolID, nil
}

func ensureHostAddToLbPool(params CloudProviderParam, hostIds []string, hostPorts []string, poolId string) error {
	subnets := []subnets.Subnet{}
	sc, err := createLbServiceClient(params)
	if err != nil {
		return err
	}

	for i, hostId := range hostIds {
		vm, err := getVmInfoById(params, hostId)
		if err != nil {
			return err
		}
		if len(subnets) == 0 {
			subnets, err = getVpcAllSubnets(params, vm.Metadata.VpcID)
			if err != nil {
				return err
			}
		}

		address, _ := getIpFromVmInfo(vm)
		subnetId, err := getSubnetIdByIpAddress(subnets, address)
		if err != nil {
			return err
		}

		port, _ := strconv.Atoi(hostPorts[i])
		opts := pools.CreateMemberOpts{
			Address:      address,
			ProtocolPort: port,
			SubnetID:     subnetId,
			Name:         "wecube_created",
		}

		if _, err = pools.CreateMember(sc, poolId, opts).Extract(); err != nil {
			return err
		}
	}

	return nil
}

func addHostToLb(input LbHostInput) (output LbHostOutput, err error) {
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

	if err = checkLbHostInputParam(input); err != nil {
		return
	}

	_, poolId, err := ensureLbListenerAndPoolCreated(input)
	if err != nil {
		return
	}

	hostIds, err := GetArrayFromString(input.HostIds, ARRAY_SIZE_REAL, 0)
	if err != nil {
		return
	}

	hostPorts, err := GetArrayFromString(input.HostPorts, ARRAY_SIZE_AS_EXPECTED, len(hostIds))
	if err != nil {
		return
	}
	err = ensureHostAddToLbPool(input.CloudProviderParam, hostIds, hostPorts, poolId)
	return
}

func (action *AddLbHostAction) Do(inputs interface{}) (interface{}, error) {
	hosts, _ := inputs.(LbHostInputs)
	outputs := LbHostOutputs{}
	var finalErr error

	for _, input := range hosts.Inputs {
		output, err := addHostToLb(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return &outputs, finalErr
}

//--------del host from loadbalancer---------//
type DelLbHostAction struct {
}

func (action *DelLbHostAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs LbHostInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func getMemberIdByIp(allMembers []pools.Member, ip string) (string, error) {
	for _, member := range allMembers {
		if member.Address == ip {
			return member.ID, nil
		}
	}
	logrus.Errorf("can't found address(%v) in pool member", ip)
	return "", fmt.Errorf("can't found address(%v) in pool member", ip)
}

func ensureDeleteHostFromPool(params CloudProviderParam, hostIds []string, hostPorts []string, poolId string) error {
	sc, err := createLbServiceClient(params)
	if err != nil {
		return err
	}
	allPages, err := pools.ListMembers(sc, poolId, pools.ListMembersOpts{}).AllPages()
	if err != nil {
		logrus.Errorf("lb pool listMembers meet err=%v", err)
		return err
	}
	allMembers, err := pools.ExtractMembers(allPages)
	if err != nil {
		logrus.Errorf("lb pool ExtractMembers meet err=%v", err)
		return err
	}

	for _, hostId := range hostIds {
		vm, err := getVmInfoById(params, hostId)
		if err != nil {
			return err
		}
		address, _ := getIpFromVmInfo(vm)
		memberId, err := getMemberIdByIp(allMembers, address)
		if err != nil {
			continue
		}

		if err = pools.DeleteMember(sc, poolId, memberId).ExtractErr(); err != nil {
			logrus.Errorf("lb pools deleteMember meet err=%v", err)
			return err
		}
	}

	return nil
}

func delHostFromLb(input LbHostInput) (output LbHostOutput, err error) {
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

	if err = checkLbHostInputParam(input); err != nil {
		return
	}

	sc, err := createLbServiceClient(input.CloudProviderParam)
	if err != nil {
		return
	}

	listener, err := getLbListener(sc, input.LbId, input.Protocol, input.Port)
	if err != nil {
		return
	}
	if listener == nil {
		err = fmt.Errorf("can't found lbListener,lbId(%v),protocol(%v),port(%v)", input.LbId, input.Protocol, input.Port)
		return
	}

	hostIds, err := GetArrayFromString(input.HostIds, ARRAY_SIZE_REAL, 0)
	if err != nil {
		return
	}

	hostPorts, err := GetArrayFromString(input.HostPorts, ARRAY_SIZE_AS_EXPECTED, len(hostIds))
	if err != nil {
		return
	}
	if nil == listener.DefaultPoolID {
		err = fmt.Errorf("listener hav no default pool")
		return
	}

	err = ensureDeleteHostFromPool(input.CloudProviderParam, hostIds, hostPorts, *listener.DefaultPoolID)
	return
}

func (action *DelLbHostAction) Do(inputs interface{}) (interface{}, error) {
	hosts, _ := inputs.(LbHostInputs)
	outputs := LbHostOutputs{}
	var finalErr error

	for _, input := range hosts.Inputs {
		output, err := delHostFromLb(input)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	return &outputs, finalErr
}

func deleteLbPools(params CloudProviderParam, id string) error {
	sc, err := createLbServiceClient(params)
	if err != nil {
		return err
	}

	//get pool and delete healthMonitor
	pool, err := pools.Get(sc, id).Extract()
	if err != nil {
		logrus.Errorf("get pool meet err=%v", err)
		return err
	}

	if err = monitors.Delete(sc, pool.MonitorID).ExtractErr(); err != nil {
		logrus.Errorf("delete monitor meet err=%v", err)
		return err
	}

	err = pools.Delete(sc, id).ExtractErr()
	if err != nil {
		logrus.Errorf("delete pool meet err=%v", err)
	}
	return err
}

func deleteLbListener(params CloudProviderParam, id string) error {
	sc, err := createLbServiceClient(params)
	if err != nil {
		return err
	}

	err = listeners.Delete(sc, id).ExtractErr()
	if err != nil {
		logrus.Errorf("delete listener meet err=%v", err)
	}
	return err
}
