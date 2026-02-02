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

import "testing"

func TestOrderPay(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("OrderPayProduct")
	owner := "admin"

	product := &Product{
		Owner:       owner,
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,

		Image:       "https://cdn.casbin.org/img/casdoor-logo_1185x256.png",
		Description: "Casdoor Website",
		Tag:         "auto_created_product_for_plan",

		Quantity:  999,
		Sold:      0,
		State:     "Published",
		Providers: []string{"provider_payment_dummy"},
		Currency:  "USD",
	}
	_, err := AddProduct(product)
	if err != nil {
		t.Fatalf("Failed to add product: %v", err)
	}

	productInfos := []ProductInfo{{
		Name:     name,
		Quantity: 1,
	}}

	order, err := PlaceOrder(productInfos, "admin")
	if err != nil {
		t.Fatalf("Failed to place order: %v", err)
	}
	if order == nil {
		t.Fatalf("Failed to place order: nil response")
	}

	payment, err := PayOrder(order.Name, "provider_payment_dummy")
	if err != nil {
		t.Fatalf("Failed to pay order: %v", err)
	}
	if payment == nil {
		t.Fatalf("Failed to pay order: nil response")
	}

	_, err = DeleteProduct(product)
	if err != nil {
		t.Fatalf("Failed to delete product: %v", err)
	}

	deletedProduct, err := GetProduct(name)
	if err != nil || deletedProduct != nil {
		t.Fatalf("Failed to delete product, it's still retrievable")
	}
}
