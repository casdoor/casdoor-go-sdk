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

func TestMfaRequestWithRecoveryCodes(t *testing.T) {
	// Test that MfaRequest properly marshals with recoveryCodes
	mfaReq := MfaRequest{
		Owner:         "test-owner",
		MfaType:       "app",
		Name:          "test-user",
		Secret:        "test-secret",
		RecoveryCodes: "code1,code2,code3",
	}

	jsonBytes, err := json.Marshal(mfaReq)
	if err != nil {
		t.Fatalf("Failed to marshal MfaRequest: %v", err)
	}

	var result map[string]string
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal MfaRequest: %v", err)
	}

	// Verify all fields are present
	if result["owner"] != "test-owner" {
		t.Errorf("Expected owner to be 'test-owner', got '%s'", result["owner"])
	}
	if result["mfaType"] != "app" {
		t.Errorf("Expected mfaType to be 'app', got '%s'", result["mfaType"])
	}
	if result["name"] != "test-user" {
		t.Errorf("Expected name to be 'test-user', got '%s'", result["name"])
	}
	if result["secret"] != "test-secret" {
		t.Errorf("Expected secret to be 'test-secret', got '%s'", result["secret"])
	}
	if result["recoveryCodes"] != "code1,code2,code3" {
		t.Errorf("Expected recoveryCodes to be 'code1,code2,code3', got '%s'", result["recoveryCodes"])
	}
}

func TestMfaRequestWithoutRecoveryCodes(t *testing.T) {
	// Test that MfaRequest properly omits recoveryCodes when empty
	mfaReq := MfaRequest{
		Owner:   "test-owner",
		MfaType: "app",
		Name:    "test-user",
		Secret:  "test-secret",
	}

	jsonBytes, err := json.Marshal(mfaReq)
	if err != nil {
		t.Fatalf("Failed to marshal MfaRequest: %v", err)
	}

	var result map[string]interface{}
	err = json.Unmarshal(jsonBytes, &result)
	if err != nil {
		t.Fatalf("Failed to unmarshal MfaRequest: %v", err)
	}

	// Verify recoveryCodes is omitted when empty
	if _, exists := result["recoveryCodes"]; exists {
		t.Error("Expected recoveryCodes to be omitted when empty")
	}
}
