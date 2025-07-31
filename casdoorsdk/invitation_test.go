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
	"testing"
	"time"
)

func TestInvitation(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("unit_test_invitation")
	code := "TEST1234"
	// Test invitation object
	invitation := &Invitation{
		Owner:       TestCasdoorOrganization,
		Name:        name,
		CreatedTime: time.Now().Format(time.RFC3339),
		DisplayName: "Test Invitation",
		Code:        code,
		DefaultCode: code,
		Quota:       10,
		UsedCount:   0,
		Application: TestCasdoorApplication,
		Email:       "test@example.com",
		SignupGroup: "test-group",
		State:       "Active",
	}

	// Test AddInvitation
	_, err := AddInvitation(invitation)
	if err != nil {
		t.Fatalf("Failed to add invitation: %v", err)
	}

	// Test GetInvitation
	invitation2, err := GetInvitation(name)
	if err != nil {
		t.Fatalf("Failed to get invitation: %v", err)
	}
	if invitation2.Code != invitation.Code {
		t.Fatalf("Retrieved invitation does not match added invitation")
	}

	// Test GetInvitations
	invitations, err := GetInvitations()
	if err != nil {
		t.Fatalf("Failed to get invitations: %v", err)
	}
	if len(invitations) == 0 {
		t.Fatalf("No invitations found")
	}

	// Test UpdateInvitation
	invitation2.State = "Suspended"
	_, err = UpdateInvitation(invitation2)
	if err != nil {
		t.Fatalf("Failed to update invitation: %v", err)
	}

	// Test UpdateInvitation to Active to check getInvitationInfo
	invitation2.State = "Active"
	_, err = UpdateInvitation(invitation2)
	if err != nil {
		t.Fatalf("Failed to update invitation: %v", err)
	}

	// Test GetInvitationInfo
	invitation, err = GetInvitationInfo(code, TestCasdoorApplication)
	if err != nil {
		t.Fatalf("Failed to get invitation info by code: %v", err)
	}
	if invitation == nil {
		t.Fatalf("Invitation not found by code")
	}

	// Test DeleteInvitation
	_, err = DeleteInvitation(invitation2)
	if err != nil {
		t.Fatalf("Failed to delete invitation: %v", err)
	}

	// Verify deletion
	deletedInvitation, err := GetInvitation(name)
	if err == nil && deletedInvitation != nil {
		t.Fatalf("Failed to delete invitation, it still exists")
	}
}
