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

func TestProvider(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Provider")

	// Add a new object
	provider := &Provider{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
		Category:    "Captcha",
		Type:        "Default",
	}
	_, err := AddProvider(provider)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	providers, err := GetProviders()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range providers {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	provider, err = GetProvider(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if provider.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", provider.Name, name)
	}

	// Update the object
	updatedDisplayName := "Updated Casdoor Website"
	provider.DisplayName = updatedDisplayName
	_, err = UpdateProvider(provider)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedProvider, err := GetProvider(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedProvider.DisplayName != updatedDisplayName {
		t.Fatalf("Failed to update object, DisplayName mismatch: %s != %s", updatedProvider.DisplayName, updatedDisplayName)
	}

	// Delete the object
	_, err = DeleteProvider(provider)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedProvider, err := GetProvider(name)
	if err != nil || deletedProvider != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
