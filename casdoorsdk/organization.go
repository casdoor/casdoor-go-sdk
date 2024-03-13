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
}

type ThemeData struct {
	ThemeType    string `xorm:"varchar(30)" json:"themeType"`
	ColorPrimary string `xorm:"varchar(10)" json:"colorPrimary"`
	BorderRadius int    `xorm:"int" json:"borderRadius"`
	IsCompact    bool   `xorm:"bool" json:"isCompact"`
	IsEnabled    bool   `xorm:"bool" json:"isEnabled"`
}

type MfaItem struct {
	Name string `json:"name"`
	Rule string `json:"rule"`
}

// Organization has the same definition as https://github.com/casdoor/casdoor/blob/master/object/organization.go#L25
type Organization struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	DisplayName        string     `xorm:"varchar(100)" json:"displayName"`
	WebsiteUrl         string     `xorm:"varchar(100)" json:"websiteUrl"`
	Favicon            string     `xorm:"varchar(100)" json:"favicon"`
	PasswordType       string     `xorm:"varchar(100)" json:"passwordType"`
	PasswordSalt       string     `xorm:"varchar(100)" json:"passwordSalt"`
	PasswordOptions    []string   `xorm:"varchar(100)" json:"passwordOptions"`
	CountryCodes       []string   `xorm:"varchar(200)"  json:"countryCodes"`
	DefaultAvatar      string     `xorm:"varchar(200)" json:"defaultAvatar"`
	DefaultApplication string     `xorm:"varchar(100)" json:"defaultApplication"`
	Tags               []string   `xorm:"mediumtext" json:"tags"`
	Languages          []string   `xorm:"varchar(255)" json:"languages"`
	ThemeData          *ThemeData `xorm:"json" json:"themeData"`
	MasterPassword     string     `xorm:"varchar(100)" json:"masterPassword"`
	InitScore          int        `json:"initScore"`
	EnableSoftDeletion bool       `json:"enableSoftDeletion"`
	IsProfilePublic    bool       `json:"isProfilePublic"`

	MfaItems     []*MfaItem     `xorm:"varchar(300)" json:"mfaItems"`
	AccountItems []*AccountItem `xorm:"varchar(5000)" json:"accountItems"`
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
