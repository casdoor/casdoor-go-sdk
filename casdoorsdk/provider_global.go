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

func GetProviders() ([]*Provider, error) {
	return globalClient.GetProviders()
}

func GetPaginationProviders(p int, pageSize int, queryMap map[string]string) ([]*Provider, int, error) {
	return globalClient.GetPaginationProviders(p, pageSize, queryMap)
}

func GetProvider(name string) (*Provider, error) {
	return globalClient.GetProvider(name)
}

func UpdateProvider(provider *Provider) (bool, error) {
	return globalClient.UpdateProvider(provider)
}

func AddProvider(provider *Provider) (bool, error) {
	return globalClient.AddProvider(provider)
}

func DeleteProvider(provider *Provider) (bool, error) {
	return globalClient.DeleteProvider(provider)
}
