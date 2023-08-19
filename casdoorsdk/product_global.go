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

func GetProducts() ([]*Product, error) {
	return globalClient.GetProducts()
}

func GetPaginationProducts(p int, pageSize int, queryMap map[string]string) ([]*Product, int, error) {
	return globalClient.GetPaginationProducts(p, pageSize, queryMap)
}

func GetProduct(name string) (*Product, error) {
	return globalClient.GetProduct(name)
}

func UpdateProduct(product *Product) (bool, error) {
	return globalClient.UpdateProduct(product)
}

func AddProduct(product *Product) (bool, error) {
	return globalClient.AddProduct(product)
}

func DeleteProduct(product *Product) (bool, error) {
	return globalClient.DeleteProduct(product)
}

func BuyProduct(name string, providerName string) (*Product, error) {
	return globalClient.BuyProduct(name, providerName)
}
