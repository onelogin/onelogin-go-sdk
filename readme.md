# Onelogin-Go-SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/onelogin/onelogin-go-sdk)](https://goreportcard.com/report/github.com/onelogin/onelogin-go-sdk)
<a href='https://github.com/dcaponi/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-88%25-brightgreen.svg?longCache=true&style=flat)</a>

  This is the Onelogin Go SDK client, supporting the following apis:

    - Apps (v2)
    - Auth (v2)

## Installation
  **Make sure you have [Go](https://golang.org/doc/install) installed.**

  1. Install the package:
    ```
    go get -u "github.com/onelogin/onelogin-go-sdk/pkg/client"
    ```
  2. Import:
    ```
    import (
      "github.com/onelogin/onelogin-go-sdk/pkg/client"
    )
    ```
## Examples

  - Starting up the client:
    ```go
    import (
      "github.com/onelogin/onelogin-go-sdk/pkg/client"
    )

    sdkClient := client.New(&client.ApiClientConfig{
      TimeoutInSeconds: 5,
      ClientId: "your_onelogin_developer_client_id",
      ClientSecret: "your_onelogin_developer_client_secret",
      Region: "us",
    })

    resp, app, err := sdkClient.Services.AppsV2.GetAppById(11111)

    if err != nil {
      // handle error
    }
    ```
