// Copyright 2022 The Casdoor Authors. All Rights Reserved.
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
// See the License for the specific language governing records and
// limitations under the License.

package casdoorsdk

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
)

type Record struct {
	Id int `xorm:"int notnull pk autoincr" json:"id"`

	Owner       string `xorm:"varchar(100) index" json:"owner"`
	Name        string `xorm:"varchar(100) index" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	Organization string `xorm:"varchar(100)" json:"organization"`
	ClientIp     string `xorm:"varchar(100)" json:"clientIp"`
	User         string `xorm:"varchar(100)" json:"user"`
	Method       string `xorm:"varchar(100)" json:"method"`
	RequestUri   string `xorm:"varchar(1000)" json:"requestUri"`
	Action       string `xorm:"varchar(1000)" json:"action"`

	Object       string `xorm:"-" json:"object"`
	ExtendedUser *User  `xorm:"-" json:"extendedUser"`

	IsTriggered bool `json:"isTriggered"`
}

// Deprecated: Use GetRecordsWithContext.
func (c *Client) GetRecords() ([]*Record, error) {
	return c.GetRecordsWithContext(context.Background())
}

func (c *Client) GetRecordsWithContext(ctx context.Context) ([]*Record, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-records", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	var records []*Record
	err = json.Unmarshal(bytes, &records)
	if err != nil {
		return nil, err
	}
	return records, nil
}

// Deprecated: Use GetPaginationRecordsWithContext.
func (c *Client) GetPaginationRecords(p int, pageSize int, queryMap map[string]string) ([]*Record, int, error) {
	return c.GetPaginationRecordsWithContext(context.Background(), p, pageSize, queryMap)
}

func (c *Client) GetPaginationRecordsWithContext(ctx context.Context, p int, pageSize int, queryMap map[string]string) ([]*Record, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-records", queryMap)

	response, err := c.DoGetResponseWithContext(ctx, url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var records []*Record
	err = json.Unmarshal(dataBytes, &records)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return records, int(response.Data2.(float64)), nil
}

// Deprecated: Use GetRecordWithContext.
func (c *Client) GetRecord(name string) (*Record, error) {
	return c.GetRecordWithContext(context.Background(), name)
}

func (c *Client) GetRecordWithContext(ctx context.Context, name string) (*Record, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-record", queryMap)

	bytes, err := c.DoGetBytesWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	var record *Record
	err = json.Unmarshal(bytes, &record)
	if err != nil {
		return nil, err
	}
	return record, nil
}

// Deprecated: Use AddRecordWithContext.
func (c *Client) AddRecord(record *Record) (bool, error) {
	return c.AddRecordWithContext(context.Background(), record)
}

func (c *Client) AddRecordWithContext(ctx context.Context, record *Record) (bool, error) {
	if record.Owner == "" {
		record.Owner = c.OrganizationName
	}
	if record.Organization == "" {
		record.Organization = c.OrganizationName
	}

	postBytes, err := json.Marshal(record)
	if err != nil {
		return false, err
	}

	resp, err := c.DoPostWithContext(ctx, "add-record", nil, postBytes, false, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}
