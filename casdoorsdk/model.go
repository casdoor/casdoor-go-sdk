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

type Model struct {
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	CreatedTime string `json:"createdTime"`
	DisplayName string `json:"displayName"`
	Description string `json:"description"`
	ModelText   string `json:"modelText"`
	// Deprecated: removed from server
	UpdatedTime string `json:"updatedTime"`
	// Deprecated: removed from server
	Manager string `json:"manager"`
	// Deprecated: removed from server
	ContactEmail string `json:"contactEmail"`
	// Deprecated: removed from server
	Type string `json:"type"`
	// Deprecated: removed from server
	ParentId string `json:"parentId"`
	// Deprecated: removed from server
	IsTopModel bool `json:"isTopModel"`
	// Deprecated: removed from server
	Users []*User `json:"users"`
	// Deprecated: removed from server
	Title string `json:"title,omitempty"`
	// Deprecated: removed from server
	Key string `json:"key,omitempty"`
	// Deprecated: removed from server
	Children []*Model `json:"children,omitempty"`
	// Deprecated: removed from server
	IsEnabled bool `json:"isEnabled"`
}

func (c *Client) GetModels() ([]*Model, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-models", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var models []*Model
	err = json.Unmarshal(bytes, &models)
	if err != nil {
		return nil, err
	}
	return models, nil
}

func (c *Client) GetPaginationModels(p int, pageSize int, queryMap map[string]string) ([]*Model, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-models", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var models []*Model
	err = json.Unmarshal(dataBytes, &models)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return models, int(response.Data2.(float64)), nil
}

func (c *Client) GetModel(name string) (*Model, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-model", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var model *Model
	err = json.Unmarshal(bytes, &model)
	if err != nil {
		return nil, err
	}
	return model, nil
}

func (c *Client) UpdateModel(model *Model) (bool, error) {
	_, affected, err := c.modifyModel("update-model", model, nil)
	return affected, err
}

func (c *Client) AddModel(model *Model) (bool, error) {
	_, affected, err := c.modifyModel("add-model", model, nil)
	return affected, err
}

func (c *Client) DeleteModel(model *Model) (bool, error) {
	_, affected, err := c.modifyModel("delete-model", model, nil)
	return affected, err
}
