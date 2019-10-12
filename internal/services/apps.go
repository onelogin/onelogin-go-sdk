package services

import (
	"net/http"
)

// AppsV2 is the apps service v2.
type AppsV2 struct {
	BaseURL      string
	client       *http.Client
	Auth         Authenticator
	ErrorContext string
}

// AppsV2Config is the config for apps service v2.
type AppsV2Config struct {
	BaseURL string
	Client  *http.Client
	Auth    Authenticator
}

// NewAppsV2 creates the new apps service v2.
func NewAppsV2(cfg *AppsV2Config) *AppsV2 {
	return &AppsV2{
		BaseURL:      cfg.BaseURL,
		client:       cfg.Client,
		Auth:         cfg.Auth,
		ErrorContext: "apps v2 service",
	}
}
