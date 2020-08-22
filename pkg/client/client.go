// Package client provides intergation with api calls.
package client

import (
	"net/http"
	"time"

	"github.com/onelogin/onelogin-go-sdk/pkg/services"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/app_rules"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/apps"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/auth_servers"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/legal_values"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/olhttp"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/session_login_tokens"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/user_mappings"
	"github.com/onelogin/onelogin-go-sdk/pkg/services/users"
)

// APIClient is used to communicate with the available api services.
type APIClient struct {
	clientID     string
	clientSecret string
	region       string
	baseURL      string
	client       *http.Client
	Services     *Services
}

// Services contains all the available api services.
type Services struct {
	HTTPService          *olhttp.OLHTTPService
	AppsV2               *apps.V2Service
	UsersV2              *users.V2Service
	AppRulesV2           *apprules.V2Service
	UserMappingsV2       *usermappings.V2Service
	SessionLoginTokensV1 *sessionlogintokens.V1Service
	AuthServersV2        *authservers.V2Service
}

// NewClient uses the config to generate the api client with services attached, and returns
// the new api client.
func NewClient(cfg *APIClientConfig) (*APIClient, error) {
	cfg, err := cfg.Initialize()
	if err != nil {
		return &APIClient{}, err
	}

	httpClient := &http.Client{
		Timeout: time.Second * time.Duration(cfg.Timeout),
	}

	repo := olhttp.New(services.HTTPServiceConfig{
		ClientID:     cfg.ClientID,
		ClientSecret: cfg.ClientSecret,
		BaseURL:      cfg.Url,
		Client:       httpClient,
	})

	legalMappingValuesService := legalvalues.New(repo, cfg.Url)
	userMappingsV2Service := usermappings.New(repo, legalMappingValuesService, cfg.Url)
	authServersV2Service := authservers.New(repo, cfg.Url)
	appsV2Service := apps.New(repo, cfg.Url)
	sessionLoginTokensV1Service := sessionlogintokens.New(repo, cfg.Url)
	appRulesV2Service := apprules.New(repo, cfg.Url)

	return &APIClient{
		clientID:     cfg.ClientID,
		clientSecret: cfg.ClientSecret,
		region:       cfg.Region,
		baseURL:      cfg.Url,
		client:       httpClient,
		Services: &Services{
			HTTPService:          repo,
			UserMappingsV2:       userMappingsV2Service,
			AuthServersV2:        authServersV2Service,
			AppsV2:               appsV2Service,
			SessionLoginTokensV1: sessionLoginTokensV1Service,
			AppRulesV2:           appRulesV2Service,
		},
	}, nil
}
