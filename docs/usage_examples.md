```go
package main

import (
	"fmt"

	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	"github.com/onelogin/onelogin-go-sdk/pkg/onelogin"
)

func main() {
	var testSSO = models.SSOOpenId{ClientID: "1234567890"}
	var testConfig = models.ConfigurationOpenId{
		RedirectURI:                   "https://example.com",
		LoginURL:                      "https://example.com/login",
		OidcApplicationType:           1,
		TokenEndpointAuthMethod:       1,
		AccessTokenExpirationMinutes:  60,
		RefreshTokenExpirationMinutes: 60,
	}

	var (
		connetorid int32  = 108419
		name       string = "Bow wow wow yippy yo yippy yay"
		descr      string = "Dog app"
	)

	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println(err)
	}
	appQ := models.AppQuery{}
	applist, err := client.GetApps(&appQ)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", applist)

	appT, err := client.CreateApp(models.App{
		ConnectorID:   &connetorid,
		Name:          &name,
		Description:   &descr,
		SSO:           testSSO,
		Configuration: testConfig,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", appT)
}

```