package casdoorsdk

func GetAdapters() ([]*Adapter, error) {
	return globalClient.GetAdapters()
}

func GetPaginationAdapters(p int, pageSize int, queryMap map[string]string) ([]*Adapter, int, error) {
	return globalClient.GetPaginationAdapters(p, pageSize, queryMap)
}

func GetAdapter(name string) (*Adapter, error) {
	return globalClient.GetAdapter(name)
}

func UpdateAdapter(adapter *Adapter) (bool, error) {
	return globalClient.UpdateAdapter(adapter)
}

func AddAdapter(adapter *Adapter) (bool, error) {
	return globalClient.AddAdapter(adapter)
}

func DeleteAdapter(adapter *Adapter) (bool, error) {
	return globalClient.DeleteAdapter(adapter)
}
