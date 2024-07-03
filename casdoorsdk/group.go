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

func (c *Client) GetGroups() ([]*Group, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-groups", queryMap)

	bytes, err := c.DoGetBytes(url)
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

func (c *Client) GetPaginationGroups(p int, pageSize int, queryMap map[string]string) ([]*Group, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-groups", queryMap)

	response, err := c.DoGetResponse(url)
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

func (c *Client) GetGroup(name string) (*Group, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-group", queryMap)

	bytes, err := c.DoGetBytes(url)
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

func (c *Client) UpdateGroup(group *Group) (bool, error) {
	_, affected, err := c.modifyGroup("update-group", group, nil)
	return affected, err
}

func (c *Client) AddGroup(group *Group) (bool, error) {
	_, affected, err := c.modifyGroup("add-group", group, nil)
	return affected, err
}

func (c *Client) DeleteGroup(group *Group) (bool, error) {
	_, affected, err := c.modifyGroup("delete-group", group, nil)
	return affected, err
}
