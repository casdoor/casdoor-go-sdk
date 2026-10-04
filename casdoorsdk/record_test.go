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

import "testing"

func TestRecord(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Record")

	// Add a new object
	record := &Record{
		Owner:        TestCasdoorOrganization,
		Name:         name,
		CreatedTime:  GetCurrentTime(),
		Organization: TestCasdoorOrganization,
		User:         "admin",
		Action:       "test-record",
	}
	_, err := AddRecord(record)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Reading the records needs the access token of an admin user
	token, err := GetOAuthTokenByPassword("admin", "123")
	if err != nil {
		t.Fatalf("Failed to get the access token: %v", err)
	}
	adminClient := WithAccessToken(token.AccessToken)

	// Get all objects, check if our added object is inside the list
	records, err := adminClient.GetRecords()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range records {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	record, err = adminClient.GetRecord(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if record == nil || record.Name != name {
		t.Fatalf("Retrieved object does not match added object: %v != %s", record, name)
	}

	// Get an object that doesn't exist
	record, err = adminClient.GetRecord(name + "_missing")
	if err != nil || record != nil {
		t.Fatalf("Failed to get a missing object: %v, %v", record, err)
	}
}
