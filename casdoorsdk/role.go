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
	"fmt"
	"strconv"
)

// Role has the same definition as https://github.com/casdoor/casdoor/blob/master/object/role.go#L24
type Role struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`
	Description string `xorm:"varchar(100)" json:"description"`

	Users     []string `xorm:"mediumtext" json:"users"`
	Groups    []string `xorm:"mediumtext" json:"groups"`
	Roles     []string `xorm:"mediumtext" json:"roles"`
	Domains   []string `xorm:"mediumtext" json:"domains"`
	IsEnabled bool     `json:"isEnabled"`
}

func (c *Client) GetRoles() ([]*Role, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-roles", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var roles []*Role
	err = json.Unmarshal(bytes, &roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

func (c *Client) GetPaginationRoles(p int, pageSize int, queryMap map[string]string) ([]*Role, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-roles", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var roles []*Role
	err = json.Unmarshal(dataBytes, &roles)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return roles, int(response.Data2.(float64)), nil
}

func (c *Client) GetRole(name string) (*Role, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-role", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var role *Role
	err = json.Unmarshal(bytes, &role)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (c *Client) UpdateRole(role *Role) (bool, error) {
	_, affected, err := c.modifyRole("update-role", role, nil)
	return affected, err
}

func (c *Client) UpdateRoleForColumns(role *Role, columns []string) (bool, error) {
	_, affected, err := c.modifyRole("update-role", role, columns)
	return affected, err
}

func (c *Client) AddRole(role *Role) (bool, error) {
	_, affected, err := c.modifyRole("add-role", role, nil)
	return affected, err
}

func (c *Client) DeleteRole(role *Role) (bool, error) {
	_, affected, err := c.modifyRole("delete-role", role, nil)
	return affected, err
}
