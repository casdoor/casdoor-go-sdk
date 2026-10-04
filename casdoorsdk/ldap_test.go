// Copyright 2026 The Casdoor Authors. All Rights Reserved.
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

func TestLdap(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	id := getRandomName("Ldap")

	// An LDAP server belongs to an organization, the synced users are added to it
	ldap := &Ldap{
		Id:          id,
		CreatedTime: GetCurrentTime(),
		ServerName:  "Test LDAP Server",
		Host:        "localhost",
		Port:        389,
		Username:    "cn=admin,dc=example,dc=com",
		Password:    "password",
		BaseDn:      "dc=example,dc=com",
	}
	_, err := AddLdap(ldap)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}
	if ldap.Owner != TestCasdoorOrganization {
		t.Fatalf("The owner should default to the organization: %s != %s", ldap.Owner, TestCasdoorOrganization)
	}

	// Get all objects, check if our added object is inside the list
	ldaps, err := GetLdaps()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range ldaps {
		if item.Id == id {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	ldap, err = GetLdap(id)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if ldap.Id != id {
		t.Fatalf("Retrieved object does not match added object: %s != %s", ldap.Id, id)
	}

	// Update the object
	updatedServerName := "Updated LDAP Server"
	ldap.ServerName = updatedServerName
	_, err = UpdateLdap(ldap)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedLdap, err := GetLdap(id)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedLdap.ServerName != updatedServerName {
		t.Fatalf("Failed to update object, serverName mismatch: %s != %s", updatedLdap.ServerName, updatedServerName)
	}

	// Delete the object
	_, err = DeleteLdap(ldap)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedLdap, err := GetLdap(id)
	if err != nil || deletedLdap != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}

func TestGetOrganizationNames(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	// The organizations are owned by "admin", not by the client's organization
	organizations, err := GetOrganizationNames()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range organizations {
		if item.Name == TestCasdoorOrganization {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("The client's organization not found in list")
	}
}
