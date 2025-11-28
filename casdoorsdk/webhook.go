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

type Header struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

// Webhook has the same definition as https://github.com/casdoor/casdoor/blob/master/object/webhook.go#L30
type Webhook struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Organization string `xorm:"varchar(100) index" json:"organization"`

	Url            string    `xorm:"varchar(200)" json:"url"`
	Method         string    `xorm:"varchar(100)" json:"method"`
	ContentType    string    `xorm:"varchar(100)" json:"contentType"`
	Headers        []*Header `xorm:"mediumtext" json:"headers"`
	Events         []string  `xorm:"varchar(1000)" json:"events"`
	TokenFields    []string  `xorm:"varchar(1000)" json:"tokenFields"`
	ObjectFields   []string  `xorm:"varchar(1000)" json:"objectFields"`
	IsUserExtended bool      `json:"isUserExtended"`
	SingleOrgOnly  bool      `json:"singleOrgOnly"`
	IsEnabled      bool      `json:"isEnabled"`
}

func (c *Client) GetWebhooks() ([]*Webhook, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-webhooks", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var webhooks []*Webhook
	err = json.Unmarshal(bytes, &webhooks)
	if err != nil {
		return nil, err
	}
	return webhooks, nil
}

func (c *Client) GetPaginationWebhooks(p int, pageSize int, queryMap map[string]string) ([]*Webhook, int, error) {
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

	var webhooks []*Webhook
	err = json.Unmarshal(dataBytes, &webhooks)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return webhooks, int(response.Data2.(float64)), nil
}

func (c *Client) GetWebhook(name string) (*Webhook, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-webhook", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var webhook *Webhook
	err = json.Unmarshal(bytes, &webhook)
	if err != nil {
		return nil, err
	}
	return webhook, nil
}

func (c *Client) AddWebhook(webhook *Webhook) (bool, error) {
	_, affected, err := c.modifyWebhook("add-webhook", webhook, nil)
	return affected, err
}

func (c *Client) UpdateWebhook(webhook *Webhook) (bool, error) {
	_, affected, err := c.modifyWebhook("update-webhook", webhook, nil)
	return affected, err
}

func (c *Client) DeleteWebhook(webhook *Webhook) (bool, error) {
	_, affected, err := c.modifyWebhook("delete-webhook", webhook, nil)
	return affected, err
}
