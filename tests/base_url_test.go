package tests

import (
	"os"
	"strings"
	"testing"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/authentication"
)

// TestBaseURL covers how the root of every request is chosen.
//
// Deriving it from the subdomain alone assumes every tenant lives at
// <subdomain>.onelogin.com. That host belongs to a real tenant, so a caller
// pointed at a development or custom-domain deployment did not fail: it sent
// its credentials to production instead.
func TestBaseURL(t *testing.T) {
	// The variable is process-wide, so put it back however the test exits.
	original, had := os.LookupEnv("ONELOGIN_API_URL")
	t.Cleanup(func() {
		if had {
			os.Setenv("ONELOGIN_API_URL", original)
		} else {
			os.Unsetenv("ONELOGIN_API_URL")
		}
	})

	t.Run("falls back to the subdomain when unset", func(t *testing.T) {
		os.Unsetenv("ONELOGIN_API_URL")

		if got, want := authentication.BaseURL("chicken"), "https://chicken.onelogin.com"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("prefers ONELOGIN_API_URL over the subdomain", func(t *testing.T) {
		os.Setenv("ONELOGIN_API_URL", "https://chicken.onelogin-dev.com")

		got := authentication.BaseURL("chicken")
		if want := "https://chicken.onelogin-dev.com"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
		if got == "https://chicken.onelogin.com" {
			t.Fatal("resolved to production while configured for development")
		}
	})

	t.Run("does not double up on the path separator", func(t *testing.T) {
		os.Setenv("ONELOGIN_API_URL", "https://chicken.onelogin-dev.com/")

		// Callers append a path beginning with "/", so a trailing slash here
		// would produce "...com//auth/oauth2/v2/token".
		if got, want := authentication.BaseURL("chicken"), "https://chicken.onelogin-dev.com"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("adds a scheme to a bare host", func(t *testing.T) {
		os.Setenv("ONELOGIN_API_URL", "chicken.onelogin-dev.com")

		// http.NewRequest rejects a URL with no scheme rather than assuming
		// one, so a bare host would otherwise fail at the request rather than
		// anywhere informative.
		if got, want := authentication.BaseURL("chicken"), "https://chicken.onelogin-dev.com"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("leaves an explicit http scheme alone", func(t *testing.T) {
		os.Setenv("ONELOGIN_API_URL", "http://localhost:8080")

		if got, want := authentication.BaseURL("chicken"), "http://localhost:8080"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("ignores a blank value", func(t *testing.T) {
		os.Setenv("ONELOGIN_API_URL", "   ")

		if got, want := authentication.BaseURL("chicken"), "https://chicken.onelogin.com"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	// Callers append absolute paths to whatever this returns, so anything past
	// the authority is added to rather than replaced.
	t.Run("drops a path", func(t *testing.T) {
		os.Setenv("ONELOGIN_API_URL", "https://chicken.onelogin-dev.com/auth")

		got := authentication.BaseURL("chicken")
		if want := "https://chicken.onelogin-dev.com"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
		// The failure this prevents.
		if bad := got + "/auth/oauth2/v2/token"; strings.Contains(bad, "/auth/auth/") {
			t.Fatalf("token URL would double the path segment: %s", bad)
		}
	})

	t.Run("drops a query and fragment", func(t *testing.T) {
		os.Setenv("ONELOGIN_API_URL", "https://chicken.onelogin-dev.com/x?a=1#frag")

		if got, want := authentication.BaseURL("chicken"), "https://chicken.onelogin-dev.com"; got != want {
			t.Fatalf("expected %q, got %q", want, got)
		}
	})

	t.Run("keeps an explicit port", func(t *testing.T) {
		os.Setenv("ONELOGIN_API_URL", "http://localhost:8080/api")

		if got, want := authentication.BaseURL("chicken"), "http://localhost:8080"; got != want {
			t.Fatalf("expected the port to survive, got %q", got)
		}
	})

	t.Run("the resulting token URL is well formed", func(t *testing.T) {
		for _, configured := range []string{
			"https://chicken.onelogin-dev.com",
			"https://chicken.onelogin-dev.com/",
			"chicken.onelogin-dev.com",
			"https://chicken.onelogin-dev.com/auth",
		} {
			os.Setenv("ONELOGIN_API_URL", configured)

			got := authentication.BaseURL("chicken") + authentication.TkPath
			if want := "https://chicken.onelogin-dev.com/auth/oauth2/v2/token"; got != want {
				t.Fatalf("configured %q: expected %q, got %q", configured, want, got)
			}
		}
	})
}
