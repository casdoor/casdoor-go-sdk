// Copyright 2025 The Casdoor Authors. All Rights Reserved.
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
	"strconv"
)

// Invitation has the same definition as https://github.com/casdoor/casdoor/blob/master/object/invitation.go
type Invitation struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	UpdatedTime string `xorm:"varchar(100)" json:"updatedTime"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`

	Code      string `xorm:"varchar(100) index" json:"code"`
	IsRegexp  bool   `json:"isRegexp"`
	Quota     int    `json:"quota"`
	UsedCount int    `json:"usedCount"`

	Application string `xorm:"varchar(100)" json:"application"`
	Username    string `xorm:"varchar(100)" json:"username"`
	Email       string `xorm:"varchar(100)" json:"email"`
	Phone       string `xorm:"varchar(100)" json:"phone"`

	SignupGroup string `xorm:"varchar(100)" json:"signupGroup"`
	DefaultCode string `xorm:"varchar(100)" json:"defaultCode"`

	State string `xorm:"varchar(100)" json:"state"`
}

func (c *Client) GetInvitations() ([]*Invitation, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-invitations", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var invitations []*Invitation
	err = json.Unmarshal(bytes, &invitations)
	if err != nil {
		return nil, err
	}
	return invitations, nil
}

func (c *Client) GetPaginationInvitations(p int, pageSize int, queryMap map[string]string) ([]*Invitation, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-invitations", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var invitations []*Invitation
	err = json.Unmarshal(dataBytes, &invitations)
	if err != nil {
		return nil, 0, err
	}

	return invitations, int(response.Data2.(float64)), nil
}

func (c *Client) GetInvitation(name string) (*Invitation, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-invitation", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var invitation *Invitation
	err = json.Unmarshal(bytes, &invitation)
	if err != nil {
		return nil, err
	}
	return invitation, nil
}

func (c *Client) GetInvitationInfo(code string, applicationName string) (*Invitation, error) {
	applicationId := fmt.Sprintf("admin/%s", applicationName)
	queryMap := map[string]string{
		"applicationId": applicationId,
		"code":          code,
	}

	url := c.GetUrl("get-invitation-info", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var invitation *Invitation
	err = json.Unmarshal(bytes, &invitation)
	if err != nil {
		return nil, err
	}
	return invitation, nil
}

func (c *Client) UpdateInvitation(invitation *Invitation) (bool, error) {
	_, affected, err := c.modifyInvitation("update-invitation", invitation, nil)
	return affected, err
}

func (c *Client) UpdateInvitationForColumns(invitation *Invitation, columns []string) (bool, error) {
	_, affected, err := c.modifyInvitation("update-invitation", invitation, columns)
	return affected, err
}

func (c *Client) AddInvitation(invitation *Invitation) (bool, error) {
	_, affected, err := c.modifyInvitation("add-invitation", invitation, nil)
	return affected, err
}

func (c *Client) DeleteInvitation(invitation *Invitation) (bool, error) {
	_, affected, err := c.modifyInvitation("delete-invitation", invitation, nil)
	return affected, err
}

func (i Invitation) GetId() string {
	return fmt.Sprintf("%s/%s", i.Owner, i.Name)
}
