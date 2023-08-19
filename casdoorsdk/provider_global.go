package casdoorsdk

func GetProviders() ([]*Provider, error) {
	return globalClient.GetProviders()
}

func GetPaginationProviders(p int, pageSize int, queryMap map[string]string) ([]*Provider, int, error) {
	return globalClient.GetPaginationProviders(p, pageSize, queryMap)
}

func GetProvider(name string) (*Provider, error) {
	return globalClient.GetProvider(name)
}

func UpdateProvider(provider *Provider) (bool, error) {
	return globalClient.UpdateProvider(provider)
}

func AddProvider(provider *Provider) (bool, error) {
	return globalClient.AddProvider(provider)
}

func DeleteProvider(provider *Provider) (bool, error) {
	return globalClient.DeleteProvider(provider)
}
