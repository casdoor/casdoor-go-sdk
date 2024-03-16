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

func GetEnforcers() ([]*Enforcer, error) {
	return globalClient.GetEnforcers()
}

func GetPaginationEnforcers(p int, pageSize int, queryMap map[string]string) ([]*Enforcer, int, error) {
	return globalClient.GetPaginationEnforcers(p, pageSize, queryMap)
}

func GetEnforcer(name string) (*Enforcer, error) {
	return globalClient.GetEnforcer(name)
}

func UpdateEnforcer(enforcer *Enforcer) (bool, error) {
	return globalClient.UpdateEnforcer(enforcer)
}

func AddEnforcer(enforcer *Enforcer) (bool, error) {
	return globalClient.AddEnforcer(enforcer)
}

func DeleteEnforcer(enforcer *Enforcer) (bool, error) {
	return globalClient.DeleteEnforcer(enforcer)
}
