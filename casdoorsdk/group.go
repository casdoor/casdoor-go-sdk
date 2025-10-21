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
)

type Group struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk unique index" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	UpdatedTime string `xorm:"varchar(100)" json:"updatedTime"`

	DisplayName  string   `xorm:"varchar(100)" json:"displayName"`
	Manager      string   `xorm:"varchar(100)" json:"manager"`
	ContactEmail string   `xorm:"varchar(100)" json:"contactEmail"`
	Type         string   `xorm:"varchar(100)" json:"type"`
	ParentId     string   `xorm:"varchar(100)" json:"parentId"`
	IsTopGroup   bool     `xorm:"bool" json:"isTopGroup"`
	Users        []string `xorm:"mediumtext" json:"users"`

	Title    string   `json:"title,omitempty"`
	Key      string   `json:"key,omitempty"`
	Children []*Group `json:"children,omitempty"`

	IsEnabled bool `json:"isEnabled"`
}

// Deprecated: Use GetGroupsWithContext.
func (c *Client) GetGroups() ([]*Group, error) {
	return c.GetGroupsWithContext(context.Background())
}

func (c *Client) GetGroupsWithContext(ctx context.Context) ([]*Group, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-groups", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	var groups []*Group
	err = json.Unmarshal(bytes, &groups)
	if err != nil {
		return nil, err
	}
	return groups, nil
}

// Deprecated: Use GetPaginationGroupsWithContext.
func (c *Client) GetPaginationGroups(p int, pageSize int, queryMap map[string]string) ([]*Group, int, error) {
	return c.GetPaginationGroupsWithContext(context.Background(), p, pageSize, queryMap)
}

func (c *Client) GetPaginationGroupsWithContext(ctx context.Context, p int, pageSize int, queryMap map[string]string) ([]*Group, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-groups", queryMap)

	response, err := c.DoGetResponseWithContext(ctx, url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var groups []*Group
	err = json.Unmarshal(dataBytes, &groups)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return groups, int(response.Data2.(float64)), nil
}

// Deprecated: Use GetGroupWithContext.
func (c *Client) GetGroup(name string) (*Group, error) {
	return c.GetGroupWithContext(context.Background(), name)
}

func (c *Client) GetGroupWithContext(ctx context.Context, name string) (*Group, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-group", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	var group *Group
	err = json.Unmarshal(bytes, &group)
	if err != nil {
		return nil, err
	}
	return group, nil
}

// Deprecated: Use UpdateGroupWithContext.
func (c *Client) UpdateGroup(group *Group) (bool, error) {
	_, affected, err := c.modifyGroup("update-group", group, nil)
	return affected, err
}

// Deprecated: Use AddGroupWithContext.
func (c *Client) AddGroup(group *Group) (bool, error) {
	_, affected, err := c.modifyGroup("add-group", group, nil)
	return affected, err
}

// Deprecated: Use DeleteGroupWithContext.
func (c *Client) DeleteGroup(group *Group) (bool, error) {
	_, affected, err := c.modifyGroup("delete-group", group, nil)
	return affected, err
}

func (c *Client) UpdateGroupWithContext(ctx context.Context, group *Group) (bool, error) {
	_, affected, err := c.modifyGroupWithContext(ctx, "update-group", group, nil)
	return affected, err
}

func (c *Client) AddGroupWithContext(ctx context.Context, group *Group) (bool, error) {
	_, affected, err := c.modifyGroupWithContext(ctx, "add-group", group, nil)
	return affected, err
}

func (c *Client) DeleteGroupWithContext(ctx context.Context, group *Group) (bool, error) {
	_, affected, err := c.modifyGroupWithContext(ctx, "delete-group", group, nil)
	return affected, err
}
