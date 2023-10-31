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

func TestCert(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("cert")

	// Add a new object
	cert := &Cert{
		Owner:           "admin",
		Name:            name,
		CreatedTime:     GetCurrentTime(),
		DisplayName:     name,
		Scope:           "JWT",
		Type:            "x509",
		CryptoAlgorithm: "RS256",
		BitSize:         4096,
		ExpireInYears:   20,
	}
	_, err := AddCert(cert)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	certs, err := GetCerts()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range certs {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	cert, err = GetCert(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if cert.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", cert.Name, name)
	}

	// Update the object
	updatedDisplayName := "Updated Casdoor Website"
	cert.DisplayName = updatedDisplayName
	_, err = UpdateCert(cert)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedcert, err := GetCert(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedcert.DisplayName != updatedDisplayName {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedcert.DisplayName, updatedDisplayName)
	}

	// Delete the object
	_, err = DeleteCert(cert)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedcert, err := GetCert(name)
	if err != nil || deletedcert != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
