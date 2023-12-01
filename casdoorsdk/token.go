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
	"encoding/json"
	"errors"
	"strconv"
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

func (c *Client) DeleteToken(token *Token) (bool, error) {
	token.Owner = "admin"
	postBytes, err := json.Marshal(token)
	if err != nil {
		return false, err
	}

	resp, err := c.DoPost("delete-token", nil, postBytes, false, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}
