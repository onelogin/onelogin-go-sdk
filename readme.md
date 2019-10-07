# Onelogin-Go-SDK
  This is the Onelogin go sdk client, supporting the following apis:
    - Apps (v2)
    - Auth (v2)

## Installation
  0) Make sure you have Go installed!

  1) Install the package
    go get -u "github.com/onelogin/onelogin-go-sdk/pkg/client"

  2) Import
    '''
    import (
      "github.com/onelogin/onelogin-go-sdk/pkg/client"
    )
    '''

## Examples

  - Starting up the client

  '''
    import (
      "github.com/onelogin/onelogin-go-sdk/pkg/client"
    )

    client.New(&client.ApiClientConfig{
      TimeoutInSeconds: 5,
      ClientId: "your_onelogin_developer_client_id",
      ClientSecret: "your_onelogin_developer_client_secret",
      Region: "us",
    })

    resp, app, err := sdkClient.Services.AppsV1.GetAppById(11111)

    if err != nil {
      // handle error
    }
  '''
