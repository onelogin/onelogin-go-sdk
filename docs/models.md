# SDK Models

The SDK provides several models that define the structure of data used for data representation and interaction with the OneLogin API. These models are utilized in the SDK's client implementation to handle requests to the OneLogin API. The `Client` struct provides methods such as `Get`, `Delete`,`Post`, and `Put` that internally use the specified models to serialize data into JSON and send HTTP requests to the API. These methods handle authentication, path construction, query parameters, and request body serialization based on the provided models.

The following is a list of models available in the SDK:

## [App](../pkg/onelogin/models/app.go)

The `App` model represents an application within the OneLogin platform. It contains information such as the application's name, description, icon URL, and other relevant details.

```go
type App struct {
    ID          int64  `json:"id,omitempty"`
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`
    // ...
}
```

## [AppRule](../pkg/onelogin/models/app_rule.go)

The `AppRule` model represents a rule associated with an application. It defines conditions and actions that determine how the application behaves in specific scenarios. The model includes properties such as rule name, match attribute, value, and actions to be performed.

```go
type AppRule struct {
    AppID      int64    `json:"id"`
    Name       string   `json:"name"`
    Match      string   `json:"match"`
    Actions    []string `json:"actions"`
    // ...
}
```

## [AuthServer](../pkg/onelogin/models/auth_server.go)

The `AuthServer` model represents an authentication server used for Single Sign-On (SSO) integration. It includes properties such as the server's name, issuer URL, token endpoint URL, and other relevant information.

```go
type AuthServer struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
    // ...
}
```

## [Brand](../pkg/onelogin/models/branding.go)

The `Brand` model represents a brand's configuration, including properties such as the brand's name, custom colors, logo, background, and login screen customization. It also includes settings for enabling features like custom support, and hiding the OneLogin footer.

```go
type Brand struct {
	ID            int32   `json:"id,omitempty"`
	Name          *string `json:"name"`
	CustomColor   *string `json:"custom_color,omitempty"`
    // ...
```

## [Group](../pkg/onelogin/models/group.go)

The `Group` model represents a user group within the OneLogin platform. It contains information about the group, such as the group name, description, and any associated custom attributes.

```go
type Group struct {
    ID        int     `json:"id"`
    Name      string  `json:"name"`
    Reference *string `json:"reference"`
}
```

## [Invite](../pkg/onelogin/models/invite_link.go)

The `Invite` model represents an invitation, with properties for the user's email to generate the invite link and an optional personal email address to send the invite link.

```go
type Invite struct {
	Email         string `json:"email"`
	PersonalEmail string `json:"personal_email,omitempty"`
}
```

## [Privilege](../pkg/onelogin/models/privilege.go)

The `Privilege` model represents a privilege assigned to a user or a group within the OneLogin platform. It includes properties such as privilege name, description, and associated roles.

```go
type Privilege struct {
    ID          int64    `json:"id"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    RoleIDs     []string `json:"role_ids"`
    // ...
}
```

## [Role](../pkg/onelogin/models/role.go)

The `Role` model represents a role within the OneLogin platform. It contains information such as the role's name and associated users, admins, and apps.

```go
type Role struct {
	ID     *int32  `json:"id,omitempty"`
	Name   *string `json:"name,omitempty"`
	Admins []int32 `json:"admins,omitempty"`
	Apps   []int32 `json:"apps,omitempty"`
	Users  []int32 `json:"users,omitempty"`
}
```

## [SmartHooks](../pkg/onelogin/models/smart_hooks.go)

The `SmartHooks` model represents a set of smart hooks associated with an application. It includes properties such as hook type, script, and configuration details.

```go
type SmartHooks struct {
    ID       *string `json:"id,omitempty"`
    Type     *string `json:"type,omitempty"`
    Disabled *bool   `json:"disabled,omitempty"`
    // ...
}
```

## [User](../pkg/onelogin/models/user.go)

The `User` model represents a user within the OneLogin platform. It contains information such as the user's email, username, first name, last name, and other relevant details.

```go


type User struct {
    Firstname *string `json:"firstname,omitempty"`
    Lastname  *string `json:"lastname,omitempty"`
    Username  *string `json:"username,omitempty"`
    Email     *string `json:"email,omitempty"`
    // ...
}
```

## [UserMapping](../pkg/onelogin/models/user_mapping.go)

The `UserMapping` model represents a mapping between a user in the OneLogin platform and an external system.

```go
type UserMapping struct {
    ID    *int32  `json:"id,omitempty"`
    Name  *string `json:"name,omitempty"`
    Match *string `json:"match,omitempty"`
    // ...
}
```

These models provide a convenient way to work with data when interacting with the OneLogin API using the SDK. You can use them to create, retrieve, update, or delete various resources within the OneLogin platform.

Feel free to explore the specific model files linked above to view the complete code and additional information.
