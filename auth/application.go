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

import "encoding/json"

// Application has the same definition as https://github.com/casdoor/casdoor/blob/master/object/application.go#L24
type Application struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100)" json:"createdTime"`

	DisplayName         string `xorm:"varchar(100)" json:"displayName"`
	Logo                string `xorm:"varchar(100)" json:"logo"`
	HomepageUrl         string `xorm:"varchar(100)" json:"homepageUrl"`
	Description         string `xorm:"varchar(100)" json:"description"`
	Organization        string `xorm:"varchar(100)" json:"organization"`
	Cert                string `xorm:"varchar(100)" json:"cert"`
	EnablePassword      bool   `json:"enablePassword"`
	EnableSignUp        bool   `json:"enableSignUp"`
	EnableSigninSession bool   `json:"enableSigninSession"`
	EnableCodeSignin    bool   `json:"enableCodeSignin"`

	ClientId             string   `xorm:"varchar(100)" json:"clientId"`
	ClientSecret         string   `xorm:"varchar(100)" json:"clientSecret"`
	RedirectUris         []string `xorm:"varchar(1000)" json:"redirectUris"`
	TokenFormat          string   `xorm:"varchar(100)" json:"tokenFormat"`
	ExpireInHours        int      `json:"expireInHours"`
	RefreshExpireInHours int      `json:"refreshExpireInHours"`
	SignupUrl            string   `xorm:"varchar(200)" json:"signupUrl"`
	SigninUrl            string   `xorm:"varchar(200)" json:"signinUrl"`
	ForgetUrl            string   `xorm:"varchar(200)" json:"forgetUrl"`
	AffiliationUrl       string   `xorm:"varchar(100)" json:"affiliationUrl"`
	TermsOfUse           string   `xorm:"varchar(100)" json:"termsOfUse"`
	SignupHtml           string   `xorm:"mediumtext" json:"signupHtml"`
	SigninHtml           string   `xorm:"mediumtext" json:"signinHtml"`
}

func AddApplication(application *Application) (bool, error) {
	if application.Owner == "" {
		application.Owner = "admin"
	}
	postBytes, err := json.Marshal(application)
	if err != nil {
		return false, err
	}

	resp, err := doPost("add-application", nil, postBytes, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}

func DeleteApplication(name string) (bool, error) {
	application := Application{
		Owner: "admin",
		Name:  name,
	}
	postBytes, err := json.Marshal(application)
	if err != nil {
		return false, err
	}

	resp, err := doPost("delete-application", nil, postBytes, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}
