# casdoor-go-sdk

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
  <a href="https://gitter.im/casbin/casdoor">
    <img alt="Gitter" src="https://badges.gitter.im/casbin/casdoor.svg?style=flat-square">
  </a>
</p>

This is Casdoor's SDK for golang, which will allow you to easily connect your application to the Casdoor authentication system without having to implement it from scratch.

Casdoor SDK is very simple to use. We will show you the steps below.

> Noted that this sdk has been applied to casnode, if you still don’t know how to use it after reading README.md, you can refer to it

## Step1. Install and Import

First in your go project, just need to run:

```bash
go get github.com/casdoor/casdoor-go-sdk@latest
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
