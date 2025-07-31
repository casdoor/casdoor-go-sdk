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

func GetInvitations() ([]*Invitation, error) {
	return globalClient.GetInvitations()
}

func GetPaginationInvitations(p int, pageSize int, queryMap map[string]string) ([]*Invitation, int, error) {
	return globalClient.GetPaginationInvitations(p, pageSize, queryMap)
}

func GetInvitation(name string) (*Invitation, error) {
	return globalClient.GetInvitation(name)
}

func GetInvitationInfo(code string, applicationName string) (*Invitation, error) {
	return globalClient.GetInvitationInfo(code, applicationName)
}

func UpdateInvitation(invitation *Invitation) (bool, error) {
	return globalClient.UpdateInvitation(invitation)
}

func UpdateInvitationForColumns(invitation *Invitation, columns []string) (bool, error) {
	return globalClient.UpdateInvitationForColumns(invitation, columns)
}

func AddInvitation(invitation *Invitation) (bool, error) {
	return globalClient.AddInvitation(invitation)
}

func DeleteInvitation(invitation *Invitation) (bool, error) {
	return globalClient.DeleteInvitation(invitation)
}
