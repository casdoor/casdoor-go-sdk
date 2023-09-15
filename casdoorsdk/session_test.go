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

func TestSession(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Session")

	// Add a new object
	Session := &Session{
		Owner:        "casbin",
		Name:         name,
		CreatedTime:  time.Now().Format(time.RFC3339),
		Application:  "app-built-in",
		SessionId:    []string{},
	}
	_, err := AddSession(Session)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	Sessions, err := GetSessions()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range Sessions {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	Session, err = GetSession(name, Session.Application)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if Session.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", Session.Name, name)
	}

	// Update the object
	UpdateTime := time.Now().Format(time.RFC3339)
	Session.CreatedTime = UpdateTime
	_, err = UpdateSession(Session)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedSession, err := GetSession(name, Session.Application)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedSession.CreatedTime != UpdateTime {
		t.Fatalf("Failed to update object, Application mismatch: %s != %s", updatedSession.CreatedTime, UpdateTime)
	}

	// Delete the object
	_, err = DeleteSession(Session)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedSession, err := GetSession(name, Session.Application)
	if err != nil || deletedSession != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
