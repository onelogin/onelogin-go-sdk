package models

// RoleQuery represents available query parameters
type RoleQuery struct {
	Limit  string
	Page   string
	Cursor string
}

// Role represents the Role resource in OneLogin
type Role struct {
	ID     *int32  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Admins []int32 `json:"admins,omitempty"`
	Apps   []int32 `json:"apps,omitempty"`
	Users  []int32 `json:"users,omitempty"`
}

func (r *Role) GetKeyValidators() map[string]func(interface{}) bool {
	return map[string]func(interface{}) bool{
		"limit":  validateString,
		"page":   validateString,
		"cursor": validateString,
	}
}
