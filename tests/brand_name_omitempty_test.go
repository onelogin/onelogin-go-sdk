package tests

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
)

func marshalBrand(t *testing.T, brand models.Brand) string {
	t.Helper()

	b, err := json.Marshal(brand)
	if err != nil {
		t.Fatalf("marshal: %v", err)
	}
	return string(b)
}

// TestBrandOmitsNameWhenUnset covers the same defect on Brand that #123 covered
// on App.
//
// A caller changing one attribute of a brand -- a colour, say -- builds a
// models.Brand with only that field set. Without omitempty on Name that request
// carried "name": null, and the branding endpoint rejects it with
// 422 "Value must be a string.", refusing an update that never meant to touch
// the name. Omitting the key entirely is answered 200. Both were confirmed
// against the API.
func TestBrandOmitsNameWhenUnset(t *testing.T) {
	t.Run("a colour change does not mention the name", func(t *testing.T) {
		colour := "#123456"

		// Exact rather than a Contains, so that any future field arriving
		// without omitempty is caught here too: the whole point is a key
		// nobody asked to send.
		got := marshalBrand(t, models.Brand{CustomColor: &colour})

		if want := `{"custom_color":"#123456"}`; got != want {
			t.Fatalf("unexpected encoding of a partial brand update:\n got %s\nwant %s", got, want)
		}
	})

	t.Run("an empty brand sends nothing at all", func(t *testing.T) {
		if got := marshalBrand(t, models.Brand{}); got != `{}` {
			t.Fatalf("expected an empty object, got %s", got)
		}
	})

	t.Run("sends the name when it is set", func(t *testing.T) {
		name := "Engineering"

		if got := marshalBrand(t, models.Brand{Name: &name}); !strings.Contains(got, `"name":"Engineering"`) {
			t.Fatalf("expected the name to be sent, got %s", got)
		}
	})

	// omitempty drops the nil pointer, not a pointer to "". A create needs a
	// name -- without one the endpoint answers 422 {"name":"Value is
	// required."} -- so a caller that explicitly asks for a blank one should
	// still reach the endpoint and be told no, rather than have the field
	// quietly dropped.
	t.Run("still sends an explicitly empty name", func(t *testing.T) {
		blank := ""

		if got := marshalBrand(t, models.Brand{Name: &blank}); !strings.Contains(got, `"name":""`) {
			t.Fatalf("expected an explicitly empty name to be sent, got %s", got)
		}
	})
}
