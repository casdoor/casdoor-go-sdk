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

func GetPricings() ([]*Pricing, error) {
	return globalClient.GetPricings()
}

func GetPaginationPricings(p int, pageSize int, queryMap map[string]string) ([]*Pricing, int, error) {
	return globalClient.GetPaginationPricings(p, pageSize, queryMap)
}

func GetPricing(name string) (*Pricing, error) {
	return globalClient.GetPricing(name)
}

func UpdatePricing(pricing *Pricing) (bool, error) {
	return globalClient.UpdatePricing(pricing)
}

func AddPricing(pricing *Pricing) (bool, error) {
	return globalClient.AddPricing(pricing)
}

func DeletePricing(pricing *Pricing) (bool, error) {
	return globalClient.DeletePricing(pricing)
}
