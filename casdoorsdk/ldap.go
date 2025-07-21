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

type Ldap struct {
	Id          string `xorm:"varchar(100) notnull pk" json:"id"`
	Owner       string `xorm:"varchar(100)" json:"owner"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	ServerName   string   `xorm:"varchar(100)" json:"serverName"`
	Host         string   `xorm:"varchar(100)" json:"host"`
	Port         int      `xorm:"int" json:"port"`
	EnableSsl    bool     `xorm:"bool" json:"enableSsl"`
	Username     string   `xorm:"varchar(100)" json:"username"`
	Password     string   `xorm:"varchar(100)" json:"password"`
	BaseDn       string   `xorm:"varchar(100)" json:"baseDn"`
	Filter       string   `xorm:"varchar(200)" json:"filter"`
	FilterFields []string `xorm:"varchar(100)" json:"filterFields"`
	DefaultGroup string   `xorm:"varchar(100)" json:"defaultGroup"`

	AutoSync int    `json:"autoSync"`
	LastSync string `xorm:"varchar(100)" json:"lastSync"`
}

func (c *Client) GetLdaps() ([]*Ldap, error) {
	queryMap := map[string]string{
		"owner": "admin",
	}

	url := c.GetUrl("get-ldaps", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var ldaps []*Ldap
	err = json.Unmarshal(bytes, &ldaps)
	if err != nil {
		return nil, err
	}
	return ldaps, nil
}

func (c *Client) GetLdap(id string) (*Ldap, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", "admin", id),
	}

	url := c.GetUrl("get-ldap", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var ldap *Ldap
	err = json.Unmarshal(bytes, &ldap)
	if err != nil {
		return nil, err
	}
	return ldap, nil
}

func (c *Client) AddLdap(ldap *Ldap) (bool, error) {
	_, affected, err := c.modifyLdap("add-ldap", ldap, nil)
	return affected, err
}

func (c *Client) DeleteLdap(ldap *Ldap) (bool, error) {
	_, affected, err := c.modifyLdap("delete-ldap", ldap, nil)
	return affected, err
}

func (c *Client) UpdateLdap(ldap *Ldap) (bool, error) {
	_, affected, err := c.modifyLdap("update-ldap", ldap, nil)
	return affected, err
}
