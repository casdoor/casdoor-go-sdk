// Copyright 2021 The Casdoor Authors. All Rights Reserved.
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

type AccountItem struct {
	Name       string `json:"name"`
	Visible    bool   `json:"visible"`
	ViewRule   string `json:"viewRule"`
	ModifyRule string `json:"modifyRule"`
	Regex      string `json:"regex"`
}

type ThemeData struct {
	ThemeType    string `json:"themeType"`
	ColorPrimary string `json:"colorPrimary"`
	BorderRadius int    `json:"borderRadius"`
	IsCompact    bool   `json:"isCompact"`
	IsEnabled    bool   `json:"isEnabled"`
}

type MfaItem struct {
	Name string `json:"name"`
	Rule string `json:"rule"`
}

// Organization has the same definition as https://github.com/casdoor/casdoor/blob/master/object/organization.go#L50
type Organization struct {
	Owner                  string         `json:"owner"`
	Name                   string         `json:"name"`
	CreatedTime            string         `json:"createdTime"`
	DisplayName            string         `json:"displayName"`
	WebsiteUrl             string         `json:"websiteUrl"`
	Logo                   string         `json:"logo"`
	LogoDark               string         `json:"logoDark"`
	Favicon                string         `json:"favicon"`
	HasPrivilegeConsent    bool           `json:"hasPrivilegeConsent"`
	PasswordType           string         `json:"passwordType"`
	PasswordSalt           string         `json:"passwordSalt"`
	PasswordOptions        []string       `json:"passwordOptions"`
	PasswordObfuscatorType string         `json:"passwordObfuscatorType"`
	PasswordObfuscatorKey  string         `json:"passwordObfuscatorKey"`
	PasswordExpireDays     int            `json:"passwordExpireDays"`
	CountryCodes           []string       `json:"countryCodes"`
	DefaultAvatar          string         `json:"defaultAvatar"`
	DefaultApplication     string         `json:"defaultApplication"`
	UserTypes              []string       `json:"userTypes"`
	Tags                   []string       `json:"tags"`
	Languages              []string       `json:"languages"`
	ThemeData              *ThemeData     `json:"themeData"`
	MasterPassword         string         `json:"masterPassword"`
	DefaultPassword        string         `json:"defaultPassword"`
	MasterVerificationCode string         `json:"masterVerificationCode"`
	IpWhitelist            string         `json:"ipWhitelist"`
	InitScore              int            `json:"initScore"`
	EnableSoftDeletion     bool           `json:"enableSoftDeletion"`
	IsProfilePublic        bool           `json:"isProfilePublic"`
	UseEmailAsUsername     bool           `json:"useEmailAsUsername"`
	EnableTour             bool           `json:"enableTour"`
	DisableSignin          bool           `json:"disableSignin"`
	IpRestriction          string         `json:"ipRestriction"`
	NavItems               []string       `json:"navItems"`
	UserNavItems           []string       `json:"userNavItems"`
	WidgetItems            []string       `json:"widgetItems"`
	MfaItems               []*MfaItem     `json:"mfaItems"`
	MfaRememberInHours     int            `json:"mfaRememberInHours"`
	AccountItems           []*AccountItem `json:"accountItems"`
	OrgBalance             float64        `json:"orgBalance"`
	UserBalance            float64        `json:"userBalance"`
	BalanceCredit          float64        `json:"balanceCredit"`
	BalanceCurrency        string         `json:"balanceCurrency"`
}

func (c *Client) GetOrganization(name string) (*Organization, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", "admin", name),
	}

	url := c.GetUrl("get-organization", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var organization *Organization
	err = json.Unmarshal(bytes, &organization)
	if err != nil {
		return nil, err
	}
	return organization, nil
}

func (c *Client) GetOrganizations() ([]*Organization, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-organizations", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var organizations []*Organization
	err = json.Unmarshal(bytes, &organizations)
	if err != nil {
		return nil, err
	}
	return organizations, nil
}

func (c *Client) GetOrganizationNames() ([]*Organization, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-organization-names", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var organizationNames []*Organization
	err = json.Unmarshal(bytes, &organizationNames)
	if err != nil {
		return nil, err
	}
	return organizationNames, nil
}

func (c *Client) AddOrganization(organization *Organization) (bool, error) {
	_, affected, err := c.modifyOrganization("add-organization", organization, nil)
	return affected, err
}

func (c *Client) DeleteOrganization(organization *Organization) (bool, error) {
	_, affected, err := c.modifyOrganization("delete-organization", organization, nil)
	return affected, err
}

func (c *Client) UpdateOrganization(organization *Organization) (bool, error) {
	_, affected, err := c.modifyOrganization("update-organization", organization, nil)
	return affected, err
}
