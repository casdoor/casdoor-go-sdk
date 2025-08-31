// Copyright 2025 The Casdoor Authors. All Rights Reserved.
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
)

type MfaType string

const (
	EMAIL string = "email"
	SMS   string = "sms"
	APP   string = "app"
)

type MfaRequest struct {
	Owner   string `json:"owner"`
	MfaType string `json:"mfaType"`
	Name    string `json:"name"`
	Secret  string `json:"secret,omitempty"`
}

type MfaInitiateResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   struct {
		Enabled            bool     `json:"enabled"`
		IsPreferred        bool     `json:"isPreferred"`
		MfaRememberInHours int      `json:"mfaRememberInHours"`
		MfaType            string   `json:"mfaType"`
		RecoveryCodes      []string `json:"recoveryCodes"`
		Secret             string   `json:"secret"`
		URL                string   `json:"url"`
	} `json:"data"`
}

type MfaVerifyResponse struct {
	Status string `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
}

func (c *Client) Initiate(owner, mfaType, name string) (*MfaInitiateResponse, error) {
	mfaReq := MfaRequest{
		Owner:   owner,
		MfaType: mfaType,
		Name:    name,
	}

	postBytes, err := json.Marshal(mfaReq)
	if err != nil {
		return nil, err
	}

	resp, err := c.DoPost("mfa/setup/initiate", nil, postBytes, true, false)
	if err != nil {
		return nil, err
	}

	var mfaResp MfaInitiateResponse
	mfaResp.Status = resp.Status
	mfaResp.Msg = resp.Msg

	if resp.Data != nil {
		dataBytes, err := json.Marshal(resp.Data)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(dataBytes, &mfaResp.Data)
		if err != nil {
			return nil, err
		}
	}

	return &mfaResp, nil
}

func (c *Client) Verify(owner, mfaType, name, secret, passcode string) (*MfaVerifyResponse, error) {
	reqData := map[string]string{
		"owner":    owner,
		"mfaType":  mfaType,
		"name":     name,
		"secret":   secret,
		"passcode": passcode,
	}

	postBytes, err := json.Marshal(reqData)
	if err != nil {
		return nil, err
	}

	resp, err := c.DoPost("mfa/setup/verify", nil, postBytes, true, false)
	if err != nil {
		return nil, err
	}

	var mfaResp MfaVerifyResponse
	mfaResp.Status = resp.Status
	mfaResp.Msg = resp.Msg
	if resp.Data != nil {
		if dataStr, ok := resp.Data.(string); ok {
			mfaResp.Data = dataStr
		}
	}

	return &mfaResp, nil
}

func (c *Client) Enable(owner, mfaType, name, secret string) (*MfaVerifyResponse, error) {
	mfaReq := MfaRequest{
		Owner:   owner,
		MfaType: mfaType,
		Name:    name,
		Secret:  secret,
	}

	postBytes, err := json.Marshal(mfaReq)
	if err != nil {
		return nil, err
	}

	resp, err := c.DoPost("mfa/setup/enable", nil, postBytes, true, false)
	if err != nil {
		return nil, err
	}

	dataBytes, err := json.Marshal(resp.Data)
	if err != nil {
		return nil, err
	}

	var mfaResp MfaVerifyResponse
	err = json.Unmarshal(dataBytes, &mfaResp)
	if err != nil {
		return nil, err
	}

	return &mfaResp, nil
}

func (c *Client) SetPreferred(owner, mfaType, name, secret string) error {
	mfaReq := MfaRequest{
		Owner:   owner,
		MfaType: mfaType,
		Name:    name,
		Secret:  secret,
	}

	postBytes, err := json.Marshal(mfaReq)
	if err != nil {
		return err
	}

	_, err = c.DoPost("set-preferred-mfa", nil, postBytes, true, false)
	return err
}

func (c *Client) Delete(owner, name string) error {
	queryMap := map[string]string{
		"owner": owner,
		"name":  name,
	}

	_, err := c.DoPost("delete-mfa", queryMap, nil, true, false)
	return err
}
