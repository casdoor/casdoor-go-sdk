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

func GetPermissions() ([]*Permission, error) {
	return globalClient.GetPermissions()
}

func GetPermissionsByRole(name string) ([]*Permission, error) {
	return globalClient.GetPermissionsByRole(name)
}

func GetPaginationPermissions(p int, pageSize int, queryMap map[string]string) ([]*Permission, int, error) {
	return globalClient.GetPaginationPermissions(p, pageSize, queryMap)
}

func GetPermission(name string) (*Permission, error) {
	return globalClient.GetPermission(name)
}

func UpdatePermission(permission *Permission) (bool, error) {
	return globalClient.UpdatePermission(permission)
}

func UpdatePermissionForColumns(permission *Permission, columns []string) (bool, error) {
	return globalClient.UpdatePermissionForColumns(permission, columns)
}

func AddPermission(permission *Permission) (bool, error) {
	return globalClient.AddPermission(permission)
}

func DeletePermission(permission *Permission) (bool, error) {
	return globalClient.DeletePermission(permission)
}
