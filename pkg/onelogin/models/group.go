package models

type Group struct {
	ID        int     `json:"id"`
	Name      string  `json:"name"`
	Reference *string `json:"reference,omitempty"`
	// PolicyID is the user policy applied to the group's members. It is a
	// pointer because CreateGroup and UpdateGroup marshal this whole struct as
	// the request body: a bare int would send "policy_id": 0 on every write
	// that did not set one, and partial updates are the normal case here --
	// renaming a group would silently blank its policy.
	//
	// Only user policies are assignable; the API rejects an app policy with
	// 422 "Policy must reference a user policy".
	PolicyID *int `json:"policy_id,omitempty"`
}
