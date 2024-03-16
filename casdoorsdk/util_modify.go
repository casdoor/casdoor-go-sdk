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
	"fmt"
	"strings"
)

// modifyOrganization is an encapsulation of permission CUD(Create, Update, Delete) operations.
// possible actions are `add-organization`, `update-organization`, `delete-organization`,
func (c *Client) modifyOrganization(action string, organization *Organization, columns []string) (*Response, bool, error) {
	organization.Owner = "admin"

	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", organization.Owner, organization.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	postBytes, err := json.Marshal(organization)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyApplication is an encapsulation of permission CUD(Create, Update, Delete) operations.
// possible actions are `add-application`, `update-application`, `delete-application`,
func (c *Client) modifyApplication(action string, application *Application, columns []string) (*Response, bool, error) {
	application.Owner = "admin"

	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", application.Owner, application.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	postBytes, err := json.Marshal(application)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyProvider is an encapsulation of permission CUD(Create, Update, Delete) operations.
// possible actions are `add-provider`, `update-provider`, `delete-provider`,
func (c *Client) modifyProvider(action string, provider *Provider, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", provider.Owner, provider.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	provider.Owner = c.OrganizationName
	postBytes, err := json.Marshal(provider)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifySession is an encapsulation of permission CUD(Create, Update, Delete) operations.
// possible actions are `add-session`, `update-session`, `delete-session`,
func (c *Client) modifySession(action string, session *Session, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", session.Owner, session.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	session.Owner = c.OrganizationName
	postBytes, err := json.Marshal(session)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyUser is an encapsulation of user CUD(Create, Update, Delete) operations.
// possible actions are `add-user`, `update-user`, `delete-user`,
func (c *Client) modifyUser(action string, user *User, columns []string) (*Response, bool, error) {
	return c.modifyUserById(action, user.GetId(), user, columns)
}

func (c *Client) modifyUserById(action string, id string, user *User, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": id,
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	user.Owner = c.OrganizationName
	postBytes, err := json.Marshal(user)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyPermission is an encapsulation of permission CUD(Create, Update, Delete) operations.
// possible actions are `add-permission`, `update-permission`, `delete-permission`,
func (c *Client) modifyPermission(action string, permission *Permission, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", permission.Owner, permission.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	permission.Owner = c.OrganizationName
	postBytes, err := json.Marshal(permission)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyRole is an encapsulation of role CUD(Create, Update, Delete) operations.
// possible actions are `add-role`, `update-role`, `delete-role`,
func (c *Client) modifyRole(action string, role *Role, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", role.Owner, role.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	role.Owner = c.OrganizationName
	postBytes, err := json.Marshal(role)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyCert is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-cert`, `update-cert`, `delete-cert`,
func (c *Client) modifyCert(action string, cert *Cert, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", cert.Owner, cert.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	cert.Owner = c.OrganizationName
	postBytes, err := json.Marshal(cert)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyEnforcer is an encapsulation of cert CUD(Create, Update, Delete) operations.
func (c *Client) modifyEnforcer(action string, enforcer *Enforcer, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", enforcer.Owner, enforcer.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	enforcer.Owner = c.OrganizationName
	postBytes, err := json.Marshal(enforcer)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyPolicy is an encapsulation of cert CUD(Create, Update, Delete) operations.
func (c *Client) modifyPolicy(action string, enforcer *Enforcer, policies []*CasbinRule, columns []string) (*Response, bool, error) {
	enforcer.Owner = c.OrganizationName
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", enforcer.Owner, enforcer.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	var postBytes []byte
	var err error
	if action == "update-policy" {
		postBytes, err = json.Marshal(policies)
	} else {
		postBytes, err = json.Marshal(policies[0])
	}

	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyEnforcer is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-group`, `update-group`, `delete-group`,
func (c *Client) modifyGroup(action string, group *Group, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", group.Owner, group.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	group.Owner = c.OrganizationName
	postBytes, err := json.Marshal(group)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyAdapter is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-adapter`, `update-adapter`, `delete-adapter`,
func (c *Client) modifyAdapter(action string, adapter *Adapter, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", adapter.Owner, adapter.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	adapter.Owner = c.OrganizationName
	postBytes, err := json.Marshal(adapter)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyModel is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-model`, `update-model`, `delete-model`,
func (c *Client) modifyModel(action string, model *Model, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", model.Owner, model.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	model.Owner = c.OrganizationName
	postBytes, err := json.Marshal(model)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyProduct is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-product`, `update-product`, `delete-product`,
func (c *Client) modifyProduct(action string, product *Product, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", product.Owner, product.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	product.Owner = c.OrganizationName
	postBytes, err := json.Marshal(product)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyPayment is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-payment`, `update-payment`, `delete-payment`,
func (c *Client) modifyPayment(action string, payment *Payment, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", payment.Owner, payment.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	payment.Owner = c.OrganizationName
	postBytes, err := json.Marshal(payment)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyPlan is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-plan`, `update-plan`, `delete-plan`,
func (c *Client) modifyPlan(action string, plan *Plan, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", plan.Owner, plan.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	plan.Owner = c.OrganizationName
	postBytes, err := json.Marshal(plan)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyPricing is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-pricing`, `update-pricing`, `delete-pricing`,
func (c *Client) modifyPricing(action string, pricing *Pricing, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", pricing.Owner, pricing.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	pricing.Owner = c.OrganizationName
	postBytes, err := json.Marshal(pricing)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifySubscription is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-subscription`, `update-subscription`, `delete-subscription`,
func (c *Client) modifySubscription(action string, subscription *Subscription, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", subscription.Owner, subscription.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	subscription.Owner = c.OrganizationName
	postBytes, err := json.Marshal(subscription)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifySyner is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-syncer`, `update-syncer`, `delete-syncer`,
func (c *Client) modifySyncer(action string, syncer *Syncer, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", syncer.Owner, syncer.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	syncer.Owner = c.OrganizationName
	postBytes, err := json.Marshal(syncer)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyTransaction is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-transaction`, `update-transaction`, `delete-transaction`,
func (c *Client) modifyTransaction(action string, transaction *Transaction, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", transaction.Owner, transaction.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	transaction.Owner = c.OrganizationName
	postBytes, err := json.Marshal(transaction)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyWebhook is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-webhook`, `update-webhook`, `delete-webhook`,
func (c *Client) modifyWebhook(action string, webhook *Webhook, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", webhook.Owner, webhook.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	webhook.Owner = c.OrganizationName
	postBytes, err := json.Marshal(webhook)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyToken is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-token`, `update-token`, `delete-token`,
func (c *Client) modifyToken(action string, token *Token, columns []string) (*Response, bool, error) {
	token.Owner = "admin"

	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", token.Owner, token.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	postBytes, err := json.Marshal(token)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPost(action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}
