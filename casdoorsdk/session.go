// Copyright 2023 The Casdoor Authors. All Rights Reserved.
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
	"strings"
)

// Session has the same definition as https://github.com/casdoor/casdoor/blob/master/object/session.go#L28
type Session struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	Application string `xorm:"varchar(100) notnull pk default ''" json:"application"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	SessionId []string `json:"sessionId"`
}

func AddUserSession(userName string, sessionId string, sessionCreateTime string) {
	session := &Session{
		Owner:       authConfig.OrganizationName,
		Name:        userName,
		Application: authConfig.ApplicationName,
		SessionId:   []string{sessionId},
		CreatedTime: sessionCreateTime,
	}

	postBytes, _ := json.Marshal(session)

	DoPost("add-user-session", nil, postBytes, false, false)
}

func ClearUserDuplicated(userName string) {
	session := &Session{
		Owner:       authConfig.OrganizationName,
		Name:        userName,
		Application: authConfig.ApplicationName,
	}

	postBytes, _ := json.Marshal(session)

	DoPost("delete-user-session", nil, postBytes, false, false)
}

func IsUserSessionDuplicated(userName string, sessionId string, sessionCreateTime string) bool {
	queryMap := map[string]string{
		"owner":       authConfig.OrganizationName,
		"name":        userName,
		"application": authConfig.ApplicationName,
		"sessionId":   sessionId,
		"createdTime": strings.Replace(strings.Replace(sessionCreateTime, "+", "%2B", -1), " ", "%20", -1),
	}

	url := GetUrl("is-user-session-duplicated", queryMap)

	resp, _ := DoGetResponse(url)

	return resp.Data == true
}
