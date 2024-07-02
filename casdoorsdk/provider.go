// Copyright 2022 The Casdoor Authors. All Rights Reserved.
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

type Provider struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk unique" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	DisplayName       string            `xorm:"varchar(100)" json:"displayName"`
	Category          string            `xorm:"varchar(100)" json:"category"`
	Type              string            `xorm:"varchar(100)" json:"type"`
	SubType           string            `xorm:"varchar(100)" json:"subType"`
	Method            string            `xorm:"varchar(100)" json:"method"`
	ClientId          string            `xorm:"varchar(100)" json:"clientId"`
	ClientSecret      string            `xorm:"varchar(2000)" json:"clientSecret"`
	ClientId2         string            `xorm:"varchar(100)" json:"clientId2"`
	ClientSecret2     string            `xorm:"varchar(100)" json:"clientSecret2"`
	Cert              string            `xorm:"varchar(100)" json:"cert"`
	CustomAuthUrl     string            `xorm:"varchar(200)" json:"customAuthUrl"`
	CustomTokenUrl    string            `xorm:"varchar(200)" json:"customTokenUrl"`
	CustomUserInfoUrl string            `xorm:"varchar(200)" json:"customUserInfoUrl"`
	CustomLogo        string            `xorm:"varchar(200)" json:"customLogo"`
	Scopes            string            `xorm:"varchar(100)" json:"scopes"`
	UserMapping       map[string]string `xorm:"varchar(500)" json:"userMapping"`

	Host       string `xorm:"varchar(100)" json:"host"`
	Port       int    `json:"port"`
	DisableSsl bool   `json:"disableSsl"` // If the provider type is WeChat, DisableSsl means EnableQRCode
	Title      string `xorm:"varchar(100)" json:"title"`
	Content    string `xorm:"varchar(1000)" json:"content"` // If provider type is WeChat, Content means QRCode string by Base64 encoding
	Receiver   string `xorm:"varchar(100)" json:"receiver"`

	RegionId     string `xorm:"varchar(100)" json:"regionId"`
	SignName     string `xorm:"varchar(100)" json:"signName"`
	TemplateCode string `xorm:"varchar(100)" json:"templateCode"`
	AppId        string `xorm:"varchar(100)" json:"appId"`

	Endpoint         string `xorm:"varchar(1000)" json:"endpoint"`
	IntranetEndpoint string `xorm:"varchar(100)" json:"intranetEndpoint"`
	Domain           string `xorm:"varchar(100)" json:"domain"`
	Bucket           string `xorm:"varchar(100)" json:"bucket"`
	PathPrefix       string `xorm:"varchar(100)" json:"pathPrefix"`

	Metadata               string `xorm:"mediumtext" json:"metadata"`
	IdP                    string `xorm:"mediumtext" json:"idP"`
	IssuerUrl              string `xorm:"varchar(100)" json:"issuerUrl"`
	EnableSignAuthnRequest bool   `json:"enableSignAuthnRequest"`

	ProviderUrl string `xorm:"varchar(200)" json:"providerUrl"`
}

func (c *Client) GetProviders() ([]*Provider, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-providers", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var providers []*Provider
	err = json.Unmarshal(bytes, &providers)
	if err != nil {
		return nil, err
	}
	return providers, nil
}

func (c *Client) GetProvider(name string) (*Provider, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-provider", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var provider *Provider
	err = json.Unmarshal(bytes, &provider)
	if err != nil {
		return nil, err
	}
	return provider, nil
}

func (c *Client) GetPaginationProviders(p int, pageSize int, queryMap map[string]string) ([]*Provider, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-providers", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var providers []*Provider
	err = json.Unmarshal(dataBytes, &providers)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return providers, int(response.Data2.(float64)), nil
}

func (c *Client) UpdateProvider(provider *Provider) (bool, error) {
	_, affected, err := c.modifyProvider("update-provider", provider, nil)
	return affected, err
}

func (c *Client) AddProvider(provider *Provider) (bool, error) {
	_, affected, err := c.modifyProvider("add-provider", provider, nil)
	return affected, err
}

func (c *Client) DeleteProvider(provider *Provider) (bool, error) {
	_, affected, err := c.modifyProvider("delete-provider", provider, nil)
	return affected, err
}
