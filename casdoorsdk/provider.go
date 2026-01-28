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
	Owner                  string            `json:"owner"`
	Name                   string            `json:"name"`
	CreatedTime            string            `json:"createdTime"`
	DisplayName            string            `json:"displayName"`
	Category               string            `json:"category"`
	Type                   string            `json:"type"`
	SubType                string            `json:"subType"`
	Method                 string            `json:"method"`
	ClientId               string            `json:"clientId"`
	ClientSecret           string            `json:"clientSecret"`
	ClientId2              string            `json:"clientId2"`
	ClientSecret2          string            `json:"clientSecret2"`
	Cert                   string            `json:"cert"`
	CustomAuthUrl          string            `json:"customAuthUrl"`
	CustomTokenUrl         string            `json:"customTokenUrl"`
	CustomUserInfoUrl      string            `json:"customUserInfoUrl"`
	CustomLogo             string            `json:"customLogo"`
	Scopes                 string            `json:"scopes"`
	UserMapping            map[string]string `json:"userMapping"`
	HttpHeaders            map[string]string `json:"httpHeaders"`
	Host                   string            `json:"host"`
	Port                   int               `json:"port"`
	DisableSsl             bool              `json:"disableSsl"`
	Title                  string            `json:"title"`
	Content                string            `json:"content"`
	Receiver               string            `json:"receiver"`
	RegionId               string            `json:"regionId"`
	SignName               string            `json:"signName"`
	TemplateCode           string            `json:"templateCode"`
	AppId                  string            `json:"appId"`
	Endpoint               string            `json:"endpoint"`
	IntranetEndpoint       string            `json:"intranetEndpoint"`
	Domain                 string            `json:"domain"`
	Bucket                 string            `json:"bucket"`
	PathPrefix             string            `json:"pathPrefix"`
	Metadata               string            `json:"metadata"`
	IdP                    string            `json:"idP"`
	IssuerUrl              string            `json:"issuerUrl"`
	EnableSignAuthnRequest bool              `json:"enableSignAuthnRequest"`
	EmailRegex             string            `json:"emailRegex"`
	ProviderUrl            string            `json:"providerUrl"`
	EnableProxy            bool              `json:"enableProxy"`
	EnablePkce             bool              `json:"enablePkce"`
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
