package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

// TestGroupPolicyID covers the field's states. CreateGroup and UpdateGroup
// marshal the whole Group as the request body, so what this produces is
// literally what OneLogin receives.
func TestGroupPolicyID(t *testing.T) {
	marshal := func(t *testing.T, g models.Group) string {
		t.Helper()
		b, err := json.Marshal(g)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		return string(b)
	}

	t.Run("omitted when unset", func(t *testing.T) {
		if got := marshal(t, models.Group{}); strings.Contains(got, "policy_id") {
			t.Fatalf("expected the key to be absent, got %s", got)
		}
	})

	// The case the pointer exists for. A bare int would put "policy_id": 0 in
	// this body and blank the group's policy as a side effect of the rename.
	t.Run("a rename does not touch the policy", func(t *testing.T) {
		if got := marshal(t, models.Group{Name: "Engineering"}); strings.Contains(got, "policy_id") {
			t.Fatalf("expected a rename to leave policy_id out of the body, got %s", got)
		}
	})

	t.Run("sends the id when set", func(t *testing.T) {
		id := 955633
		if got := marshal(t, models.Group{Name: "Engineering", PolicyID: &id}); !strings.Contains(got, `"policy_id":955633`) {
			t.Fatalf("expected the id to be sent, got %s", got)
		}
	})

	// The response shape core-api returns from #show, #create and #update: a
	// flat policy_id next to the nested policy object the admin UI renders
	// from. The nested object is not what populates the field.
	t.Run("reads the flat id the API returns", func(t *testing.T) {
		body := `{"id":590390,"name":"Engineering","policy_id":955633,"policy":{"id":955633,"name":"MFA Required"}}`

		var g models.Group
		if err := json.Unmarshal([]byte(body), &g); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if g.PolicyID == nil {
			t.Fatal("expected a policy id, got nil")
		}
		if *g.PolicyID != 955633 {
			t.Fatalf("expected 955633, got %d", *g.PolicyID)
		}
	})

	t.Run("a group with no policy reads back as nil", func(t *testing.T) {
		body := `{"id":578136,"name":"Disney","policy_id":null,"policy":{}}`

		var g models.Group
		if err := json.Unmarshal([]byte(body), &g); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if g.PolicyID != nil {
			t.Fatalf("expected nil, got %d", *g.PolicyID)
		}
	})
}
