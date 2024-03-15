// Copyright 2024 The Casdoor Authors. All Rights Reserved.
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

func GetTransactions() ([]*Transaction, error) {
	return globalClient.GetTransactions()
}

func GetPaginationTransactions(p int, pageSize int, queryMap map[string]string) ([]*Transaction, int, error) {
	return globalClient.GetPaginationTransactions(p, pageSize, queryMap)
}

func GetTransaction(name string) (*Transaction, error) {
	return globalClient.GetTransaction(name)
}

func GetUserTransactions(userName string) ([]*Transaction, error) {
	return globalClient.GetUserTransactions(userName)
}

func UpdateTransaction(transaction *Transaction) (bool, error) {
	return globalClient.UpdateTransaction(transaction)
}

func AddTransaction(transaction *Transaction) (bool, error) {
	return globalClient.AddTransaction(transaction)
}

func DeleteTransaction(transaction *Transaction) (bool, error) {
	return globalClient.DeleteTransaction(transaction)
}
