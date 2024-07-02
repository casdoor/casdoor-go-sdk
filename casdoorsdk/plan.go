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

// Plan has the same definition as https://github.com/casdoor/casdoor/blob/master/object/plan.go#L24
type Plan struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`
	Description string `xorm:"varchar(100)" json:"description"`

	PricePerMonth float64 `json:"pricePerMonth"`
	PricePerYear  float64 `json:"pricePerYear"`
	Currency      string  `xorm:"varchar(100)" json:"currency"`
	IsEnabled     bool    `json:"isEnabled"`

	Role    string   `xorm:"varchar(100)" json:"role"`
	Options []string `xorm:"-" json:"options"`
}

func (c *Client) GetPlans() ([]*Plan, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-plans", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var plans []*Plan
	err = json.Unmarshal(bytes, &plans)
	if err != nil {
		return nil, err
	}
	return plans, nil
}

func (c *Client) GetPaginationPlans(p int, pageSize int, queryMap map[string]string) ([]*Plan, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-payments", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var plans []*Plan
	err = json.Unmarshal(dataBytes, &plans)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return plans, int(response.Data2.(float64)), nil
}

func (c *Client) GetPlan(name string) (*Plan, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-plan", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var plan *Plan
	err = json.Unmarshal(bytes, &plan)
	if err != nil {
		return nil, err
	}
	return plan, nil
}

func (c *Client) AddPlan(plan *Plan) (bool, error) {
	_, affected, err := c.modifyPlan("add-plan", plan, nil)
	return affected, err
}

func (c *Client) UpdatePlan(plan *Plan) (bool, error) {
	_, affected, err := c.modifyPlan("update-plan", plan, nil)
	return affected, err
}

func (c *Client) DeletePlan(plan *Plan) (bool, error) {
	_, affected, err := c.modifyPlan("delete-plan", plan, nil)
	return affected, err
}
