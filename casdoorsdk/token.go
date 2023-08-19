// Copyright 2023 The Casdoor Authors. All Rights Reserved.
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
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"golang.org/x/oauth2"
)

// Token has the same definition as https://github.com/casdoor/casdoor/blob/master/object/token.go#L45
type Token struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Application  string `xorm:"varchar(100)" json:"application"`
	Organization string `xorm:"varchar(100)" json:"organization"`
	User         string `xorm:"varchar(100)" json:"user"`

	Code          string `xorm:"varchar(100) index" json:"code"`
	AccessToken   string `xorm:"mediumtext" json:"accessToken"`
	RefreshToken  string `xorm:"mediumtext" json:"refreshToken"`
	ExpiresIn     int    `json:"expiresIn"`
	Scope         string `xorm:"varchar(100)" json:"scope"`
	TokenType     string `xorm:"varchar(100)" json:"tokenType"`
	CodeChallenge string `xorm:"varchar(100)" json:"codeChallenge"`
	CodeIsUsed    bool   `json:"codeIsUsed"`
	CodeExpireIn  int64  `json:"codeExpireIn"`
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

func (c *Client) GetTokens(p int, pageSize int) ([]*Token, int, error) {
	queryMap := map[string]string{
		"owner":    c.OrganizationName,
		"p":        strconv.Itoa(p),
		"pageSize": strconv.Itoa(pageSize),
	}

	url := c.GetUrl("get-tokens", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	tokens, ok := response.Data.([]*Token)
	if !ok {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return tokens, int(response.Data2.(float64)), nil
}

func (c *Client) DeleteToken(name string) (bool, error) {
	organization := Organization{
		Owner: "admin",
		Name:  name,
	}
	postBytes, err := json.Marshal(organization)
	if err != nil {
		return false, err
	}

	resp, err := c.DoPost("delete-token", nil, postBytes, false, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}
