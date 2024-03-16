// Copyright 2024 The Casdoor Authors. All Rights Reserved.
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

import "testing"

func TestPolicy(t *testing.T) {
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

	//Add a new policy
	policy := &CasbinRule{
		Ptype: "p",
		V0:    "1",
		V1:    "2",
		V2:    "4",
	}
	_, err = AddPolicy(enforcer, policy)
	if err != nil {
		t.Fatalf("Failed to add policy: %v", err)
	}
	//
	// Get all objects, check if our added object is inside the list
	Policies, err := GetPolicies(name, "")
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range Policies {
		if item.Ptype == "p" && item.V2 == "4" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	newPolicy := &CasbinRule{
		Ptype: "p",
		V0:    "1",
		V1:    "2",
		V2:    "5",
	}

	// Update the object
	_, err = UpdatePolicy(enforcer, policy, newPolicy)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	Policies, err = GetPolicies(name, "")
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	found = false
	for _, item := range Policies {
		if item.Ptype == "p" && item.V2 == "5" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Update object not found in list")
	}

	// Delete the object
	_, err = RemovePolicy(enforcer, newPolicy)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	Policies, err = GetPolicies(name, "")
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	found = false
	for _, item := range Policies {
		if item.Ptype == "p" && item.V2 == "3" {
			found = true
			break
		}
	}
	if found {
		t.Fatalf("Delete object found in list")
	}

}
