package casdoorsdk

func GetModels() ([]*Model, error) {
	return globalClient.GetModels()
}

func GetPaginationModels(p int, pageSize int, queryMap map[string]string) ([]*Model, int, error) {
	return globalClient.GetPaginationModels(p, pageSize, queryMap)
}

func GetModel(name string) (*Model, error) {
	return globalClient.GetModel(name)
}

func UpdateModel(model *Model) (bool, error) {
	return globalClient.UpdateModel(model)
}

func AddModel(model *Model) (bool, error) {
	return globalClient.AddModel(model)
}

func DeleteModel(model *Model) (bool, error) {
	return globalClient.DeleteModel(model)
}
