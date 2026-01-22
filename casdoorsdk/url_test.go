// Copyright 2021 The Casdoor Authors. All Rights Reserved.
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
	"strings"
	"testing"
)

func TestGetLogoutUrl(t *testing.T) {
	client := NewClient(
		"http://localhost:8000",
		"test-client-id",
		"test-client-secret",
		"test-certificate",
		"test-org",
		"test-app",
	)

	// Test without redirect URI
	logoutUrl := client.GetLogoutUrl("")
	expected := "http://localhost:8000/api/logout"
	if logoutUrl != expected {
		t.Fatalf("GetLogoutUrl without redirect failed: expected %s, got %s", expected, logoutUrl)
	}

	// Test with redirect URI
	redirectUri := "http://example.com/home"
	logoutUrl = client.GetLogoutUrl(redirectUri)
	if !strings.HasPrefix(logoutUrl, "http://localhost:8000/api/logout?redirect_uri=") {
		t.Fatalf("GetLogoutUrl with redirect failed: expected to start with 'http://localhost:8000/api/logout?redirect_uri=', got %s", logoutUrl)
	}
	if !strings.Contains(logoutUrl, "http%3A%2F%2Fexample.com%2Fhome") {
		t.Fatalf("GetLogoutUrl redirect URI not properly encoded: %s", logoutUrl)
	}
}

func TestGetLogoutUrlGlobal(t *testing.T) {
	InitConfig(
		"http://localhost:8000",
		"test-client-id",
		"test-client-secret",
		"test-certificate",
		"test-org",
		"test-app",
	)

	// Test without redirect URI
	logoutUrl := GetLogoutUrl("")
	expected := "http://localhost:8000/api/logout"
	if logoutUrl != expected {
		t.Fatalf("GetLogoutUrl (global) without redirect failed: expected %s, got %s", expected, logoutUrl)
	}

	// Test with redirect URI
	redirectUri := "http://example.com/home"
	logoutUrl = GetLogoutUrl(redirectUri)
	if !strings.HasPrefix(logoutUrl, "http://localhost:8000/api/logout?redirect_uri=") {
		t.Fatalf("GetLogoutUrl (global) with redirect failed: expected to start with 'http://localhost:8000/api/logout?redirect_uri=', got %s", logoutUrl)
	}
	if !strings.Contains(logoutUrl, "http%3A%2F%2Fexample.com%2Fhome") {
		t.Fatalf("GetLogoutUrl (global) redirect URI not properly encoded: %s", logoutUrl)
	}
}
