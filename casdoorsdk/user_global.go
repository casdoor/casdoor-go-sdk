// Copyright 2021 The Casdoor Authors. All Rights Reserved.
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

func GetGlobalUsers() ([]*User, error) {
	return globalClient.GetGlobalUsers()
}

func GetUsers() ([]*User, error) {
	return globalClient.GetUsers()
}

func GetSortedUsers(sorter string, limit int) ([]*User, error) {
	return globalClient.GetSortedUsers(sorter, limit)
}

func GetPaginationUsers(p int, pageSize int, queryMap map[string]string) ([]*User, int, error) {
	return globalClient.GetPaginationUsers(p, pageSize, queryMap)
}

func GetUserCount(isOnline string) (int, error) {
	return globalClient.GetUserCount(isOnline)
}

func GetUser(name string) (*User, error) {
	return globalClient.GetUser(name)
}

func GetUserByEmail(email string) (*User, error) {
	return globalClient.GetUserByEmail(email)
}

func GetUserByPhone(phone string) (*User, error) {
	return globalClient.GetUserByPhone(phone)
}

func GetUserByUserId(userId string) (*User, error) {
	return globalClient.GetUserByUserId(userId)
}

// note: oldPassword is not required, if you don't need, just pass a empty string
func SetPassword(owner, name, oldPassword, newPassword string) (bool, error) {
	return globalClient.SetPassword(owner, name, oldPassword, newPassword)
}

func UpdateUserById(id string, user *User) (bool, error) {
	return globalClient.UpdateUserById(id, user)
}

func UpdateUser(user *User) (bool, error) {
	return globalClient.UpdateUser(user)
}

func UpdateUserForColumns(user *User, columns []string) (bool, error) {
	return globalClient.UpdateUserForColumns(user, columns)
}

func AddUser(user *User) (bool, error) {
	return globalClient.AddUser(user)
}

func DeleteUser(user *User) (bool, error) {
	return globalClient.DeleteUser(user)
}

func CheckUserPassword(user *User) (bool, error) {
	return globalClient.CheckUserPassword(user)
}
