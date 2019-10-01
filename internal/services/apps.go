package services

import (
	"fmt"
	"net/http"
)

type AppsV1 struct {
	Url    string
	client *http.Client
}

type AppsV1Config struct {
	BaseUrl string
	Client  *http.Client
}

func NewAppsV1(cfg *AppsV1Config) *AppsV1 {
	return &AppsV1{
		Url:    fmt.Sprintf("%s/api", cfg.BaseUrl),
		client: cfg.Client,
	}
}
