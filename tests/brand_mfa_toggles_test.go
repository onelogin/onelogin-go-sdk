package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

// TestBrandMFAToggles covers the three-state pointer/omitempty contract on
// each per-brand MFA link toggle. A bare bool could not express the three
// states: with omitempty a false would be dropped, leaving no way to turn a
// link off once it was on. A misspelled tag or a dropped explicit false would
// otherwise slip through the existing whole-struct encoding test.
func TestBrandMFAToggles(t *testing.T) {
	marshal := func(t *testing.T, b models.Brand) string {
		t.Helper()
		out, err := json.Marshal(b)
		if err != nil {
			t.Fatalf("marshal: %v", err)
		}
		return string(out)
	}

	cases := []struct {
		name   string
		key    string
		assign func(*models.Brand, *bool)
		read   func(models.Brand) *bool
	}{
		{
			name:   "show_help_on_mfa",
			key:    "show_help_on_mfa",
			assign: func(b *models.Brand, v *bool) { b.ShowHelpOnMFA = v },
			read:   func(b models.Brand) *bool { return b.ShowHelpOnMFA },
		},
		{
			name:   "show_support_on_mfa",
			key:    "show_support_on_mfa",
			assign: func(b *models.Brand, v *bool) { b.ShowSupportOnMFA = v },
			read:   func(b models.Brand) *bool { return b.ShowSupportOnMFA },
		},
		{
			name:   "show_additional_links_on_mfa",
			key:    "show_additional_links_on_mfa",
			assign: func(b *models.Brand, v *bool) { b.ShowAdditionalLinksOnMFA = v },
			read:   func(b models.Brand) *bool { return b.ShowAdditionalLinksOnMFA },
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			t.Run("omitted when unset", func(t *testing.T) {
				if got := marshal(t, models.Brand{}); strings.Contains(got, tc.key) {
					t.Fatalf("expected %q to be absent, got %s", tc.key, got)
				}
			})

			t.Run("sends true", func(t *testing.T) {
				v := true
				var b models.Brand
				tc.assign(&b, &v)
				want := `"` + tc.key + `":true`
				if got := marshal(t, b); !strings.Contains(got, want) {
					t.Fatalf("expected %s in %s", want, got)
				}
			})

			t.Run("sends false rather than dropping it", func(t *testing.T) {
				v := false
				var b models.Brand
				tc.assign(&b, &v)
				want := `"` + tc.key + `":false`
				if got := marshal(t, b); !strings.Contains(got, want) {
					t.Fatalf("expected %s in %s -- an explicit false must reach the API to turn the link off", want, got)
				}
			})

			t.Run("reads back what the API returns", func(t *testing.T) {
				var b models.Brand
				payload := `{"` + tc.key + `":true}`
				if err := json.Unmarshal([]byte(payload), &b); err != nil {
					t.Fatalf("unmarshal: %v", err)
				}
				got := tc.read(b)
				if got == nil || !*got {
					t.Fatalf("expected true from %s, got %v", payload, got)
				}
			})
		})
	}
}
