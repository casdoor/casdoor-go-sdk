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

import (
	"testing"
)

func TestProduct(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Product")

	// Add a new object
	product := &Product{
		Owner:       "admin",
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
	}
	_, err := AddProduct(product)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	products, err := GetProducts()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range products {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	product, err = GetProduct(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if product.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", product.Name, name)
	}

	// Update the object
	updatedDescription := "Updated Casdoor Website"
	product.Description = updatedDescription
	_, err = UpdateProduct(product)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedProduct, err := GetProduct(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedProduct.Description != updatedDescription {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedProduct.Description, updatedDescription)
	}

	boughtProduct, err := BuyProduct(name, "provider_payment_dummy", "admin")
	if err != nil {
		t.Fatalf("Failed to buy product: %v", err)
	}
	if boughtProduct == nil {
		t.Fatalf("Failed to buy product: nil response")
	}

	// Delete the object
	_, err = DeleteProduct(product)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedProduct, err := GetProduct(name)
	if err != nil || deletedProduct != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
