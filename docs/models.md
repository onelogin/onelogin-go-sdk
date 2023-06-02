# SDK Models

The SDK provides several models that are used for data representation and interaction with the API. These models define the structure of the data that can be sent or received when using the SDK methods.

The following is a list of models available in the SDK:

## App

The `App` model represents an application within the OneLogin platform. It contains information such as the application's name, description, icon URL, and other relevant details.

```go
type App struct {
    ID          int64  `json:"id"`
    Name        string `json:"name"`
    Description string `json:"description,omitempty"`
    // ...
}
```

## AppRule

The `AppRule` model represents a rule associated with an application. It defines conditions and actions that determine how the application behaves in specific scenarios. The model includes properties such as rule name, match attribute, value, and actions to be performed.

```go
type AppRule struct {
    AppID           int64    `json:"id"`
    Name         string   `json:"name"`
    Match        string   `json:"match"`
    Actions      []string `json:"actions"`
    // ...
}
```

## AuthServer

The `AuthServer` model represents an authentication server used for Single Sign-On (SSO) integration. It includes properties such as the server's name, issuer URL, token endpoint URL, and other relevant information.

```go
type AuthServer struct {
    ID                int64  `json:"id"`
    Name              string `json:"name"`
    // ...
}
```

## Group

The `Group` model represents a user group within the OneLogin platform. It contains information about the group, such as the group name, description, and any associated custom attributes.

```go
type Group struct {
    ID        int     `json:"id"`
    Name      string  `json:"name"`
    Reference *string `json:"reference"`
}
```

## Privilege

The `Privilege` model represents a privilege assigned to a user or a group within the OneLogin platform. It includes properties such as privilege name, description, and associated roles.

```go
type Privilege struct {
    ID          int64    `json:"id"`
    Name        string   `json:"name"`
    Description string   `json:"description"`
    RoleIDs       []string `json:"role_ids"`
    // ...
}
```

## Role

The `Role` model represents a role within the OneLogin platform. It contains information such as the role's name, description, and any associated privileges.

```go
type Role struct {
    ID          int64    `json:"id"`
    Name        string   `json:"name"`
    // ...
}
```

## SmartHooks

The `SmartHooks` model represents a set of smart hooks associated with an application. It includes

 properties such as hook type, script, and configuration details.

```go
type SmartHooks struct {
    ID             *string           `json:"id,omitempty"`
    Type           *string           `json:"type,omitempty"`
    Disabled       *bool             `json:"disabled,omitempty"`
    // ...
}
```

## User

The `User` model represents a user within the OneLogin platform. It contains information such as the user's email, username, first name, last name, and other relevant details.

```go
type User struct {
    Firstname    *string        `json:"firstname,omitempty"`
    Lastname     *string        `json:"lastname,omitempty"`
    Username     *string        `json:"username,omitempty"`
    Email        *string        `json:"email,omitempty"`
    // ...
}
```

## UserMapping

The `UserMapping` model represents a mapping between a user in the OneLogin platform and an external system. It includes properties such as the external system's ID, user's ID, and any associated attributes.

```go
type UserMapping struct {
    ID         *int32              `json:"id,omitempty"`
    Name       *string             `json:"name,omitempty"`
    Match      *string             `json:"match,omitempty"`
    // ...
}
```

These models provide a convenient way to work with data when interacting with the OneLogin API using the SDK. You can use them to create, retrieve, update, or delete various resources within the OneLogin platform.

Please refer to the corresponding sections of the documentation for more details on how to use these models in conjunction with the SDK methods.