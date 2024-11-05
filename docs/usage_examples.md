1. **User model**

```go
package main

import (
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
)

func main() {
	var (
		email    string = "test@example.com"
		username string = "testuser"
	)

	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println(err)
	}

	user, err := client.CreateUser(models.User{
		Email:    &email,
		Username: &username,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", user)
}
```

2. **Role model**

```go
package main

import (
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
)

func main() {
	var (
		name string = "Admin"
	)

	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println(err)
	}

	role, err := client.CreateRole(models.Role{
		Name: &name,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", role)
}
```

3. **Event model**

```go
package main

import (
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
)

func main() {
	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println(err)
	}

	eventQ := models.EventQuery{}
	eventList, err := client.GetEvents(&eventQ)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", eventList)
}
```


4. **App model**

```go
package main

import (
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
)

func main() {
	var (
		connetorid int32  = 108419
		name       string = "Bow wow wow yippy yo yippy yay"
		descr      string = "Dog app"
	)

	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println(err)
	}

	app, err := client.CreateApp(models.App{
		ConnectorID: &connetorid,
		Name:        &name,
		Description: &descr,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", app)
}
```

5. **UserMappings model**

```go
package main

import (
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
)

func main() {
	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println(err)
	}

	userMappingsQuery := models.UserMappingsQuery{}
	userMappings, err := client.GetUserMappings(&userMappingsQuery)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", userMappings)
}
```

6. **AppRule model**

```go
package main

import (
	"fmt"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin/models"
	"github.com/onelogin/onelogin-go-sdk/v4/pkg/onelogin"
)

func main() {
	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println(err)
	}

	appRuleQuery := models.AppRuleQuery{}
	appRules, err := client.GetAppRules(&appRuleQuery)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", appRules)
}
```

Please note that these are basic examples and may not work as expected without proper setup and context. You may need to adjust them according to your needs.
