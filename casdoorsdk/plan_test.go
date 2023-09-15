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
	"time"
)

func TestPlan(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Plan")

	// Add a new object
	plan := &Plan{
		Owner:        "admin",
		Name:         name,
		CreatedTime:  time.Now().Format(time.RFC3339),
		DisplayName:  name,
		Description:  "Casdoor Website",
	}
	_, err := AddPlan(plan)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	Plans, err := GetPlans()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range Plans {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	plan, err = GetPlan(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if plan.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", plan.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	plan.Description = updatedDescription
	_, err = UpdatePlan(plan)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedPlan, err := GetPlan(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedPlan.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedPlan.Description, updatedDescription)
	}

	// Delete the object
	_, err = DeletePlan(plan)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedPlan, err := GetPlan(name)
	if err == nil || deletedPlan != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
