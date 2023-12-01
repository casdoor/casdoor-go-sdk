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

import (
	"golang.org/x/oauth2"
)

func GetOAuthToken(code string, state string) (*oauth2.Token, error) {
	return globalClient.GetOAuthToken(code, state)
}

func RefreshOAuthToken(refreshToken string) (*oauth2.Token, error) {
	return globalClient.RefreshOAuthToken(refreshToken)
}

func GetTokens(p int, pageSize int) ([]*Token, int, error) {
	return globalClient.GetTokens(p, pageSize)
}

func GetToken(tokenID string) (*Token, error) {
	return globalClient.GetToken(tokenID)
}

func DeleteToken(token *Token) (bool, error) {
	return globalClient.DeleteToken(token)
}
