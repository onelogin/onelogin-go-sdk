package models

type Condition struct {
	Source   string `json:"source"`
	Operator string `json:"operator"`
	Value    string `json:"value"`
}

type Action struct {
	Action     string   `json:"action"`
	Value      []string `json:"value,omitempty"`
	Expression string   `json:"expression,omitempty"`
	Scriplet   string   `json:"scriplet,omitempty"`
	Macro      string   `json:"macro,omitempty"`
}

type AppRule struct {
	AppID      int         `json:"app_id"`
	Name       string      `json:"name"`
	Enabled    bool        `json:"enabled"`
	Match      string      `json:"match"`
	Position   int         `json:"position,omitempty"`
	Conditions []Condition `json:"conditions"`
	Actions    []Action    `json:"actions"`
}
