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

func TestWebhook(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Webhook")

	// Add a new object
	Webhook := &Webhook{
		Owner:        "casbin",
		Name:         name,
		CreatedTime:  GetCurrentTime(),
		Organization: "casbin",
	}
	_, err := AddWebhook(Webhook)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	Webhooks, err := GetWebhooks()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range Webhooks {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	Webhook, err = GetWebhook(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if Webhook.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", Webhook.Name, name)
	}

	// Update the object
	updatedOrganization := "admin"
	Webhook.Organization = updatedOrganization
	_, err = UpdateWebhook(Webhook)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedWebhook, err := GetWebhook(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedWebhook.Organization != updatedOrganization {
		t.Fatalf("Failed to update object, Port mismatch: %s != %s", updatedWebhook.Organization, updatedOrganization)
	}

	// Delete the object
	_, err = DeleteWebhook(Webhook)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedWebhook, err := GetWebhook(name)
	if err != nil || deletedWebhook != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
