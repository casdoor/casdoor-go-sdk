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

import (
	"encoding/json"
	"fmt"
	"strconv"
)

// TicketMessage has the same definition as https://github.com/casdoor/casdoor/blob/master/object/ticket.go
type TicketMessage struct {
	Author    string `json:"author"`
	Text      string `json:"text"`
	Timestamp string `json:"timestamp"`
	IsAdmin   bool   `json:"isAdmin"`
}

// Ticket has the same definition as https://github.com/casdoor/casdoor/blob/master/object/ticket.go
//
// The ticket APIs act on behalf of a signed-in user, Casdoor rejects the application's own
// credentials with "Please sign in first". Call them on a client from WithAccessToken(), with
// the access token of the user who owns the tickets, or of an admin.
type Ticket struct {
	Owner       string           `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string           `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string           `xorm:"varchar(100)" json:"createdTime"`
	UpdatedTime string           `xorm:"varchar(100)" json:"updatedTime"`
	DisplayName string           `xorm:"varchar(100)" json:"displayName"`
	User        string           `xorm:"varchar(100) index" json:"user"`
	Title       string           `xorm:"varchar(200)" json:"title"`
	Content     string           `xorm:"mediumtext" json:"content"`
	State       string           `xorm:"varchar(50)" json:"state"`
	Messages    []*TicketMessage `xorm:"mediumtext json" json:"messages"`
}

// GetTickets returns the tickets of the organization of the client. An admin gets all the
// tickets, the other users only get their own ones.
func (c *Client) GetTickets() ([]*Ticket, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-tickets", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var tickets []*Ticket
	err = json.Unmarshal(bytes, &tickets)
	if err != nil {
		return nil, err
	}
	return tickets, nil
}

// GetPaginationTickets returns a page of the tickets of the organization of the client and the
// total number of the matched tickets. Only an admin gets paging, searching ("field", "value")
// and sorting ("sortField", "sortOrder"), the other users always get all their own tickets.
func (c *Client) GetPaginationTickets(p int, pageSize int, queryMap map[string]string) ([]*Ticket, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-tickets", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var tickets []*Ticket
	err = json.Unmarshal(dataBytes, &tickets)
	if err != nil {
		return nil, 0, err
	}

	return tickets, int(response.Data2.(float64)), nil
}

// GetTicket returns the ticket in the organization of the client, or nil if it doesn't exist.
// A user who isn't an admin can only get their own tickets.
func (c *Client) GetTicket(name string) (*Ticket, error) {
	queryMap := map[string]string{
		"id": c.GetId(name),
	}

	url := c.GetUrl("get-ticket", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var ticket *Ticket
	err = json.Unmarshal(bytes, &ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

// AddTicket adds a ticket, Casdoor sets its user to the signed-in user.
func (c *Client) AddTicket(ticket *Ticket) (bool, error) {
	_, affected, err := c.modifyTicket("add-ticket", ticket)
	return affected, err
}

// UpdateTicket updates a ticket. A user who isn't an admin can only close their own tickets.
func (c *Client) UpdateTicket(ticket *Ticket) (bool, error) {
	_, affected, err := c.modifyTicket("update-ticket", ticket)
	return affected, err
}

// DeleteTicket deletes a ticket, only an admin can do it.
func (c *Client) DeleteTicket(ticket *Ticket) (bool, error) {
	_, affected, err := c.modifyTicket("delete-ticket", ticket)
	return affected, err
}

// AddTicketMessage appends a message to the ticket in the organization of the client. Casdoor
// sets the author and the isAdmin of the message according to the signed-in user.
func (c *Client) AddTicketMessage(name string, message *TicketMessage) (bool, error) {
	queryMap := map[string]string{
		"id": c.GetId(name),
	}

	postBytes, err := json.Marshal(message)
	if err != nil {
		return false, err
	}

	resp, err := c.DoPost("add-ticket-message", queryMap, postBytes, false, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}

func (t Ticket) GetId() string {
	return fmt.Sprintf("%s/%s", t.Owner, t.Name)
}
