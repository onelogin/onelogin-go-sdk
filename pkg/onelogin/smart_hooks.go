package onelogin

import (
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	utl "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/utilities"
)

const (
	SmartHooksPath string = "api/2/hooks"
)

func (sdk *OneloginSDK) CreateHook(hook models.SmartHook) (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, hook)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteHook(hookID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath, hookID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetHook(hookID int, query models.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath, hookID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, query)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) ListHooks(query models.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, query)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateSmartHook(hookID int, hook models.SmartHook) (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath, hookID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, hook)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) ListEnvironmentVariables() (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath, "envs")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) CreateEnvironmentVariable(name, value string) (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath, "envs")
	if err != nil {
		return nil, err
	}
	envVar := map[string]string{
		"name":  name,
		"value": value,
	}
	resp, err := sdk.Client.Post(&p, envVar)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetEnvironmentVariable(envVarID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath, "envs", envVarID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) UpdateEnvironmentVariable(envVarID int, name, value string) (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath, "envs", envVarID)
	if err != nil {
		return nil, err
	}
	envVar := map[string]string{
		"name":  name,
		"value": value,
	}
	resp, err := sdk.Client.Put(&p, envVar)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) DeleteEnvironmentVariable(envVarID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath, "envs", envVarID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetHookLogs(hookID int, query models.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(SmartHooksPath, hookID, "logs")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, query)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
