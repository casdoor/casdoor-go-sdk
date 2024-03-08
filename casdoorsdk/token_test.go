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

func TestToken(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Token")

	// Add a new object
	token := &Token{
		Owner:        "admin",
		Name:         name,
		CreatedTime:  GetCurrentTime(),
		Organization: "casbin",
		Code:         "abc",
		AccessToken:  "123456",
	}
	_, err := AddToken(token)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	tokens, err := GetTokens()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range tokens {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	token, err = GetToken(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if token.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", token.Name, name)
	}

	// Update the object
	updatedCode := "Updated Code"
	token.Code = updatedCode
	_, err = UpdateToken(token)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedToken, err := GetToken(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedToken.Code != updatedCode {
		t.Fatalf("Failed to update object, code mismatch: %s != %s", updatedToken.Code, updatedCode)
	}

	// Delete the object
	_, err = DeleteToken(token)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedToken, err := GetToken(name)
	if err != nil || deletedToken != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
