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

func TestSyncer(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Syncer")

	// Add a new object
	Syncer := &Syncer{
		Owner:        "admin",
		Name:         name,
		CreatedTime:  time.Now().Format(time.RFC3339),
		Organization: "casbin",
		Host:         "localhost",
		Port:         3306,
		User:         "root",
		Password:     "123",
		DatabaseType: "mysql",
		Database:     "syncer_db",
		Table:        "user-table",
		SyncInterval: 1,
	}
	_, err := AddSyncer(Syncer)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	Syncers, err := GetSyncers()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range Syncers {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	Syncer, err = GetSyncer(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if Syncer.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", Syncer.Name, name)
	}

	// Update the object
	updatedPassword := "123456"
	Syncer.Password = updatedPassword
	_, err = UpdateSyncer(Syncer)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedSyncer, err := GetSyncer(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedSyncer.Password != updatedPassword {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedSyncer.Password, updatedPassword)
	}

	// Delete the object
	_, err = DeleteSyncer(Syncer)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedSyncer, err := GetSyncer(name)
	if err != nil || deletedSyncer != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
