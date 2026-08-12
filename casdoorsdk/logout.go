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
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// Logout logs out the user that owns the accessToken by calling the "/api/sso-logout" API.
// It performs a single sign-out: all the sessions of the user across all the applications
// of the organization are deleted and all the access tokens issued to the user are expired.
// The accessToken is the user's access token returned by GetOAuthToken().
func (c *Client) Logout(accessToken string) error {
	return c.ssoLogout(accessToken, true)
}

// LogoutCurrentSession is like Logout() but only ends the session that the accessToken
// belongs to, so the user stays signed in on their other devices and browsers.
func (c *Client) LogoutCurrentSession(accessToken string) error {
	return c.ssoLogout(accessToken, false)
}

func (c *Client) ssoLogout(accessToken string, logoutAll bool) error {
	if accessToken == "" {
		return errors.New("Logout() error: the accessToken should not be empty")
	}

	url := c.GetUrl("sso-logout", map[string]string{
		"logoutAll": strconv.FormatBool(logoutAll),
	})

	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return err
	}

	// the "/api/sso-logout" API identifies the user by their own access token,
	// so the Bearer token is used here instead of the application's Basic Auth
	req.Header.Set("Authorization", "Bearer "+accessToken)

	// Add custom headers
	for key, value := range c.CustomHeaders {
		req.Header.Set(key, value)
	}

	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {
		return fmt.Errorf("status code: %d, status: %s, body: %s", resp.StatusCode, resp.Status, string(respBytes))
	}

	var response Response
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return err
	}

	if response.Status != "ok" {
		return errors.New(response.Msg)
	}

	return nil
}
