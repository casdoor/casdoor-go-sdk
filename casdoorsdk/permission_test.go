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

func TestPermission(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Permission")

	// Add a new object
	permission := &Permission{
		Owner:        "casbin",
		Name:         name,
		CreatedTime:  GetCurrentTime(),
		DisplayName:  name,
		Description:  "Casdoor Website",
		Users:        []string{"casbin/*"},
		Groups:       []string{},
		Roles:        []string{},
		Domains:      []string{},
		Model:        "user-model-built-in",
		ResourceType: "Application",
		Resources:    []string{"app-casbin"},
		Actions:      []string{"Read", "Write"},
		Effect:       "Allow",
		IsEnabled:    true,
	}
	_, err := AddPermission(permission)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	permissions, err := GetPermissions()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range permissions {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	permission, err = GetPermission(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if permission.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", permission.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	permission.Description = updatedDescription
	_, err = UpdatePermission(permission)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedPermission, err := GetPermission(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedPermission.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedPermission.Description, updatedDescription)
	}

	// Delete the object
	_, err = DeletePermission(permission)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedPermission, err := GetPermission(name)
	if err != nil || deletedPermission != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
