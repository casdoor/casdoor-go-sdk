// Copyright 2024 The Casdoor Authors. All Rights Reserved.
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
	"context"
	"encoding/json"
	"fmt"
)

type CasbinRule struct {
	Id    int64  `xorm:"pk autoincr"`
	Ptype string `xorm:"varchar(100) index not null default ''"`
	V0    string `xorm:"varchar(100) index not null default ''"`
	V1    string `xorm:"varchar(100) index not null default ''"`
	V2    string `xorm:"varchar(100) index not null default ''"`
	V3    string `xorm:"varchar(100) index not null default ''"`
	V4    string `xorm:"varchar(100) index not null default ''"`
	V5    string `xorm:"varchar(100) index not null default ''"`

	tableName string `xorm:"-"`
}

func (c *Client) AddPolicy(enforcer *Enforcer, policy *CasbinRule) (bool, error) {
	var policies []*CasbinRule
	policies = make([]*CasbinRule, 1)
	policies[0] = policy
	_, affected, err := c.modifyPolicy("add-policy", enforcer, policies, nil)
	return affected, err
}

func (c *Client) UpdatePolicy(enforcer *Enforcer, oldpolicy *CasbinRule, newpolicy *CasbinRule) (bool, error) {
	var policies []*CasbinRule
	policies = make([]*CasbinRule, 2)
	policies[0] = oldpolicy
	policies[1] = newpolicy
	_, affected, err := c.modifyPolicy("update-policy", enforcer, policies, nil)
	return affected, err
}

func (c *Client) RemovePolicy(enforcer *Enforcer, policy *CasbinRule) (bool, error) {
	var policies []*CasbinRule
	policies = make([]*CasbinRule, 1)
	policies[0] = policy
	_, affected, err := c.modifyPolicy("remove-policy", enforcer, policies, nil)
	return affected, err
}

// Deprecated: Use GetPoliciesWithContext.
func (c *Client) GetPolicies(enforcerName string, adapterId string) ([]*CasbinRule, error) {
	return c.GetPoliciesWithContext(context.Background(), enforcerName, adapterId)
}

func (c *Client) GetPoliciesWithContext(ctx context.Context, enforcerName string, adapterId string) ([]*CasbinRule, error) {
	queryMap := map[string]string{
		"id":        fmt.Sprintf("%s/%s", c.OrganizationName, enforcerName),
		"adapterId": adapterId,
	}

	url := c.GetUrl("get-policies", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	var policies []*CasbinRule
	err = json.Unmarshal(bytes, &policies)
	if err != nil {
		return nil, err
	}
	return policies, nil
}

// PolicyFilter represents a filter for getting policies
type PolicyFilter struct {
	Ptype       string   `json:"ptype"`
	FieldIndex  *int     `json:"fieldIndex,omitempty"`
	FieldValues []string `json:"fieldValues,omitempty"`
}

// GetFilteredPolicies gets policies with filtering capabilities based on field index and values
// Deprecated: Use GetFilteredPoliciesWithContext.
func (c *Client) GetFilteredPolicies(enforcerId string, filters []*PolicyFilter) ([]*CasbinRule, error) {
	return c.GetFilteredPoliciesWithContext(context.Background(), enforcerId, filters)
}

func (c *Client) GetFilteredPoliciesWithContext(ctx context.Context, enforcerId string, filters []*PolicyFilter) ([]*CasbinRule, error) {
	queryMap := map[string]string{
		"id": enforcerId,
	}

	// Convert filters to JSON
	postBytes, err := json.Marshal(filters)
	if err != nil {
		return nil, err
	}

	// Make POST request with filters in body
	resp, err := c.DoPostWithContext(ctx, "get-filtered-policies", queryMap, postBytes, false, false)
	if err != nil {
		return nil, err
	}

	// Extract data from response
	res, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	var policies []*CasbinRule
	err = json.Unmarshal(res, &policies)
	if err != nil {
		return nil, err
	}
	return policies, nil
}
