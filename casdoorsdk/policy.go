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

func (c *Client) GetPolicies(enforcerName string, adapterId string) ([]*CasbinRule, error) {
	queryMap := map[string]string{
		"id":        fmt.Sprintf("%s/%s", c.OrganizationName, enforcerName),
		"adapterId": adapterId,
	}

	url := c.GetUrl("get-policies", queryMap)

	bytes, err := c.DoGetBytes(url)
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
