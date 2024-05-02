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
// See the License for the specific language governing records and
// limitations under the License.

package casdoorsdk

import (
	"encoding/json"
	"errors"
)

type PermissionRule struct {
	Ptype string `xorm:"varchar(100) index not null default ''" json:"ptype"`
	V0    string `xorm:"varchar(100) index not null default ''" json:"v0"`
	V1    string `xorm:"varchar(100) index not null default ''" json:"v1"`
	V2    string `xorm:"varchar(100) index not null default ''" json:"v2"`
	V3    string `xorm:"varchar(100) index not null default ''" json:"v3"`
	V4    string `xorm:"varchar(100) index not null default ''" json:"v4"`
	V5    string `xorm:"varchar(100) index not null default ''" json:"v5"`
	Id    string `xorm:"varchar(100) index not null default ''" json:"id"`
}

type CasbinRequest = []interface{}

func (c *Client) Enforce(permissionId string, modelId string, resourceId string, enforcerId string, owner string, casbinRequest CasbinRequest) (bool, error) {
	postBytes, err := json.Marshal(casbinRequest)
	if err != nil {
		return false, err
	}

	res, err := c.doEnforce("enforce", permissionId, modelId, resourceId, enforcerId, owner, postBytes)
	if err != nil {
		return false, err
	}

	results, ok := res.Data.([]interface{})
	if !ok {
		return false, errors.New("invalid data")
	}

	for _, result := range results {
		isAllow, ok := result.(bool)
		if !ok {
			return false, errors.New("invalid data")
		}

		if isAllow {
			return isAllow, nil
		}
	}

	return false, nil
}

func Enforce(permissionId string, modelId string, resourceId string, enforcerId string, owner string, casbinRequest CasbinRequest) (bool, error) {
	return globalClient.Enforce(permissionId, modelId, resourceId, enforcerId, owner, casbinRequest)
}

func (c *Client) BatchEnforce(permissionId string, modelId string, resourceId string, enforcerId string, owner string, casbinRequests []CasbinRequest) ([][]bool, error) {
	postBytes, err := json.Marshal(casbinRequests)
	if err != nil {
		return nil, err
	}

	res, err := c.doEnforce("batch-enforce", permissionId, modelId, resourceId, enforcerId, owner, postBytes)
	if err != nil {
		return nil, err
	}

	var allows [][]bool
	data, ok := res.Data.([]interface{})
	if !ok {
		return nil, errors.New("invalid data")
	}

	for _, d := range data {
		elems, ok := d.([]interface{})
		if !ok {
			return nil, errors.New("invalid data")
		}
		var permRes []bool
		for _, el := range elems {
			r, ok := el.(bool)
			if !ok {
				return nil, errors.New("invalid data")
			}
			permRes = append(permRes, r)
		}
		allows = append(allows, permRes)
	}

	return allows, nil
}

func BatchEnforce(permissionId string, modelId string, resourceId string, enforcerId string, owner string, casbinRequests []CasbinRequest) ([][]bool, error) {
	return globalClient.BatchEnforce(permissionId, modelId, resourceId, enforcerId, owner, casbinRequests)
}

func (c *Client) doEnforce(action string, permissionId string, modelId string, resourceId string, enforcerId string, owner string, postBytes []byte) (*Response, error) {
	queryMap := map[string]string{
		"permissionId": permissionId,
		"modelId":      modelId,
		"resourceId":   resourceId,
		"enforcerId":   enforcerId,
		"owner":        owner,
	}

	// bytes, err := DoPostBytesRaw(url, "", bytes.NewBuffer(postBytes))
	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
