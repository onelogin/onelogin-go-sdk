package services

import (
	"net/http"
)

type AppsV2 struct {
	BaseUrl      string
	client       *http.Client
	Auth         Authenticator
	ErrorContext string
}

// config
type AppsV2Config struct {
	BaseUrl string
	Client  *http.Client
	Auth    Authenticator
}

// apps version 2 service
func NewAppsV2(cfg *AppsV2Config) *AppsV2 {
	return &AppsV2{
		BaseUrl:      cfg.BaseUrl,
		client:       cfg.Client,
		Auth:         cfg.Auth,
		ErrorContext: "apps v2 service",
	}
}
