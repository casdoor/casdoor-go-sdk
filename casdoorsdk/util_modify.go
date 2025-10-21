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
	"context"
	"encoding/json"
	"fmt"
	"strings"
)

// modifyOrganization is an encapsulation of permission CUD(Create, Update, Delete) operations.
// possible actions are `add-organization`, `update-organization`, `delete-organization`,
// Deprecated: Use modifyOrganizationWithContext.
func (c *Client) modifyOrganization(action string, organization *Organization, columns []string) (*Response, bool, error) {
	return c.modifyOrganizationWithContext(context.Background(), action, organization, columns)
}

func (c *Client) modifyOrganizationWithContext(ctx context.Context, action string, organization *Organization, columns []string) (*Response, bool, error) {
	if organization.Owner == "" {
		organization.Owner = "admin"
	}

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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyApplication is an encapsulation of permission CUD(Create, Update, Delete) operations.
// possible actions are `add-application`, `update-application`, `delete-application`,
// Deprecated: Use modifyApplicationWithContext.
func (c *Client) modifyApplication(action string, application *Application, columns []string) (*Response, bool, error) {
	return c.modifyApplicationWithContext(context.Background(), action, application, columns)
}

func (c *Client) modifyApplicationWithContext(ctx context.Context, action string, application *Application, columns []string) (*Response, bool, error) {
	if application.Owner == "" {
		application.Owner = "admin"
	}

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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyProvider is an encapsulation of permission CUD(Create, Update, Delete) operations.
// possible actions are `add-provider`, `update-provider`, `delete-provider`,
// Deprecated: Use modifyProviderWithContext.
func (c *Client) modifyProvider(action string, provider *Provider, columns []string) (*Response, bool, error) {
	return c.modifyProviderWithContext(context.Background(), action, provider, columns)
}

func (c *Client) modifyProviderWithContext(ctx context.Context, action string, provider *Provider, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifySession is an encapsulation of permission CUD(Create, Update, Delete) operations.
// possible actions are `add-session`, `update-session`, `delete-session`,
// Deprecated: Use modifySessionWithContext.
func (c *Client) modifySession(action string, session *Session, columns []string) (*Response, bool, error) {
	return c.modifySessionWithContext(context.Background(), action, session, columns)
}

func (c *Client) modifySessionWithContext(ctx context.Context, action string, session *Session, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyUser is an encapsulation of user CUD(Create, Update, Delete) operations.
// possible actions are `add-user`, `update-user`, `delete-user`,
// Deprecated: Use modifyUserWithContext.
func (c *Client) modifyUser(action string, user *User, columns []string) (*Response, bool, error) {
	return c.modifyUserWithContext(context.Background(), action, user, columns)
}

func (c *Client) modifyUserWithContext(ctx context.Context, action string, user *User, columns []string) (*Response, bool, error) {
	return c.modifyUserByIdWithContext(ctx, action, user.GetId(), user, columns)
}

// Deprecated: Use modifyUserByIdWithContext.
func (c *Client) modifyUserById(action string, id string, user *User, columns []string) (*Response, bool, error) {
	return c.modifyUserByIdWithContext(context.Background(), action, id, user, columns)
}

func (c *Client) modifyUserByIdWithContext(ctx context.Context, action string, id string, user *User, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": id,
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	if user.Owner == "" {
		user.Owner = c.OrganizationName
	}
	postBytes, err := json.Marshal(user)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	if action == "check-user-password" {
		return resp, resp.Status == "ok", nil
	}

	return resp, resp.Data == "Affected", nil
}

// Deprecated: Use modifyUserByUserIdWithContext.
func (c *Client) modifyUserByUserId(action string, owner string, userId string, user *User, columns []string) (*Response, bool, error) {
	return c.modifyUserByUserIdWithContext(context.Background(), action, owner, userId, user, columns)
}

func (c *Client) modifyUserByUserIdWithContext(ctx context.Context, action string, owner string, userId string, user *User, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"owner":  owner,
		"userId": userId,
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	postBytes, err := json.Marshal(user)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	if action == "check-user-password" {
		return resp, resp.Status == "ok", nil
	}

	return resp, resp.Data == "Affected", nil
}

// modifyPermission is an encapsulation of permission CUD(Create, Update, Delete) operations.
// possible actions are `add-permission`, `update-permission`, `delete-permission`,
// Deprecated: Use modifyPermissionWithContext.
func (c *Client) modifyPermission(action string, permission *Permission, columns []string) (*Response, bool, error) {
	return c.modifyPermissionWithContext(context.Background(), action, permission, columns)
}

func (c *Client) modifyPermissionWithContext(ctx context.Context, action string, permission *Permission, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyRole is an encapsulation of role CUD(Create, Update, Delete) operations.
// possible actions are `add-role`, `update-role`, `delete-role`,
// Deprecated: Use modifyRoleWithContext.
func (c *Client) modifyRole(action string, role *Role, columns []string) (*Response, bool, error) {
	return c.modifyRoleWithContext(context.Background(), action, role, columns)
}

func (c *Client) modifyRoleWithContext(ctx context.Context, action string, role *Role, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyCert is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-cert`, `update-cert`, `delete-cert`,
// Deprecated: Use modifyCertWithContext.
func (c *Client) modifyCert(action string, cert *Cert, columns []string) (*Response, bool, error) {
	return c.modifyCertWithContext(context.Background(), action, cert, columns)
}

func (c *Client) modifyCertWithContext(ctx context.Context, action string, cert *Cert, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", cert.Owner, cert.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	if cert.Owner == "" {
		cert.Owner = c.OrganizationName
	}
	postBytes, err := json.Marshal(cert)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyEnforcer is an encapsulation of cert CUD(Create, Update, Delete) operations.
// Deprecated: Use modifyEnforcerWithContext.
func (c *Client) modifyEnforcer(action string, enforcer *Enforcer, columns []string) (*Response, bool, error) {
	return c.modifyEnforcerWithContext(context.Background(), action, enforcer, columns)
}

func (c *Client) modifyEnforcerWithContext(ctx context.Context, action string, enforcer *Enforcer, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyPolicy is an encapsulation of cert CUD(Create, Update, Delete) operations.
// Deprecated: Use modifyPolicyWithContext.
func (c *Client) modifyPolicy(action string, enforcer *Enforcer, policies []*CasbinRule, columns []string) (*Response, bool, error) {
	return c.modifyPolicyWithContext(context.Background(), action, enforcer, policies, columns)
}

func (c *Client) modifyPolicyWithContext(ctx context.Context, action string, enforcer *Enforcer, policies []*CasbinRule, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyEnforcer is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-group`, `update-group`, `delete-group`,
func (c *Client) modifyGroup(action string, group *Group, columns []string) (*Response, bool, error) {
	return c.modifyGroupWithContext(context.Background(), action, group, columns)
}

// modifyGroupWithContext is like modifyGroup but accepts a context.
func (c *Client) modifyGroupWithContext(ctx context.Context, action string, group *Group, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyAdapter is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-adapter`, `update-adapter`, `delete-adapter`,
// Deprecated: Use modifyAdapterWithContext.
func (c *Client) modifyAdapter(action string, adapter *Adapter, columns []string) (*Response, bool, error) {
	return c.modifyAdapterWithContext(context.Background(), action, adapter, columns)
}

func (c *Client) modifyAdapterWithContext(ctx context.Context, action string, adapter *Adapter, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyModel is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-model`, `update-model`, `delete-model`,
// Deprecated: Use modifyModelWithContext.
func (c *Client) modifyModel(action string, model *Model, columns []string) (*Response, bool, error) {
	return c.modifyModelWithContext(context.Background(), action, model, columns)
}

func (c *Client) modifyModelWithContext(ctx context.Context, action string, model *Model, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyProduct is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-product`, `update-product`, `delete-product`,
// Deprecated: Use modifyProductWithContext.
func (c *Client) modifyProduct(action string, product *Product, columns []string) (*Response, bool, error) {
	return c.modifyProductWithContext(context.Background(), action, product, columns)
}

func (c *Client) modifyProductWithContext(ctx context.Context, action string, product *Product, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyPayment is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-payment`, `update-payment`, `delete-payment`,
// Deprecated: Use modifyPaymentWithContext.
func (c *Client) modifyPayment(action string, payment *Payment, columns []string) (*Response, bool, error) {
	return c.modifyPaymentWithContext(context.Background(), action, payment, columns)
}

func (c *Client) modifyPaymentWithContext(ctx context.Context, action string, payment *Payment, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyPlan is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-plan`, `update-plan`, `delete-plan`,
// Deprecated: Use modifyPlanWithContext.
func (c *Client) modifyPlan(action string, plan *Plan, columns []string) (*Response, bool, error) {
	return c.modifyPlanWithContext(context.Background(), action, plan, columns)
}

func (c *Client) modifyPlanWithContext(ctx context.Context, action string, plan *Plan, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyPricing is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-pricing`, `update-pricing`, `delete-pricing`,
// Deprecated: Use modifyPricingWithContext.
func (c *Client) modifyPricing(action string, pricing *Pricing, columns []string) (*Response, bool, error) {
	return c.modifyPricingWithContext(context.Background(), action, pricing, columns)
}

func (c *Client) modifyPricingWithContext(ctx context.Context, action string, pricing *Pricing, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifySubscription is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-subscription`, `update-subscription`, `delete-subscription`,
// Deprecated: Use modifySubscriptionWithContext.
func (c *Client) modifySubscription(action string, subscription *Subscription, columns []string) (*Response, bool, error) {
	return c.modifySubscriptionWithContext(context.Background(), action, subscription, columns)
}

func (c *Client) modifySubscriptionWithContext(ctx context.Context, action string, subscription *Subscription, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifySyner is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-syncer`, `update-syncer`, `delete-syncer`,
// Deprecated: Use modifySyncerWithContext.
func (c *Client) modifySyncer(action string, syncer *Syncer, columns []string) (*Response, bool, error) {
	return c.modifySyncerWithContext(context.Background(), action, syncer, columns)
}

func (c *Client) modifySyncerWithContext(ctx context.Context, action string, syncer *Syncer, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyTransaction is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-transaction`, `update-transaction`, `delete-transaction`,
// Deprecated: Use modifyTransactionWithContext.
func (c *Client) modifyTransaction(action string, transaction *Transaction, columns []string) (*Response, bool, error) {
	return c.modifyTransactionWithContext(context.Background(), action, transaction, columns)
}

func (c *Client) modifyTransactionWithContext(ctx context.Context, action string, transaction *Transaction, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyWebhook is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-webhook`, `update-webhook`, `delete-webhook`,
// Deprecated: Use modifyWebhookWithContext.
func (c *Client) modifyWebhook(action string, webhook *Webhook, columns []string) (*Response, bool, error) {
	return c.modifyWebhookWithContext(context.Background(), action, webhook, columns)
}

func (c *Client) modifyWebhookWithContext(ctx context.Context, action string, webhook *Webhook, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyToken is an encapsulation of cert CUD(Create, Update, Delete) operations.
// possible actions are `add-token`, `update-token`, `delete-token`,
// Deprecated: Use modifyTokenWithContext.
func (c *Client) modifyToken(action string, token *Token, columns []string) (*Response, bool, error) {
	return c.modifyTokenWithContext(context.Background(), action, token, columns)
}

func (c *Client) modifyTokenWithContext(ctx context.Context, action string, token *Token, columns []string) (*Response, bool, error) {
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

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyLdap is an encapsulation of LDAP CUD(Create, Update, Delete) operations.
// possible actions are `add-ldap`, `update-ldap`, `delete-ldap`,
// Deprecated: Use modifyLdapWithContext.
func (c *Client) modifyLdap(action string, ldap *Ldap, columns []string) (*Response, bool, error) {
	return c.modifyLdapWithContext(context.Background(), action, ldap, columns)
}

func (c *Client) modifyLdapWithContext(ctx context.Context, action string, ldap *Ldap, columns []string) (*Response, bool, error) {
	if ldap.Owner == "" {
		ldap.Owner = "admin"
	}

	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", ldap.Owner, ldap.Id),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	postBytes, err := json.Marshal(ldap)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}

// modifyInvitation is an encapsulation of invitation CUD(Create, Update, Delete) operations.
// possible actions are `add-invitation`, `update-invitation`, `delete-invitation`,
// Deprecated: Use modifyInvitationWithContext.
func (c *Client) modifyInvitation(action string, invitation *Invitation, columns []string) (*Response, bool, error) {
	return c.modifyInvitationWithContext(context.Background(), action, invitation, columns)
}

func (c *Client) modifyInvitationWithContext(ctx context.Context, action string, invitation *Invitation, columns []string) (*Response, bool, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", invitation.Owner, invitation.Name),
	}

	if len(columns) != 0 {
		queryMap["columns"] = strings.Join(columns, ",")
	}

	if invitation.Owner == "" {
		invitation.Owner = c.OrganizationName
	}
	postBytes, err := json.Marshal(invitation)
	if err != nil {
		return nil, false, err
	}

	resp, err := c.DoPostWithContext(ctx, action, queryMap, postBytes, false, false)
	if err != nil {
		return nil, false, err
	}

	return resp, resp.Data == "Affected", nil
}
