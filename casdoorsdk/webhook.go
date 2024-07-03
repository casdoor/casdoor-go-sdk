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

// Webhook has the same definition as https://github.com/casdoor/casdoor/blob/master/object/webhook.go#L24
type Webhook struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Organization string `xorm:"varchar(100)" json:"organization"`
	Type         string `xorm:"varchar(100)" json:"type"`

	Host             string         `xorm:"varchar(100)" json:"host"`
	Port             int            `json:"port"`
	User             string         `xorm:"varchar(100)" json:"user"`
	Password         string         `xorm:"varchar(100)" json:"password"`
	DatabaseType     string         `xorm:"varchar(100)" json:"databaseType"`
	Database         string         `xorm:"varchar(100)" json:"database"`
	Table            string         `xorm:"varchar(100)" json:"table"`
	TablePrimaryKey  string         `xorm:"varchar(100)" json:"tablePrimaryKey"`
	TableColumns     []*TableColumn `xorm:"mediumtext" json:"tableColumns"`
	AffiliationTable string         `xorm:"varchar(100)" json:"affiliationTable"`
	AvatarBaseUrl    string         `xorm:"varchar(100)" json:"avatarBaseUrl"`
	ErrorText        string         `xorm:"mediumtext" json:"errorText"`
	SyncInterval     int            `json:"syncInterval"`
	IsReadOnly       bool           `json:"isReadOnly"`
	IsEnabled        bool           `json:"isEnabled"`

	// Ormer *Ormer `xorm:"-" json:"-"`
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
