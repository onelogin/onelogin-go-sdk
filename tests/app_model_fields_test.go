package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

// TestConfigurationOpenIdIncludeAmrClaims covers the field's three states. A
// bare bool could not express them: with omitempty a false would be dropped,
// leaving no way to turn the claim off once it was on.
func TestConfigurationOpenIdIncludeAmrClaims(t *testing.T) {
	marshal := func(t *testing.T, c models.ConfigurationOpenId) string {
		t.Helper()
		b, err := json.Marshal(c)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		return string(b)
	}

	t.Run("omitted when unset", func(t *testing.T) {
		if got := marshal(t, models.ConfigurationOpenId{}); strings.Contains(got, "include_amr_claims") {
			t.Fatalf("expected the key to be absent, got %s", got)
		}
	})

	t.Run("sends true", func(t *testing.T) {
		v := true
		if got := marshal(t, models.ConfigurationOpenId{IncludeAmrClaims: &v}); !strings.Contains(got, `"include_amr_claims":true`) {
			t.Fatalf("expected true to be sent, got %s", got)
		}
	})

	t.Run("sends false", func(t *testing.T) {
		v := false
		got := marshal(t, models.ConfigurationOpenId{IncludeAmrClaims: &v})
		if !strings.Contains(got, `"include_amr_claims":false`) {
			t.Fatalf("expected false to be sent rather than dropped, got %s", got)
		}
	})

	t.Run("reads back what the API returns", func(t *testing.T) {
		var c models.ConfigurationOpenId
		if err := json.Unmarshal([]byte(`{"include_amr_claims":true}`), &c); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if c.IncludeAmrClaims == nil || !*c.IncludeAmrClaims {
			t.Fatalf("expected true, got %v", c.IncludeAmrClaims)
		}
	})
}

// TestParameterSafeEntitlementsEnabled covers the same three states. The
// neighbouring flags on Parameter are bare bools and cannot send a false; this
// one deliberately differs.
func TestParameterSafeEntitlementsEnabled(t *testing.T) {
	marshal := func(t *testing.T, p models.Parameter) string {
		t.Helper()
		b, err := json.Marshal(p)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		return string(b)
	}

	t.Run("omitted when unset", func(t *testing.T) {
		if got := marshal(t, models.Parameter{}); strings.Contains(got, "safe_entitlements_enabled") {
			t.Fatalf("expected the key to be absent, got %s", got)
		}
	})

	t.Run("sends false rather than dropping it", func(t *testing.T) {
		v := false
		got := marshal(t, models.Parameter{SafeEntitlementsEnabled: &v})
		if !strings.Contains(got, `"safe_entitlements_enabled":false`) {
			t.Fatalf("expected false to be sent, got %s", got)
		}
	})

	t.Run("round-trips a value from the API", func(t *testing.T) {
		var p models.Parameter
		if err := json.Unmarshal([]byte(`{"safe_entitlements_enabled":true}`), &p); err != nil {
			t.Fatalf("unmarshal: %v", err)
		}
		if p.SafeEntitlementsEnabled == nil || !*p.SafeEntitlementsEnabled {
			t.Fatalf("expected true, got %v", p.SafeEntitlementsEnabled)
		}

		// The round trip is the point of #235: the value has to survive back
		// out again, not merely be readable.
		if got := marshal(t, p); !strings.Contains(got, `"safe_entitlements_enabled":true`) {
			t.Fatalf("expected the value to survive the round trip, got %s", got)
		}
	})

	t.Run("a bare bool would have dropped the false", func(t *testing.T) {
		// Guards the reasoning rather than the code: this is what the field
		// would do if it matched its neighbours.
		bare := struct {
			SafeEntitlementsEnabled bool `json:"safe_entitlements_enabled,omitempty"`
		}{SafeEntitlementsEnabled: false}

		b, _ := json.Marshal(bare)
		if strings.Contains(string(b), "safe_entitlements_enabled") {
			t.Fatal("expected a bare bool to drop false — if this fails, the pointer is unnecessary")
		}
	})
}
