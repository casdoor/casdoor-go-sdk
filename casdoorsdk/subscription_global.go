package casdoorsdk

func GetSubscriptions() ([]*Subscription, error) {
	return globalClient.GetSubscriptions()
}
func GetPaginationSubscriptions(p int, pageSize int, queryMap map[string]string) ([]*Subscription, int, error) {
	return globalClient.GetPaginationSubscriptions(p, pageSize, queryMap)
}
func GetSubscription(name string) (*Subscription, error) {
	return globalClient.GetSubscription(name)
}

func UpdateSubscription(subscription *Subscription) (bool, error) {
	return globalClient.UpdateSubscription(subscription)
}

func AddSubscription(subscription *Subscription) (bool, error) {
	return globalClient.AddSubscription(subscription)
}

func DeleteSubscription(subscription *Subscription) (bool, error) {
	return globalClient.DeleteSubscription(subscription)
}
