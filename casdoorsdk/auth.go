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
}

// HttpClient interface has the method required to use a type as custom http client.
// The net/*http.Client type satisfies this interface.
type HttpClient interface {
	Do(*http.Request) (*http.Response, error)
}

type Response struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Data2  interface{} `json:"data2"`
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
		*config,
	}
}

// SetHttpClient sets custom http Client.
func SetHttpClient(httpClient HttpClient) {
	client = httpClient
}

// GetOAuthToken gets the pivotal and necessary secret to interact with the Casdoor server
func (c *Client) GetOAuthToken(code string, state string) (*oauth2.Token, error) {
	config := oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   fmt.Sprintf("%s/api/login/oauth/authorize", c.Endpoint),
			TokenURL:  fmt.Sprintf("%s/api/login/oauth/access_token", c.Endpoint),
			AuthStyle: oauth2.AuthStyleInParams,
		},
		// RedirectURL: redirectUri,
		Scopes: nil,
	}

	token, err := config.Exchange(context.Background(), code)
	if err != nil {
		return token, err
	}

	if strings.HasPrefix(token.AccessToken, "error:") {
		return nil, errors.New(strings.TrimPrefix(token.AccessToken, "error: "))
	}

	return token, err
}

// RefreshOAuthToken refreshes the OAuth token
func (c *Client) RefreshOAuthToken(refreshToken string) (*oauth2.Token, error) {
	config := oauth2.Config{
		ClientID:     c.ClientId,
		ClientSecret: c.ClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   fmt.Sprintf("%s/api/login/oauth/authorize", c.Endpoint),
			TokenURL:  fmt.Sprintf("%s/api/login/oauth/refresh_token", c.Endpoint),
			AuthStyle: oauth2.AuthStyleInParams,
		},
		// RedirectURL: redirectUri,
		Scopes: nil,
	}

	token, err := config.TokenSource(context.Background(), &oauth2.Token{RefreshToken: refreshToken}).Token()
	if err != nil {
		return token, err
	}

	if strings.HasPrefix(token.AccessToken, "error:") {
		return nil, errors.New(strings.TrimPrefix(token.AccessToken, "error: "))
	}

	return token, err
}
