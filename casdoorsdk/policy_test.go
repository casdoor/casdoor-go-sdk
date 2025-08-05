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

func TestGetFilteredPolicies(t *testing.T) {
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

	// Add test policies for filtering tests
	testPolicy1 := &CasbinRule{
		Ptype: "g",
		V0:    "built-in/Test1",
		V1:    "group:built-in/Test1",
		V2:    "",
	}
	_, err = AddPolicy(enforcer, testPolicy1)
	if err != nil {
		t.Fatalf("Failed to add test policy 1: %v", err)
	}

	testPolicy2 := &CasbinRule{
		Ptype: "g",
		V0:    "built-in/Test2",
		V1:    "group:built-in/Test2",
		V2:    "",
	}
	_, err = AddPolicy(enforcer, testPolicy2)
	if err != nil {
		t.Fatalf("Failed to add test policy 2: %v", err)
	}

	testPolicy3 := &CasbinRule{
		Ptype: "p",
		V0:    "1",
		V1:    "2",
		V2:    "4",
	}
	_, err = AddPolicy(enforcer, testPolicy3)
	if err != nil {
		t.Fatalf("Failed to add test policy 3: %v", err)
	}
	enforcerId := globalClient.OrganizationName + "/" + name
	// Test filtered policies functionality
	fieldIndex := 0
	filters := []*PolicyFilter{
		{
			Ptype:       "g",
			FieldIndex:  &fieldIndex,
			FieldValues: []string{"built-in/Test1"},
		},
	}
	policies, err := GetFilteredPolicies(enforcerId, filters)
	if err != nil {
		t.Fatalf("GetFilteredPolicies failed: %v", err)
	}
	found := false
	for _, policy := range policies {
		if policy.Ptype == "g" && policy.V0 == "built-in/Test1" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Filtered policy not found in results")
	}
	t.Logf("Successfully retrieved %d policies", len(policies))

	// Test with fieldIndex 0 and multiple values
	fieldIndex = 0
	filters2 := []*PolicyFilter{
		{
			Ptype:       "g",
			FieldIndex:  &fieldIndex,
			FieldValues: []string{"built-in/Test1", "built-in/Test2"},
		},
	}
	policies, err = GetFilteredPolicies(enforcerId, filters2)
	if err != nil {
		t.Fatalf("GetFilteredPolicies failed: %v", err)
	}
	if len(policies) < 2 {
		t.Fatalf("Expected at least 2 policies, got %d", len(policies))
	}
	t.Logf("Successfully retrieved %d policies", len(policies))

	// Test with fieldIndex 1
	fieldIndex = 1
	filters3 := []*PolicyFilter{
		{
			Ptype:       "g",
			FieldIndex:  &fieldIndex,
			FieldValues: []string{"group:built-in/Test1"},
		},
	}
	policies, err = GetFilteredPolicies(enforcerId, filters3)
	if err != nil {
		t.Fatalf("GetFilteredPolicies failed: %v", err)
	}
	found = false
	for _, policy := range policies {
		if policy.Ptype == "g" && policy.V1 == "group:built-in/Test1" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Filtered policy not found in results")
	}
	t.Logf("Successfully retrieved %d policies", len(policies))

	// Test without fieldIndex (all policies of type)
	filters4 := []*PolicyFilter{
		{
			Ptype: "g",
		},
	}
	policies, err = GetFilteredPolicies(enforcerId, filters4)
	if err != nil {
		t.Fatalf("GetFilteredPolicies failed: %v", err)
	}
	if len(policies) < 2 {
		t.Fatalf("Expected at least 2 policies of type 'g', got %d", len(policies))
	}
	t.Logf("Successfully retrieved %d policies", len(policies))

	// Test with different ptype
	fieldIndex = 0
	fieldIndex2 := 1
	filters5 := []*PolicyFilter{
		{
			Ptype:       "p",
			FieldIndex:  &fieldIndex,
			FieldValues: []string{"1"},
		},
		{
			Ptype:       "p",
			FieldIndex:  &fieldIndex2,
			FieldValues: []string{"2"},
		},
	}
	policies, err = GetFilteredPolicies(enforcerId, filters5)
	if err != nil {
		t.Fatalf("GetFilteredPolicies failed: %v", err)
	}
	found = false
	for _, policy := range policies {
		if policy.Ptype == "p" && policy.V0 == "1" && policy.V1 == "2" {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Filtered policy not found in results")
	}
	t.Logf("Successfully retrieved %d policies", len(policies))

}
