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
