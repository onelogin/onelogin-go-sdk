package groups

// GroupQuery represents available query parameters
type GroupQuery struct {
	Limit  string
	Page   string
	Cursor string
}

// Group represents the Group resource in OneLogin
type Group struct {
	ID        *int32  `json:"id,omitempty"`
	Name      *string `json:"name,omitempty"`
	Reference *string `json:"reference,omitempty"`
}
