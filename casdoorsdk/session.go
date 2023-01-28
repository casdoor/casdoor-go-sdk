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
// See the License for the specific language governing permissions and
// limitations under the License.

package casdoorsdk

import (
	"encoding/json"
	"strconv"
)

// Session has the same definition as https://github.com/casdoor/casdoor/blob/master/object/session.go#L28
type Session struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	Application string `xorm:"varchar(100) notnull pk default ''" json:"application"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	SessionId []string `json:"sessionId"`
}

func AddUserSession(claims *Claims, applicationName string, sessionId string) {
	session := &Session{
		Owner:       claims.Owner,
		Name:        claims.Name,
		Application: applicationName,
		SessionId:   []string{sessionId},
		CreatedTime: strconv.FormatInt(claims.IssuedAt.Unix(), 10),
	}

	postBytes, _ := json.Marshal(session)

	DoPost("add-user-session", nil, postBytes, false, false)
}

func ClearUserDuplicated(claims *Claims, applicationName string) {
	session := &Session{
		Owner:       claims.Owner,
		Name:        claims.Name,
		Application: applicationName,
	}

	postBytes, _ := json.Marshal(session)

	DoPost("delete-user-session", nil, postBytes, false, false)
}

func IsUserDuplicated(claims *Claims, applicationName string, sessionId string) bool {

	session := &Session{
		Owner:       claims.Owner,
		Name:        claims.Name,
		Application: applicationName,
		SessionId:   []string{sessionId},
		CreatedTime: strconv.FormatInt(claims.IssuedAt.Unix(), 10),
	}

	postBytes, _ := json.Marshal(session)

	resp, _ := DoPost("is-user-session-duplicated", nil, postBytes, false, false)

	return resp.Data == true
}
