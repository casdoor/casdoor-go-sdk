// Copyright 2024 The Casdoor Authors. All Rights Reserved.
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

// Transaction has the same definition as https://github.com/casdoor/casdoor/blob/master/object/transaction.go#L24
type Transaction struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`
	// Transaction Provider Info
	Provider string `xorm:"varchar(100)" json:"provider"`
	Category string `xorm:"varchar(100)" json:"category"`
	Type     string `xorm:"varchar(100)" json:"type"`
	// Product Info
	ProductName        string  `xorm:"varchar(100)" json:"productName"`
	ProductDisplayName string  `xorm:"varchar(100)" json:"productDisplayName"`
	Detail             string  `xorm:"varchar(255)" json:"detail"`
	Tag                string  `xorm:"varchar(100)" json:"tag"`
	Currency           string  `xorm:"varchar(100)" json:"currency"`
	Amount             float64 `json:"amount"`
	ReturnUrl          string  `xorm:"varchar(1000)" json:"returnUrl"`
	// User Info
	User        string `xorm:"varchar(100)" json:"user"`
	Application string `xorm:"varchar(100)" json:"application"`
	Payment     string `xorm:"varchar(100)" json:"payment"`

	State string `xorm:"varchar(100)" json:"state"`
}

func (c *Client) GetTransactions() ([]*Transaction, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-transactions", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var transactions []*Transaction
	err = json.Unmarshal(bytes, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (c *Client) GetPaginationTransactions(p int, pageSize int, queryMap map[string]string) ([]*Transaction, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-transactions", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	transactions, ok := response.Data.([]*Transaction)
	if !ok {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return transactions, int(response.Data2.(float64)), nil
}

func (c *Client) GetTransaction(name string) (*Transaction, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-transaction", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var transaction *Transaction
	err = json.Unmarshal(bytes, &transaction)
	if err != nil {
		return nil, err
	}
	return transaction, nil
}

func (c *Client) GetUserTransactions(userName string) ([]*Transaction, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
		"user":  userName,
	}

	url := c.GetUrl("get-user-transactions", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var transactions []*Transaction
	err = json.Unmarshal(bytes, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (c *Client) UpdateTransaction(transaction *Transaction) (bool, error) {
	_, affected, err := c.modifyTransaction("update-transaction", transaction, nil)
	return affected, err
}

func (c *Client) AddTransaction(transaction *Transaction) (bool, error) {
	_, affected, err := c.modifyTransaction("add-transaction", transaction, nil)
	return affected, err
}

func (c *Client) DeleteTransaction(transaction *Transaction) (bool, error) {
	_, affected, err := c.modifyTransaction("delete-transaction", transaction, nil)
	return affected, err
}
