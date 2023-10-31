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

func TestApplication(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("application")

	// Add a new object
	application := &Application{
		Owner:        "admin",
		Name:         name,
		CreatedTime:  GetCurrentTime(),
		DisplayName:  name,
		Logo:         "https://cdn.casbin.org/img/casdoor-logo_1185x256.png",
		HomepageUrl:  "https://casdoor.org",
		Description:  "Casdoor Website",
		Organization: "casbin",
	}
	_, err := AddApplication(application)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	applications, err := GetApplications()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range applications {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	application, err = GetApplication(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if application.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", application.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	application.Description = updatedDescription
	_, err = UpdateApplication(application)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedApplication, err := GetApplication(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedApplication.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedApplication.Description, updatedDescription)
	}

	// Delete the object
	_, err = DeleteApplication(application)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedApplication, err := GetApplication(name)
	if err != nil || deletedApplication != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
