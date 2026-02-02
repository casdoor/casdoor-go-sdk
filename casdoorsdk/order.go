// Copyright 2025 The Casdoor Authors. All Rights Reserved.
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

type Order struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	UpdateTime  string `xorm:"varchar(100)" json:"updateTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`

	// Product Info
	Products     []string      `xorm:"varchar(1000)" json:"products"` // Support for multiple products per order. Using varchar(1000) for simple JSON array storage; can be refactored to separate table if needed
	ProductInfos []ProductInfo `xorm:"mediumtext" json:"productInfos"`

	// User Info
	User string `xorm:"varchar(100)" json:"user"`

	// Payment Info
	Payment  string  `xorm:"varchar(100)" json:"payment"`
	Price    float64 `json:"price"`
	Currency string  `xorm:"varchar(100)" json:"currency"`

	// Order State
	State   string `xorm:"varchar(100)" json:"state"`
	Message string `xorm:"varchar(2000)" json:"message"`
}

type ProductInfo struct {
	Owner       string  `json:"owner"`
	Name        string  `json:"name"`
	DisplayName string  `json:"displayName"`
	Image       string  `json:"image,omitempty"`
	Detail      string  `json:"detail,omitempty"`
	Price       float64 `json:"price"`
	Currency    string  `json:"currency,omitempty"`
	IsRecharge  bool    `json:"isRecharge,omitempty"`
	Quantity    int     `json:"quantity,omitempty"`
	PricingName string  `json:"pricingName,omitempty"`
	PlanName    string  `json:"planName,omitempty"`
}

func (c *Client) GetOrders() ([]*Order, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-orders", queryMap)
	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var orders []*Order
	err = json.Unmarshal(bytes, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (c *Client) GetPaginationOrders(p int, pageSize int, queryMap map[string]string) ([]*Order, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-orders", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var orders []*Order
	err = json.Unmarshal(dataBytes, &orders)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return orders, int(response.Data2.(float64)), nil
}

func (c *Client) GetUserOrders(userName string) ([]*Order, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
		"user":  userName,
	}

	url := c.GetUrl("get-user-orders", queryMap)
	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var orders []*Order
	err = json.Unmarshal(bytes, &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}

func (c *Client) GetOrder(name string) (*Order, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-order", queryMap)
	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var order *Order
	err = json.Unmarshal(bytes, &order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (c *Client) UpdateOrder(order *Order) (bool, error) {
	_, affected, err := c.modifyOrder("update-order", order, nil)
	return affected, err
}

func (c *Client) AddOrder(order *Order) (bool, error) {
	_, affected, err := c.modifyOrder("add-order", order, nil)
	return affected, err
}

func (c *Client) DeleteOrder(order *Order) (bool, error) {
	_, affected, err := c.modifyOrder("delete-order", order, nil)
	return affected, err
}

func (c *Client) CancelOrder(name string) (bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	resp, err := c.DoPost("cancel-order", queryMap, []byte(""), false, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}

func (order *Order) GetId() string {
	return fmt.Sprintf("%s/%s", order.Owner, order.Name)
}
