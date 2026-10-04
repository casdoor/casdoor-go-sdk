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
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("User")

	// Add a new object
	user := &User{
		Owner:       TestCasdoorOrganization,
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
	}
	_, err := AddUser(user)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	users, err := GetUsers()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range users {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	user, err = GetUser(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if user.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", user.Name, name)
	}

	// Update the object
	updatedDisplayName := "Updated Casdoor Website"
	user.DisplayName = updatedDisplayName
	_, err = UpdateUser(user)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedUser, err := GetUser(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedUser.DisplayName != updatedDisplayName {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedUser.DisplayName, updatedDisplayName)
	}

	// Delete the object
	_, err = DeleteUser(user)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedUser, err := GetUser(name)
	if err != nil || deletedUser != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}

func TestRemoveUserFromGroup(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	groupName := getRandomName("RemoveGroup")
	otherGroupName := getRandomName("RemoveOtherGroup")
	userName := getRandomName("RemoveUser")

	for _, name := range []string{groupName, otherGroupName} {
		group := &Group{
			Owner:       TestCasdoorOrganization,
			Name:        name,
			CreatedTime: GetCurrentTime(),
			DisplayName: name,
		}
		if _, err := AddGroup(group); err != nil {
			t.Fatalf("Failed to add group: %v", err)
		}
		defer func(group *Group) {
			if _, err := DeleteGroup(group); err != nil {
				t.Errorf("Failed to delete group: %v", err)
			}
		}(group)
	}

	// Add a user that belongs to two groups
	user := &User{
		Owner:       TestCasdoorOrganization,
		Name:        userName,
		CreatedTime: GetCurrentTime(),
		DisplayName: userName,
		Groups: []string{
			fmt.Sprintf("%s/%s", TestCasdoorOrganization, groupName),
			fmt.Sprintf("%s/%s", TestCasdoorOrganization, otherGroupName),
		},
	}
	if _, err := AddUser(user); err != nil {
		t.Fatalf("Failed to add user: %v", err)
	}
	defer func() {
		if _, err := DeleteUser(user); err != nil {
			t.Errorf("Failed to delete user: %v", err)
		}
	}()

	user, err := GetUser(userName)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}
	if len(user.Groups) != 2 {
		t.Fatalf("Expected the user to be in 2 groups, got %v", user.Groups)
	}

	// Remove the user from one group, the owner defaults to the organization of the client
	removed, err := RemoveUserFromGroup("", userName, groupName)
	if err != nil {
		t.Fatalf("Failed to remove user from group: %v", err)
	}
	if !removed {
		t.Fatalf("Expected the user to be removed from the group %s", groupName)
	}

	// Validate that only that group is removed
	user, err = GetUser(userName)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}
	otherGroupId := fmt.Sprintf("%s/%s", TestCasdoorOrganization, otherGroupName)
	if len(user.Groups) != 1 || user.Groups[0] != otherGroupId {
		t.Fatalf("Expected the user to be only in the group %s, got %v", otherGroupId, user.Groups)
	}

	// Removing the user from a group that it is not in changes nothing
	removed, err = RemoveUserFromGroup("", userName, groupName)
	if err != nil {
		t.Fatalf("Failed to remove user from group: %v", err)
	}
	if removed {
		t.Fatalf("Expected nothing to be removed from the group %s again", groupName)
	}

	// Remove the user from the other group with an explicit owner
	removed, err = RemoveUserFromGroup(TestCasdoorOrganization, userName, otherGroupName)
	if err != nil {
		t.Fatalf("Failed to remove user from group: %v", err)
	}
	if !removed {
		t.Fatalf("Expected the user to be removed from the group %s", otherGroupName)
	}

	user, err = GetUser(userName)
	if err != nil {
		t.Fatalf("Failed to get user: %v", err)
	}
	if len(user.Groups) != 0 {
		t.Fatalf("Expected the user to be in no group, got %v", user.Groups)
	}
}
