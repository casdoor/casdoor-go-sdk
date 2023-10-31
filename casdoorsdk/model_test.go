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

func TestModel(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Model")

	// Add a new object
	model := &Model{
		Owner:       "casbin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
		ModelText: `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act`,
	}
	_, err := AddModel(model)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	models, err := GetModels()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range models {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	model, err = GetModel(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if model.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", model.Name, name)
	}

	// Update the object
	updatedDisplayName := "UpdatedName"
	model.DisplayName = updatedDisplayName
	_, err = UpdateModel(model)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedModel, err := GetModel(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedModel.DisplayName != updatedDisplayName {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedModel.DisplayName, updatedDisplayName)
	}

	// Delete the object
	_, err = DeleteModel(model)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedModel, err := GetModel(name)
	if err != nil || deletedModel != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
