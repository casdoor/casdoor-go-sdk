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

func Enforce(permissionId, modelId, resourceId string, casbinRequest CasbinRequest) (bool, error) {
	postBytes, err := json.Marshal(casbinRequest)
	if err != nil {
		return false, err
	}

	res, err := doEnforce("enforce", permissionId, modelId, resourceId, postBytes)
	if err != nil {
		return false, err
	}

	data, ok := res.Data.([]interface{})
	if !ok {
		return false, errors.New("invalid data")
	}

	allow, ok := data[0].(bool)
	if !ok {
		return false, errors.New("invalid data")
	}

	return allow, nil
}

func BatchEnforce(permissionId, modelId, resourceId string, casbinRequests []CasbinRequest) ([]bool, error) {
	postBytes, err := json.Marshal(casbinRequests)
	if err != nil {
		return nil, err
	}

	res, err := doEnforce("batch-enforce", permissionId, modelId, resourceId, postBytes)
	if err != nil {
		return nil, err
	}

	var allows []bool
	data, ok := res.Data.([]interface{})
	if !ok {
		return nil, errors.New("invalid data")
	}

	for _, d := range data {
		el, ok := d.([]interface{})
		if !ok {
			return nil, errors.New("invalid data")
		}
		allows = append(allows, el[0].(bool))
	}

	return allows, nil
}

func doEnforce(action string, permissionId, modelId, resourceId string, postBytes []byte) (*Response, error) {
	queryMap := map[string]string{
		"permissionId": permissionId,
		"modelId":      modelId,
		"resourceId":   resourceId,
	}

	//bytes, err := DoPostBytesRaw(url, "", bytes.NewBuffer(postBytes))
	resp, err := DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
