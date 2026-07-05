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

import (
	"testing"
	"time"
)

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
	affected, err := AddEnforcer(enforcer)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}
	if !affected {
		t.Fatalf("Failed to add object")
	}
	defer DeleteEnforcer(enforcer)

	oldPolicyValue := getRandomName("PolicyOld")
	newPolicyValue := getRandomName("PolicyNew")

	//Add a new policy
	policy := &CasbinRule{
		Ptype: "p",
		V0:    "1",
		V1:    "2",
		V2:    oldPolicyValue,
	}
	_, err = AddPolicy(enforcer, policy)
	if err != nil {
		t.Fatalf("Failed to add policy: %v", err)
	}
	//
	// Get all objects, check if our added object is inside the list
	found, err := waitForPolicy(name, "p", oldPolicyValue, true)
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	newPolicy := &CasbinRule{
		Ptype: "p",
		V0:    "1",
		V1:    "2",
		V2:    newPolicyValue,
	}

	// Update the object
	_, err = UpdatePolicy(enforcer, policy, newPolicy)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	found, err = waitForPolicy(name, "p", newPolicyValue, true)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
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
	found, err = waitForPolicy(name, "p", newPolicyValue, false)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if !found {
		t.Fatalf("Delete object found in list")
	}
}

func waitForPolicy(name string, ptype string, v2 string, expected bool) (bool, error) {
	var err error
	for i := 0; i < 10; i++ {
		var policies []*CasbinRule
		policies, err = GetPolicies(name, "")
		if err == nil && hasPolicy(policies, ptype, v2) == expected {
			return true, nil
		}
		time.Sleep(time.Second)
	}
	return false, err
}

func hasPolicy(policies []*CasbinRule, ptype string, v2 string) bool {
	for _, item := range policies {
		if item.Ptype == ptype && item.V2 == v2 {
			return true
		}
	}
	return false
}

func addEnforcerWithRetry(enforcer *Enforcer) error {
	var err error
	for i := 0; i < 10; i++ {
		var affected bool
		affected, err = AddEnforcer(enforcer)
		if err == nil && affected {
			return waitForEnforcer(enforcer.Name)
		}
		time.Sleep(time.Second)
	}
	return err
}

func waitForEnforcer(name string) error {
	var err error
	for i := 0; i < 10; i++ {
		var enforcer *Enforcer
		enforcer, err = GetEnforcer(name)
		if err == nil && enforcer != nil {
			return nil
		}
		time.Sleep(time.Second)
	}
	return err
}

func addPolicyWithRetry(enforcer *Enforcer, policy *CasbinRule) error {
	var err error
	for i := 0; i < 10; i++ {
		var affected bool
		affected, err = AddPolicy(enforcer, policy)
		if err == nil && affected {
			return nil
		}
		time.Sleep(time.Second)
	}
	return err
}

func waitForFilteredPolicies(enforcerId string, filters []*PolicyFilter, predicate func([]*CasbinRule) bool) ([]*CasbinRule, error) {
	var policies []*CasbinRule
	var err error
	for i := 0; i < 10; i++ {
		policies, err = GetFilteredPolicies(enforcerId, filters)
		if err == nil && predicate(policies) {
			return policies, nil
		}
		time.Sleep(time.Second)
	}
	return policies, err
}

func hasFilteredPolicy(policies []*CasbinRule, ptype string, v0 string, v1 string, v2 string) bool {
	for _, policy := range policies {
		if policy.Ptype == ptype && policy.V0 == v0 && policy.V1 == v1 && policy.V2 == v2 {
			return true
		}
	}
	return false
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
	if err := addEnforcerWithRetry(enforcer); err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}
	defer DeleteEnforcer(enforcer)

	groupUser1 := getRandomName("Test1")
	groupUser2 := getRandomName("Test2")
	groupName1 := "group:" + groupUser1
	groupName2 := "group:" + groupUser2
	policySubject := getRandomName("PolicySubject")
	policyObject := getRandomName("PolicyObject")
	policyAction := getRandomName("PolicyAction")

	var err error

	// Add test policies for filtering tests
	testPolicy1 := &CasbinRule{
		Ptype: "g",
		V0:    groupUser1,
		V1:    groupName1,
		V2:    "",
	}
	if err = addPolicyWithRetry(enforcer, testPolicy1); err != nil {
		t.Fatalf("Failed to add test policy 1: %v", err)
	}

	testPolicy2 := &CasbinRule{
		Ptype: "g",
		V0:    groupUser2,
		V1:    groupName2,
		V2:    "",
	}
	if err = addPolicyWithRetry(enforcer, testPolicy2); err != nil {
		t.Fatalf("Failed to add test policy 2: %v", err)
	}

	testPolicy3 := &CasbinRule{
		Ptype: "p",
		V0:    policySubject,
		V1:    policyObject,
		V2:    policyAction,
	}
	if err = addPolicyWithRetry(enforcer, testPolicy3); err != nil {
		t.Fatalf("Failed to add test policy 3: %v", err)
	}
	enforcerId := globalClient.OrganizationName + "/" + name
	// Test filtered policies functionality
	fieldIndex := 0
	filters := []*PolicyFilter{
		{
			Ptype:       "g",
			FieldIndex:  &fieldIndex,
			FieldValues: []string{groupUser1},
		},
	}
	policies, err := waitForFilteredPolicies(enforcerId, filters, func(policies []*CasbinRule) bool {
		return hasFilteredPolicy(policies, "g", groupUser1, groupName1, "")
	})
	if err != nil {
		t.Fatalf("GetFilteredPolicies failed: %v", err)
	}
	if !hasFilteredPolicy(policies, "g", groupUser1, groupName1, "") {
		t.Fatalf("Filtered policy not found in results")
	}
	t.Logf("Successfully retrieved %d policies", len(policies))

	// Test with fieldIndex 0 and multiple values
	fieldIndex = 0
	filters2 := []*PolicyFilter{
		{
			Ptype:       "g",
			FieldIndex:  &fieldIndex,
			FieldValues: []string{groupUser1, groupUser2},
		},
	}
	policies, err = waitForFilteredPolicies(enforcerId, filters2, func(policies []*CasbinRule) bool {
		return hasFilteredPolicy(policies, "g", groupUser1, groupName1, "") &&
			hasFilteredPolicy(policies, "g", groupUser2, groupName2, "")
	})
	if err != nil {
		t.Fatalf("GetFilteredPolicies failed: %v", err)
	}
	if !hasFilteredPolicy(policies, "g", groupUser1, groupName1, "") || !hasFilteredPolicy(policies, "g", groupUser2, groupName2, "") {
		t.Fatalf("Expected filtered policies not found in results")
	}
	t.Logf("Successfully retrieved %d policies", len(policies))

	// Test with fieldIndex 1
	fieldIndex = 1
	filters3 := []*PolicyFilter{
		{
			Ptype:       "g",
			FieldIndex:  &fieldIndex,
			FieldValues: []string{groupName1},
		},
	}
	policies, err = waitForFilteredPolicies(enforcerId, filters3, func(policies []*CasbinRule) bool {
		return hasFilteredPolicy(policies, "g", groupUser1, groupName1, "")
	})
	if err != nil {
		t.Fatalf("GetFilteredPolicies failed: %v", err)
	}
	if !hasFilteredPolicy(policies, "g", groupUser1, groupName1, "") {
		t.Fatalf("Filtered policy not found in results")
	}
	t.Logf("Successfully retrieved %d policies", len(policies))

	// Test without fieldIndex (all policies of type)
	filters4 := []*PolicyFilter{
		{
			Ptype: "g",
		},
	}
	policies, err = waitForFilteredPolicies(enforcerId, filters4, func(policies []*CasbinRule) bool {
		return hasFilteredPolicy(policies, "g", groupUser1, groupName1, "") &&
			hasFilteredPolicy(policies, "g", groupUser2, groupName2, "")
	})
	if err != nil {
		t.Fatalf("GetFilteredPolicies failed: %v", err)
	}
	if !hasFilteredPolicy(policies, "g", groupUser1, groupName1, "") || !hasFilteredPolicy(policies, "g", groupUser2, groupName2, "") {
		t.Fatalf("Expected filtered policies of type 'g' not found in results")
	}
	t.Logf("Successfully retrieved %d policies", len(policies))

	// Test with different ptype
	fieldIndex = 0
	fieldIndex2 := 1
	filters5 := []*PolicyFilter{
		{
			Ptype:       "p",
			FieldIndex:  &fieldIndex,
			FieldValues: []string{policySubject},
		},
		{
			Ptype:       "p",
			FieldIndex:  &fieldIndex2,
			FieldValues: []string{policyObject},
		},
	}
	policies, err = waitForFilteredPolicies(enforcerId, filters5, func(policies []*CasbinRule) bool {
		return hasFilteredPolicy(policies, "p", policySubject, policyObject, policyAction)
	})
	if err != nil {
		t.Fatalf("GetFilteredPolicies failed: %v", err)
	}
	if !hasFilteredPolicy(policies, "p", policySubject, policyObject, policyAction) {
		t.Fatalf("Filtered policy not found in results")
	}
	t.Logf("Successfully retrieved %d policies", len(policies))

}
