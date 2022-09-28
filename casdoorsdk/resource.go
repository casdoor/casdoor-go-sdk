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

import (
	"encoding/json"
	"fmt"
)

// Resource has the same definition as https://github.com/casdoor/casdoor/blob/master/object/resource.go#L24
type Resource struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(250) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	User        string `xorm:"varchar(100)" json:"user"`
	Provider    string `xorm:"varchar(100)" json:"provider"`
	Application string `xorm:"varchar(100)" json:"application"`
	Tag         string `xorm:"varchar(100)" json:"tag"`
	Parent      string `xorm:"varchar(100)" json:"parent"`
	FileName    string `xorm:"varchar(1000)" json:"fileName"`
	FileType    string `xorm:"varchar(100)" json:"fileType"`
	FileFormat  string `xorm:"varchar(100)" json:"fileFormat"`
	FileSize    int    `json:"fileSize"`
	Url         string `xorm:"varchar(1000)" json:"url"`
	Description string `xorm:"varchar(1000)" json:"description"`
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

	resp, err := DoPost("upload-resource", queryMap, fileBytes, true, true)
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

	resp, err := DoPost("upload-resource", queryMap, fileBytes, true, true)
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

	resp, err := DoPost("delete-resource", nil, postBytes, false, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}
