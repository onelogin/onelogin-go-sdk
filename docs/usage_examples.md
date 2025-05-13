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
	// Initialize the SDK
	client, err := onelogin.NewOneloginSDK()
	if err != nil {
		fmt.Println(err)
	}

	// Get authentication token
	_, err = client.GetToken()
	if err != nil {
		fmt.Println(err)
	}

	// List all user mappings with query parameters
	userMappingsQuery := &models.UserMappingsQuery{
		Limit: "10",
		Page:  "1",
	}
	
	mappings, err := client.ListUserMappings(userMappingsQuery)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Found %d user mappings\n", len(mappings))
	
	// Create a new user mapping
	name := "Test Mapping via SDK"
	match := "all"
	enabled := true
	position := int32(1)
	
	// Set up conditions
	sourceCondition := "firstname"
	equalsOperator := "eq"
	valueCondition := "Test"
	
	// Set up actions
	actionName := "set_status"
	actionValue := []string{"2"}
	
	newMapping := models.UserMapping{
		Name:     &name,
		Match:    &match,
		Enabled:  &enabled,
		Position: &position,
		Conditions: []models.UserMappingConditions{
			{
				Source:   &sourceCondition,
				Operator: &equalsOperator,
				Value:    &valueCondition,
			},
		},
		Actions: []models.UserMappingActions{
			{
				Action: &actionName,
				Value:  actionValue,
			},
		},
	}
	
	createdMapping, err := client.CreateUserMapping(newMapping)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Created new mapping with ID: %d\n", *createdMapping.ID)
	
	// Get a specific mapping
	mapping, err := client.GetUserMapping(*createdMapping.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Retrieved mapping: %s\n", *mapping.Name)
	
	// Update a mapping
	updatedName := "Updated Test Mapping"
	mapping.Name = &updatedName
	
	updatedMapping, err := client.UpdateUserMapping(*mapping.ID, *mapping)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Updated mapping: %s\n", *updatedMapping.Name)
	
	// Delete a mapping
	err = client.DeleteUserMapping(*updatedMapping.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Deleted mapping successfully")
}
```

For a more detailed example, see the `cmd/user_mappings_test/main.go` file.

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
