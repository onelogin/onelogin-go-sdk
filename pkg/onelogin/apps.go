package onelogin

import (
	mod "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	utl "github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/utilities"
)

const (
	AppPath string = "api/2/apps"
)

func (sdk *OneloginSDK) CreateApp(app mod.App) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, app)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)

}

// was ListApps
func (sdk *OneloginSDK) GetApps(queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetAppByID(id int, queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)

}

func (sdk *OneloginSDK) UpdateApp(id int, app mod.App) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, app)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)

}

func (sdk *OneloginSDK) DeleteApp(id int) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)

}

func (sdk *OneloginSDK) CreateAppRule(id int, appRule mod.AppRule) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Post(&p, appRule)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)

}

func (sdk *OneloginSDK) GetAppRules(id int, queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)

}

func (sdk *OneloginSDK) GetAppRuleByID(id, ruleID int, queryParams mod.Queryable) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath, id)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, queryParams)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)

}

func (sdk *OneloginSDK) UpdateAppRule(id, ruleID int, appRule mod.AppRule, queryParams map[string]string) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath, id, "rules", ruleID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Put(&p, appRule)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)

}

func (sdk *OneloginSDK) DeleteAppRule(id, ruleID int, queryParams map[string]string) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath, id, "rules", ruleID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Delete(&p)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetAppUsers(appID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(AppPath, appID, "users")
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
