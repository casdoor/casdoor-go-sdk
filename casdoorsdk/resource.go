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
	"context"
	"encoding/json"
	"fmt"
	"strconv"
)

// Resource has the same definition as https://github.com/casdoor/casdoor/blob/master/object/resource.go#L24
type Resource struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(180) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	User        string `xorm:"varchar(100)" json:"user"`
	Provider    string `xorm:"varchar(100)" json:"provider"`
	Application string `xorm:"varchar(100)" json:"application"`
	Tag         string `xorm:"varchar(100)" json:"tag"`
	Parent      string `xorm:"varchar(100)" json:"parent"`
	FileName    string `xorm:"varchar(255)" json:"fileName"`
	FileType    string `xorm:"varchar(100)" json:"fileType"`
	FileFormat  string `xorm:"varchar(100)" json:"fileFormat"`
	FileSize    int    `json:"fileSize"`
	Url         string `xorm:"varchar(255)" json:"url"`
	Description string `xorm:"varchar(255)" json:"description"`
}

// Deprecated: Use GetResourceWithContext.
func (c *Client) GetResource(id string) (*Resource, error) {
	return c.GetResourceWithContext(context.Background(), id)
}

func (c *Client) GetResourceWithContext(ctx context.Context, id string) (*Resource, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
		"id":    id,
	}

	url := c.GetUrl("get-resource", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	var resource *Resource
	err = json.Unmarshal(bytes, &resource)
	if err != nil {
		return nil, err
	}

	return resource, nil
}

func (c *Client) GetResourceEx(owner, name string) (*Resource, error) {
	return c.GetResource(fmt.Sprintf("%s/%s", owner, name))
}

// Deprecated: Use GetResourcesWithContext.
func (c *Client) GetResources(owner, user, field, value, sortField, sortOrder string) ([]*Resource, error) {
	return c.GetResourcesWithContext(context.Background(), owner, user, field, value, sortField, sortOrder)
}

func (c *Client) GetResourcesWithContext(ctx context.Context, owner, user, field, value, sortField, sortOrder string) ([]*Resource, error) {
	queryMap := map[string]string{
		"owner":     owner,
		"user":      user,
		"field":     field,
		"value":     value,
		"sortField": sortField,
		"sortOrder": sortOrder,
	}

	url := c.GetUrl("get-resources", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	var resources []*Resource
	err = json.Unmarshal(bytes, &resources)
	if err != nil {
		return nil, err
	}
	return resources, nil
}

// Deprecated: Use GetPaginationResourcesWithContext.
func (c *Client) GetPaginationResources(owner, user, field, value string, pageSize, page int, sortField, sortOrder string) ([]*Resource, error) {
	return c.GetPaginationResourcesWithContext(context.Background(), owner, user, field, value, pageSize, page, sortField, sortOrder)
}

func (c *Client) GetPaginationResourcesWithContext(ctx context.Context, owner, user, field, value string, pageSize, page int, sortField, sortOrder string) ([]*Resource, error) {
	queryMap := map[string]string{
		"owner":     owner,
		"user":      user,
		"field":     field,
		"value":     value,
		"p":         strconv.Itoa(page),
		"pageSize":  strconv.Itoa(pageSize),
		"sortField": sortField,
		"sortOrder": sortOrder,
	}

	url := c.GetUrl("get-resources", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	var resources []*Resource
	err = json.Unmarshal(bytes, &resources)
	if err != nil {
		return nil, err
	}
	return resources, nil
}

// Deprecated: Use UploadResourceWithContext.
func (c *Client) UploadResource(user string, tag string, parent string, fullFilePath string, fileBytes []byte) (string, string, error) {
	return c.UploadResourceWithContext(context.Background(), user, tag, parent, fullFilePath, fileBytes)
}

func (c *Client) UploadResourceWithContext(ctx context.Context, user string, tag string, parent string, fullFilePath string, fileBytes []byte) (string, string, error) {
	queryMap := map[string]string{
		"owner":        c.OrganizationName,
		"user":         user,
		"application":  c.ApplicationName,
		"tag":          tag,
		"parent":       parent,
		"fullFilePath": fullFilePath,
	}

	resp, err := c.DoPostWithContext(ctx, "upload-resource", queryMap, fileBytes, true, true)
	if err != nil {
		return "", "", err
	}

	fileUrl := resp.Data.(string)
	name := resp.Data2.(string)
	return fileUrl, name, nil
}

// Deprecated: Use UploadResourceExWithContext.
func (c *Client) UploadResourceEx(user string, tag string, parent string, fullFilePath string, fileBytes []byte, createdTime string, description string) (string, string, error) {
	return c.UploadResourceExWithContext(context.Background(), user, tag, parent, fullFilePath, fileBytes, createdTime, description)
}

func (c *Client) UploadResourceExWithContext(ctx context.Context, user string, tag string, parent string, fullFilePath string, fileBytes []byte, createdTime string, description string) (string, string, error) {
	queryMap := map[string]string{
		"owner":        c.OrganizationName,
		"user":         user,
		"application":  c.ApplicationName,
		"tag":          tag,
		"parent":       parent,
		"fullFilePath": fullFilePath,
		"createdTime":  createdTime,
		"description":  description,
	}

	resp, err := c.DoPostWithContext(ctx, "upload-resource", queryMap, fileBytes, true, true)
	if err != nil {
		return "", "", err
	}

	fileUrl := resp.Data.(string)
	name := resp.Data2.(string)
	return fileUrl, name, nil
}

// Deprecated: Use DeleteResourceWithTagWithContext (or DeleteResourceWithTag).
func (c *Client) DeleteResource(resource *Resource) (bool, error) {
	return c.DeleteResourceWithTag(resource, "")
}

// Deprecated: Use DeleteResourceWithTagWithContext.
func (c *Client) DeleteResourceWithTag(resource *Resource, tag string) (bool, error) {
	return c.DeleteResourceWithTagWithContext(context.Background(), resource, tag)
}

func (c *Client) DeleteResourceWithTagWithContext(ctx context.Context, resource *Resource, tag string) (bool, error) {
	if resource.Owner == "" {
		resource.Owner = c.OrganizationName
	}

	queryMap := map[string]string{
		"tag": tag,
	}

	postBytes, err := json.Marshal(resource)
	if err != nil {
		return false, err
	}

	resp, err := c.DoPostWithContext(ctx, "delete-resource", queryMap, postBytes, false, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}
