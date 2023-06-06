package onelogin

import (
	"fmt"

	mod "github.com/onelogin/onelogin-go-sdk/internal/models"
)

const (
	AppPath      string = "/api/2/apps"
	AppPathI     string = "/api/2/apps/%d"
	AppRulePath  string = "/api/2/apps/%d/rules"
	AppRulePathI string = "/api/2/apps/%d/rules/%d"
)

func (sdk *OneloginSDK) CreateApp(app mod.App) ([]byte, error) {
	return sdk.ApiClient.Post(AppPath, nil, app)
}

func (sdk *OneloginSDK) GetApps(queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Get(AppPath, &queryParams)
}

func (sdk *OneloginSDK) GetAppByID(id int, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Get(fmt.Sprintf(AppPathI, id), &queryParams)
}

func (sdk *OneloginSDK) UpdateApp(id int, app mod.App, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Put(fmt.Sprintf(AppPathI, id), &queryParams, app)
}

func (sdk *OneloginSDK) DeleteApp(id int, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Delete(fmt.Sprintf(AppPathI, id), &queryParams)
}
func (sdk *OneloginSDK) CreateAppRule(id int, appRule mod.AppRule) ([]byte, error) {
	return sdk.ApiClient.Post(fmt.Sprintf(AppRulePath, id), nil, appRule)
}

func (sdk *OneloginSDK) GetAppRules(id int, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Get(fmt.Sprintf(AppRulePath, id), &queryParams)
}

func (sdk *OneloginSDK) GetAppRuleByID(id, ruleID int, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Get(fmt.Sprintf(AppRulePathI, id, ruleID), &queryParams)
}

func (sdk *OneloginSDK) UpdateAppRule(id, ruleID int, appRule mod.AppRule, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Put(fmt.Sprintf(AppRulePathI, id, ruleID), &queryParams, appRule)
}

func (sdk *OneloginSDK) DeleteAppRule(id, ruleID int, queryParams map[string]string) ([]byte, error) {
	return sdk.ApiClient.Delete(fmt.Sprintf(AppRulePathI, id, ruleID), &queryParams)
}
