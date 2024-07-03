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

type Enforcer struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	UpdatedTime string `xorm:"varchar(100) updated" json:"updatedTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`
	Description string `xorm:"varchar(100)" json:"description"`

	Model     string `xorm:"varchar(100)" json:"model"`
	Adapter   string `xorm:"varchar(100)" json:"adapter"`
	IsEnabled bool   `json:"isEnabled"`

	//*casbin.Enforcer
}

func (c *Client) GetEnforcers() ([]*Enforcer, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-enforcers", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var enforcers []*Enforcer
	err = json.Unmarshal(bytes, &enforcers)
	if err != nil {
		return nil, err
	}
	return enforcers, nil
}

func (c *Client) GetPaginationEnforcers(p int, pageSize int, queryMap map[string]string) ([]*Enforcer, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-enforcers", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var enforcers []*Enforcer
	err = json.Unmarshal(dataBytes, &enforcers)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return enforcers, int(response.Data2.(float64)), nil
}

func (c *Client) GetEnforcer(name string) (*Enforcer, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-enforcer", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var enforcer *Enforcer
	err = json.Unmarshal(bytes, &enforcer)
	if err != nil {
		return nil, err
	}
	return enforcer, nil
}

func (c *Client) UpdateEnforcer(enforcer *Enforcer) (bool, error) {
	_, affected, err := c.modifyEnforcer("update-enforcer", enforcer, nil)
	return affected, err
}

func (c *Client) AddEnforcer(enforcer *Enforcer) (bool, error) {
	_, affected, err := c.modifyEnforcer("add-enforcer", enforcer, nil)
	return affected, err
}

func (c *Client) DeleteEnforcer(enforcer *Enforcer) (bool, error) {
	_, affected, err := c.modifyEnforcer("delete-enforcer", enforcer, nil)
	return affected, err
}
