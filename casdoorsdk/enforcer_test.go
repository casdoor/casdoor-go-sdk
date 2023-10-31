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

func TestEnforcer(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Enforcer")

	// Add a new object
	enforcer := &Enforcer{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
		Model:       "built-in/user-model-built-in",
		Adapter:     "built-in/user-adapter-built-in",
		Description: "Casdoor Website",
	}
	_, err := AddEnforcer(enforcer)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	Enforcers, err := GetEnforcers()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range Enforcers {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	enforcer, err = GetEnforcer(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if enforcer.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", enforcer.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	enforcer.Description = updatedDescription
	_, err = UpdateEnforcer(enforcer)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedEnforcer, err := GetEnforcer(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedEnforcer.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedEnforcer.Description, updatedDescription)
	}

	// Delete the object
	_, err = DeleteEnforcer(enforcer)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedEnforcer, err := GetEnforcer(name)
	if err != nil || deletedEnforcer != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
