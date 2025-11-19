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

import "golang.org/x/oauth2"

func GetOAuthToken(code string, state string, opts ...OAuthOption) (*oauth2.Token, error) {
	return globalClient.GetOAuthToken(code, state, opts...)
}

func RefreshOAuthToken(refreshToken string, opts ...OAuthOption) (*oauth2.Token, error) {
	return globalClient.RefreshOAuthToken(refreshToken, opts...)
}

// SetCustomHeaders sets custom headers for the global client.
// These headers will be included in all API requests.
// Common use cases include setting Accept-Language, custom tenant headers, or trace IDs.
func SetCustomHeaders(headers map[string]string) {
	globalClient.SetCustomHeaders(headers)
}

// SetCustomHeader sets a single custom header for the global client.
func SetCustomHeader(key, value string) {
	globalClient.SetCustomHeader(key, value)
}

// GetCustomHeaders returns a copy of the custom headers from the global client.
func GetCustomHeaders() map[string]string {
	return globalClient.GetCustomHeaders()
}

// ClearCustomHeaders removes all custom headers from the global client.
func ClearCustomHeaders() {
	globalClient.ClearCustomHeaders()
}
