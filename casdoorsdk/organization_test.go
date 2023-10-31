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

func TestOrganization(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Organization")

	// Add a new object
	organization := &Organization{
		Owner:              "admin",
		Name:               name,
		CreatedTime:        GetCurrentTime(),
		DisplayName:        name,
		WebsiteUrl:         "https://example.com",
		PasswordType:       "plain",
		PasswordOptions:    []string{"AtLeast6"},
		CountryCodes:       []string{"US", "ES", "FR", "DE", "GB", "CN", "JP", "KR", "VN", "ID", "SG", "IN"},
		Tags:               []string{},
		Languages:          []string{"en", "zh", "es", "fr", "de", "id", "ja", "ko", "ru", "vi", "pt"},
		InitScore:          2000,
		EnableSoftDeletion: false,
		IsProfilePublic:    false,
	}
	_, err := AddOrganization(organization)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	//organizations, err := GetOrganizations()
	//if err != nil {
	//	t.Fatalf("Failed to get objects: %v", err)
	//}
	//found := false
	//for _, item := range organizations {
	//	if item.Name == name {
	//		found = true
	//		break
	//	}
	//}
	//if !found {
	//	t.Fatalf("Added object not found in list")
	//}

	// Get the object
	organization, err = GetOrganization(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if organization.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", organization.Name, name)
	}

	// Update the object
	updatedDisplayName := "Updated Casdoor Website"
	organization.DisplayName = updatedDisplayName
	_, err = UpdateOrganization(organization)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedOrganization, err := GetOrganization(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedOrganization.DisplayName != updatedDisplayName {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedOrganization.DisplayName, updatedDisplayName)
	}

	// Delete the object
	_, err = DeleteOrganization(organization)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedOrganization, err := GetOrganization(name)
	if err != nil || deletedOrganization != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
