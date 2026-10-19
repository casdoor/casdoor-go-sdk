// Copyright 2026 The Casdoor Authors. All Rights Reserved.
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

func GetTickets() ([]*Ticket, error) {
	return globalClient.GetTickets()
}

func GetPaginationTickets(p int, pageSize int, queryMap map[string]string) ([]*Ticket, int, error) {
	return globalClient.GetPaginationTickets(p, pageSize, queryMap)
}

func GetTicket(name string) (*Ticket, error) {
	return globalClient.GetTicket(name)
}

func AddTicket(ticket *Ticket) (bool, error) {
	return globalClient.AddTicket(ticket)
}

func UpdateTicket(ticket *Ticket) (bool, error) {
	return globalClient.UpdateTicket(ticket)
}

func DeleteTicket(ticket *Ticket) (bool, error) {
	return globalClient.DeleteTicket(ticket)
}

func AddTicketMessage(name string, message *TicketMessage) (bool, error) {
	return globalClient.AddTicketMessage(name, message)
}
