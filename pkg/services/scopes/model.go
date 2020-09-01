package scopes

type ScopesQuery struct {
	AuthServerID string
}

type Scope struct {
	ID           *int32
	AuthServerID *int32
	Value        *string
	Description  *string
}
