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
	"fmt"
)

func (c *Client) PlaceOrder(productInfos []ProductInfo, userName string) (*Order, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}
	if userName != "" {
		queryMap["userName"] = userName
	}

	requestBody := struct {
		ProductInfos []ProductInfo `json:"productInfos"`
	}{
		ProductInfos: productInfos,
	}
	postBytes, err := json.Marshal(requestBody)
	if err != nil {
		return nil, err
	}

	resp, err := c.DoPost("place-order", queryMap, postBytes, false, false)
	if err != nil {
		return nil, err
	}

	orderBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	var order Order
	err = json.Unmarshal(orderBytes, &order)
	if err != nil {
		return nil, err
	}

	return &order, nil
}

func (c *Client) PayOrder(orderName string, providerName string) (*Payment, error) {
	queryMap := map[string]string{
		"id":           fmt.Sprintf("%s/%s", c.OrganizationName, orderName),
		"providerName": providerName,
	}

	resp, err := c.DoPost("pay-order", queryMap, []byte(""), false, false)
	if err != nil {
		return nil, err
	}

	paymentBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	var payment Payment
	err = json.Unmarshal(paymentBytes, &payment)
	if err != nil {
		return nil, err
	}

	return &payment, nil
}
