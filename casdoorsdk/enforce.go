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
	var allow bool

	postBytes, err := json.Marshal(permissionRule)
	if err != nil {
		return false, err
	}

	url := GetUrl("enforce", nil)
	bytes, err := DoPostBytesRaw(url, "", bytes.NewBuffer(postBytes))
	if err != nil {
		return false, err
	}

	err = json.Unmarshal(bytes, &allow)
	if err != nil {
		return false, err
	}

	return allow, nil
}

func BatchEnforce(permissionRules []PermissionRule) ([]bool, error) {
	var allow []bool

	postBytes, err := json.Marshal(permissionRules)
	if err != nil {
		return nil, err
	}

	url := GetUrl("batch-enforce", nil)
	bytes, err := DoPostBytesRaw(url, "", bytes.NewBuffer(postBytes))
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &allow)
	if err != nil {
		return nil, err
	}

	return allow, nil
}
