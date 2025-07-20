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

func TestSubscription(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Subscription")

	// Add a new object
	subscription := &Subscription{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		StartTime:   time.Now().AddDate(-1, 0, 0), // 1 year ago
		DisplayName: name,
		Description: "Casdoor Website",
	}
	_, err := AddSubscription(subscription)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	subscriptions, err := GetSubscriptions()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range subscriptions {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	subscription, err = GetSubscription(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if subscription.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", subscription.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	subscription.Description = updatedDescription
	_, err = UpdateSubscription(subscription)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedSubscription, err := GetSubscription(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedSubscription.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedSubscription.Description, updatedDescription)
	}

	// Delete the object
	_, err = DeleteSubscription(subscription)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedSubscription, err := GetSubscription(name)
	if err != nil || deletedSubscription != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
