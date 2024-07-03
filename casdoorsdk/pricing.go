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

// Pricing has the same definition as https://github.com/casdoor/casdoor/blob/master/object/pricing.go#L24
type Pricing struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`
	Description string `xorm:"varchar(100)" json:"description"`

	Plans         []string `xorm:"mediumtext" json:"plans"`
	IsEnabled     bool     `json:"isEnabled"`
	TrialDuration int      `json:"trialDuration"`
	Application   string   `xorm:"varchar(100)" json:"application"`

	Submitter   string `xorm:"varchar(100)" json:"submitter"`
	Approver    string `xorm:"varchar(100)" json:"approver"`
	ApproveTime string `xorm:"varchar(100)" json:"approveTime"`

	State string `xorm:"varchar(100)" json:"state"`
}

func (c *Client) GetPricings() ([]*Pricing, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-pricings", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var pricings []*Pricing
	err = json.Unmarshal(bytes, &pricings)
	if err != nil {
		return nil, err
	}
	return pricings, nil
}

func (c *Client) GetPaginationPricings(p int, pageSize int, queryMap map[string]string) ([]*Pricing, int, error) {
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

	var pricings []*Pricing
	err = json.Unmarshal(dataBytes, &pricings)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return pricings, int(response.Data2.(float64)), nil
}

func (c *Client) GetPricing(name string) (*Pricing, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-pricing", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var pricing *Pricing
	err = json.Unmarshal(bytes, &pricing)
	if err != nil {
		return nil, err
	}
	return pricing, nil
}

func (c *Client) AddPricing(pricing *Pricing) (bool, error) {
	_, affected, err := c.modifyPricing("add-pricing", pricing, nil)
	return affected, err
}

func (c *Client) UpdatePricing(pricing *Pricing) (bool, error) {
	_, affected, err := c.modifyPricing("update-pricing", pricing, nil)
	return affected, err
}

func (c *Client) DeletePricing(pricing *Pricing) (bool, error) {
	_, affected, err := c.modifyPricing("delete-pricing", pricing, nil)
	return affected, err
}
