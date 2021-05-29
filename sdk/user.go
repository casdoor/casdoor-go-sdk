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

package sdk

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type AuthConfig struct {
	Endpoint     string
	ClientId     string
	ClientSecret string
	JwtSecret    string
}

var authConfig AuthConfig

type User struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Id            string `xorm:"varchar(100)" json:"id"`
	Type          string `xorm:"varchar(100)" json:"type"`
	Password      string `xorm:"varchar(100)" json:"password"`
	PasswordType  string `xorm:"varchar(100)" json:"passwordType"`
	DisplayName   string `xorm:"varchar(100)" json:"displayName"`
	Avatar        string `xorm:"varchar(255)" json:"avatar"`
	Email         string `xorm:"varchar(100)" json:"email"`
	PhonePrefix   string `xorm:"varchar(10)" json:"phonePrefix"`
	Phone         string `xorm:"varchar(100)" json:"phone"`
	Affiliation   string `xorm:"varchar(100)" json:"affiliation"`
	Tag           string `xorm:"varchar(100)" json:"tag"`
	IsAdmin       bool   `json:"isAdmin"`
	IsGlobalAdmin bool   `json:"isGlobalAdmin"`

	Github string `xorm:"varchar(100)" json:"github"`
	Google string `xorm:"varchar(100)" json:"google"`
}

func InitConfig(endpoint string, clientId string, clientSecret string, jwtSecret string) {
	authConfig = AuthConfig{
		Endpoint:     endpoint,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		JwtSecret:    jwtSecret,
	}
}

func GetUsers(owner string) []*User {
	url := fmt.Sprintf("%s/api/get-users?owner=%s", authConfig.Endpoint, owner)
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var users []*User
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		panic(err)
	}
	return users
}
