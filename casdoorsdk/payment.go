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

type Payment struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`
	// Payment Provider Info
	Provider string `xorm:"varchar(100)" json:"provider"`
	Type     string `xorm:"varchar(100)" json:"type"`
	// Product Info
	ProductName        string  `xorm:"varchar(100)" json:"productName"`
	ProductDisplayName string  `xorm:"varchar(100)" json:"productDisplayName"`
	Detail             string  `xorm:"varchar(255)" json:"detail"`
	Tag                string  `xorm:"varchar(100)" json:"tag"`
	Currency           string  `xorm:"varchar(100)" json:"currency"`
	Price              float64 `json:"price"`
	ReturnUrl          string  `xorm:"varchar(1000)" json:"returnUrl"`
	// Payer Info
	User         string `xorm:"varchar(100)" json:"user"`
	PersonName   string `xorm:"varchar(100)" json:"personName"`
	PersonIdCard string `xorm:"varchar(100)" json:"personIdCard"`
	PersonEmail  string `xorm:"varchar(100)" json:"personEmail"`
	PersonPhone  string `xorm:"varchar(100)" json:"personPhone"`
	// Invoice Info
	InvoiceType   string `xorm:"varchar(100)" json:"invoiceType"`
	InvoiceTitle  string `xorm:"varchar(100)" json:"invoiceTitle"`
	InvoiceTaxId  string `xorm:"varchar(100)" json:"invoiceTaxId"`
	InvoiceRemark string `xorm:"varchar(100)" json:"invoiceRemark"`
	InvoiceUrl    string `xorm:"varchar(255)" json:"invoiceUrl"`
	// Order Info
	OutOrderId string `xorm:"varchar(100)" json:"outOrderId"`
	PayUrl     string `xorm:"varchar(2000)" json:"payUrl"`
	// State      pp.PaymentState `xorm:"varchaFr(100)" json:"state"`
	State   string `xorm:"varchar(100)" json:"state"`
	Message string `xorm:"varchar(2000)" json:"message"`
}

func (c *Client) GetPayments() ([]*Payment, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-payments", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var payments []*Payment
	err = json.Unmarshal(bytes, &payments)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (c *Client) GetPaginationPayments(p int, pageSize int, queryMap map[string]string) ([]*Payment, int, error) {
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

	var payments []*Payment
	err = json.Unmarshal(dataBytes, &payments)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return payments, int(response.Data2.(float64)), nil
}

func (c *Client) GetPayment(name string) (*Payment, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-payment", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var payment *Payment
	err = json.Unmarshal(bytes, &payment)
	if err != nil {
		return nil, err
	}
	return payment, nil
}

func (c *Client) GetUserPayments(userName string) ([]*Payment, error) {
	queryMap := map[string]string{
		"owner":        c.OrganizationName,
		"organization": c.OrganizationName,
		"user":         userName,
	}

	url := c.GetUrl("get-user-payments", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var payments []*Payment
	err = json.Unmarshal(bytes, &payments)
	if err != nil {
		return nil, err
	}
	return payments, nil
}

func (c *Client) UpdatePayment(payment *Payment) (bool, error) {
	_, affected, err := c.modifyPayment("update-payment", payment, nil)
	return affected, err
}

func (c *Client) AddPayment(payment *Payment) (bool, error) {
	_, affected, err := c.modifyPayment("add-payment", payment, nil)
	return affected, err
}

func (c *Client) DeletePayment(payment *Payment) (bool, error) {
	_, affected, err := c.modifyPayment("delete-payment", payment, nil)
	return affected, err
}

func (c *Client) NotifyPayment(payment *Payment) (bool, error) {
	_, affected, err := c.modifyPayment("notify-payment", payment, nil)
	return affected, err
}

func (c *Client) InvoicePayment(payment *Payment) (bool, error) {
	_, affected, err := c.modifyPayment("invoice-payment", payment, nil)
	return affected, err
}
