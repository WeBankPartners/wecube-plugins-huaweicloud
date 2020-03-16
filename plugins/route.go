package plugins

import (
	"fmt"
	"strings"

	"github.com/gophercloud/gophercloud"
	"github.com/gophercloud/gophercloud/openstack/vpc/v2.0/routes"
	"github.com/sirupsen/logrus"
)

var routeActions = make(map[string]Action)
var routeTypeMap = map[string]string{
	"PEERING": "peering",
}

func init() {
	routeActions["create"] = new(RouteCreateAction)
	routeActions["delete"] = new(RouteDeleteAction)
}

type RoutePlugin struct {
}

func (plugin *RoutePlugin) GetActionByName(actionName string) (Action, error) {
	action, found := routeActions[actionName]
	if !found {
		logrus.Errorf("Route plugin,action = %s not found", actionName)
		return nil, fmt.Errorf("Route plugin,action = %s not found", actionName)
	}

	return action, nil
}

type RouteCreateInputs struct {
	Inputs []RouteCreateInput `json:"inputs,omitempty"`
}

type RouteCreateInput struct {
	CallBackParameter
	CloudProviderParam
	Guid        string `json:"guid,omitempty"`
	Id          string `json:"id,omitempty"`
	Destination string `json:"destination,omitempty"`
	Nexthop     string `json:"nexthop,omitempty"`
	Type        string `json:"type,omitempty"`
	VpcId       string `json:"vpc_id,omitempty"`
}

type RouteCreateOutputs struct {
	Outputs []RouteCreateOutput `json:"outputs,omitempty"`
}

type RouteCreateOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type RouteCreateAction struct {
}

func (action *RouteCreateAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs RouteCreateInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (acton *RouteCreateAction) checkCreaterouteParams(input RouteCreateInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}

	if input.Destination == "" {
		return fmt.Errorf("Destination is empty")
	}
	if input.Nexthop == "" {
		return fmt.Errorf("Nexthop is empty")
	}

	if input.Type == "" {
		return fmt.Errorf("Type is empty")
	} else {
		if _, ok := routeTypeMap[strings.ToUpper(input.Type)]; !ok {
			return fmt.Errorf("Type is wrong")
		}
	}
	if input.VpcId == "" {
		return fmt.Errorf("VpcId is empty")
	}
	return nil
}

func (action *RouteCreateAction) createRoute(input *RouteCreateInput) (output RouteCreateOutput, err error) {
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

	if err = action.checkCreaterouteParams(*input); err != nil {
		logrus.Errorf("checkCreaterouteParams meet error=%v", err)
		return
	}

	// create vpc service client
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}

	// check whether is exist.
	if input.Id != "" {
		var routeInfo *routes.Route
		routeInfo, _, err = isRouteExist(sc, input.Id)
		if err != nil {
			logrus.Errorf("check route[routeId=%v] meet error=%v", input.Id, err)
			return
		}
		if routeInfo != nil {
			output.Id = routeInfo.ID
			return
		}
	}

	// create route
	request := routes.CreateOpts{
		Type:        routeTypeMap[strings.ToUpper(input.Type)],
		Nexthop:     input.Nexthop,
		Destination: input.Destination,
		VpcID:       input.VpcId,
	}

	response, err := routes.Create(sc, request).Extract()
	if err != nil {
		logrus.Errorf("create route meet error=%v", err)
		return
	}

	output.Id = response.ID
	return
}

func isRouteExist(sc *gophercloud.ServiceClient, routeId string) (*routes.Route, bool, error) {
	routeInfo, err := routes.Get(sc, routeId).Extract()
	if err != nil {
		if ue, ok := err.(*gophercloud.UnifiedError); ok {
			if strings.Contains(ue.Message(), "could not be found") {
				return nil, false, nil
			}
		}
		return nil, false, err
	}
	return routeInfo, true, nil
}

func (action *RouteCreateAction) Do(inputs interface{}) (interface{}, error) {
	routes, _ := inputs.(RouteCreateInputs)
	outputs := RouteCreateOutputs{}
	var finalErr error
	for _, route := range routes.Inputs {
		output, err := action.createRoute(&route)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all routes = %v are created", routes)
	return &outputs, finalErr
}

type RouteDeleteInputs struct {
	Inputs []RouteDeleteInput `json:"inputs,omitempty"`
}

type RouteDeleteInput struct {
	CallBackParameter
	CloudProviderParam
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type RouteDeleteOutputs struct {
	Outputs []RouteDeleteOutput `json:"outputs,omitempty"`
}

type RouteDeleteOutput struct {
	CallBackParameter
	Result
	Guid string `json:"guid,omitempty"`
	Id   string `json:"id,omitempty"`
}

type RouteDeleteAction struct {
}

func (action *RouteDeleteAction) ReadParam(param interface{}) (interface{}, error) {
	var inputs RouteDeleteInputs
	err := UnmarshalJson(param, &inputs)
	if err != nil {
		return nil, err
	}
	return inputs, nil
}

func (action *RouteDeleteAction) checkDeleteRouteParams(input RouteDeleteInput) error {
	if err := isCloudProviderParamValid(input.CloudProviderParam); err != nil {
		return err
	}
	if input.Id == "" {
		return fmt.Errorf("Route id is empty")
	}

	return nil
}

func (action *RouteDeleteAction) deleteRoute(input *RouteDeleteInput) (output RouteDeleteOutput, err error) {
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

	if err = action.checkDeleteRouteParams(*input); err != nil {
		logrus.Errorf("RouteDeleteAction checkDeleteRouteParams meet error=%v", err)
		return
	}

	// Create vpc service client
	sc, err := CreateVpcServiceClientV1(input.CloudProviderParam)
	if err != nil {
		return
	}

	// check whether route is exist.
	_, ok, err := isRouteExist(sc, input.Id)
	if err != nil {
		logrus.Errorf("check route[routeId=%v] meet error=%v", input.Id, err)
		return
	}
	if !ok {
		logrus.Infof("the route[routeId= %v] is not exist", input.Id)
		return
	}

	// delete route
	response := routes.Delete(sc, input.Id)
	if response.Err != nil {
		err = response.Err
		logrus.Errorf("delete route[routeId=%v] meet error=%v", input.Id, err)
	}
	return
}

func (action *RouteDeleteAction) Do(inputs interface{}) (interface{}, error) {
	routes, _ := inputs.(RouteDeleteInputs)
	outputs := RouteDeleteOutputs{}
	var finalErr error
	for _, route := range routes.Inputs {
		output, err := action.deleteRoute(&route)
		if err != nil {
			finalErr = err
		}
		outputs.Outputs = append(outputs.Outputs, output)
	}

	logrus.Infof("all routes = %v are deleted", routes)
	return &outputs, finalErr
}
