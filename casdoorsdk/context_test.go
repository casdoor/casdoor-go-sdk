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
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"
)

// These tests don't need a Casdoor server: they run against a local fake server.

type contextTestKey struct{}

// newContextTestClient starts a fake Casdoor server that answers every API with "ok" and
// returns a client pointing to it, along with the number of requests the server received.
func newContextTestClient(t *testing.T) (*Client, *int32) {
	t.Helper()

	var hits int32
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		atomic.AddInt32(&hits, 1)
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"status":"ok","data":[]}`))
	}))
	t.Cleanup(server.Close)

	return NewClient(server.URL, "client-id", "client-secret", "", "org", "app"), &hits
}

func cancelledContext() context.Context {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	return ctx
}

// recordingHttpClient is a custom http client (see SetHttpClient) that records the
// context of the last request it sent.
type recordingHttpClient struct {
	lastCtx context.Context
}

func (r *recordingHttpClient) Do(req *http.Request) (*http.Response, error) {
	r.lastCtx = req.Context()
	return http.DefaultClient.Do(req)
}

func useRecordingHttpClient(t *testing.T) *recordingHttpClient {
	t.Helper()

	previous := client
	t.Cleanup(func() { SetHttpClient(previous) })

	recorder := &recordingHttpClient{}
	SetHttpClient(recorder)
	return recorder
}

func TestWithContextCancelledGet(t *testing.T) {
	c, hits := newContextTestClient(t)

	_, err := c.WithContext(cancelledContext()).GetUsers()
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("GetUsers() error = %v, want context.Canceled", err)
	}
	if n := atomic.LoadInt32(hits); n != 0 {
		t.Fatalf("the server received %d requests, want 0", n)
	}
}

func TestWithContextCancelledPost(t *testing.T) {
	c, hits := newContextTestClient(t)

	_, err := c.WithContext(cancelledContext()).AddUser(&User{Name: "alice"})
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("AddUser() error = %v, want context.Canceled", err)
	}
	if n := atomic.LoadInt32(hits); n != 0 {
		t.Fatalf("the server received %d requests, want 0", n)
	}
}

func TestWithContextDeadlineAbortsRequestInFlight(t *testing.T) {
	// the server answers too late, unless the request is aborted by the client before
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		select {
		case <-r.Context().Done():
		case <-time.After(5 * time.Second):
			_, _ = w.Write([]byte(`{"status":"ok","data":[]}`))
		}
	}))
	t.Cleanup(server.Close)

	c := NewClient(server.URL, "client-id", "client-secret", "", "org", "app")

	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel()

	start := time.Now()
	_, err := c.WithContext(ctx).GetUsers()
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("GetUsers() error = %v, want context.DeadlineExceeded", err)
	}
	if elapsed := time.Since(start); elapsed > 2*time.Second {
		t.Fatalf("GetUsers() returned after %v, it should have been aborted at the deadline", elapsed)
	}
}

func TestWithContextCancelledLogout(t *testing.T) {
	c, _ := newContextTestClient(t)

	err := c.WithContext(cancelledContext()).Logout("access-token")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("Logout() error = %v, want context.Canceled", err)
	}
}

func TestWithContextCancelledOAuth(t *testing.T) {
	c, hits := newContextTestClient(t)
	cc := c.WithContext(cancelledContext())

	if _, err := cc.GetOAuthToken("code", "state"); !errors.Is(err, context.Canceled) {
		t.Fatalf("GetOAuthToken() error = %v, want context.Canceled", err)
	}
	if _, err := cc.RefreshOAuthToken("refresh-token"); !errors.Is(err, context.Canceled) {
		t.Fatalf("RefreshOAuthToken() error = %v, want context.Canceled", err)
	}
	if _, err := cc.GetOAuthTokenByPassword("alice", "password"); !errors.Is(err, context.Canceled) {
		t.Fatalf("GetOAuthTokenByPassword() error = %v, want context.Canceled", err)
	}

	// WithHTTPClient() must not drop the client's context
	_, err := cc.GetOAuthToken("code", "state", WithHTTPClient(&http.Client{}))
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("GetOAuthToken(WithHTTPClient()) error = %v, want context.Canceled", err)
	}

	if n := atomic.LoadInt32(hits); n != 0 {
		t.Fatalf("the server received %d requests, want 0", n)
	}
}

func TestWithContextPropagatesValues(t *testing.T) {
	c, _ := newContextTestClient(t)
	recorder := useRecordingHttpClient(t)

	ctx := context.WithValue(context.Background(), contextTestKey{}, "trace-123")
	if _, err := c.WithContext(ctx).GetUsers(); err != nil {
		t.Fatalf("GetUsers() error = %v", err)
	}
	if got := recorder.lastCtx.Value(contextTestKey{}); got != "trace-123" {
		t.Fatalf("request context value = %v, want %q", got, "trace-123")
	}

	// without WithContext(), the requests keep using context.Background()
	if _, err := c.GetUsers(); err != nil {
		t.Fatalf("GetUsers() error = %v", err)
	}
	if got := recorder.lastCtx.Value(contextTestKey{}); got != nil {
		t.Fatalf("request context value = %v, want nil", got)
	}
}

func TestWithContextDoesNotChangeOriginalClient(t *testing.T) {
	c, _ := newContextTestClient(t)
	c.CustomHeaders["X-Original"] = "1"

	cc := c.WithContext(cancelledContext())
	cc.CustomHeaders["X-Copy"] = "1"

	if _, ok := c.CustomHeaders["X-Copy"]; ok {
		t.Fatal("changing the headers of the copy changed the original client")
	}
	if cc.CustomHeaders["X-Original"] != "1" {
		t.Fatal("the copy lost the custom headers of the original client")
	}

	// the original client still works, although the copy's context is cancelled
	if _, err := c.GetUsers(); err != nil {
		t.Fatalf("GetUsers() on the original client error = %v", err)
	}
}

func TestWithContextAndWithAccessTokenCompose(t *testing.T) {
	c, _ := newContextTestClient(t)
	ctx := context.WithValue(context.Background(), contextTestKey{}, "value")

	for name, cc := range map[string]*Client{
		"WithContext().WithAccessToken()": c.WithContext(ctx).WithAccessToken("token"),
		"WithAccessToken().WithContext()": c.WithAccessToken("token").WithContext(ctx),
	} {
		if cc.AccessToken != "token" {
			t.Errorf("%s: AccessToken = %q, want %q", name, cc.AccessToken, "token")
		}
		if cc.getContext() != ctx {
			t.Errorf("%s: the context was lost", name)
		}
	}

	if c.AccessToken != "" || c.ctx != nil {
		t.Fatal("the original client was changed")
	}
}

func TestWithContextGlobal(t *testing.T) {
	c, _ := newContextTestClient(t)

	previous := globalClient
	t.Cleanup(func() { globalClient = previous })
	globalClient = c

	_, err := WithContext(cancelledContext()).GetUsers()
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("WithContext().GetUsers() error = %v, want context.Canceled", err)
	}

	// the global client itself is not changed
	if _, err := GetUsers(); err != nil {
		t.Fatalf("GetUsers() error = %v", err)
	}
}

func TestWithContextNilPanics(t *testing.T) {
	c, _ := newContextTestClient(t)

	defer func() {
		if recover() == nil {
			t.Fatal("WithContext(nil) didn't panic")
		}
	}()

	//nolint:staticcheck // passing a nil context on purpose
	c.WithContext(nil)
}
