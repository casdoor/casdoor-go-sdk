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
// See the License for the specific language governing records and
// limitations under the License.

package casdoorsdk

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

var (
	CasdoorApplication  = "app-built-in"
	CasdoorOrganization = "built-in"
)

type Session struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	Application string `xorm:"varchar(100) notnull pk" json:"application"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	SessionId []string `json:"sessionId"`
}

func (c *Client) GetSessions() ([]*Session, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-sessions", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var sessions []*Session
	err = json.Unmarshal(bytes, &sessions)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func (c *Client) GetPaginationSessions(p int, pageSize int, queryMap map[string]string) ([]*Session, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-sessions", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var sessions []*Session
	err = json.Unmarshal(dataBytes, &sessions)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}
	return sessions, int(response.Data2.(float64)), nil
}

func (c *Client) GetSession(name string, application string) (*Session, error) {
	queryMap := map[string]string{
		"sessionPkId": fmt.Sprintf("%s/%s/%s", c.OrganizationName, name, application),
	}

	url := c.GetUrl("get-session", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var session *Session
	err = json.Unmarshal(bytes, &session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func (c *Client) UpdateSession(session *Session) (bool, error) {
	_, affected, err := c.modifySession("update-session", session, nil)
	return affected, err
}

func (c *Client) UpdateSessionForColumns(session *Session, columns []string) (bool, error) {
	_, affected, err := c.modifySession("update-session", session, columns)
	return affected, err
}

func (c *Client) AddSession(session *Session) (bool, error) {
	_, affected, err := c.modifySession("add-session", session, nil)
	return affected, err
}

func (c *Client) DeleteSession(session *Session) (bool, error) {
	_, affected, err := c.modifySession("delete-session", session, nil)
	return affected, err
}
