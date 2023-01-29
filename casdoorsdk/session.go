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
	"fmt"
)

// Session has the same definition as https://github.com/casdoor/casdoor/blob/master/object/session.go#L28
type Session struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	Application string `xorm:"varchar(100) notnull pk default ''" json:"application"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	SessionId []string `json:"sessionId"`
}

func GetSessions() ([]*Session, error) {
	queryMap := map[string]string{
		"owner": authConfig.OrganizationName,
	}

	url := GetUrl("get-sessions", queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var sessions []*Session
	err = json.Unmarshal(bytes, &sessions)
	if err != nil {
		return nil, err
	}
	return sessions, nil
}

func GetSession(userName string) (*Session, error) {
	queryMap := map[string]string{
		"sessionPkId": fmt.Sprintf("%s/%s/%s", authConfig.OrganizationName, userName, authConfig.ApplicationName),
	}

	url := GetUrl("get-session", queryMap)

	bytes, err := DoGetBytesRaw(url)
	if err != nil {
		return nil, err
	}

	var session *Session
	err = json.Unmarshal(bytes, &session)
	if err != nil {
		return nil, err
	}
	return session, nil
}

func UpdateSession(userName string, sessionId string) (bool, error) {
	session := &Session{
		Owner:       authConfig.OrganizationName,
		Name:        userName,
		Application: authConfig.ApplicationName,
		SessionId:   []string{sessionId},
	}

	postBytes, _ := json.Marshal(session)

	resp, err := DoPost("update-session", nil, postBytes, false, false)
	if err != nil {
		return false, err
	}
	return resp.Data == "Affected", nil
}

func AddSession(userName string, sessionId string) (bool, error) {
	session := &Session{
		Owner:       authConfig.OrganizationName,
		Name:        userName,
		Application: authConfig.ApplicationName,
		SessionId:   []string{sessionId},
	}

	postBytes, _ := json.Marshal(session)

	resp, err := DoPost("add-session", nil, postBytes, false, false)
	if err != nil {
		return false, err
	}
	return resp.Data == "Affected", nil
}

func DeleteSession(userName string) (bool, error) {
	session := &Session{
		Owner:       authConfig.OrganizationName,
		Name:        userName,
		Application: authConfig.ApplicationName,
	}

	postBytes, _ := json.Marshal(session)

	resp, err := DoPost("delete-session", nil, postBytes, false, false)
	if err != nil {
		return false, err
	}
	return resp.Data == "Affected", nil
}

func IsSessionDuplicated(userName string, sessionId string) bool {
	queryMap := map[string]string{
		"sessionPkId": fmt.Sprintf("%s/%s/%s", authConfig.OrganizationName, userName, authConfig.ApplicationName),
		"sessionId":   sessionId,
	}

	url := GetUrl("is-session-duplicated", queryMap)

	resp, _ := DoGetResponse(url)
	return resp.Data == true
}
