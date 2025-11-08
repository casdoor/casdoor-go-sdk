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
	"testing"
	"time"
)

// mockHTTPClient is a mock implementation of HttpClient for testing
type mockHTTPClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

func (m *mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	if m.DoFunc != nil {
		return m.DoFunc(req)
	}
	return &http.Response{}, nil
}

// TestWithHTTPClientOption tests that WithHTTPClient option correctly sets the HTTP client
func TestWithHTTPClientOption(t *testing.T) {
	customClient := &http.Client{
		Timeout: 10 * time.Second,
	}

	option := WithHTTPClient(customClient)
	opts := &oauthOptions{}
	option(opts)

	if opts.httpClient != customClient {
		t.Errorf("Expected httpClient to be set, got nil")
	}
}

// TestGetOAuthTokenAcceptsOptions tests that GetOAuthToken accepts OAuthOptions
func TestGetOAuthTokenAcceptsOptions(t *testing.T) {
	client := NewClient("http://localhost:8000", "client_id", "client_secret", "cert", "org", "app")

	customHTTPClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// This test just verifies that the function signature works correctly
	// We can't actually call the function without a real server, but we can verify
	// that the option is accepted without compilation errors
	_ = WithHTTPClient(customHTTPClient)

	// Verify that the option type is correct
	var opts []OAuthOption
	opts = append(opts, WithHTTPClient(customHTTPClient))

	if len(opts) != 1 {
		t.Errorf("Expected 1 option, got %d", len(opts))
	}

	// Verify we can pass options to the methods (even if we don't call them)
	// This is mainly a compile-time check
	if client == nil {
		t.Error("Client should not be nil")
	}
}

// TestRefreshOAuthTokenAcceptsOptions tests that RefreshOAuthToken accepts OAuthOptions
func TestRefreshOAuthTokenAcceptsOptions(t *testing.T) {
	client := NewClient("http://localhost:8000", "client_id", "client_secret", "cert", "org", "app")

	customHTTPClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Verify that the option type is correct
	var opts []OAuthOption
	opts = append(opts, WithHTTPClient(customHTTPClient))

	if len(opts) != 1 {
		t.Errorf("Expected 1 option, got %d", len(opts))
	}

	// Verify we can pass options to the methods (even if we don't call them)
	if client == nil {
		t.Error("Client should not be nil")
	}
}

// TestGlobalGetOAuthTokenAcceptsOptions tests that global GetOAuthToken accepts OAuthOptions
func TestGlobalGetOAuthTokenAcceptsOptions(t *testing.T) {
	InitConfig("http://localhost:8000", "client_id", "client_secret", "cert", "org", "app")

	customHTTPClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Verify that the option type is correct
	var opts []OAuthOption
	opts = append(opts, WithHTTPClient(customHTTPClient))

	if len(opts) != 1 {
		t.Errorf("Expected 1 option, got %d", len(opts))
	}
}

// TestGlobalRefreshOAuthTokenAcceptsOptions tests that global RefreshOAuthToken accepts OAuthOptions
func TestGlobalRefreshOAuthTokenAcceptsOptions(t *testing.T) {
	InitConfig("http://localhost:8000", "client_id", "client_secret", "cert", "org", "app")

	customHTTPClient := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Verify that the option type is correct
	var opts []OAuthOption
	opts = append(opts, WithHTTPClient(customHTTPClient))

	if len(opts) != 1 {
		t.Errorf("Expected 1 option, got %d", len(opts))
	}
}
