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
	"bytes"
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

func Enforce(permissionRule *PermissionRule) (bool, error) {
	postBytes, err := json.Marshal(permissionRule)
	if err != nil {
		return false, err
	}

	bytes, err := doEnforce("enforce", postBytes)
	if err != nil {
		return false, err
	}

	var allow bool

	err = unmarshalResponse(bytes, &allow)
	if err != nil {
		return false, err
	}

	return allow, nil
}

func BatchEnforce(permissionRules []PermissionRule) ([]bool, error) {
	postBytes, err := json.Marshal(permissionRules)
	if err != nil {
		return nil, err
	}

	bytes, err := doEnforce("batch-enforce", postBytes)
	if err != nil {
		return nil, err
	}

	var allow []bool

	err = unmarshalResponse(bytes, &allow)
	if err != nil {
		return nil, err
	}

	return allow, nil
}

func doEnforce(action string, postBytes []byte) ([]byte, error) {
	url := GetUrl(action, nil)
	bytes, err := DoPostBytesRaw(url, "", bytes.NewBuffer(postBytes))
	if err != nil {
		return nil, err
	}

	if len(bytes) == 0 {
		return nil, errors.New("response is empty")
	}

	if bytes[0] == '{' {
		var res Response

		err = unmarshalResponse(bytes, &res)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(res.Msg)
	}

	return bytes, nil
}
