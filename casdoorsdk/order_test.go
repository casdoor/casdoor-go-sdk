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

func TestOrder(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	productName := getRandomName("OrderProduct")
	orderName := getRandomName("Order")
	owner := "admin"

	product := &Product{
		Owner:       owner,
		Name:        productName,
		CreatedTime: GetCurrentTime(),
		DisplayName: productName,

		Image:       "https://cdn.casbin.org/img/casdoor-logo_1185x256.png",
		Description: "Casdoor Website",
		Tag:         "auto_created_product_for_plan",

		Quantity:  999,
		Sold:      0,
		State:     "Published",
		Providers: []string{"provider_payment_dummy"},
		Price:     1,
		Currency:  "USD",
	}
	_, err := AddProduct(product)
	if err != nil {
		t.Fatalf("Failed to add product: %v", err)
	}

	order := &Order{
		Owner:       owner,
		Name:        orderName,
		CreatedTime: GetCurrentTime(),
		DisplayName: orderName,
		Products:    []string{productName},
		ProductInfos: []ProductInfo{{
			Owner:       owner,
			Name:        productName,
			DisplayName: productName,
			Price:       1,
			Currency:    "USD",
			Quantity:    1,
		}},
		User:     owner,
		Price:    1,
		Currency: "USD",
		State:    "Created",
		Message:  "",
	}
	_, err = AddOrder(order)
	if err != nil {
		t.Fatalf("Failed to add order: %v", err)
	}

	orders, err := GetOrders()
	if err != nil {
		t.Fatalf("Failed to get orders: %v", err)
	}
	found := false
	for _, item := range orders {
		if item.Name == orderName {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added order not found in list")
	}

	userOrders, err := GetUserOrders(owner)
	if err != nil {
		t.Fatalf("Failed to get user orders: %v", err)
	}
	found = false
	for _, item := range userOrders {
		if item.Name == orderName {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added order not found in user list")
	}

	order, err = GetOrder(orderName)
	if err != nil {
		t.Fatalf("Failed to get order: %v", err)
	}
	if order.Name != orderName {
		t.Fatalf("Retrieved order does not match added order: %s != %s", order.Name, orderName)
	}

	updatedMessage := "Updated order message"
	order.Message = updatedMessage
	_, err = UpdateOrder(order)
	if err != nil {
		t.Fatalf("Failed to update order: %v", err)
	}

	updatedOrder, err := GetOrder(orderName)
	if err != nil {
		t.Fatalf("Failed to get updated order: %v", err)
	}
	if updatedOrder.Message != updatedMessage {
		t.Fatalf("Failed to update order, message mismatch: %s != %s", updatedOrder.Message, updatedMessage)
	}

	canceled, err := CancelOrder(orderName)
	if err != nil {
		t.Fatalf("Failed to cancel order: %v", err)
	}
	if !canceled {
		t.Fatalf("Failed to cancel order: not affected")
	}

	_, err = DeleteOrder(order)
	if err != nil {
		t.Fatalf("Failed to delete order: %v", err)
	}

	deletedOrder, err := GetOrder(orderName)
	if err != nil || deletedOrder != nil {
		t.Fatalf("Failed to delete order, it's still retrievable")
	}

	_, err = DeleteProduct(product)
	if err != nil {
		t.Fatalf("Failed to delete product: %v", err)
	}

	deletedProduct, err := GetProduct(productName)
	if err != nil || deletedProduct != nil {
		t.Fatalf("Failed to delete product, it's still retrievable")
	}
}
