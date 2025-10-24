// Copyright 2025 The Casdoor Authors. All Rights Reserved.
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
	"testing"
)

// TestMfaRequestMarshal tests that MfaRequest can be properly marshaled to JSON
// This ensures the Delete function will send valid JSON body
func TestMfaRequestMarshal(t *testing.T) {
	req := MfaRequest{
		Owner: "test-owner",
		Name:  "test-user",
	}

	jsonBytes, err := json.Marshal(req)
	if err != nil {
		t.Fatalf("Failed to marshal MfaRequest: %v", err)
	}

	// Verify the JSON contains expected fields
	var unmarshaled MfaRequest
	err = json.Unmarshal(jsonBytes, &unmarshaled)
	if err != nil {
		t.Fatalf("Failed to unmarshal MfaRequest: %v", err)
	}

	if unmarshaled.Owner != "test-owner" {
		t.Errorf("Owner mismatch: got %s, want test-owner", unmarshaled.Owner)
	}

	if unmarshaled.Name != "test-user" {
		t.Errorf("Name mismatch: got %s, want test-user", unmarshaled.Name)
	}
}
