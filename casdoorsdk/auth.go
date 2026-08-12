// Copyright 2021 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package casdoorsdk

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/oauth2"
)

// AuthConfig is the core configuration.
// The first step to use this SDK is to use the InitConfig function to initialize the global authConfig.
type AuthConfig struct {
	Endpoint         string
	ClientId         string
	ClientSecret     string
	Certificate      string
	OrganizationName string
	ApplicationName  string
}

type Client struct {
	AuthConfig
	CustomHeaders map[string]string
	// AccessToken is the user's access token. If it's not empty, all the API requests
	// sent by this client will be authenticated as the user who owns the access token
	// (via the "Authorization: Bearer" header), instead of as the application itself
	// (via the client ID and client secret's Basic Auth).
	// Use WithAccessToken() to get such a client.
	AccessToken string
}

// HttpClient interface has the method required to use a type as custom http client.
// The net/*http.Client type satisfies this interface.
type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Response struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Sub    string      `json:"sub"`
	Name   string      `json:"name"`
	Data   interface{} `json:"data"`
	Data2  interface{} `json:"data2"`
	Data3  interface{} `json:"data3"`
}

// client is a shared http Client.
var client HttpClient = &http.Client{}
var globalClient *Client

func InitConfig(endpoint string, clientId string, clientSecret string, certificate string, organizationName string, applicationName string) {
	globalClient = NewClient(endpoint, clientId, clientSecret, certificate, organizationName, applicationName)
}

func NewClient(endpoint string, clientId string, clientSecret string, certificate string, organizationName string, applicationName string) *Client {
	return NewClientWithConf(
		&AuthConfig{
			Endpoint:         endpoint,
			ClientId:         clientId,
			ClientSecret:     clientSecret,
			Certificate:      certificate,
			OrganizationName: organizationName,
			ApplicationName:  applicationName,
		})
}

func NewClientWithConf(config *AuthConfig) *Client {
	return &Client{
		AuthConfig:    *config,
		CustomHeaders: make(map[string]string),
	}
}

// WithAccessToken returns a copy of the client that calls all the Casdoor APIs as the user
// who owns the given access token, rather than as the application. The access token is the
// user's OAuth access token returned by GetOAuthToken(), RefreshOAuthToken(),
// GetOAuthTokenByPassword() or ImpersonateUser().
// The returned client is independent from the original one, so the original client keeps
// using the application's client ID and client secret, and it's safe to create one client
// per user request:
//
//	token, err := client.GetOAuthToken(code, state)
//	user, err := client.WithAccessToken(token.AccessToken).GetUserByAccessToken()
//
// Note that the APIs are still subject to Casdoor's permission check, so a non-admin user
// can only access their own data.
func (c *Client) WithAccessToken(accessToken string) *Client {
	customHeaders := make(map[string]string, len(c.CustomHeaders))
	for key, value := range c.CustomHeaders {
		customHeaders[key] = value
	}

	return &Client{
		AuthConfig:    c.AuthConfig,
		CustomHeaders: customHeaders,
		AccessToken:   accessToken,
	}
}

// SetHttpClient sets custom http Client.
func SetHttpClient(httpClient HttpClient) {
	client = httpClient
}

// OAuthOption is a function type for configuring OAuth requests.
type OAuthOption func(*oauthOptions)

// oauthOptions holds configuration options for OAuth operations.
type oauthOptions struct {
	httpClient *http.Client
}

// WithHTTPClient sets a custom http client for oauth operations.
func WithHTTPClient(httpClient *http.Client) OAuthOption {
	return func(opts *oauthOptions) {
		opts.httpClient = httpClient
	}
}

// getOAuthConfig returns the OAuth config, tokenAction is the action of the token API,
// like "access_token" or "refresh_token"
func (c *Client) getOAuthConfig(tokenAction string) oauth2.Config {
	return oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   fmt.Sprintf("%s/api/login/oauth/authorize", c.Endpoint),
			TokenURL:  fmt.Sprintf("%s/api/login/oauth/%s", c.Endpoint, tokenAction),
			AuthStyle: oauth2.AuthStyleInParams,
		},
		// RedirectURL: redirectUri,
		Scopes: nil,
	}
}

func getOAuthContext(opts ...OAuthOption) context.Context {
	options := &oauthOptions{}
	for _, opt := range opts {
		opt(options)
	}

	ctx := context.Background()
	if options.httpClient != nil {
		ctx = context.WithValue(ctx, oauth2.HTTPClient, options.httpClient)
	}

	return ctx
}

// checkOAuthToken converts the "error: xxx" access token returned by the Casdoor server into a real error
func checkOAuthToken(token *oauth2.Token, err error) (*oauth2.Token, error) {
	if err != nil {
		return token, err
	}

	if strings.HasPrefix(token.AccessToken, "error:") {
		return nil, errors.New(strings.TrimPrefix(token.AccessToken, "error: "))
	}

	return token, nil
}

// GetOAuthToken gets the pivotal and necessary secret to interact with the Casdoor server
func (c *Client) GetOAuthToken(code string, state string, opts ...OAuthOption) (*oauth2.Token, error) {
	config := c.getOAuthConfig("access_token")

	token, err := config.Exchange(getOAuthContext(opts...), code)
	return checkOAuthToken(token, err)
}

// RefreshOAuthToken refreshes the OAuth token
func (c *Client) RefreshOAuthToken(refreshToken string, opts ...OAuthOption) (*oauth2.Token, error) {
	config := c.getOAuthConfig("refresh_token")

	token, err := config.TokenSource(getOAuthContext(opts...), &oauth2.Token{RefreshToken: refreshToken}).Token()
	return checkOAuthToken(token, err)
}

// GetOAuthTokenByPassword gets the OAuth token via the "password" grant type, i.e., the
// "Resource Owner Password Credentials Grant" of OAuth 2.0. The "password" grant type
// must be enabled in the application's "Grant types" in Casdoor.
// The username is the user's name inside the application's organization, like "alice"
// instead of "my-org/alice".
func (c *Client) GetOAuthTokenByPassword(username string, password string, opts ...OAuthOption) (*oauth2.Token, error) {
	config := c.getOAuthConfig("access_token")

	token, err := config.PasswordCredentialsToken(getOAuthContext(opts...), username, password)
	return checkOAuthToken(token, err)
}

// ImpersonateUser gets an OAuth token which acts as the given user, so that an admin can
// call the APIs on behalf of the user, without knowing the user's own password. It's the
// SDK equivalent of the "Impersonation" button in Casdoor's Web UI.
// The masterPassword is the "Master password" of the user's organization, it needs to be
// set in Casdoor first: "Organizations" -> Edit the organization -> "Master password".
// See: https://casdoor.org/docs/user/impersonation
func (c *Client) ImpersonateUser(username string, masterPassword string, opts ...OAuthOption) (*oauth2.Token, error) {
	return c.GetOAuthTokenByPassword(username, masterPassword, opts...)
}
