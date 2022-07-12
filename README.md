# casdoor-go-sdk

This is Casdoor's SDK for golang, which will allow you to easily connect your application to the Casdoor authentication system without having to implement it from scratch.

Casdoor SDK is very simple to use. We will show you the steps below.

> Noted that this sdk has been applied to casnode, if you still don’t know how to use it after reading README.md, you can refer to it

## Step1. Init Config

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

## Step2. Get token and parse

After casdoor verification passed, it will be redirected to your application with code and state, like `https://forum.casbin.com?code=xxx&state=yyyy`.

Your web application can get the `code`,`state` and call `GetOAuthToken(code, state)`, then parse out jwt token.

The general process is as follows:

```go
token, err := auth.GetOAuthToken(code, state)
if err != nil {
	panic(err)
}

claims, err := auth.ParseJwtToken(token.AccessToken)
if err != nil {
	panic(err)
}

claims.AccessToken = token.AccessToken
```

## Step3. Set Session in your app

`auth.Claims` contains the basic information about the user provided by casdoor, you can use it as a keyword to set the session in your application, like this:

```go
data, _ := json.Marshal(claims)
c.setSession("user", data)
```

## Step4. Interact with the users

Casdoor-go-sdk support basic user operations, like:

- `GetUser(name string)`, get one user by user name.
- `GetUsers()`, get all users.
- `UpdateUser(auth.User)/AddUser(auth.User)/DeleteUser(auth.User)`, write user to database.
