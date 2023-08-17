package casdoorsdk

func GetGroups() ([]*Group, error) {
	return globalClient.GetGroups()
}

func GetPaginationGroups(p int, pageSize int, queryMap map[string]string) ([]*Group, int, error) {
	return globalClient.GetPaginationGroups(p, pageSize, queryMap)
}

func GetGroup(name string) (*Group, error) {
	return globalClient.GetGroup(name)
}

func UpdateGroup(group *Group) (bool, error) {
	return globalClient.UpdateGroup(group)
}

func AddGroup(group *Group) (bool, error) {
	return globalClient.AddGroup(group)
}

func DeleteGroup(group *Group) (bool, error) {
	return globalClient.DeleteGroup(group)
}
