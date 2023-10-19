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
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func (resource *Resource) GetId() string {
	return fmt.Sprintf("%s/%s", resource.Owner, resource.Name)
}

func TestResource(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	filename := "resource.go"
	file, err := os.Open(filename)

	if err != nil {
		t.Fatalf("Failed to open the file: %v\n", err)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read data from the file: %v\n", err)
	}

	name := fmt.Sprintf("/casdoor/%s", filename)
	// Add a new object
	resource := &Resource{
		Owner:       "casbin",
		Name:        name,
		CreatedTime: time.Now().Format(time.RFC3339),
		Description: "Casdoor Website",
		User:        "casbin",
		FileName:    filename,
		FileSize:    len(data),
		Tag:         name,
	}
	_, _, err = UploadResource(resource.User, resource.Tag, "", resource.FileName, data)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	Resources, err := GetResources(resource.Owner, resource.User, "", "", "", "")
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range Resources {
		if item.Tag == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	resource, err = GetResource(resource.GetId())
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if resource.Tag != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", resource.Name, name)
	}

	// Delete the object
	_, err = DeleteResource(resource)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedResource, err := GetResource(name)
	if err != nil || deletedResource != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
