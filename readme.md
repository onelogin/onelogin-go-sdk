# Onelogin-Go-SDK

[![Go Report Card](https://goreportcard.com/badge/github.com/onelogin/onelogin-go-sdk)](https://goreportcard.com/report/github.com/onelogin/onelogin-go-sdk)
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-95%25-brightgreen.svg?longCache=true&style=flat)</a>

  This is the Onelogin Go SDK client, supporting the following apis:

    - [Apps](https://developers.onelogin.com/api-docs/2/apps/overview) (v2)
    - [App Rules](https://developers.onelogin.com/api-docs/2/app-rules/overview)
    - [Auth Servers](https://developers.onelogin.com/api-docs/2/api-authorization/overview) (v2)
    - [Roles](https://developers.onelogin.com/api-docs/2/roles/list-roles) (v2)
    - [Privileges](https://developers.onelogin.com/api-docs/1/privileges/list-privileges) (v1)
    - [Session Login Tokens](https://developers.onelogin.com/api-docs/1/login-page/create-session-login-token) (v1)
    - [Smart Hooks](https://developers.onelogin.com/api-docs/2/smart-hooks/overview) (v2)
    - [User Mappings](https://developers.onelogin.com/api-docs/2/user-mappings/overview) (v2)
    - [Users](https://developers.onelogin.com/api-docs/2/users/list-users) (v2)

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

    sdkClient, err := client.NewClient(&client.APIClientConfig{
      Timeout:      5,
      ClientID:     "your_onelogin_developer_client_id",
      ClientSecret: "your_onelogin_developer_client_secret",
      Region:       "us",
    })
    if err != nil {
      // handle error
    }

    app, err := sdkClient.Services.AppsV2.GetOne(12345)
    if err != nil {
      // handle error
    }
    ```

  - Requesting Apps:
  ```go
  awsApps, err := sdkClient.AppsV2.Query(apps.AppsQuery{ConnectorID: 9, Limit: 10}) // get 10 aws apps
  ```

  - Creating Apps:
  ```go
  newApp := apps.App{Name: "new app", ConnectorID: 9}
  err := sdkClient.AppsV2.Create(newApp) // Saves the app to OneLogin. Updates app in place with new state as represented in OneLogin
  if err != nil {
    return err
  }
  fmt.Println("my app", *newApp.ID)
  ```

  - Updating an App:
  ```go
  someApp, err := sdkClient.AppsV2.GetOne(123)
  someApp.Name = "updated name"
  sdkClient.AppsV2.Update(someApp) // saves the new app state to OneLogin
  ```

  - Destroying an App:
  ```go
  err := sdkClient.AppsV2.Destroy(*someApp.ID)
  ```  
