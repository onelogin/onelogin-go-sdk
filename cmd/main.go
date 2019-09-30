package main

import (
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/pkg/client"
)

func main() {
	fmt.Println("Starting Client SDK")

	sdkClient := client.New(&client.ApiClientConfig{
		TimeoutInSeconds: 5,
		ClientId:         "test",
		ClientSecret:     "test",
		Region:           "us",
	})

	fmt.Println(sdkClient.Services.AppsV1.Url)
}
