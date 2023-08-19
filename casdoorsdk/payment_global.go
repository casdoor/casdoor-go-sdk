package casdoorsdk

func GetPayments() ([]*Payment, error) {
	return globalClient.GetPayments()
}

func GetPaginationPayments(p int, pageSize int, queryMap map[string]string) ([]*Payment, int, error) {
	return globalClient.GetPaginationPayments(p, pageSize, queryMap)
}

func GetPayment(name string) (*Payment, error) {
	return globalClient.GetPayment(name)
}
func GetUserPayments() ([]*Payment, error) {
	return globalClient.GetUserPayments()
}

func UpdatePayment(payment *Payment) (bool, error) {
	return globalClient.UpdatePayment(payment)
}

func AddPayment(payment *Payment) (bool, error) {
	return globalClient.AddPayment(payment)
}

func DeletePayment(payment *Payment) (bool, error) {
	return globalClient.DeletePayment(payment)
}
func NotifyPayment(payment *Payment) (bool, error) {
	return globalClient.NotifyPayment(payment)
}

func InvoicePayment(payment *Payment) (bool, error) {
	return globalClient.NotifyPayment(payment)
}
