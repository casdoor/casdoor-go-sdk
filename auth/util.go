// Copyright 2021 The casbin Authors. All Rights Reserved.
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

package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Status string      `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
	Data2  interface{} `json:"data2"`
}

// getBytes is a general function to get response from param url through HTTP Get method.
func getBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

// modifyUser is an encapsulation of user CUD(Create, Update, Delete) operations.
// allowable values of parameter method are `add-user`, `update-user`, `delete-user`,
// get one user information directly through the GetUser function.
func modifyUser(method string, user *User) (*Response, bool, error) {
	user.Owner = authConfig.OrganizationName

	url := fmt.Sprintf("%s/api/%s?id=%s/%s&clientId=%s&clientSecret=%s", authConfig.Endpoint, method, user.Owner, user.Name, authConfig.ClientId, authConfig.ClientSecret)
	userByte, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	resp, err := http.Post(url, "text/plain;charset=UTF-8", bytes.NewReader(userByte))
	if err != nil {
		return nil, false, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	respByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, false, err
	}

	var response Response
	err = json.Unmarshal(respByte, &response)
	if err != nil {
		return nil, false, err
	}

	if response.Data == "Affected" {
		return &response, true, nil
	}
	return &response, false, nil
}
