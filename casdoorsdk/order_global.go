// Copyright 2025 The Casdoor Authors. All Rights Reserved.
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

func GetOrders() ([]*Order, error) {
	return globalClient.GetOrders()
}

func GetPaginationOrders(p int, pageSize int, queryMap map[string]string) ([]*Order, int, error) {
	return globalClient.GetPaginationOrders(p, pageSize, queryMap)
}

func GetUserOrders(userName string) ([]*Order, error) {
	return globalClient.GetUserOrders(userName)
}

func GetOrder(name string) (*Order, error) {
	return globalClient.GetOrder(name)
}

func UpdateOrder(order *Order) (bool, error) {
	return globalClient.UpdateOrder(order)
}

func AddOrder(order *Order) (bool, error) {
	return globalClient.AddOrder(order)
}

func DeleteOrder(order *Order) (bool, error) {
	return globalClient.DeleteOrder(order)
}

func CancelOrder(name string) (bool, error) {
	return globalClient.CancelOrder(name)
}
