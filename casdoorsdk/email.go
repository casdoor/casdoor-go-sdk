// Copyright 2021 The Casdoor Authors. All Rights Reserved.
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

import "encoding/json"

type emailForm struct {
	Title     string   `json:"title"`
	Content   string   `json:"content"`
	Sender    string   `json:"sender"`
	Receivers []string `json:"receivers"`
}

func (c *Client) SendEmail(title string, content string, sender string, receivers ...string) error {
	form := emailForm{
		Title:     title,
		Content:   content,
		Sender:    sender,
		Receivers: receivers,
	}
	postBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	_, err = c.DoPost("send-email", nil, postBytes, false, false)
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) SendEmailByProvider(title string, content string, sender string, provider string, receivers ...string) error {
	form := emailForm{
		Title:     title,
		Content:   content,
		Sender:    sender,
		Receivers: receivers,
	}
	postBytes, err := json.Marshal(form)
	if err != nil {
		return err
	}

	providerMap := map[string]string{
		"provider": provider,
	}
	_, err = c.DoPost("send-email", providerMap, postBytes, false, false)
	if err != nil {
		return err
	}
	return nil
}
