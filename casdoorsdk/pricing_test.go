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

func TestPricing(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Pricing")

	// Add a new object
	pricing := &Pricing{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
		Application: "app-admin",
		Description: "Casdoor Website",
	}
	_, err := AddPricing(pricing)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	pricings, err := GetPricings()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range pricings {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	pricing, err = GetPricing(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if pricing.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", pricing.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	pricing.Description = updatedDescription
	_, err = UpdatePricing(pricing)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedPricing, err := GetPricing(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedPricing.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedPricing.Description, updatedDescription)
	}

	// Delete the object
	_, err = DeletePricing(pricing)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedPricing, err := GetPricing(name)
	if err != nil || deletedPricing != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
