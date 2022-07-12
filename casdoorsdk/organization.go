// Copyright 2021 The casbin Authors. All Rights Reserved.
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

import "encoding/json"

// Organization has the same definition as https://github.com/casdoor/casdoor/blob/master/object/organization.go#L25
type Organization struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	DisplayName        string `xorm:"varchar(100)" json:"displayName"`
	WebsiteUrl         string `xorm:"varchar(100)" json:"websiteUrl"`
	Favicon            string `xorm:"varchar(100)" json:"favicon"`
	PasswordType       string `xorm:"varchar(100)" json:"passwordType"`
	PasswordSalt       string `xorm:"varchar(100)" json:"passwordSalt"`
	PhonePrefix        string `xorm:"varchar(10)"  json:"phonePrefix"`
	DefaultAvatar      string `xorm:"varchar(100)" json:"defaultAvatar"`
	MasterPassword     string `xorm:"varchar(100)" json:"masterPassword"`
	EnableSoftDeletion bool   `json:"enableSoftDeletion"`
}

func AddOrganization(organization *Organization) (bool, error) {
	if organization.Owner == "" {
		organization.Owner = "admin"
	}
	postBytes, err := json.Marshal(organization)
	if err != nil {
		return false, err
	}

	resp, err := doPost("add-organization", nil, postBytes, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}

func DeleteOrganization(name string) (bool, error) {
	organization := Organization{
		Owner: "admin",
		Name:  name,
	}
	postBytes, err := json.Marshal(organization)
	if err != nil {
		return false, err
	}

	resp, err := doPost("delete-organization", nil, postBytes, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}
