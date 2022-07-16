# casdoor-go-sdk

This is Casdoor's SDK for golang, which will allow you to easily connect your application to the Casdoor authentication system without having to implement it from scratch.

Casdoor SDK is very simple to use. We will show you the steps below.

> Noted that this sdk has been applied to casnode, if you still don’t know how to use it after reading README.md, you can refer to it

## Step1. Install and Import

First in your go project, just need to run:

```bash
go get https://github.com/casdoor/casdoor-go-sdk@latest
```

and import this when you need:

```go
import "github.com/casdoor/casdoor-go-sdk/casdoorsdk"
```

## Step2. Init Config

Initialization requires 6 parameters, which are all string type:

| Name (in order)  | Must | Description                                         |
| ---------------- | ---- | --------------------------------------------------- |
| endpoint         | Yes  | Casdoor server URL, such as `http://localhost:8000` |
| clientId         | Yes  | Application.clientId                                |
| clientSecret     | Yes  | Application.clientSecret                            |
| certificate      | Yes  | x509 certificate content of Application.cert        |
| organizationName | Yes  | Application.organization                            |
| applicationName  | Yes  | Application.applicationName                         |

```go
func InitConfig(endpoint string, clientId string, clientSecret string, certificate string, organizationName string, applicationName string)
```

## Step3. Get token and parse

After casdoor verification passed, it will be redirected to your application with code and state, like `https://forum.casbin.com?code=xxx&state=yyyy`.

Your web application can get the `code`,`state` and call `GetOAuthToken(code, state)`, then parse out jwt token.

The general process is as follows:

```go
token, err := casdoorsdk.GetOAuthToken(code, state)
if err != nil {
	panic(err)
}

claims, err := casdoorsdk.ParseJwtToken(token.AccessToken)
if err != nil {
	panic(err)
}

claims.AccessToken = token.AccessToken
```

## Step4. Set Session in your app

`auth.Claims` contains the basic information about the user provided by casdoor, you can use it as a keyword to set the session in your application, like this:

```go
data, _ := json.Marshal(claims)
c.setSession("user", data)
```

## Step5. Interact with the users

Casdoor-go-sdk support basic user operations, like:

- `GetUser(name string)`, get one user by user name.
- `GetUsers()`, get all users.
- `UpdateUser(casdoorsdk.User)/AddUser(casdoorsdk.User)/DeleteUser(casdoorsdk.User)`, write user to database.
