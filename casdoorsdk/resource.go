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

package casdoorsdk

import (
	"encoding/json"
	"fmt"
)

// Resource has the same definition as https://github.com/casdoor/casdoor/blob/master/object/resource.go#L24
type Resource struct {
	Owner string `json:"owner"`
	Name  string `json:"name"`
}

func UploadResource(user string, tag string, parent string, fullFilePath string, fileBytes []byte) (string, string, error) {
	queryMap := map[string]string{
		"owner":        authConfig.OrganizationName,
		"user":         user,
		"application":  authConfig.ApplicationName,
		"tag":          tag,
		"parent":       parent,
		"fullFilePath": fullFilePath,
	}

	resp, err := doPost("upload-resource", queryMap, fileBytes, true)
	if err != nil {
		return "", "", err
	}

	if resp.Status != "ok" {
		return "", "", fmt.Errorf(resp.Msg)
	}

	fileUrl := resp.Data.(string)
	name := resp.Data2.(string)
	return fileUrl, name, nil
}

func UploadResourceEx(user string, tag string, parent string, fullFilePath string, fileBytes []byte, createdTime string, description string) (string, string, error) {
	queryMap := map[string]string{
		"owner":        authConfig.OrganizationName,
		"user":         user,
		"application":  authConfig.ApplicationName,
		"tag":          tag,
		"parent":       parent,
		"fullFilePath": fullFilePath,
		"createdTime":  createdTime,
		"description":  description,
	}

	resp, err := doPost("upload-resource", queryMap, fileBytes, true)
	if err != nil {
		return "", "", err
	}

	if resp.Status != "ok" {
		return "", "", fmt.Errorf(resp.Msg)
	}

	fileUrl := resp.Data.(string)
	name := resp.Data2.(string)
	return fileUrl, name, nil
}

func DeleteResource(name string) (bool, error) {
	resource := Resource{
		Owner: authConfig.OrganizationName,
		Name:  name,
	}
	postBytes, err := json.Marshal(resource)
	if err != nil {
		return false, err
	}

	resp, err := doPost("delete-resource", nil, postBytes, false)
	if err != nil {
		return false, err
	}

	return resp.Status == "ok", nil
}
