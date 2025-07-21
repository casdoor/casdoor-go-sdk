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

import (
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type SubscriptionState string

const (
	SubStatePending   SubscriptionState = "Pending"
	SubStateError     SubscriptionState = "Error"
	SubStateSuspended SubscriptionState = "Suspended" // suspended by the admin

	SubStateActive   SubscriptionState = "Active"
	SubStateUpcoming SubscriptionState = "Upcoming"
	SubStateExpired  SubscriptionState = "Expired"
)

// Subscription has the same definition as https://github.com/casdoor/casdoor/blob/master/object/subscription.go#L39
type Subscription struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	DisplayName string `xorm:"varchar(100)" json:"displayName"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`
	Description string `xorm:"varchar(100)" json:"description"`

	User    string `xorm:"varchar(100)" json:"user"`
	Pricing string `xorm:"varchar(100)" json:"pricing"`
	Plan    string `xorm:"varchar(100)" json:"plan"`
	Payment string `xorm:"varchar(100)" json:"payment"`

	StartTime string            `xorm:"varchar(100)" json:"startTime"`
	EndTime   string            `xorm:"varchar(100)" json:"endTime"`
	Period    string            `xorm:"varchar(100)" json:"period"`
	State     SubscriptionState `xorm:"varchar(100)" json:"state"`
}

func (c *Client) GetSubscriptions() ([]*Subscription, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-subscriptions", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var subscriptions []*Subscription
	err = json.Unmarshal(bytes, &subscriptions)
	if err != nil {
		return nil, err
	}
	return subscriptions, nil
}

func (c *Client) GetPaginationSubscriptions(p int, pageSize int, queryMap map[string]string) ([]*Subscription, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-subscriptions", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var subscriptions []*Subscription
	err = json.Unmarshal(dataBytes, &subscriptions)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return subscriptions, int(response.Data2.(float64)), nil
}

func (c *Client) GetSubscription(name string) (*Subscription, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-subscription", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var subscription *Subscription
	err = json.Unmarshal(bytes, &subscription)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func (c *Client) AddSubscription(subscription *Subscription) (bool, error) {
	_, affected, err := c.modifySubscription("add-subscription", subscription, nil)
	return affected, err
}

func (c *Client) UpdateSubscription(subscription *Subscription) (bool, error) {
	_, affected, err := c.modifySubscription("update-subscription", subscription, nil)
	return affected, err
}

func (c *Client) DeleteSubscription(subscription *Subscription) (bool, error) {
	_, affected, err := c.modifySubscription("delete-subscription", subscription, nil)
	return affected, err
}
