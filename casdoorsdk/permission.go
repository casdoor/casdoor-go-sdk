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

import (
	"encoding/json"
	"time"
)

type Permission struct {
	Action       string   `json:"action"`
	Actions      []string `json:"actions"`
	CreatedTime  string   `xorm:"varchar(100)" json:"createdTime"`
	DisplayName  string   `xorm:"varchar(100)" json:"displayName"`
	Effect       string   `xorm:"varchar(100)" json:"effect"`
	IsEnabled    bool     `json:"isEnabled"`
	Name         string   `xorm:"varchar(100)" json:"name"`
	Owner        string   `xorm:"varchar(100)" json:"owner"`
	ResourceType string   `xorm:"varchar(100)" json:"resourceType"`
	Resources    []string `json:"resources"`
	Roles        []string `json:"roles"`
	Users        []string `json:"users"`
}

func GetPermission() ([]*Permission, error) {
	queryMap := map[string]string{
		"owner":       authConfig.OrganizationName,
		"application": authConfig.ApplicationName,
	}
	url := GetUrl("get-permissions", queryMap)
	bytes, err := DoGetBytes(url)
	if err != nil {
		return nil, err
	}
	var permission []*Permission
	err = json.Unmarshal(bytes, &permission)
	if err != nil {
		return nil, err
	}
	return permission, nil
}

func AddPermission(q Permission) (*Response, error) {
	data := Permission{
		Action:       "Read",
		Actions:      q.Actions,
		CreatedTime:  time.Now().UTC().String(),
		DisplayName:  q.Name,
		Effect:       "Allow",
		IsEnabled:    true,
		Name:         q.Name,
		Owner:        authConfig.OrganizationName,
		ResourceType: "Api",
		Resources:    q.Resources,
		Roles:        []string{},
		Users:        []string{},
	}

	postBytes, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	resp, err := doPost("add-permission", nil, postBytes, false)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
