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
