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

type Permission struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`
	Description string `xorm:"varchar(100)" json:"description"`

	Users   []string `xorm:"mediumtext" json:"users"`
	Groups  []string `xorm:"mediumtext" json:"groups"`
	Roles   []string `xorm:"mediumtext" json:"roles"`
	Domains []string `xorm:"mediumtext" json:"domains"`

	Model        string   `xorm:"varchar(100)" json:"model"`
	Adapter      string   `xorm:"varchar(100)" json:"adapter"`
	ResourceType string   `xorm:"varchar(100)" json:"resourceType"`
	Resources    []string `xorm:"mediumtext" json:"resources"`
	Actions      []string `xorm:"mediumtext" json:"actions"`
	Effect       string   `xorm:"varchar(100)" json:"effect"`
	IsEnabled    bool     `json:"isEnabled"`

	Submitter   string `xorm:"varchar(100)" json:"submitter"`
	Approver    string `xorm:"varchar(100)" json:"approver"`
	ApproveTime string `xorm:"varchar(100)" json:"approveTime"`
	State       string `xorm:"varchar(100)" json:"state"`
}

func (c *Client) GetPermissions() ([]*Permission, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-permissions", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var permissions []*Permission
	err = json.Unmarshal(bytes, &permissions)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (c *Client) GetPermissionsByRole(name string) ([]*Permission, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-permissions-by-role", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var permissions []*Permission
	err = json.Unmarshal(bytes, &permissions)
	if err != nil {
		return nil, err
	}
	return permissions, nil
}

func (c *Client) GetPaginationPermissions(p int, pageSize int, queryMap map[string]string) ([]*Permission, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-permissions", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var permissions []*Permission
	err = json.Unmarshal(dataBytes, &permissions)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return permissions, int(response.Data2.(float64)), nil
}

func (c *Client) GetPermission(name string) (*Permission, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-permission", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var permission *Permission
	err = json.Unmarshal(bytes, &permission)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func (c *Client) UpdatePermission(permission *Permission) (bool, error) {
	_, affected, err := c.modifyPermission("update-permission", permission, nil)
	return affected, err
}

func (c *Client) UpdatePermissionForColumns(permission *Permission, columns []string) (bool, error) {
	_, affected, err := c.modifyPermission("update-permission", permission, columns)
	return affected, err
}

func (c *Client) AddPermission(permission *Permission) (bool, error) {
	_, affected, err := c.modifyPermission("add-permission", permission, nil)
	return affected, err
}

func (c *Client) DeletePermission(permission *Permission) (bool, error) {
	_, affected, err := c.modifyPermission("delete-permission", permission, nil)
	return affected, err
}
