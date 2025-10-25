# Casdoor Go SDK

<p align="center">
  <a href="#badge">
    <img alt="semantic-release" src="https://img.shields.io/badge/%20%20%F0%9F%93%A6%F0%9F%9A%80-semantic--release-e10079.svg">
  </a>
  <a href="https://github.com/casdoor/casdoor-go-sdk/actions/workflows/ci.yml">
    <img alt="GitHub Workflow Status (branch)" src="https://img.shields.io/github/actions/workflow/status/casdoor/casdoor-go-sdk/ci.yml?branch=master">
  </a>
  <a href="https://github.com/casdoor/casdoor-go-sdk/releases/latest">
    <img alt="GitHub Release" src="https://img.shields.io/github/v/release/casdoor/casdoor-go-sdk.svg">
  </a>
</p>

<p align="center">
  <a href="https://goreportcard.com/report/github.com/casdoor/casdoor-go-sdk">
    <img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/casdoor/casdoor-go-sdk?style=flat-square">
  </a>
  <a href="https://github.com/casdoor/casdoor-go-sdk/blob/master/LICENSE">
    <img src="https://img.shields.io/github/license/casdoor/casdoor-go-sdk?style=flat-square" alt="license">
  </a>
  <a href="https://github.com/casdoor/casdoor-go-sdk/issues">
    <img alt="GitHub issues" src="https://img.shields.io/github/issues/casdoor/casdoor-go-sdk?style=flat-square">
  </a>
  <a href="#">
    <img alt="GitHub stars" src="https://img.shields.io/github/stars/casdoor/casdoor-go-sdk?style=flat-square">
  </a>
  <a href="https://github.com/casdoor/casdoor-go-sdk/network">
    <img alt="GitHub forks" src="https://img.shields.io/github/forks/casdoor/casdoor-go-sdk?style=flat-square">
  </a>
  <a href="https://discord.gg/5rPsrAzK7S">
    <img alt="Casdoor" src="https://img.shields.io/discord/1022748306096537660?style=flat-square&logo=discord&label=discord&color=5865F2">
  </a>
</p>

Casdoor Go SDK is the official Go client library for [Casdoor](https://casdoor.org/), which allows you to easily integrate Casdoor authentication and authorization into your Go applications. This SDK provides a comprehensive set of APIs to interact with Casdoor server, enabling you to manage users, organizations, applications, roles, permissions, and much more.

## 📋 Table of Contents

- [Features](#-features)
- [Installation](#-installation)
- [Quick Start](#-quick-start)
- [Configuration](#-configuration)
- [Authentication](#-authentication)
- [Resource Management](#-resource-management)
- [API Reference](#-api-reference)
- [Examples](#-examples)
- [Documentation](#-documentation)
- [Contributing](#-contributing)
- [License](#-license)

## ✨ Features

- **OAuth 2.0 Authentication**: Complete OAuth 2.0 flow implementation with token refresh
- **User Management**: Create, read, update, and delete users with comprehensive profile support
- **Organization Management**: Manage organizations and organizational structures
- **Role-Based Access Control (RBAC)**: Full support for roles, permissions, and policies
- **Resource Management**: Manage applications, certificates, providers, and more
- **Session Management**: Handle user sessions and authentication states
- **Multi-Factor Authentication (MFA)**: Support for TOTP and other MFA methods
- **Email & SMS**: Send verification codes and notifications
- **Payment & Subscriptions**: Handle user payments and subscription management
- **Webhook Support**: Configure and manage webhooks for event notifications
- **Global and Client-Specific Configuration**: Flexible configuration options

## 📦 Installation

To install the Casdoor Go SDK, you need Go 1.17 or higher. Run the following command in your Go project:

```bash
go get github.com/casdoor/casdoor-go-sdk@latest
```

Then import the SDK in your Go files:

```go
import "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
```

## 🚀 Quick Start

Here's a minimal example to get you started with Casdoor Go SDK:

```go
package main

import (
    "fmt"
    "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func main() {
    // Initialize the SDK with your Casdoor instance configuration
    casdoorsdk.InitConfig(
        "http://localhost:8000",           // endpoint
        "CLIENT_ID",                        // clientId
        "CLIENT_SECRET",                    // clientSecret
        "CERTIFICATE_CONTENT",              // certificate (x509 format)
        "my-organization",                  // organizationName
        "my-application",                   // applicationName
    )

    // Get all users
    users, err := casdoorsdk.GetUsers()
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Found %d users\n", len(users))
}
```

## ⚙️ Configuration

The SDK requires six configuration parameters. You can configure the SDK in two ways:

### Method 1: Global Configuration

Initialize once and use throughout your application:

```go
casdoorsdk.InitConfig(endpoint, clientId, clientSecret, certificate, organizationName, applicationName)

// Then use SDK functions directly
users, err := casdoorsdk.GetUsers()
token, err := casdoorsdk.GetOAuthToken(code, state)
```

### Method 2: Client-Specific Configuration

Create multiple clients with different configurations:

```go
client1 := casdoorsdk.NewClient(endpoint, clientId, clientSecret, certificate, organizationName, applicationName)
users, err := client1.GetUsers()

client2 := casdoorsdk.NewClient(endpoint2, clientId2, clientSecret2, certificate2, organizationName2, applicationName2)
users2, err := client2.GetUsers()
```

### Configuration Parameters

| Parameter        | Required | Description                                                  |
|------------------|----------|--------------------------------------------------------------|
| endpoint         | Yes      | Casdoor server URL (e.g., `http://localhost:8000`)          |
| clientId         | Yes      | Application client ID from Casdoor                           |
| clientSecret     | Yes      | Application client secret from Casdoor                       |
| certificate      | Yes      | x509 certificate content of your application (PEM format)    |
| organizationName | Yes      | Organization name in Casdoor                                 |
| applicationName  | Yes      | Application name in Casdoor                                  |

### Getting Configuration Parameters from Casdoor

1. **endpoint**: Your Casdoor server URL
2. **clientId** and **clientSecret**: Found in your application settings in Casdoor admin panel
3. **certificate**: Copy the certificate content from your application's "Cert" field (must be in x509 PEM format)
4. **organizationName**: The organization that owns your application
5. **applicationName**: Your application's name in Casdoor

## 🔐 Authentication

### OAuth 2.0 Flow

The SDK provides complete OAuth 2.0 authentication flow support.

#### Step 1: Redirect User to Casdoor Login

Direct users to your Casdoor login page:

```
https://your-casdoor-instance.com/login/oauth/authorize?client_id=CLIENT_ID&response_type=code&redirect_uri=REDIRECT_URI&scope=read&state=STATE
```

#### Step 2: Handle OAuth Callback

After successful authentication, Casdoor redirects back to your application with `code` and `state` parameters:

```go
// Extract code and state from the callback URL
code := r.URL.Query().Get("code")
state := r.URL.Query().Get("state")

// Exchange code for access token
token, err := casdoorsdk.GetOAuthToken(code, state)
if err != nil {
    panic(err)
}

// Parse the JWT token to get user information
claims, err := casdoorsdk.ParseJwtToken(token.AccessToken)
if err != nil {
    panic(err)
}

// Store the access token in claims for future use
claims.AccessToken = token.AccessToken
```

#### Step 3: Store User Session

After getting user information, store it in your session:

```go
import "encoding/json"

data, err := json.Marshal(claims)
if err != nil {
    panic(err)
}
// Store in session (implementation depends on your session management)
session.Set("user", data)
```

### Token Refresh

Refresh an expired access token using the refresh token:

```go
newToken, err := casdoorsdk.RefreshOAuthToken(refreshToken)
if err != nil {
    panic(err)
}
```

### JWT Token Parsing

Parse and validate JWT tokens:

```go
claims, err := casdoorsdk.ParseJwtToken(accessToken)
if err != nil {
    panic(err)
}

// Access user information
fmt.Printf("User: %s\n", claims.Name)
fmt.Printf("Email: %s\n", claims.Email)
fmt.Printf("Organization: %s\n", claims.Owner)
```

## 📦 Resource Management

The SDK provides comprehensive APIs to manage various resources in Casdoor.

### User Management

```go
// Get all users in your organization
users, err := casdoorsdk.GetUsers()

// Get a specific user by name
user, err := casdoorsdk.GetUser("username")

// Get user by email
user, err := casdoorsdk.GetUserByEmail("user@example.com")

// Get user by phone number
user, err := casdoorsdk.GetUserByPhone("+1234567890")

// Get paginated users
users, totalCount, err := casdoorsdk.GetPaginationUsers(
    1,          // page number
    10,         // page size
    map[string]string{},  // query filters
)

// Create a new user
user := &casdoorsdk.User{
    Owner:       "my-organization",
    Name:        "new-user",
    DisplayName: "New User",
    Email:       "newuser@example.com",
    Phone:       "+1234567890",
    Password:    "password123",
}
success, err := casdoorsdk.AddUser(user)

// Update an existing user
user.DisplayName = "Updated Name"
success, err := casdoorsdk.UpdateUser(user)

// Update specific user fields
success, err := casdoorsdk.UpdateUserForColumns(user, []string{"displayName", "email"})

// Delete a user
success, err := casdoorsdk.DeleteUser(user)

// Set/change user password
success, err := casdoorsdk.SetPassword("owner", "username", "oldPassword", "newPassword")

// Get user count
count, err := casdoorsdk.GetUserCount("1") // "1" for online users, "0" for all users
```

### Organization Management

```go
// Get all organizations
orgs, err := casdoorsdk.GetOrganizations()

// Get a specific organization
org, err := casdoorsdk.GetOrganization("org-name")

// Add a new organization
org := &casdoorsdk.Organization{
    Name:        "new-org",
    DisplayName: "New Organization",
    // ... other fields
}
success, err := casdoorsdk.AddOrganization(org)

// Update an organization
success, err := casdoorsdk.UpdateOrganization(org)

// Delete an organization
success, err := casdoorsdk.DeleteOrganization(org)
```

### Role Management

```go
// Get all roles
roles, err := casdoorsdk.GetRoles()

// Get a specific role
role, err := casdoorsdk.GetRole("role-name")

// Add a new role
role := &casdoorsdk.Role{
    Owner:       "my-organization",
    Name:        "admin",
    DisplayName: "Administrator",
    Users:       []string{"user1", "user2"},
    IsEnabled:   true,
}
success, err := casdoorsdk.AddRole(role)

// Update a role
success, err := casdoorsdk.UpdateRole(role)

// Delete a role
success, err := casdoorsdk.DeleteRole(role)
```

### Permission Management

```go
// Get all permissions
permissions, err := casdoorsdk.GetPermissions()

// Get a specific permission
permission, err := casdoorsdk.GetPermission("permission-name")

// Add a new permission
permission := &casdoorsdk.Permission{
    Owner:        "my-organization",
    Name:         "read-permission",
    DisplayName:  "Read Permission",
    Resources:    []string{"resource1", "resource2"},
    Actions:      []string{"read", "list"},
    Effect:       "Allow",
    IsEnabled:    true,
}
success, err := casdoorsdk.AddPermission(permission)

// Update a permission
success, err := casdoorsdk.UpdatePermission(permission)

// Delete a permission
success, err := casdoorsdk.DeletePermission(permission)

// Enforce permission check
allowed, err := casdoorsdk.Enforce("user", "resource", "action")
```

### Application Management

```go
// Get all applications
apps, err := casdoorsdk.GetApplications()

// Get a specific application
app, err := casdoorsdk.GetApplication("app-name")

// Add a new application
app := &casdoorsdk.Application{
    Owner:       "my-organization",
    Name:        "my-app",
    DisplayName: "My Application",
    // ... other configuration
}
success, err := casdoorsdk.AddApplication(app)

// Update an application
success, err := casdoorsdk.UpdateApplication(app)

// Delete an application
success, err := casdoorsdk.DeleteApplication(app)
```

### Session Management

```go
// Get all sessions
sessions, err := casdoorsdk.GetSessions()

// Get a specific session
session, err := casdoorsdk.GetSession("session-id")

// Delete a session
success, err := casdoorsdk.DeleteSession(session)
```

### Provider Management

Manage third-party authentication providers (OAuth, SAML, etc.):

```go
// Get all providers
providers, err := casdoorsdk.GetProviders()

// Get a specific provider
provider, err := casdoorsdk.GetProvider("provider-name")

// Add, update, or delete providers
success, err := casdoorsdk.AddProvider(provider)
success, err := casdoorsdk.UpdateProvider(provider)
success, err := casdoorsdk.DeleteProvider(provider)
```

### Certificate Management

```go
// Get all certificates
certs, err := casdoorsdk.GetCertificates()

// Get a specific certificate
cert, err := casdoorsdk.GetCertificate("cert-name")

// Add, update, or delete certificates
success, err := casdoorsdk.AddCertificate(cert)
success, err := casdoorsdk.UpdateCertificate(cert)
success, err := casdoorsdk.DeleteCertificate(cert)
```

### Email and SMS

```go
// Send email
err := casdoorsdk.SendEmail("Email Title", "email-content", "sender@example.com", "receiver@example.com")

// Send SMS
err := casdoorsdk.SendSms("randomCode", "+1234567890")
```

### Token Management

```go
// Get all tokens
tokens, err := casdoorsdk.GetTokens()

// Get a specific token
token, err := casdoorsdk.GetToken("owner", "token-name")

// Update or delete tokens
success, err := casdoorsdk.UpdateToken(token)
success, err := casdoorsdk.DeleteToken(token)
```

### Resource Management

Upload and manage resources (files, images, etc.):

```go
// Upload a resource
success, resource, err := casdoorsdk.UploadResource(
    "user",
    "tag",
    "parent",
    "fullFilePath",
    fileBuffer,
)

// Delete a resource
success, err := casdoorsdk.DeleteResource(resource)
```

### Webhook Management

```go
// Get all webhooks
webhooks, err := casdoorsdk.GetWebhooks()

// Get a specific webhook
webhook, err := casdoorsdk.GetWebhook("webhook-name")

// Add, update, or delete webhooks
success, err := casdoorsdk.AddWebhook(webhook)
success, err := casdoorsdk.UpdateWebhook(webhook)
success, err := casdoorsdk.DeleteWebhook(webhook)
```

## 📚 API Reference

### Available Resources

The SDK provides comprehensive support for managing the following Casdoor resources:

| Resource      | Description                                   | Key Operations                    |
|---------------|-----------------------------------------------|-----------------------------------|
| **User**      | User accounts and profiles                    | CRUD, authentication, password    |
| **Organization** | Organization entities                      | CRUD operations                   |
| **Application** | Application configurations                  | CRUD operations                   |
| **Role**      | User roles                                    | CRUD, assignment                  |
| **Permission** | Access permissions                           | CRUD, enforcement                 |
| **Provider**  | Third-party authentication providers          | CRUD operations                   |
| **Token**     | Access and refresh tokens                     | Get, update, delete               |
| **Certificate** | SSL/TLS certificates                        | CRUD operations                   |
| **Session**   | User sessions                                 | Get, delete                       |
| **Webhook**   | Event webhooks                                | CRUD operations                   |
| **Group**     | User groups                                   | CRUD operations                   |
| **Syncer**    | User synchronization from external systems    | CRUD operations                   |
| **Adapter**   | Policy adapters                               | CRUD operations                   |
| **Enforcer**  | Policy enforcers                              | CRUD operations                   |
| **Model**     | Policy models                                 | CRUD operations                   |
| **Policy**    | Access control policies                       | CRUD operations                   |
| **Payment**   | Payment records                               | CRUD operations                   |
| **Product**   | Products/services                             | CRUD operations                   |
| **Subscription** | User subscriptions                         | CRUD operations                   |
| **Plan**      | Subscription plans                            | CRUD operations                   |
| **Pricing**   | Pricing configurations                        | CRUD operations                   |
| **Transaction** | Payment transactions                        | CRUD operations                   |
| **Resource**  | File and media resources                      | Upload, delete                    |
| **Invitation** | User invitations                             | CRUD operations                   |
| **Record**    | Audit and activity records                    | Get operations                    |

### Method Patterns

Most resources follow consistent method patterns:

- `Get{Resource}s()` - Get all resources
- `Get{Resource}(name)` - Get a specific resource
- `Add{Resource}(resource)` - Create a new resource
- `Update{Resource}(resource)` - Update an existing resource
- `Delete{Resource}(resource)` - Delete a resource

## 💡 Examples

### Complete Authentication Example

```go
package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func main() {
    // Initialize SDK
    casdoorsdk.InitConfig(
        "http://localhost:8000",
        "CLIENT_ID",
        "CLIENT_SECRET",
        "CERTIFICATE",
        "my-organization",
        "my-app",
    )

    // Handle OAuth callback
    http.HandleFunc("/callback", func(w http.ResponseWriter, r *http.Request) {
        code := r.URL.Query().Get("code")
        state := r.URL.Query().Get("state")

        token, err := casdoorsdk.GetOAuthToken(code, state)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        claims, err := casdoorsdk.ParseJwtToken(token.AccessToken)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        // Store in session
        response, err := json.Marshal(claims)
        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }
        w.Header().Set("Content-Type", "application/json")
        w.Write(response)
    })

    http.ListenAndServe(":8080", nil)
}
```

### User Management Example

```go
package main

import (
    "fmt"
    "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func main() {
    casdoorsdk.InitConfig("http://localhost:8000", "CLIENT_ID", "CLIENT_SECRET", "CERT", "org", "app")

    // Create a new user
    newUser := &casdoorsdk.User{
        Owner:       "my-organization",
        Name:        "john_doe",
        DisplayName: "John Doe",
        Email:       "john@example.com",
        Password:    "secure_password",
    }

    success, err := casdoorsdk.AddUser(newUser)
    if err != nil {
        panic(err)
    }
    fmt.Printf("User created: %v\n", success)

    // Get and update user
    user, err := casdoorsdk.GetUser("john_doe")
    if err != nil {
        panic(err)
    }

    user.DisplayName = "John D."
    success, err = casdoorsdk.UpdateUser(user)
    if err != nil {
        panic(err)
    }
    fmt.Printf("User updated: %v\n", success)
}
```

### Permission Enforcement Example

```go
package main

import (
    "fmt"
    "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
)

func main() {
    casdoorsdk.InitConfig("http://localhost:8000", "CLIENT_ID", "CLIENT_SECRET", "CERT", "org", "app")

    // Check if user has permission
    allowed, err := casdoorsdk.Enforce("user123", "resource1", "read")
    if err != nil {
        panic(err)
    }

    if allowed {
        fmt.Println("Access granted")
    } else {
        fmt.Println("Access denied")
    }
}
```

## 📖 Documentation

For more detailed information, please refer to:

- [Casdoor Official Documentation](https://casdoor.org/docs/overview)
- [Casdoor GitHub Repository](https://github.com/casdoor/casdoor)
- [API Documentation](https://door.casdoor.com/swagger)
- [GoDoc Reference](https://pkg.go.dev/github.com/casdoor/casdoor-go-sdk/casdoorsdk)

## 📄 License

This project is licensed under the Apache License 2.0 - see the [LICENSE](LICENSE) file for details.

## 🌟 Star History

If you find this project useful, please consider giving it a star! ⭐

[![Star History Chart](https://api.star-history.com/svg?repos=casdoor/casdoor-go-sdk&type=Date)](https://star-history.com/#casdoor/casdoor-go-sdk&Date)
