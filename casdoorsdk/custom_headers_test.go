// Copyright 2024 The Casdoor Authors. All Rights Reserved.
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
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCustomHeaders(t *testing.T) {
	// Create a test server that echoes back the headers it received
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Check for custom headers
		acceptLanguage := r.Header.Get("Accept-Language")
		customTenant := r.Header.Get("X-Tenant-ID")
		traceID := r.Header.Get("X-Trace-ID")

		// Write response with status 200
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")

		// Echo back headers for verification
		response := `{"status":"ok","msg":"success","data":{"Accept-Language":"` + acceptLanguage +
			`","X-Tenant-ID":"` + customTenant +
			`","X-Trace-ID":"` + traceID + `"}}`
		w.Write([]byte(response))
	}))
	defer server.Close()

	// Create a client with custom headers
	client := NewClient(
		server.URL,
		"test-client-id",
		"test-client-secret",
		"test-cert",
		"test-org",
		"test-app",
	)

	// Test setting custom headers
	client.SetCustomHeader("Accept-Language", "de")
	client.SetCustomHeader("X-Tenant-ID", "tenant-123")
	client.SetCustomHeader("X-Trace-ID", "trace-abc-123")

	// Verify headers were set
	headers := client.GetCustomHeaders()
	if headers["Accept-Language"] != "de" {
		t.Errorf("Accept-Language header not set correctly: got %s, want de", headers["Accept-Language"])
	}
	if headers["X-Tenant-ID"] != "tenant-123" {
		t.Errorf("X-Tenant-ID header not set correctly: got %s, want tenant-123", headers["X-Tenant-ID"])
	}
	if headers["X-Trace-ID"] != "trace-abc-123" {
		t.Errorf("X-Trace-ID header not set correctly: got %s, want trace-abc-123", headers["X-Trace-ID"])
	}

	// Make a request to verify headers are sent
	resp, err := client.DoGetResponse(server.URL + "/test")
	if err != nil {
		t.Fatalf("Request failed: %v", err)
	}

	if resp.Status != "ok" {
		t.Errorf("Response status not ok: %s", resp.Status)
	}
}

func TestSetCustomHeaders(t *testing.T) {
	client := NewClient(
		"http://localhost:8000",
		"test-client-id",
		"test-client-secret",
		"test-cert",
		"test-org",
		"test-app",
	)

	// Test setting multiple headers at once
	headers := map[string]string{
		"Accept-Language": "fr",
		"X-Custom-Header": "custom-value",
		"X-API-Version":   "v2",
	}
	client.SetCustomHeaders(headers)

	// Verify headers were set
	retrievedHeaders := client.GetCustomHeaders()
	if len(retrievedHeaders) != 3 {
		t.Errorf("Expected 3 headers, got %d", len(retrievedHeaders))
	}
	if retrievedHeaders["Accept-Language"] != "fr" {
		t.Errorf("Accept-Language not set correctly")
	}
	if retrievedHeaders["X-Custom-Header"] != "custom-value" {
		t.Errorf("X-Custom-Header not set correctly")
	}
	if retrievedHeaders["X-API-Version"] != "v2" {
		t.Errorf("X-API-Version not set correctly")
	}
}

func TestClearCustomHeaders(t *testing.T) {
	client := NewClient(
		"http://localhost:8000",
		"test-client-id",
		"test-client-secret",
		"test-cert",
		"test-org",
		"test-app",
	)

	// Set some headers
	client.SetCustomHeader("Accept-Language", "es")
	client.SetCustomHeader("X-Test", "test")

	// Verify headers were set
	headers := client.GetCustomHeaders()
	if len(headers) != 2 {
		t.Errorf("Expected 2 headers before clear, got %d", len(headers))
	}

	// Clear headers
	client.ClearCustomHeaders()

	// Verify headers were cleared
	headers = client.GetCustomHeaders()
	if len(headers) != 0 {
		t.Errorf("Expected 0 headers after clear, got %d", len(headers))
	}
}

func TestGlobalCustomHeaders(t *testing.T) {
	// Initialize global config
	InitConfig(
		"http://localhost:8000",
		"test-client-id",
		"test-client-secret",
		"test-cert",
		"test-org",
		"test-app",
	)

	// Test setting global custom headers
	SetCustomHeader("Accept-Language", "ja")
	SetCustomHeader("X-Global-Header", "global-value")

	// Verify headers were set
	headers := GetCustomHeaders()
	if headers["Accept-Language"] != "ja" {
		t.Errorf("Global Accept-Language not set correctly")
	}
	if headers["X-Global-Header"] != "global-value" {
		t.Errorf("Global X-Global-Header not set correctly")
	}

	// Test setting multiple headers
	SetCustomHeaders(map[string]string{
		"X-Another-Header": "another-value",
		"X-API-Key":        "key-123",
	})

	headers = GetCustomHeaders()
	if len(headers) != 4 {
		t.Errorf("Expected 4 headers, got %d", len(headers))
	}

	// Clear headers
	ClearCustomHeaders()
	headers = GetCustomHeaders()
	if len(headers) != 0 {
		t.Errorf("Expected 0 headers after clear, got %d", len(headers))
	}
}

func TestCustomHeadersInPostRequest(t *testing.T) {
	// Create a test server that echoes back the headers it received
	receivedHeaders := make(map[string]string)
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Store received headers
		receivedHeaders["Accept-Language"] = r.Header.Get("Accept-Language")
		receivedHeaders["X-Custom-Post"] = r.Header.Get("X-Custom-Post")

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok","msg":"success","data":null}`))
	}))
	defer server.Close()

	// Create a client with custom headers
	client := NewClient(
		server.URL,
		"test-client-id",
		"test-client-secret",
		"test-cert",
		"test-org",
		"test-app",
	)

	// Set custom headers
	client.SetCustomHeader("Accept-Language", "pt")
	client.SetCustomHeader("X-Custom-Post", "post-value")

	// Make a POST request via DoPost method which is safer
	postData := []byte(`{"test":"data"}`)
	queryMap := make(map[string]string)
	_, err := client.DoPost("test", queryMap, postData, false, false)
	if err != nil {
		t.Fatalf("POST request failed: %v", err)
	}

	// Verify headers were received
	if receivedHeaders["Accept-Language"] != "pt" {
		t.Errorf("Accept-Language header not received correctly: got %s, want pt", receivedHeaders["Accept-Language"])
	}
	if receivedHeaders["X-Custom-Post"] != "post-value" {
		t.Errorf("X-Custom-Post header not received correctly: got %s, want post-value", receivedHeaders["X-Custom-Post"])
	}
}
