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

func GetRecords() ([]*Record, error) {
	return globalClient.GetRecords()
}

func GetPaginationRecords(p int, pageSize int, queryMap map[string]string) ([]*Record, int, error) {
	return globalClient.GetPaginationRecords(p, pageSize, queryMap)
}

func GetRecord(name string) (*Record, error) {
	return globalClient.GetRecord(name)
}

func AddRecord(record *Record) (bool, error) {
	return globalClient.AddRecord(record)
}
