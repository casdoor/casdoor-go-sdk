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
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func (c *Client) GetUrl(action string, queryMap map[string]string) string {
	query := ""
	for k, v := range queryMap {
		query += fmt.Sprintf("%s=%s&", url.QueryEscape(k), url.QueryEscape(v))
	}
	query = strings.TrimRight(query, "&")

	return fmt.Sprintf("%s/api/%s?%s", c.Endpoint, action, query)
}

func (c *Client) GetId(name string) string {
	return c.OrganizationName + "/" + name
}

func createFormFile(formData map[string][]byte) (string, io.Reader, error) {
	// https://tonybai.com/2021/01/16/upload-and-download-file-using-multipart-form-over-http/

	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	defer w.Close()

	for k, v := range formData {
		pw, err := w.CreateFormFile(k, "file")
		if err != nil {
			panic(err)
		}

		_, err = pw.Write(v)
		if err != nil {
			panic(err)
		}
	}

	return w.FormDataContentType(), body, nil
}

func createForm(formData map[string]string) (string, io.Reader, error) {
	body := new(bytes.Buffer)
	w := multipart.NewWriter(body)
	for k, v := range formData {
		if err := w.WriteField(k, v); err != nil {
			return "", nil, err
		}
	}
	if err := w.Close(); err != nil {
		return "", nil, err
	}

	return w.FormDataContentType(), body, nil
}

func GetCurrentTime() string {
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	return tm.Format(time.RFC3339)
}

// DoGetResponse is a general function to get response from param url through HTTP Get method.

// Deprecated: Use DoGetResponseWithContext.
func (c *Client) DoGetResponse(url string) (*Response, error) {
	return c.DoGetResponseWithContext(context.Background(), url)
}

// DoGetResponseWithContext is a general function to get response from param url through HTTP Get method with context.
func (c *Client) DoGetResponseWithContext(ctx context.Context, url string) (*Response, error) {
	respBytes, err := c.doGetBytesRawWithoutCheckWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf(response.Msg)
	}

	return &response, nil
}

// DoGetBytes is a general function to get response data in bytes from param url through HTTP Get method.
// Deprecated: Use DoGetBytesWithContext.
func (c *Client) DoGetBytes(url string) ([]byte, error) {
	return c.DoGetBytesWithContext(context.Background(), url)
}

// DoGetBytesWithContext is a general function to get response data in bytes from param url through HTTP Get method with context.
func (c *Client) DoGetBytesWithContext(ctx context.Context, url string) ([]byte, error) {
	response, err := c.DoGetResponseWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	res, err := json.Marshal(response.Data)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// DoGetBytesRaw is a general function to get response from param url through HTTP Get method.
// Deprecated: Use DoGetBytesRawWithContext.
func (c *Client) DoGetBytesRaw(url string) ([]byte, error) {
	return c.DoGetBytesRawWithContext(context.Background(), url)
}

// DoGetBytesRawWithContext is a general function to get response from param url through HTTP Get method with context.
func (c *Client) DoGetBytesRawWithContext(ctx context.Context, url string) ([]byte, error) {
	respBytes, err := c.doGetBytesRawWithoutCheckWithContext(ctx, url)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(respBytes, &response)
	if err == nil && response.Status == "error" {
		return nil, errors.New(response.Msg)
	}

	return respBytes, nil
}

// Deprecated: Use DoPostWithContext.
func (c *Client) DoPost(action string, queryMap map[string]string, postBytes []byte, isForm, isFile bool) (*Response, error) {
	return c.DoPostWithContext(context.Background(), action, queryMap, postBytes, isForm, isFile)
}

// DoPostWithContext is a general function to post a request with context support.
func (c *Client) DoPostWithContext(ctx context.Context, action string, queryMap map[string]string, postBytes []byte, isForm, isFile bool) (*Response, error) {
	url := c.GetUrl(action, queryMap)

	var err error
	var contentType string
	var body io.Reader
	if isForm {
		if isFile {
			contentType, body, err = createFormFile(map[string][]byte{"file": postBytes})
			if err != nil {
				return nil, err
			}
		} else {
			var params map[string]string
			err = json.Unmarshal(postBytes, &params)
			if err != nil {
				return nil, err
			}

			contentType, body, err = createForm(params)
			if err != nil {
				return nil, err
			}
		}
	} else {
		contentType = "text/plain;charset=UTF-8"
		body = bytes.NewReader(postBytes)
	}

	respBytes, err := c.DoPostBytesRawWithContext(ctx, url, contentType, body)
	if err != nil {
		return nil, err
	}

	var response Response
	err = json.Unmarshal(respBytes, &response)
	if err != nil {
		return nil, err
	}

	if response.Status != "ok" {
		return nil, fmt.Errorf(response.Msg)
	}

	return &response, nil
}

// DoPostBytesRaw is a general function to post a request from url, body through HTTP Post method.
// Deprecated: Use DoPostBytesRawWithContext.
func (c *Client) DoPostBytesRaw(url string, contentType string, body io.Reader) ([]byte, error) {
	return c.DoPostBytesRawWithContext(context.Background(), url, contentType, body)
}

// DoPostBytesRawWithContext is a general function to post a request from url, body through HTTP Post method with context.
func (c *Client) DoPostBytesRawWithContext(ctx context.Context, url string, contentType string, body io.Reader) ([]byte, error) {
	if contentType == "" {
		contentType = "text/plain;charset=UTF-8"
	}

	var resp *http.Response

	req, err := http.NewRequestWithContext(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.ClientId, c.ClientSecret)
	req.Header.Set("Content-Type", contentType)

	resp, err = client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {
		return nil, fmt.Errorf("%s", string(respBytes))
	}

	return respBytes, nil
}

// doGetBytesRawWithoutCheck is a general function to get response from param url through HTTP Get method without checking response status
// Deprecated: Use doGetBytesRawWithoutCheckWithContext.
func (c *Client) doGetBytesRawWithoutCheck(url string) ([]byte, error) {
	return c.doGetBytesRawWithoutCheckWithContext(context.Background(), url)
}

// doGetBytesRawWithoutCheckWithContext is a general function to get response from param url through HTTP Get method without checking response status, with context support.
func (c *Client) doGetBytesRawWithoutCheckWithContext(ctx context.Context, url string) ([]byte, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.ClientId, c.ClientSecret)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)

	respBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusForbidden {
		return nil, fmt.Errorf("%s", string(respBytes))
	}

	return respBytes, nil
}
