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

func GetAdapters() ([]*Adapter, error) {
	return globalClient.GetAdapters()
}

func GetPaginationAdapters(p int, pageSize int, queryMap map[string]string) ([]*Adapter, int, error) {
	return globalClient.GetPaginationAdapters(p, pageSize, queryMap)
}

func GetAdapter(name string) (*Adapter, error) {
	return globalClient.GetAdapter(name)
}

func UpdateAdapter(adapter *Adapter) (bool, error) {
	return globalClient.UpdateAdapter(adapter)
}

func AddAdapter(adapter *Adapter) (bool, error) {
	return globalClient.AddAdapter(adapter)
}

func DeleteAdapter(adapter *Adapter) (bool, error) {
	return globalClient.DeleteAdapter(adapter)
}
