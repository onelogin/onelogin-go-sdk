package onelogin

import utl "github.com/onelogin/onelogin-go-sdk/internal/utilities"

const (
	GroupsPath = "api/1/groups"
)

func (sdk *OneloginSDK) GetGroupByID(groupID int) (interface{}, error) {
	p, err := utl.BuildAPIPath(GroupsPath, groupID)
	if err != nil {
		return nil, err
	}
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}

func (sdk *OneloginSDK) GetGroups() (interface{}, error) {
	p := GroupsPath
	resp, err := sdk.Client.Get(&p, nil)
	if err != nil {
		return nil, err
	}
	return utl.CheckHTTPResponse(resp)
}
