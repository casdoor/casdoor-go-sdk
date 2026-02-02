// Copyright 2023 The Casdoor Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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

func GetUserPayments(userName string) ([]*Payment, error) {
	return globalClient.GetUserPayments(userName)
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
