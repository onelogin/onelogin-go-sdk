# SDK Models

The SDK provides several models that define the structure of data used for data representation and interaction with the OneLogin API. These models are utilized in the SDK's client implementation to handle requests to the OneLogin API. The `Client` struct provides methods such as `Get`, `Delete`,`Post`, and `Put` that internally use the specified models to serialize data into JSON and send HTTP requests to the API. These methods handle authentication, path construction, query parameters, and request body serialization based on the provided models.

The following is a list of models available in the SDK:

## [App](../internal/models/app.go)

The `App` model represents an application within the OneLogin platform. It contains information such as the application's name, description, icon URL, and other relevant details.

```go
type App struct {
    ID          int64  `json:"id,omitempty"`
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`
    // ...
}
```

## [AppRule](../internal/models/app_rule.go)

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

## [AuthServer](../internal/models/auth_server.go)

The `AuthServer` model represents an authentication server used for Single Sign-On (SSO) integration. It includes properties such as the server's name, issuer URL, token endpoint URL, and other relevant information.

```go
type AuthServer struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
    // ...
}
```

## [Group](../internal/models/group.go)

The `Group` model represents a user group within the OneLogin platform. It contains information about the group, such as the group name, description, and any associated custom attributes.

```go
type Group struct {
    ID        int     `json:"id"`
    Name      string  `json:"name"`
    Reference *string `json:"reference"`
}
```

## [Privilege](../internal/models/privilege.go)

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

## [Role](../internal/models/role.go)

The `Role` model represents a role within the OneLogin platform. It contains information such as the role's name, description, and any associated privileges.

```go
type Role struct {
    ID   int64  `json:"id"`
    Name string `json:"name"`
    // ...
}
```

## [SmartHooks](../internal/models/smart_hooks.go)

The `SmartHooks` model represents a set of smart hooks associated with an application. It includes properties such as hook type, script, and configuration details.

```go
type SmartHooks struct {
    ID       *string `json:"id,omitempty"`
    Type     *string `json:"type,omitempty"`
    Disabled *bool   `json:"disabled,omitempty"`
    // ...
}
```

## [User](../internal/models/user.go)

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

## [UserMapping](../internal/models/user_mapping.go)

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
