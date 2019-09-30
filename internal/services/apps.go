package services

import (
	"fmt"
	"net/http"
)

type AppsV1 struct {
	Url    string
	client *http.Client
}

func NewAppsV1(baseUrl string, client *http.Client) *AppsV1 {
	return &AppsV1{
		Url:    fmt.Sprintf("%s/api", baseUrl),
		client: client,
	}
}
