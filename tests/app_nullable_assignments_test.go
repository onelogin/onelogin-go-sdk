package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

func marshalApp(t *testing.T, app models.App) string {
	t.Helper()

	b, err := json.Marshal(app)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return string(b)
}

func intPtr(v int) *int { return &v }

// TestAppAssignmentsAreThreeState covers policy_id and brand_id, which name a
// record the app is assigned to and so have three meanings on the wire rather
// than two: leave the assignment alone, change it, or take it off.
//
// A *int tagged omitempty reaches only the first two. The third needs a JSON
// null, and no pointer can produce one. 0 does not do it either: the endpoint
// answers it with 422, "The associated Policy with ID 0 could not be found"
// for policy_id and "The associated AccountBrand with ID 0 could not be found"
// for brand_id. Both were confirmed against the API.
func TestAppAssignmentsAreThreeState(t *testing.T) {
	for _, field := range []struct {
		key    string
		assign func(*models.App, int)
		clear  func(*models.App)
	}{
		{
			key:    "policy_id",
			assign: func(a *models.App, id int) { a.PolicyID = intPtr(id) },
			clear:  func(a *models.App) { a.ClearPolicyID = true },
		},
		{
			key:    "brand_id",
			assign: func(a *models.App, id int) { a.BrandID = intPtr(id) },
			clear:  func(a *models.App) { a.ClearBrandID = true },
		},
	} {
		t.Run(field.key, func(t *testing.T) {
			t.Run("omitted when untouched", func(t *testing.T) {
				if got := marshalApp(t, models.App{}); strings.Contains(got, field.key) {
					t.Fatalf("expected %s to be absent, got %s", field.key, got)
				}
			})

			t.Run("sends the id it is assigned", func(t *testing.T) {
				var app models.App
				field.assign(&app, 955633)

				if got := marshalApp(t, app); !strings.Contains(got, `"`+field.key+`":955633`) {
					t.Fatalf("expected %s to be sent, got %s", field.key, got)
				}
			})

			// The point of the exercise: a null, which is the only thing the
			// endpoint accepts as "no policy" or "no brand".
			t.Run("sends null when cleared", func(t *testing.T) {
				var app models.App
				field.clear(&app)

				if got := marshalApp(t, app); !strings.Contains(got, `"`+field.key+`":null`) {
					t.Fatalf("expected %s to be sent as null, got %s", field.key, got)
				}
			})

			// Assign and unassign in one request is a contradiction. Resolving
			// it either way would silently do half of what was asked.
			t.Run("refuses to assign and clear at once", func(t *testing.T) {
				var app models.App
				field.assign(&app, 955633)
				field.clear(&app)

				_, err := json.Marshal(app)
				if err == nil {
					t.Fatalf("expected assigning and clearing %s together to be an error", field.key)
				}
				// Naming the field matters: an app can carry both
				// assignments, so "one of them conflicts" would leave the
				// caller to work out which.
				if !strings.Contains(err.Error(), field.key) {
					t.Fatalf("expected the error to name %s, got %v", field.key, err)
				}
			})
		})
	}
}

// TestAppAlwaysSendsConnectorID pins the one field still not governed by
// omitempty.
//
// Every other field on App is dropped when nil, which is what lets a request
// mention only what it means to change. connector_id carries no omitempty, so
// a nil one is sent as an explicit null instead. The endpoint answers that
// with a 200, so it is left as it is; name, which was answered with a 422, is
// covered by TestAppOmitsNameWhenUnset below.
//
// docs/models.md states this, and the assertion is here so that adding
// omitempty to connector_id fails rather than quietly making the docs wrong.
func TestAppAlwaysSendsConnectorID(t *testing.T) {
	if got := marshalApp(t, models.App{}); got != `{"connector_id":null}` {
		t.Fatalf("unexpected encoding of an empty App: %s", got)
	}
}

// TestAppOmitsNameWhenUnset covers the regression behind issue #123.
//
// terraform-provider-onelogin's onelogin_app_role_attachments builds a bare
// models.App{RoleIDs: &ids} to attach or detach a role. Without omitempty on
// Name that request carried "name": null, and the endpoint rejected the whole
// update with 422 "Validation failed: Name can't be blank" -- so a resource
// that never mentioned the name could not be created, updated or deleted.
func TestAppOmitsNameWhenUnset(t *testing.T) {
	t.Run("a role attachment sends only the roles", func(t *testing.T) {
		roleIDs := []int{955633}

		// Exact rather than a Contains, so that any future field arriving
		// without omitempty is caught here too: the whole point of the issue
		// is a key nobody asked to send.
		got := marshalApp(t, models.App{RoleIDs: &roleIDs})

		if want := `{"connector_id":null,"role_ids":[955633]}`; got != want {
			t.Fatalf("unexpected encoding of a role attachment:\n got %s\nwant %s", got, want)
		}
	})

	t.Run("sends the name when it is set", func(t *testing.T) {
		name := "my OIDC APP"

		if got := marshalApp(t, models.App{Name: &name}); !strings.Contains(got, `"name":"my OIDC APP"`) {
			t.Fatalf("expected the name to be sent, got %s", got)
		}
	})

	// omitempty drops the nil pointer, not a pointer to "". A caller that
	// explicitly asks for a blank name should still reach the endpoint and be
	// told no, rather than have the field quietly dropped and the request
	// succeed as a no-op.
	t.Run("still sends an explicitly empty name", func(t *testing.T) {
		blank := ""

		if got := marshalApp(t, models.App{Name: &blank}); !strings.Contains(got, `"name":""`) {
			t.Fatalf("expected an explicitly empty name to be sent, got %s", got)
		}
	})
}

// TestAppClearsBothAssignmentsTogether covers the two flags not treading on
// each other, since they share one pass over the encoded object.
func TestAppClearsBothAssignmentsTogether(t *testing.T) {
	got := marshalApp(t, models.App{ClearPolicyID: true, ClearBrandID: true})

	if !strings.Contains(got, `"policy_id":null`) {
		t.Errorf("expected policy_id null, got %s", got)
	}
	if !strings.Contains(got, `"brand_id":null`) {
		t.Errorf("expected brand_id null, got %s", got)
	}
}

// TestAppMarshalUnchangedWithoutClearFlags is the compatibility guard. The
// flags are additive, so an app that does not use them has to encode exactly as
// it did before MarshalJSON existed -- byte for byte, not merely equivalently.
func TestAppMarshalUnchangedWithoutClearFlags(t *testing.T) {
	name, connector := "my OIDC APP", int32(38568)
	notes := "example"
	app := models.App{
		Name:        &name,
		ConnectorID: &connector,
		Notes:       &notes,
		PolicyID:    intPtr(955633),
		BrandID:     intPtr(44528),
	}

	// What the struct tags alone produce, obtained the same way MarshalJSON
	// does before it considers the flags.
	type plain models.App
	want, err := json.Marshal(plain(app))
	if err != nil {
		t.Fatalf("marshal baseline: %v", err)
	}

	if got := marshalApp(t, app); got != string(want) {
		t.Fatalf("encoding changed for an app that sets no clear flag:\n got %s\nwant %s", got, want)
	}
}

// TestAppClearedFieldsSurviveRoundTrip guards the map rebuild: every other
// field has to come through it untouched, and not be re-encoded on the way.
func TestAppClearedFieldsSurviveRoundTrip(t *testing.T) {
	name, connector := "my OIDC APP", int32(38568)
	visible := true
	app := models.App{
		Name:          &name,
		ConnectorID:   &connector,
		Visible:       &visible,
		BrandID:       intPtr(44528),
		ClearPolicyID: true,
	}

	got := marshalApp(t, app)

	var decoded models.App
	if err := json.Unmarshal([]byte(got), &decoded); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	if decoded.Name == nil || *decoded.Name != name {
		t.Errorf("name did not survive: %s", got)
	}
	if decoded.ConnectorID == nil || *decoded.ConnectorID != connector {
		t.Errorf("connector_id did not survive: %s", got)
	}
	if decoded.Visible == nil || !*decoded.Visible {
		t.Errorf("visible did not survive: %s", got)
	}
	if decoded.BrandID == nil || *decoded.BrandID != 44528 {
		t.Errorf("brand_id was disturbed by clearing policy_id: %s", got)
	}
	// A null decodes to a nil pointer, indistinguishable from an absent key
	// once decoded, so the wire form is what the assertion has to be on.
	if !strings.Contains(got, `"policy_id":null`) {
		t.Errorf("expected policy_id null, got %s", got)
	}
}
