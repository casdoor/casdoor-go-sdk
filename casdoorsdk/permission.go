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

// Deprecated: Use GetPermissionsWithContext.
func (c *Client) GetPermissions() ([]*Permission, error) {
	return c.GetPermissionsWithContext(context.Background())
}

func (c *Client) GetPermissionsWithContext(ctx context.Context) ([]*Permission, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-permissions", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
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

// Deprecated: Use GetPermissionsByRoleWithContext.
func (c *Client) GetPermissionsByRole(name string) ([]*Permission, error) {
	return c.GetPermissionsByRoleWithContext(context.Background(), name)
}

func (c *Client) GetPermissionsByRoleWithContext(ctx context.Context, name string) ([]*Permission, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-permissions-by-role", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
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

// Deprecated: Use GetPaginationPermissionsWithContext.
func (c *Client) GetPaginationPermissions(p int, pageSize int, queryMap map[string]string) ([]*Permission, int, error) {
	return c.GetPaginationPermissionsWithContext(context.Background(), p, pageSize, queryMap)
}

func (c *Client) GetPaginationPermissionsWithContext(ctx context.Context, p int, pageSize int, queryMap map[string]string) ([]*Permission, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-permissions", queryMap)

	response, err := c.DoGetResponseWithContext(ctx, url)
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

// Deprecated: Use GetPermissionWithContext.
func (c *Client) GetPermission(name string) (*Permission, error) {
	return c.GetPermissionWithContext(context.Background(), name)
}

func (c *Client) GetPermissionWithContext(ctx context.Context, name string) (*Permission, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-permission", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
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

// Deprecated: Use UpdatePermissionWithContext.
func (c *Client) UpdatePermission(permission *Permission) (bool, error) {
	_, affected, err := c.modifyPermission("update-permission", permission, nil)
	return affected, err
}

// Deprecated: Use UpdatePermissionForColumnsWithContext.
func (c *Client) UpdatePermissionForColumns(permission *Permission, columns []string) (bool, error) {
	_, affected, err := c.modifyPermission("update-permission", permission, columns)
	return affected, err
}

// Deprecated: Use AddPermissionWithContext.
func (c *Client) AddPermission(permission *Permission) (bool, error) {
	_, affected, err := c.modifyPermission("add-permission", permission, nil)
	return affected, err
}

// Deprecated: Use DeletePermissionWithContext.
func (c *Client) DeletePermission(permission *Permission) (bool, error) {
	_, affected, err := c.modifyPermission("delete-permission", permission, nil)
	return affected, err
}

func (c *Client) UpdatePermissionWithContext(ctx context.Context, permission *Permission) (bool, error) {
	_, affected, err := c.modifyPermissionWithContext(ctx, "update-permission", permission, nil)
	return affected, err
}

func (c *Client) UpdatePermissionForColumnsWithContext(ctx context.Context, permission *Permission, columns []string) (bool, error) {
	_, affected, err := c.modifyPermissionWithContext(ctx, "update-permission", permission, columns)
	return affected, err
}

func (c *Client) AddPermissionWithContext(ctx context.Context, permission *Permission) (bool, error) {
	_, affected, err := c.modifyPermissionWithContext(ctx, "add-permission", permission, nil)
	return affected, err
}

func (c *Client) DeletePermissionWithContext(ctx context.Context, permission *Permission) (bool, error) {
	_, affected, err := c.modifyPermissionWithContext(ctx, "delete-permission", permission, nil)
	return affected, err
}
