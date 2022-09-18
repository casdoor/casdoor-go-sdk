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

// Application has the same definition as https://github.com/casdoor/casdoor/blob/master/object/application.go#L24
type Application struct {
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	CreatedTime string `json:"createdTime"`

	DisplayName         string `json:"displayName"`
	Logo                string `json:"logo"`
	HomepageUrl         string `json:"homepageUrl"`
	Description         string `json:"description"`
	Organization        string `json:"organization"`
	Cert                string `json:"cert"`
	EnablePassword      bool   `json:"enablePassword"`
	EnableSignUp        bool   `json:"enableSignUp"`
	EnableSigninSession bool   `json:"enableSigninSession"`
	EnableCodeSignin    bool   `json:"enableCodeSignin"`

	ClientId             string   `json:"clientId"`
	ClientSecret         string   `json:"clientSecret"`
	RedirectUris         []string `json:"redirectUris"`
	TokenFormat          string   `json:"tokenFormat"`
	ExpireInHours        int      `json:"expireInHours"`
	RefreshExpireInHours int      `json:"refreshExpireInHours"`
	SignupUrl            string   `json:"signupUrl"`
	SigninUrl            string   `json:"signinUrl"`
	ForgetUrl            string   `json:"forgetUrl"`
	AffiliationUrl       string   `json:"affiliationUrl"`
	TermsOfUse           string   `json:"termsOfUse"`
	SignupHtml           string   `json:"signupHtml"`
	SigninHtml           string   `json:"signinHtml"`
}

func AddApplication(application *Application) (bool, error) {
	if application.Owner == "" {
		application.Owner = "admin"
	}
	postBytes, err := json.Marshal(application)
	if err != nil {
		return false, err
	}

	resp, err := DoPost("add-application", nil, postBytes, false, false)
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

	resp, err := DoPost("delete-application", nil, postBytes, false, false)
	if err != nil {
		return false, err
	}

	return resp.Data == "Affected", nil
}
