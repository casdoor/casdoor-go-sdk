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
	"log"
	"testing"
)

func TestGroup(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Group")

	// Add a new object
	group := &Group{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
	}
	_, err := AddGroup(group)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	id := fmt.Sprintf("%s/%s", group.Owner, group.Name)
	log.Printf("group id: %s\n", id)
	log.Printf("update group before name : %s \n", name)
	name = getRandomName("Group")
	group.Name = name
	_, err = UpdateGroupById(id, group)
	if err != nil {
		t.Fatalf("Failed to update by id object: %v", err)
	}
	log.Printf("update group after name : %s \n", name)

	// Get all objects, check if our added object is inside the list
	groups, err := GetGroups()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range groups {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	group, err = GetGroup(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if group.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", group.Name, name)
	}

	// Update the object
	updatedDisplayName := "Updated Casdoor Website"
	group.DisplayName = updatedDisplayName
	_, err = UpdateGroup(group)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedGroup, err := GetGroup(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedGroup.DisplayName != updatedDisplayName {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedGroup.DisplayName, updatedDisplayName)
	}

	// Delete the object
	_, err = DeleteGroup(group)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedGroup, err := GetGroup(name)
	if err != nil || deletedGroup != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
