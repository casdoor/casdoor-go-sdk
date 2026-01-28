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

type Adapter struct {
	Owner        string `json:"owner"`
	Name         string `json:"name"`
	CreatedTime  string `json:"createdTime"`
	Table        string `json:"table"`
	UseSameDb    bool   `json:"useSameDb"`
	Type         string `json:"type"`
	DatabaseType string `json:"databaseType"`
	Host         string `json:"host"`
	Port         int    `json:"port"`
	User         string `json:"user"`
	Password     string `json:"password"`
	Database     string `json:"database"`
	// Deprecated: removed from server
	TableNamePrefix string `json:"tableNamePrefix"`
	// Deprecated: removed from server
	IsEnabled bool `json:"isEnabled"`
}

func (c *Client) GetAdapters() ([]*Adapter, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-adapters", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var adapters []*Adapter
	err = json.Unmarshal(bytes, &adapters)
	if err != nil {
		return nil, err
	}
	return adapters, nil
}

func (c *Client) GetPaginationAdapters(p int, pageSize int, queryMap map[string]string) ([]*Adapter, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-adapters", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var adapters []*Adapter
	err = json.Unmarshal(dataBytes, &adapters)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return adapters, int(response.Data2.(float64)), nil
}

func (c *Client) GetAdapter(name string) (*Adapter, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-adapter", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var adapter *Adapter
	err = json.Unmarshal(bytes, &adapter)
	if err != nil {
		return nil, err
	}
	return adapter, nil
}

func (c *Client) UpdateAdapter(adapter *Adapter) (bool, error) {
	_, affected, err := c.modifyAdapter("update-adapter", adapter, nil)
	return affected, err
}

func (c *Client) AddAdapter(adapter *Adapter) (bool, error) {
	_, affected, err := c.modifyAdapter("add-adapter", adapter, nil)
	return affected, err
}

func (c *Client) DeleteAdapter(adapter *Adapter) (bool, error) {
	_, affected, err := c.modifyAdapter("delete-adapter", adapter, nil)
	return affected, err
}
