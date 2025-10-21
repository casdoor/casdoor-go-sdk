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

import "io"

func GetUrl(action string, queryMap map[string]string) string {
	return globalClient.GetUrl(action, queryMap)
}

// Deprecated: Use client.DoGetResponseWithContext.
func DoGetResponse(url string) (*Response, error) {
	return globalClient.DoGetResponse(url)
}

// Deprecated: Use client.DoGetBytesWithContext.
func DoGetBytes(url string) ([]byte, error) {
	return globalClient.DoGetBytes(url)
}

// Deprecated: Use client.DoGetBytesRawWithContext.
func DoGetBytesRaw(url string) ([]byte, error) {
	return globalClient.DoGetBytesRaw(url)
}

// Deprecated: Use client.DoPostWithContext.
func DoPost(action string, queryMap map[string]string, postBytes []byte, isForm, isFile bool) (*Response, error) {
	return globalClient.DoPost(action, queryMap, postBytes, isForm, isFile)
}

// Deprecated: Use client.DoPostBytesRawWithContext.
func DoPostBytesRaw(url string, contentType string, body io.Reader) ([]byte, error) {
	return globalClient.DoPostBytesRaw(url, contentType, body)
}
