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
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestModifyResource checks the requests sent by AddResource() and UpdateResource()
// against a fake server, because the real add-resource API is restricted to global admins.
func TestModifyResource(t *testing.T) {
	var gotPath, gotId string
	var gotResource Resource

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotId = r.URL.Query().Get("id")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Failed to read request body: %v", err)
		}
		if err = json.Unmarshal(body, &gotResource); err != nil {
			t.Errorf("Failed to parse request body: %v", err)
		}

		_, _ = w.Write([]byte(`{"status":"ok","msg":"","data":"Affected"}`))
	}))
	defer server.Close()

	c := NewClient(server.URL, "clientId", "clientSecret", "", "casbin", "app")

	tests := []struct {
		name   string
		action string
		call   func(*Resource) (bool, error)
	}{
		{"add", "add-resource", c.AddResource},
		{"update", "update-resource", c.UpdateResource},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			resource := &Resource{Name: "avatar.png", User: "alice", Description: "Avatar"}
			affected, err := tt.call(resource)
			if err != nil {
				t.Fatalf("Failed to %s resource: %v", tt.name, err)
			}
			if !affected {
				t.Fatalf("Expected the %s to be affected", tt.name)
			}

			if gotPath != "/api/"+tt.action {
				t.Fatalf("Unexpected path: %s != /api/%s", gotPath, tt.action)
			}
			if gotId != "casbin/avatar.png" {
				t.Fatalf("Unexpected id: %s != casbin/avatar.png", gotId)
			}
			if gotResource.Owner != "casbin" || gotResource.Name != "avatar.png" || gotResource.Description != "Avatar" {
				t.Fatalf("Unexpected request body: %+v", gotResource)
			}
		})
	}

	// the owner set by the caller is kept
	resource := &Resource{Owner: "other", Name: "avatar.png"}
	if _, err := c.UpdateResource(resource); err != nil {
		t.Fatalf("Failed to update resource: %v", err)
	}
	if gotId != "other/avatar.png" || gotResource.Owner != "other" {
		t.Fatalf("Expected the caller-provided owner to be kept, got id %s and owner %s", gotId, gotResource.Owner)
	}
}
