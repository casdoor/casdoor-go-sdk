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

Casdoor Go SDK is the official Go client for [Casdoor](https://casdoor.org), a UI-first Identity and Access Management (IAM) / Single-Sign-On (SSO) platform supporting OAuth 2.0, OIDC, SAML and CAS. This SDK allows you to easily integrate Casdoor authentication and authorization into your Go applications without implementing these complex protocols from scratch.

## Table of Contents

- [Features](#features)
- [Installation](#installation)
- [Quick Start](#quick-start)
  - [Configuration](#configuration)
  - [Authentication Flow](#authentication-flow)
  - [Session Management](#session-management)
- [Usage Examples](#usage-examples)
  - [User Management](#user-management)
  - [Organization Management](#organization-management)
  - [Role Management](#role-management)
  - [Permission Management](#permission-management)
- [Supported Resources](#supported-resources)
- [Error Handling](#error-handling)
- [Documentation](#documentation)
- [Community](#community)
- [Contributing](#contributing)
- [License](#license)

## Features

- 🔐 **OAuth 2.0 & OIDC Authentication** - Full support for modern authentication protocols
- 👥 **User Management** - Create, read, update, and delete user accounts
- 🏢 **Organization Management** - Multi-tenant organization support
- 🔑 **Role-Based Access Control (RBAC)** - Manage roles and permissions
- 📜 **Policy Enforcement** - Integration with Casbin for fine-grained access control
- 🎫 **Token Management** - Handle access tokens, refresh tokens, and JWT validation
- 📧 **Email & SMS** - Send verification emails and SMS messages
- 🔒 **Multi-Factor Authentication (MFA)** - Support for TOTP and other MFA methods
- 💳 **Payment & Subscription** - Manage user subscriptions and payments
- 🌐 **Multi-Language Support** - Built-in internationalization
- 🔗 **LDAP/AD Integration** - Sync users from LDAP and Active Directory
- 📊 **Audit & Analytics** - Track user activities and generate reports

## Installation

To use Casdoor Go SDK, you need Go 1.17 or later.

```bash
go get github.com/casdoor/casdoor-go-sdk@latest
```

For production use, we recommend pinning to a specific version:

```bash
go get github.com/casdoor/casdoor-go-sdk@v0.x.x
```

Then import the SDK in your Go code:

```go
import "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
```

## Quick Start

### Configuration

To use the Casdoor SDK, you first need to initialize it with your Casdoor server configuration. The SDK requires 6 parameters:

| Parameter        | Required | Description                                         |
|------------------|----------|-----------------------------------------------------|
| endpoint         | Yes      | Casdoor server URL, such as `http://localhost:8000` |
| clientId         | Yes      | Application.clientId                                |
| clientSecret     | Yes      | Application.clientSecret                            |
| certificate      | Yes      | x509 certificate content of Application.cert        |
| organizationName | Yes      | Application.organization                            |
| applicationName  | Yes      | Application.name                                    |

**Option 1: Global Configuration (Recommended for single configuration)**

```go
func main() {
    casdoorsdk.InitConfig(
        "http://localhost:8000",
        "clientId",
        "clientSecret",
        "certificate",
        "organizationName",
        "applicationName",
    )

    // Use SDK functions directly
    users, err := casdoorsdk.GetUsers()
    if err != nil {
        panic(err)
    }
}
```

**Option 2: Custom Client (Recommended for multiple configurations)**

```go
func main() {
    client := casdoorsdk.NewClient(
        "http://localhost:8000",
        "clientId",
        "clientSecret",
        "certificate",
        "organizationName",
        "applicationName",
    )

    // Use client methods
    users, err := client.GetUsers()
    if err != nil {
        panic(err)
    }
}
```

### Authentication Flow

After a successful Casdoor authentication, the user will be redirected back to your application with a `code` and `state` parameter (e.g., `https://your-app.com/callback?code=xxx&state=yyyy`).

```go
// In your callback handler
func callbackHandler(w http.ResponseWriter, r *http.Request) {
    code := r.URL.Query().Get("code")
    state := r.URL.Query().Get("state")

    // Exchange code for token
    token, err := casdoorsdk.GetOAuthToken(code, state)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Parse JWT token to get user claims
    claims, err := casdoorsdk.ParseJwtToken(token.AccessToken)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Store the access token in claims for future API calls
    claims.AccessToken = token.AccessToken
    
    // Now you have the authenticated user information
    // claims.User contains user details
}
```

### Session Management

`claims` contains the user information provided by Casdoor. You should store this in your application's session:

```go
// Serialize claims to JSON
data, err := json.Marshal(claims)
if err != nil {
    // handle error
}

// Store in session (example using gorilla/sessions)
session.Values["user"] = data
session.Save(r, w)
```

## Usage Examples

### User Management

The SDK provides comprehensive user management capabilities:

```go
// Get all users
users, err := casdoorsdk.GetUsers()
if err != nil {
    log.Fatal(err)
}

// Get a specific user by name
user, err := casdoorsdk.GetUser("admin")
if err != nil {
    log.Fatal(err)
}

// Get paginated users
users, count, err := casdoorsdk.GetPaginationUsers(1, 10, map[string]string{
    "owner": "built-in",
})

// Create a new user
newUser := &casdoorsdk.User{
    Owner:       "built-in",
    Name:        "user123",
    DisplayName: "User 123",
    Email:       "user123@example.com",
    Password:    "password123",
    Phone:       "+1234567890",
}
success, err := casdoorsdk.AddUser(newUser)

// Update a user
user.DisplayName = "Updated Name"
success, err = casdoorsdk.UpdateUser(user)

// Delete a user
success, err = casdoorsdk.DeleteUser(user)
```

### Organization Management

Manage multi-tenant organizations:

```go
// Get all organizations
orgs, err := casdoorsdk.GetOrganizations()

// Get a specific organization
org, err := casdoorsdk.GetOrganization("my-org")

// Create a new organization
newOrg := &casdoorsdk.Organization{
    Owner:       "admin",
    Name:        "my-organization",
    DisplayName: "My Organization",
    WebsiteUrl:  "https://example.com",
}
success, err := casdoorsdk.AddOrganization(newOrg)

// Update an organization
org.DisplayName = "Updated Organization Name"
success, err = casdoorsdk.UpdateOrganization(org)

// Delete an organization
success, err = casdoorsdk.DeleteOrganization(org)
```

### Role Management

Implement role-based access control:

```go
// Get all roles
roles, err := casdoorsdk.GetRoles()

// Get a specific role
role, err := casdoorsdk.GetRole("admin")

// Create a new role
newRole := &casdoorsdk.Role{
    Owner:       "built-in",
    Name:        "editor",
    DisplayName: "Editor",
    Description: "Can edit content",
    Users:       []string{"user1", "user2"},
    IsEnabled:   true,
}
success, err := casdoorsdk.AddRole(newRole)

// Update a role
role.Users = append(role.Users, "user3")
success, err = casdoorsdk.UpdateRole(role)

// Delete a role
success, err = casdoorsdk.DeleteRole(role)
```

### Permission Management

Define fine-grained permissions:

```go
// Get all permissions
permissions, err := casdoorsdk.GetPermissions()

// Get a specific permission
permission, err := casdoorsdk.GetPermission("permission1")

// Create a new permission
newPermission := &casdoorsdk.Permission{
    Owner:        "built-in",
    Name:         "read-documents",
    DisplayName:  "Read Documents",
    Description:  "Permission to read documents",
    Users:        []string{"user1"},
    Roles:        []string{"editor"},
    ResourceType: "document",
    Resources:    []string{"doc1", "doc2"},
    Actions:      []string{"read"},
}
success, err := casdoorsdk.AddPermission(newPermission)

// Update a permission
permission.Actions = append(permission.Actions, "write")
success, err = casdoorsdk.UpdatePermission(permission)

// Delete a permission
success, err = casdoorsdk.DeletePermission(permission)

// Enforce permission
allowed, err := casdoorsdk.Enforce(permission, "user1", "doc1", "read")
```

## Supported Resources

The Casdoor Go SDK supports managing the following resources:

| Resource | Description | Key Operations |
|----------|-------------|----------------|
| **Users** | User accounts and profiles | Get, Add, Update, Delete, GetSortedUsers, GetUserCount |
| **Organizations** | Multi-tenant organizations | Get, Add, Update, Delete |
| **Applications** | OAuth/OIDC applications | Get, Add, Update, Delete |
| **Roles** | RBAC roles | Get, Add, Update, Delete |
| **Permissions** | Fine-grained permissions | Get, Add, Update, Delete, Enforce |
| **Groups** | User groups | Get, Add, Update, Delete |
| **Providers** | Authentication providers (OAuth, SAML, etc.) | Get, Add, Update, Delete |
| **Tokens** | Access and refresh tokens | Get, Add, Update, Delete, GetOAuthToken |
| **Certificates** | x509 certificates for JWT | Get, Add, Update, Delete |
| **Products** | Subscription products | Get, Add, Update, Delete |
| **Payments** | Payment records | Get, Add, Update, Delete |
| **Plans** | Subscription plans | Get, Add, Update, Delete |
| **Pricings** | Pricing models | Get, Add, Update, Delete |
| **Subscriptions** | User subscriptions | Get, Add, Update, Delete |
| **Sessions** | User sessions | Get, Add, Update, Delete |
| **Records** | Audit logs | Get operations |
| **Webhooks** | Event webhooks | Get, Add, Update, Delete |
| **Syncers** | Data synchronization | Get, Add, Update, Delete |
| **Models** | Casbin models | Get, Add, Update, Delete |
| **Adapters** | Casbin adapters | Get, Add, Update, Delete |
| **Enforcers** | Casbin enforcers | Get, Add, Update, Delete |
| **Policies** | Casbin policies | Get, Add, Update, Delete, AddPolicy, RemovePolicy |
| **Resources** | File resources | Upload, Delete |
| **LDAP** | LDAP/AD integration | Get, Add, Update, Delete, SyncUsers |
| **Invitations** | User invitations | Get, Add, Update, Delete |

## Error Handling

The SDK functions return errors that should be properly handled:

```go
user, err := casdoorsdk.GetUser("username")
if err != nil {
    // Handle error
    log.Printf("Failed to get user: %v", err)
    return
}

if user == nil {
    // User not found
    log.Println("User not found")
    return
}

// Process user
fmt.Printf("User: %s\n", user.DisplayName)
```

For operations that modify data (Add, Update, Delete), the SDK returns a boolean indicating success:

```go
success, err := casdoorsdk.AddUser(newUser)
if err != nil {
    log.Printf("Error adding user: %v", err)
    return
}

if !success {
    log.Println("Failed to add user")
    return
}

log.Println("User added successfully")
```

## Documentation

- **Casdoor Official Website**: https://casdoor.org
- **Casdoor Documentation**: https://casdoor.org/docs/overview
- **Go SDK Documentation**: https://pkg.go.dev/github.com/casdoor/casdoor-go-sdk
- **Casdoor GitHub**: https://github.com/casdoor/casdoor
- **Examples and Demos**: https://github.com/casdoor/casdoor/tree/master/examples

## Community

- **Forum**: https://forum.casbin.com
- **Discord**: https://discord.gg/5rPsrAzK7S
- **Gitter**: https://gitter.im/casbin/casdoor
- **QQ Group**: 645200447
- **WeChat**: Contact us to be added to the group

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/AmazingFeature`)
3. Commit your changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

## License

This project is licensed under the [Apache 2.0 License](LICENSE).
