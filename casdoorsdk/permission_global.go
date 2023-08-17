package casdoorsdk

func GetPermissions() ([]*Permission, error) {
	return globalClient.GetPermissions()
}

func GetPermissionsByRole(name string) ([]*Permission, error) {
	return globalClient.GetPermissionsByRole(name)
}

func GetPaginationPermissions(p int, pageSize int, queryMap map[string]string) ([]*Permission, int, error) {
	return globalClient.GetPaginationPermissions(p, pageSize, queryMap)
}

func GetPermission(name string) (*Permission, error) {
	return globalClient.GetPermission(name)
}

func UpdatePermission(permission *Permission) (bool, error) {
	return globalClient.UpdatePermission(permission)
}

func UpdatePermissionForColumns(permission *Permission, columns []string) (bool, error) {
	return globalClient.UpdatePermissionForColumns(permission, columns)
}

func AddPermission(permission *Permission) (bool, error) {
	return globalClient.AddPermission(permission)
}

func DeletePermission(permission *Permission) (bool, error) {
	return globalClient.DeletePermission(permission)
}
