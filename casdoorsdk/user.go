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
	"errors"
	"fmt"
	"strconv"
)

const MfaRecoveryCodesSession = "mfa_recovery_codes"

type ManagedAccount struct {
	Application string `xorm:"varchar(100)" json:"application"`
	Username    string `xorm:"varchar(100)" json:"username"`
	Password    string `xorm:"varchar(100)" json:"password"`
	SigninUrl   string `xorm:"varchar(200)" json:"signinUrl"`
}

type MfaProps struct {
	Enabled       bool     `json:"enabled"`
	IsPreferred   bool     `json:"isPreferred"`
	MfaType       string   `json:"mfaType" form:"mfaType"`
	Secret        string   `json:"secret,omitempty"`
	CountryCode   string   `json:"countryCode,omitempty"`
	URL           string   `json:"url,omitempty"`
	RecoveryCodes []string `json:"recoveryCodes,omitempty"`
}

type Userinfo struct {
	Sub         string   `json:"sub"`
	Iss         string   `json:"iss"`
	Aud         string   `json:"aud"`
	Name        string   `json:"preferred_username,omitempty"`
	DisplayName string   `json:"name,omitempty"`
	Email       string   `json:"email,omitempty"`
	Avatar      string   `json:"picture,omitempty"`
	Address     string   `json:"address,omitempty"`
	Phone       string   `json:"phone,omitempty"`
	Groups      []string `json:"groups,omitempty"`
}

// User has the same definition as https://github.com/casdoor/casdoor/blob/master/object/user.go#L24
type User struct {
	Owner       string `xorm:"varchar(100) notnull pk" json:"owner"`
	Name        string `xorm:"varchar(100) notnull pk" json:"name"`
	CreatedTime string `xorm:"varchar(100) index" json:"createdTime"`
	UpdatedTime string `xorm:"varchar(100)" json:"updatedTime"`

	Id                string   `xorm:"varchar(100) index" json:"id"`
	ExternalId        string   `xorm:"varchar(100) index" json:"externalId"`
	Type              string   `xorm:"varchar(100)" json:"type"`
	Password          string   `xorm:"varchar(100)" json:"password"`
	PasswordSalt      string   `xorm:"varchar(100)" json:"passwordSalt"`
	PasswordType      string   `xorm:"varchar(100)" json:"passwordType"`
	DisplayName       string   `xorm:"varchar(100)" json:"displayName"`
	FirstName         string   `xorm:"varchar(100)" json:"firstName"`
	LastName          string   `xorm:"varchar(100)" json:"lastName"`
	Avatar            string   `xorm:"varchar(500)" json:"avatar"`
	AvatarType        string   `xorm:"varchar(100)" json:"avatarType"`
	PermanentAvatar   string   `xorm:"varchar(500)" json:"permanentAvatar"`
	Email             string   `xorm:"varchar(100) index" json:"email"`
	EmailVerified     bool     `json:"emailVerified"`
	Phone             string   `xorm:"varchar(20) index" json:"phone"`
	CountryCode       string   `xorm:"varchar(6)" json:"countryCode"`
	Region            string   `xorm:"varchar(100)" json:"region"`
	Location          string   `xorm:"varchar(100)" json:"location"`
	Address           []string `json:"address"`
	Affiliation       string   `xorm:"varchar(100)" json:"affiliation"`
	Title             string   `xorm:"varchar(100)" json:"title"`
	IdCardType        string   `xorm:"varchar(100)" json:"idCardType"`
	IdCard            string   `xorm:"varchar(100) index" json:"idCard"`
	Homepage          string   `xorm:"varchar(100)" json:"homepage"`
	Bio               string   `xorm:"varchar(100)" json:"bio"`
	Tag               string   `xorm:"varchar(100)" json:"tag"`
	Language          string   `xorm:"varchar(100)" json:"language"`
	Gender            string   `xorm:"varchar(100)" json:"gender"`
	Birthday          string   `xorm:"varchar(100)" json:"birthday"`
	Education         string   `xorm:"varchar(100)" json:"education"`
	Score             int      `json:"score"`
	Karma             int      `json:"karma"`
	Ranking           int      `json:"ranking"`
	IsDefaultAvatar   bool     `json:"isDefaultAvatar"`
	IsOnline          bool     `json:"isOnline"`
	IsAdmin           bool     `json:"isAdmin"`
	IsForbidden       bool     `json:"isForbidden"`
	IsDeleted         bool     `json:"isDeleted"`
	SignupApplication string   `xorm:"varchar(100)" json:"signupApplication"`
	Hash              string   `xorm:"varchar(100)" json:"hash"`
	PreHash           string   `xorm:"varchar(100)" json:"preHash"`
	AccessKey         string   `xorm:"varchar(100)" json:"accessKey"`
	AccessSecret      string   `xorm:"varchar(100)" json:"accessSecret"`

	CreatedIp      string `xorm:"varchar(100)" json:"createdIp"`
	LastSigninTime string `xorm:"varchar(100)" json:"lastSigninTime"`
	LastSigninIp   string `xorm:"varchar(100)" json:"lastSigninIp"`

	GitHub          string `xorm:"github varchar(100)" json:"github"`
	Google          string `xorm:"varchar(100)" json:"google"`
	QQ              string `xorm:"qq varchar(100)" json:"qq"`
	WeChat          string `xorm:"wechat varchar(100)" json:"wechat"`
	Facebook        string `xorm:"facebook varchar(100)" json:"facebook"`
	DingTalk        string `xorm:"dingtalk varchar(100)" json:"dingtalk"`
	Weibo           string `xorm:"weibo varchar(100)" json:"weibo"`
	Gitee           string `xorm:"gitee varchar(100)" json:"gitee"`
	LinkedIn        string `xorm:"linkedin varchar(100)" json:"linkedin"`
	Wecom           string `xorm:"wecom varchar(100)" json:"wecom"`
	Lark            string `xorm:"lark varchar(100)" json:"lark"`
	Gitlab          string `xorm:"gitlab varchar(100)" json:"gitlab"`
	Adfs            string `xorm:"adfs varchar(100)" json:"adfs"`
	Baidu           string `xorm:"baidu varchar(100)" json:"baidu"`
	Alipay          string `xorm:"alipay varchar(100)" json:"alipay"`
	Casdoor         string `xorm:"casdoor varchar(100)" json:"casdoor"`
	Infoflow        string `xorm:"infoflow varchar(100)" json:"infoflow"`
	Apple           string `xorm:"apple varchar(100)" json:"apple"`
	AzureAD         string `xorm:"azuread varchar(100)" json:"azuread"`
	Slack           string `xorm:"slack varchar(100)" json:"slack"`
	Steam           string `xorm:"steam varchar(100)" json:"steam"`
	Bilibili        string `xorm:"bilibili varchar(100)" json:"bilibili"`
	Okta            string `xorm:"okta varchar(100)" json:"okta"`
	Douyin          string `xorm:"douyin varchar(100)" json:"douyin"`
	Line            string `xorm:"line varchar(100)" json:"line"`
	Amazon          string `xorm:"amazon varchar(100)" json:"amazon"`
	Auth0           string `xorm:"auth0 varchar(100)" json:"auth0"`
	BattleNet       string `xorm:"battlenet varchar(100)" json:"battlenet"`
	Bitbucket       string `xorm:"bitbucket varchar(100)" json:"bitbucket"`
	Box             string `xorm:"box varchar(100)" json:"box"`
	CloudFoundry    string `xorm:"cloudfoundry varchar(100)" json:"cloudfoundry"`
	Dailymotion     string `xorm:"dailymotion varchar(100)" json:"dailymotion"`
	Deezer          string `xorm:"deezer varchar(100)" json:"deezer"`
	DigitalOcean    string `xorm:"digitalocean varchar(100)" json:"digitalocean"`
	Discord         string `xorm:"discord varchar(100)" json:"discord"`
	Dropbox         string `xorm:"dropbox varchar(100)" json:"dropbox"`
	EveOnline       string `xorm:"eveonline varchar(100)" json:"eveonline"`
	Fitbit          string `xorm:"fitbit varchar(100)" json:"fitbit"`
	Gitea           string `xorm:"gitea varchar(100)" json:"gitea"`
	Heroku          string `xorm:"heroku varchar(100)" json:"heroku"`
	InfluxCloud     string `xorm:"influxcloud varchar(100)" json:"influxcloud"`
	Instagram       string `xorm:"instagram varchar(100)" json:"instagram"`
	Intercom        string `xorm:"intercom varchar(100)" json:"intercom"`
	Kakao           string `xorm:"kakao varchar(100)" json:"kakao"`
	Lastfm          string `xorm:"lastfm varchar(100)" json:"lastfm"`
	Mailru          string `xorm:"mailru varchar(100)" json:"mailru"`
	Meetup          string `xorm:"meetup varchar(100)" json:"meetup"`
	MicrosoftOnline string `xorm:"microsoftonline varchar(100)" json:"microsoftonline"`
	Naver           string `xorm:"naver varchar(100)" json:"naver"`
	Nextcloud       string `xorm:"nextcloud varchar(100)" json:"nextcloud"`
	OneDrive        string `xorm:"onedrive varchar(100)" json:"onedrive"`
	Oura            string `xorm:"oura varchar(100)" json:"oura"`
	Patreon         string `xorm:"patreon varchar(100)" json:"patreon"`
	Paypal          string `xorm:"paypal varchar(100)" json:"paypal"`
	SalesForce      string `xorm:"salesforce varchar(100)" json:"salesforce"`
	Shopify         string `xorm:"shopify varchar(100)" json:"shopify"`
	Soundcloud      string `xorm:"soundcloud varchar(100)" json:"soundcloud"`
	Spotify         string `xorm:"spotify varchar(100)" json:"spotify"`
	Strava          string `xorm:"strava varchar(100)" json:"strava"`
	Stripe          string `xorm:"stripe varchar(100)" json:"stripe"`
	TikTok          string `xorm:"tiktok varchar(100)" json:"tiktok"`
	Tumblr          string `xorm:"tumblr varchar(100)" json:"tumblr"`
	Twitch          string `xorm:"twitch varchar(100)" json:"twitch"`
	Twitter         string `xorm:"twitter varchar(100)" json:"twitter"`
	Typetalk        string `xorm:"typetalk varchar(100)" json:"typetalk"`
	Uber            string `xorm:"uber varchar(100)" json:"uber"`
	VK              string `xorm:"vk varchar(100)" json:"vk"`
	Wepay           string `xorm:"wepay varchar(100)" json:"wepay"`
	Xero            string `xorm:"xero varchar(100)" json:"xero"`
	Yahoo           string `xorm:"yahoo varchar(100)" json:"yahoo"`
	Yammer          string `xorm:"yammer varchar(100)" json:"yammer"`
	Yandex          string `xorm:"yandex varchar(100)" json:"yandex"`
	Zoom            string `xorm:"zoom varchar(100)" json:"zoom"`
	MetaMask        string `xorm:"metamask varchar(100)" json:"metamask"`
	Web3Onboard     string `xorm:"web3onboard varchar(100)" json:"web3onboard"`
	Custom          string `xorm:"custom varchar(100)" json:"custom"`

	// WebauthnCredentials []webauthn.Credential `xorm:"webauthnCredentials blob" json:"webauthnCredentials"`
	PreferredMfaType string   `xorm:"varchar(100)" json:"preferredMfaType"`
	RecoveryCodes    []string `xorm:"varchar(1000)" json:"recoveryCodes"`
	TotpSecret       string   `xorm:"varchar(100)" json:"totpSecret"`
	MfaPhoneEnabled  bool     `json:"mfaPhoneEnabled"`
	MfaEmailEnabled  bool     `json:"mfaEmailEnabled"`
	// MultiFactorAuths    []*MfaProps           `xorm:"-" json:"multiFactorAuths,omitempty"`

	Ldap       string            `xorm:"ldap varchar(100)" json:"ldap"`
	Properties map[string]string `json:"properties"`

	Roles       []*Role       `json:"roles"`
	Permissions []*Permission `json:"permissions"`
	Groups      []string      `xorm:"groups varchar(1000)" json:"groups"`

	LastSigninWrongTime string `xorm:"varchar(100)" json:"lastSigninWrongTime"`
	SigninWrongTimes    int    `json:"signinWrongTimes"`

	ManagedAccounts []ManagedAccount `xorm:"managedAccounts blob" json:"managedAccounts"`
}

func (c *Client) GetGlobalUsers() ([]*User, error) {
	url := c.GetUrl("get-global-users", nil)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var users []*User
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *Client) GetUsers() ([]*User, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
	}

	url := c.GetUrl("get-users", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var users []*User
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *Client) GetSortedUsers(sorter string, limit int) ([]*User, error) {
	queryMap := map[string]string{
		"owner":  c.OrganizationName,
		"sorter": sorter,
		"limit":  strconv.Itoa(limit),
	}

	url := c.GetUrl("get-sorted-users", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var users []*User
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (c *Client) GetPaginationUsers(p int, pageSize int, queryMap map[string]string) ([]*User, int, error) {
	queryMap["owner"] = c.OrganizationName
	queryMap["p"] = strconv.Itoa(p)
	queryMap["pageSize"] = strconv.Itoa(pageSize)

	url := c.GetUrl("get-users", queryMap)

	response, err := c.DoGetResponse(url)
	if err != nil {
		return nil, 0, err
	}

	dataBytes, err := json.Marshal(response.Data)
	if err != nil {
		return nil, 0, err
	}

	var users []*User
	err = json.Unmarshal(dataBytes, &users)
	if err != nil {
		return nil, 0, errors.New("response data format is incorrect")
	}

	return users, int(response.Data2.(float64)), nil
}

func (c *Client) GetUserCount(isOnline string) (int, error) {
	queryMap := map[string]string{
		"owner":    c.OrganizationName,
		"isOnline": isOnline,
	}

	url := c.GetUrl("get-user-count", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return -1, err
	}

	var count int
	err = json.Unmarshal(bytes, &count)
	if err != nil {
		return -1, err
	}
	return count, nil
}

func (c *Client) GetUser(name string) (*User, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", c.OrganizationName, name),
	}

	url := c.GetUrl("get-user", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Client) GetUserByEmail(email string) (*User, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
		"email": email,
	}

	url := c.GetUrl("get-user", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Client) GetUserByPhone(phone string) (*User, error) {
	queryMap := map[string]string{
		"owner": c.OrganizationName,
		"phone": phone,
	}

	url := c.GetUrl("get-user", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Client) GetUserByUserId(userId string) (*User, error) {
	queryMap := map[string]string{
		"owner":  c.OrganizationName,
		"userId": userId,
	}

	url := c.GetUrl("get-user", queryMap)

	bytes, err := c.DoGetBytes(url)
	if err != nil {
		return nil, err
	}

	var user *User
	err = json.Unmarshal(bytes, &user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// note: oldPassword is not required, if you don't need, just pass a empty string
func (c *Client) SetPassword(owner, name, oldPassword, newPassword string) (bool, error) {
	param := map[string]string{
		"userOwner":   owner,
		"userName":    name,
		"oldPassword": oldPassword,
		"newPassword": newPassword,
	}

	bytes, err := json.Marshal(param)
	if err != nil {
		return false, err
	}

	resp, err := c.DoPost("set-password", nil, bytes, true, false)
	if err != nil {
		return false, err
	}

	return resp.Status == "ok", nil
}

func (c *Client) UpdateUserById(id string, user *User) (bool, error) {
	_, affected, err := c.modifyUserById("update-user", id, user, nil)
	return affected, err
}

func (c *Client) UpdateUser(user *User) (bool, error) {
	_, affected, err := c.modifyUser("update-user", user, nil)
	return affected, err
}

func (c *Client) UpdateUserForColumns(user *User, columns []string) (bool, error) {
	_, affected, err := c.modifyUser("update-user", user, columns)
	return affected, err
}

func (c *Client) AddUser(user *User) (bool, error) {
	_, affected, err := c.modifyUser("add-user", user, nil)
	return affected, err
}

func (c *Client) DeleteUser(user *User) (bool, error) {
	_, affected, err := c.modifyUser("delete-user", user, nil)
	return affected, err
}

func (c *Client) CheckUserPassword(user *User) (bool, error) {
	_, affected, err := c.modifyUser("check-user-password", user, nil)
	return affected, err
}

func (u User) GetId() string {
	return fmt.Sprintf("%s/%s", u.Owner, u.Name)
}
