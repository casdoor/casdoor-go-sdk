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
	"encoding/json"
	"fmt"
)

// Cert has the same definition as https://github.com/casdoor/casdoor/blob/master/object/cert.go#L24
type Cert struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	DisplayName     string `xorm:"varchar(100)" json:"displayName"`
	Scope           string `xorm:"varchar(100)" json:"scope"`
	Type            string `xorm:"varchar(100)" json:"type"`
	CryptoAlgorithm string `xorm:"varchar(100)" json:"cryptoAlgorithm"`
	BitSize         int    `json:"bitSize"`
	ExpireInYears   int    `json:"expireInYears"`

	Certificate            string `xorm:"mediumtext" json:"certificate"`
	PrivateKey             string `xorm:"mediumtext" json:"privateKey"`
	AuthorityPublicKey     string `xorm:"mediumtext" json:"authorityPublicKey"`
	AuthorityRootPublicKey string `xorm:"mediumtext" json:"authorityRootPublicKey"`
}

func (c *Client) GetGlobalCerts() ([]*Cert, error) {
	url := c.GetUrl("get-global-certs", nil)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var certs []*Cert
	err = json.Unmarshal(bytes, &certs)
	if err != nil {
		return nil, err
	}
	return certs, nil
}

func (c *Client) GetCerts() ([]*Cert, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-certs", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var certs []*Cert
	err = json.Unmarshal(bytes, &certs)
	if err != nil {
		return nil, err
	}
	return certs, nil
}

func (c *Client) GetCert(name string) (*Cert, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-cert", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var cert *Cert
	err = json.Unmarshal(bytes, &cert)
	if err != nil {
		return nil, err
	}
	return cert, nil
}

func (c *Client) AddCert(cert *Cert) (bool, error) {
	_, affected, err := c.modifyCert("add-cert", cert, nil)
	return affected, err
}

func (c *Client) UpdateCert(cert *Cert) (bool, error) {
	_, affected, err := c.modifyCert("update-cert", cert, nil)
	return affected, err
}

func (c *Client) DeleteCert(cert *Cert) (bool, error) {
	_, affected, err := c.modifyCert("delete-cert", cert, nil)
	return affected, err
}
