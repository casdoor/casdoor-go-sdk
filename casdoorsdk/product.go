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

type Product struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`

	Image       string   `xorm:"varchar(100)" json:"image"`
	Detail      string   `xorm:"varchar(255)" json:"detail"`
	Description string   `xorm:"varchar(100)" json:"description"`
	Tag         string   `xorm:"varchar(100)" json:"tag"`
	Currency    string   `xorm:"varchar(100)" json:"currency"`
	Price       float64  `json:"price"`
	Quantity    int      `json:"quantity"`
	Sold        int      `json:"sold"`
	Providers   []string `xorm:"varchar(100)" json:"providers"`
	ReturnUrl   string   `xorm:"varchar(1000)" json:"returnUrl"`

	State string `xorm:"varchar(100)" json:"state"`

	ProviderObjs []*Provider `xorm:"-" json:"providerObjs"`
}

func (c *Client) GetProducts() ([]*Product, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-products", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var products []*Product
	err = json.Unmarshal(bytes, &products)
	if err != nil {
		return nil, err
	}
	return products, nil
}

func (c *Client) GetPaginationProducts(p int, pageSize int, queryMap map[string]string) ([]*Product, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-products", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var products []*Product
	err = json.Unmarshal(dataBytes, &products)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return products, int(response.Data2.(float64)), nil
}

func (c *Client) GetProduct(name string) (*Product, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-product", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var product *Product
	err = json.Unmarshal(bytes, &product)
	if err != nil {
		return nil, err
	}
	return product, nil
}

func (c *Client) UpdateProduct(product *Product) (bool, error) {
	_, affected, err := c.modifyProduct("update-product", product, nil)
	return affected, err
}

func (c *Client) AddProduct(product *Product) (bool, error) {
	_, affected, err := c.modifyProduct("add-product", product, nil)
	return affected, err
}

func (c *Client) DeleteProduct(product *Product) (bool, error) {
	_, affected, err := c.modifyProduct("delete-product", product, nil)
	return affected, err
}

func (c *Client) BuyProduct(name string, providerName string) (*Product, error) {
	queryMap := map[string]string{
		"id":           fmt.Sprintf("%s/%s", c.OrganizationName, name),
		"providerName": providerName,
	}

	url := c.GetUrl("buy-product", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var product *Product
	err = json.Unmarshal(bytes, &product)
	if err != nil {
		return nil, err
	}
	return product, nil
}
