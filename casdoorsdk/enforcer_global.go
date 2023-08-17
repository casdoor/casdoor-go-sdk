package casdoorsdk

func GetEnforcers() ([]*Enforcer, error) {
	return globalClient.GetEnforcers()
}

func GetPaginationEnforcers(p int, pageSize int, queryMap map[string]string) ([]*Enforcer, int, error) {
	return globalClient.GetPaginationEnforcers(p, pageSize, queryMap)
}

func GetEnforcer(name string) (*Enforcer, error) {
	return globalClient.GetEnforcer(name)
}

func UpdateEnforcer(enforcer *Enforcer) (bool, error) {
	return globalClient.UpdateEnforcer(enforcer)
}

func AddEnforcer(enforcer *Enforcer) (bool, error) {
	return globalClient.AddEnforcer(enforcer)
}

func DeleteEnforcer(enforcer *Enforcer) (bool, error) {
	return globalClient.DeleteEnforcer(enforcer)
}
