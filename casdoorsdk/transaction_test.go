package casdoorsdk

import "testing"

func TestTransaction(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("Transaction")

	// Add a new object
	transaction := &Transaction{
		Owner:       "admin",
		Name:        name,
		CreatedTime: GetCurrentTime(),
		DisplayName: name,
		ProductName: "casbin",
	}
	_, err := AddTransaction(transaction)
	if err != nil {
		t.Fatalf("Failed to add object: %v", err)
	}

	// Get all objects, check if our added object is inside the list
	transactions, err := GetTransactions()
	if err != nil {
		t.Fatalf("Failed to get objects: %v", err)
	}
	found := false
	for _, item := range transactions {
		if item.Name == name {
			found = true
			break
		}
	}
	if !found {
		t.Fatalf("Added object not found in list")
	}

	// Get the object
	transaction, err = GetTransaction(name)
	if err != nil {
		t.Fatalf("Failed to get object: %v", err)
	}
	if transaction.Name != name {
		t.Fatalf("Retrieved object does not match added object: %s != %s", transaction.Name, name)
	}

	// Update the object
	updatedProductName := "Updated Casdoor Website"
	transaction.ProductName = updatedProductName
	_, err = UpdateTransaction(transaction)
	if err != nil {
		t.Fatalf("Failed to update object: %v", err)
	}

	// Validate the update
	updatedTransaction, err := GetTransaction(name)
	if err != nil {
		t.Fatalf("Failed to get updated object: %v", err)
	}
	if updatedTransaction.ProductName != updatedProductName {
		t.Fatalf("Failed to update object, description mismatch: %s != %s", updatedTransaction.ProductName, updatedProductName)
	}

	// Delete the object
	_, err = DeleteTransaction(transaction)
	if err != nil {
		t.Fatalf("Failed to delete object: %v", err)
	}

	// Validate the deletion
	deletedTransaction, err := GetTransaction(name)
	if err != nil || deletedTransaction != nil {
		t.Fatalf("Failed to delete object, it's still retrievable")
	}
}
