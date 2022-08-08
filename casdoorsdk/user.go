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
	"strconv"
)

const (
	UserPropertiesWechatUnionId = "wechatUnionId"
	UserPropertiesWechatOpenId  = "wechatOpenId"
)

// User has the same definition as https://github.com/casdoor/casdoor/blob/master/object/user.go#L24
type User struct {
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	CreatedTime string `json:"createdTime"`
	UpdatedTime string `json:"updatedTime"`

	Id                string   `json:"id"`
	Type              string   `json:"type"`
	Password          string   `json:"password"`
	PasswordSalt      string   `json:"passwordSalt"`
	DisplayName       string   `json:"displayName"`
	Avatar            string   `json:"avatar"`
	PermanentAvatar   string   `json:"permanentAvatar"`
	Email             string   `json:"email"`
	Phone             string   `json:"phone"`
	Location          string   `json:"location"`
	Address           []string `json:"address"`
	Affiliation       string   `json:"affiliation"`
	Title             string   `json:"title"`
	IdCardType        string   `json:"idCardType"`
	IdCard            string   `json:"idCard"`
	Homepage          string   `json:"homepage"`
	Bio               string   `json:"bio"`
	Tag               string   `json:"tag"`
	Region            string   `json:"region"`
	Language          string   `json:"language"`
	Gender            string   `json:"gender"`
	Birthday          string   `json:"birthday"`
	Education         string   `json:"education"`
	Score             int      `json:"score"`
	Karma             int      `json:"karma"`
	Ranking           int      `json:"ranking"`
	IsDefaultAvatar   bool     `json:"isDefaultAvatar"`
	IsOnline          bool     `json:"isOnline"`
	IsAdmin           bool     `json:"isAdmin"`
	IsGlobalAdmin     bool     `json:"isGlobalAdmin"`
	IsForbidden       bool     `json:"isForbidden"`
	IsDeleted         bool     `json:"isDeleted"`
	SignupApplication string   `json:"signupApplication"`
	Hash              string   `json:"hash"`
	PreHash           string   `json:"preHash"`

	CreatedIp      string `json:"createdIp"`
	LastSigninTime string `json:"lastSigninTime"`
	LastSigninIp   string `json:"lastSigninIp"`

	Github   string `json:"github"`
	Google   string `json:"google"`
	QQ       string `json:"qq"`
	WeChat   string `json:"wechat"`
	Facebook string `json:"facebook"`
	DingTalk string `json:"dingtalk"`
	Weibo    string `json:"weibo"`
	Gitee    string `json:"gitee"`
	LinkedIn string `json:"linkedin"`
	Wecom    string `json:"wecom"`
	Lark     string `json:"lark"`
	Gitlab   string `json:"gitlab"`

	Ldap       string            `json:"ldap"`
	Properties map[string]string `json:"properties"`

	Roles       []*Role       `json:"roles"`
	Permissions []*Permission `json:"permissions"`
}

func GetUsers() ([]*User, error) {
	queryMap := map[string]string{
		"owner": authConfig.OrganizationName,
	}

	url := GetUrl("get-users", queryMap)

	bytes, err := DoGetBytesRaw(url)
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

func GetSortedUsers(sorter string, limit int) ([]*User, error) {
	queryMap := map[string]string{
		"owner":  authConfig.OrganizationName,
		"sorter": sorter,
		"limit":  strconv.Itoa(limit),
	}

	url := GetUrl("get-sorted-users", queryMap)

	bytes, err := DoGetBytesRaw(url)
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

func GetUserCount(isOnline string) (int, error) {
	queryMap := map[string]string{
		"owner":    authConfig.OrganizationName,
		"isOnline": isOnline,
	}

	url := GetUrl("get-user-count", queryMap)

	bytes, err := DoGetBytesRaw(url)
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

func GetUser(name string) (*User, error) {
	queryMap := map[string]string{
		"id": fmt.Sprintf("%s/%s", authConfig.OrganizationName, name),
	}

	url := GetUrl("get-user", queryMap)

	bytes, err := DoGetBytesRaw(url)
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

func GetUserByEmail(email string) (*User, error) {
	queryMap := map[string]string{
		"owner": authConfig.OrganizationName,
		"email": email,
	}

	url := GetUrl("get-user", queryMap)

	bytes, err := DoGetBytesRaw(url)
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

func GetUserByPhone(phone string) (*User, error) {
	queryMap := map[string]string{
		"owner": authConfig.OrganizationName,
		"phone": phone,
	}

	url := GetUrl("get-user", queryMap)

	bytes, err := DoGetBytesRaw(url)
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

func GetUserByUserId(userId string) (*User, error) {
	queryMap := map[string]string{
		"owner":  authConfig.OrganizationName,
		"userId": userId,
	}

	url := GetUrl("get-user", queryMap)

	bytes, err := DoGetBytesRaw(url)
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

func UpdateUser(user *User) (bool, error) {
	_, affected, err := modifyUser("update-user", user, nil)
	return affected, err
}

func UpdateUserForColumns(user *User, columns []string) (bool, error) {
	_, affected, err := modifyUser("update-user", user, columns)
	return affected, err
}

func AddUser(user *User) (bool, error) {
	_, affected, err := modifyUser("add-user", user, nil)
	return affected, err
}

func DeleteUser(user *User) (bool, error) {
	_, affected, err := modifyUser("delete-user", user, nil)
	return affected, err
}

func CheckUserPassword(user *User) (bool, error) {
	response, _, err := modifyUser("check-user-password", user, nil)
	return response.Status == "ok", err
}

// GetWeChatIDFromUser return WeChat OpenId and UnionId
func GetWeChatIDFromUser(user *User) (string, string) {
	if user.Properties == nil {
		return "", ""
	}
	return user.Properties[UserPropertiesWechatOpenId], user.Properties[UserPropertiesWechatUnionId]
}
