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

type ProviderItem struct {
	Owner        string    `json:"owner"`
	Name         string    `json:"name"`
	CanSignUp    bool      `json:"canSignUp"`
	CanSignIn    bool      `json:"canSignIn"`
	CanUnlink    bool      `json:"canUnlink"`
	CountryCodes []string  `json:"countryCodes"`
	Prompted     bool      `json:"prompted"`
	SignupGroup  string    `json:"signupGroup"`
	Rule         string    `json:"rule"`
	Provider     *Provider `json:"provider"`
	// Deprecated: removed from server
	AlertType string `json:"alertType"`
}

type SignupItem struct {
	Name        string   `json:"name"`
	Visible     bool     `json:"visible"`
	Required    bool     `json:"required"`
	Prompted    bool     `json:"prompted"`
	Type        string   `json:"type"`
	CustomCss   string   `json:"customCss"`
	Label       string   `json:"label"`
	Placeholder string   `json:"placeholder"`
	Options     []string `json:"options"`
	Regex       string   `json:"regex"`
	Rule        string   `json:"rule"`
}

type SigninMethod struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	Rule        string `json:"rule"`
}

type SigninItem struct {
	Name        string `json:"name"`
	Visible     bool   `json:"visible"`
	Label       string `json:"label"`
	CustomCss   string `json:"customCss"`
	Placeholder string `json:"placeholder"`
	Rule        string `json:"rule"`
	IsCustom    bool   `json:"isCustom"`
}

type SamlItem struct {
	Name       string `json:"name"`
	NameFormat string `json:"nameFormat"`
	Value      string `json:"value"`
}

type JwtItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

// Application has the same definition as https://github.com/casdoor/casdoor/blob/master/object/application.go#L61
type Application struct {
	Owner                        string          `json:"owner"`
	Name                         string          `json:"name"`
	CreatedTime                  string          `json:"createdTime"`
	DisplayName                  string          `json:"displayName"`
	Logo                         string          `json:"logo"`
	Title                        string          `json:"title"`
	Favicon                      string          `json:"favicon"`
	Order                        int             `json:"order"`
	HomepageUrl                  string          `json:"homepageUrl"`
	Description                  string          `json:"description"`
	Organization                 string          `json:"organization"`
	Cert                         string          `json:"cert"`
	DefaultGroup                 string          `json:"defaultGroup"`
	HeaderHtml                   string          `json:"headerHtml"`
	EnablePassword               bool            `json:"enablePassword"`
	EnableSignUp                 bool            `json:"enableSignUp"`
	DisableSignin                bool            `json:"disableSignin"`
	EnableSigninSession          bool            `json:"enableSigninSession"`
	EnableAutoSignin             bool            `json:"enableAutoSignin"`
	EnableCodeSignin             bool            `json:"enableCodeSignin"`
	EnableExclusiveSignin        bool            `json:"enableExclusiveSignin"`
	EnableSamlCompress           bool            `json:"enableSamlCompress"`
	EnableSamlC14n10             bool            `json:"enableSamlC14n10"`
	EnableSamlPostBinding        bool            `json:"enableSamlPostBinding"`
	DisableSamlAttributes        bool            `json:"disableSamlAttributes"`
	EnableSamlAssertionSignature bool            `json:"enableSamlAssertionSignature"`
	UseEmailAsSamlNameId         bool            `json:"useEmailAsSamlNameId"`
	EnableWebAuthn               bool            `json:"enableWebAuthn"`
	EnableLinkWithEmail          bool            `json:"enableLinkWithEmail"`
	OrgChoiceMode                string          `json:"orgChoiceMode"`
	SamlReplyUrl                 string          `json:"samlReplyUrl"`
	Providers                    []*ProviderItem `json:"providers"`
	SigninMethods                []*SigninMethod `json:"signinMethods"`
	SignupItems                  []*SignupItem   `json:"signupItems"`
	SigninItems                  []*SigninItem   `json:"signinItems"`
	GrantTypes                   []string        `json:"grantTypes"`
	OrganizationObj              *Organization   `json:"organizationObj"`
	CertPublicKey                string          `json:"certPublicKey"`
	Tags                         []string        `json:"tags"`
	SamlAttributes               []*SamlItem     `json:"samlAttributes"`
	SamlHashAlgorithm            string          `json:"samlHashAlgorithm"`
	IsShared                     bool            `json:"isShared"`
	IpRestriction                string          `json:"ipRestriction"`
	ClientId                     string          `json:"clientId"`
	ClientSecret                 string          `json:"clientSecret"`
	RedirectUris                 []string        `json:"redirectUris"`
	ForcedRedirectOrigin         string          `json:"forcedRedirectOrigin"`
	TokenFormat                  string          `json:"tokenFormat"`
	TokenSigningMethod           string          `json:"tokenSigningMethod"`
	TokenFields                  []string        `json:"tokenFields"`
	TokenAttributes              []*JwtItem      `json:"tokenAttributes"`
	ExpireInHours                float64         `json:"expireInHours"`
	RefreshExpireInHours         float64         `json:"refreshExpireInHours"`
	CookieExpireInHours          int64           `json:"cookieExpireInHours"`
	SignupUrl                    string          `json:"signupUrl"`
	SigninUrl                    string          `json:"signinUrl"`
	ForgetUrl                    string          `json:"forgetUrl"`
	AffiliationUrl               string          `json:"affiliationUrl"`
	IpWhitelist                  string          `json:"ipWhitelist"`
	TermsOfUse                   string          `json:"termsOfUse"`
	SignupHtml                   string          `json:"signupHtml"`
	SigninHtml                   string          `json:"signinHtml"`
	ThemeData                    *ThemeData      `json:"themeData"`
	FooterHtml                   string          `json:"footerHtml"`
	FormCss                      string          `json:"formCss"`
	FormCssMobile                string          `json:"formCssMobile"`
	FormOffset                   int             `json:"formOffset"`
	FormSideHtml                 string          `json:"formSideHtml"`
	FormBackgroundUrl            string          `json:"formBackgroundUrl"`
	FormBackgroundUrlMobile      string          `json:"formBackgroundUrlMobile"`
	FailedSigninLimit            int             `json:"failedSigninLimit"`
	FailedSigninFrozenTime       int             `json:"failedSigninFrozenTime"`
	CodeResendTimeout            int             `json:"codeResendTimeout"`
	// Deprecated: removed from server
	CertObj *Cert `json:"certObj"`
}

func (c *Client) GetApplications() ([]*Application, error) {
	queryMap := map[string]string{
		"owner": "admin",
	}

	url := c.GetUrl("get-applications", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var applications []*Application
	err = json.Unmarshal(bytes, &applications)
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func (c *Client) GetOrganizationApplications() ([]*Application, error) {
	queryMap := map[string]string{
		"owner":        "admin",
		"organization": c.OrganizationName,
	}

	url := c.GetUrl("get-organization-applications", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var applications []*Application
	err = json.Unmarshal(bytes, &applications)
	if err != nil {
		return nil, err
	}
	return applications, nil
}

func (c *Client) GetApplication(name string) (*Application, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", "admin", name),
	}

	url := c.GetUrl("get-application", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var application *Application
	err = json.Unmarshal(bytes, &application)
	if err != nil {
		return nil, err
	}
	return application, nil
}

func (c *Client) AddApplication(application *Application) (bool, error) {
	_, affected, err := c.modifyApplication("add-application", application, nil)
	return affected, err
}

func (c *Client) DeleteApplication(application *Application) (bool, error) {
	_, affected, err := c.modifyApplication("delete-application", application, nil)
	return affected, err
}

func (c *Client) UpdateApplication(application *Application) (bool, error) {
	_, affected, err := c.modifyApplication("update-application", application, nil)
	return affected, err
}
