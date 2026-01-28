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
	Owner               string   `json:"owner"`
	Name                string   `json:"name"`
	CreatedTime         string   `json:"createdTime"`
	DisplayName         string   `json:"displayName"` // Payment Provider Info
	Provider            string   `json:"provider"`
	Type                string   `json:"type"` // Product Info
	Products            []string `json:"products"`
	ProductsDisplayName string   `json:"productsDisplayName"`
	Detail              string   `json:"detail"`
	Currency            string   `json:"currency"`
	Price               float64  `json:"price"` // Payer Info
	User                string   `json:"user"`
	PersonName          string   `json:"personName"`
	PersonIdCard        string   `json:"personIdCard"`
	PersonEmail         string   `json:"personEmail"`
	PersonPhone         string   `json:"personPhone"` // Invoice Info
	InvoiceType         string   `json:"invoiceType"`
	InvoiceTitle        string   `json:"invoiceTitle"`
	InvoiceTaxId        string   `json:"invoiceTaxId"`
	InvoiceRemark       string   `json:"invoiceRemark"`
	InvoiceUrl          string   `json:"invoiceUrl"` // Order Info
	Order               string   `json:"order"`
	OutOrderId          string   `json:"outOrderId"`
	PayUrl              string   `json:"payUrl"`
	SuccessUrl          string   `json:"successUrl"`
	State               string   `json:"state"`
	Message             string   `json:"message"`
	// Deprecated: removed from server
	ProductName string `json:"productName"`
	// Deprecated: removed from server
	ProductDisplayName string `json:"productDisplayName"`
	// Deprecated: removed from server
	Tag string `json:"tag"`
	// Deprecated: removed from server
	ReturnUrl string `json:"returnUrl"`
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

func (c *Client) PlaceOrder(productName string, providerName string, userName string) (*Payment, error) {
	queryMap := map[string]string{
		"productId":    fmt.Sprintf("%s/%s", c.OrganizationName, productName),
		"providerName": providerName,
		"userName":     userName,
	}

	resp, err := c.DoPost("place-order", queryMap, []byte(""), false, false)
	if err != nil {
		return nil, err
	}

	paymentJson, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	var payment Payment
	err = json.Unmarshal(paymentJson, &payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (c *Client) PayOrder(paymentName string, providerName string) (*Payment, error) {
	queryMap := map[string]string{
		"id":           fmt.Sprintf("%s/%s", c.OrganizationName, paymentName),
		"providerName": providerName,
	}

	resp, err := c.DoPost("pay-order", queryMap, []byte(""), false, false)
	if err != nil {
		return nil, err
	}

	paymentJson, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	var payment Payment
	err = json.Unmarshal(paymentJson, &payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
