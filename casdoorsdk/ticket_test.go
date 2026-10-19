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
	"testing"
	"time"
)

func TestTicket(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	// The ticket APIs act on behalf of a user, not of the application
	token, err := GetOAuthTokenByPassword("admin", "123")
	if err != nil {
		t.Fatalf("Failed to get the token of the admin: %v", err)
	}
	userClient := WithAccessToken(token.AccessToken)

	name := getRandomName("unit_test_ticket")
	ticket := &Ticket{
		Owner:       TestCasdoorOrganization,
		Name:        name,
		CreatedTime: time.Now().Format(time.RFC3339),
		DisplayName: "Test Ticket",
		User:        "somebody/else",
		Title:       "Test title",
		Content:     "Test content",
		State:       "Open",
	}

	if _, err = AddTicket(&Ticket{Owner: TestCasdoorOrganization, Name: name}); err == nil {
		t.Fatalf("Expected the application's credentials to be rejected")
	}

	// Test AddTicket
	affected, err := userClient.AddTicket(ticket)
	if err != nil {
		t.Fatalf("Failed to add ticket: %v", err)
	}
	if !affected {
		t.Fatalf("Expected AddTicket to affect a row")
	}
	defer func() {
		_, _ = userClient.DeleteTicket(&Ticket{Owner: TestCasdoorOrganization, Name: name})
	}()

	// Test GetTicket, Casdoor sets the user to the signed-in user
	got, err := userClient.GetTicket(name)
	if err != nil {
		t.Fatalf("Failed to get ticket: %v", err)
	}
	if got == nil || got.Title != ticket.Title || got.Content != ticket.Content || got.State != "Open" {
		t.Fatalf("Retrieved ticket does not match added ticket: %+v", got)
	}
	if got.User != "built-in/admin" {
		t.Fatalf("Expected the user of the ticket to be the signed-in user, got %q", got.User)
	}

	// Test GetTickets
	tickets, err := userClient.GetTickets()
	if err != nil {
		t.Fatalf("Failed to get tickets: %v", err)
	}
	if findTicket(tickets, name) == nil {
		t.Fatalf("The added ticket is not in GetTickets")
	}

	// Test GetPaginationTickets
	pageTickets, total, err := userClient.GetPaginationTickets(1, 100, map[string]string{"field": "name", "value": name})
	if err != nil {
		t.Fatalf("Failed to get pagination tickets: %v", err)
	}
	if total != 1 || len(pageTickets) != 1 || pageTickets[0].Name != name {
		t.Fatalf("Unexpected pagination result: total=%d tickets=%d", total, len(pageTickets))
	}

	// Test AddTicketMessage, Casdoor sets the author and the isAdmin
	affected, err = userClient.AddTicketMessage(name, &TicketMessage{Author: "somebody/else", Text: "first message", Timestamp: time.Now().Format(time.RFC3339)})
	if err != nil {
		t.Fatalf("Failed to add ticket message: %v", err)
	}
	if !affected {
		t.Fatalf("Expected AddTicketMessage to affect a row")
	}
	if _, err = userClient.AddTicketMessage(name, &TicketMessage{Text: "second message"}); err != nil {
		t.Fatalf("Failed to add the second ticket message: %v", err)
	}
	got, err = userClient.GetTicket(name)
	if err != nil {
		t.Fatalf("Failed to get ticket: %v", err)
	}
	if len(got.Messages) != 2 || got.Messages[0].Text != "first message" || got.Messages[1].Text != "second message" {
		t.Fatalf("Unexpected messages: %+v", got.Messages)
	}
	if got.Messages[0].Author != "built-in/admin" || !got.Messages[0].IsAdmin {
		t.Fatalf("Expected the author and isAdmin of the message to be set by Casdoor: %+v", got.Messages[0])
	}

	if _, err = userClient.AddTicketMessage(getRandomName("not_exist"), &TicketMessage{Text: "lost"}); err == nil {
		t.Fatalf("Expected an error for the ticket that doesn't exist")
	}

	// Test UpdateTicket
	got.State = "Closed"
	got.Title = "Test title updated"
	if _, err = userClient.UpdateTicket(got); err != nil {
		t.Fatalf("Failed to update ticket: %v", err)
	}
	got, err = userClient.GetTicket(name)
	if err != nil {
		t.Fatalf("Failed to get ticket: %v", err)
	}
	if got.State != "Closed" || got.Title != "Test title updated" || len(got.Messages) != 2 {
		t.Fatalf("The ticket was not updated: %+v", got)
	}

	// Test DeleteTicket
	affected, err = userClient.DeleteTicket(got)
	if err != nil {
		t.Fatalf("Failed to delete ticket: %v", err)
	}
	if !affected {
		t.Fatalf("Expected DeleteTicket to affect a row")
	}
	deleted, err := userClient.GetTicket(name)
	if err != nil {
		t.Fatalf("Failed to get the deleted ticket: %v", err)
	}
	if deleted != nil {
		t.Fatalf("Failed to delete ticket, it still exists")
	}
}

func findTicket(tickets []*Ticket, name string) *Ticket {
	for _, ticket := range tickets {
		if ticket.Name == name {
			return ticket
		}
	}
	return nil
}
