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

func GetTokens() ([]*Token, error) {
	return globalClient.GetTokens()
}

func GetPaginationTokens(p int, pageSize int, queryMap map[string]string) ([]*Token, int, error) {
	return globalClient.GetPaginationTokens(p, pageSize, queryMap)
}

func GetToken(name string) (*Token, error) {
	return globalClient.GetToken(name)
}

func UpdateToken(token *Token) (bool, error) {
	return globalClient.UpdateToken(token)
}

func UpdateTokenForColumns(token *Token, columns []string) (bool, error) {
	return globalClient.UpdateTokenForColumns(token, columns)
}

func AddToken(token *Token) (bool, error) {
	return globalClient.AddToken(token)
}

func DeleteToken(token *Token) (bool, error) {
	return globalClient.DeleteToken(token)
}
