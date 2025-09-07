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
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	User
	AccessToken string `json:"accessToken"`
	jwt.RegisteredClaims
	TokenType        string `json:"tokenType"`
	RefreshTokenType string `json:"TokenType"`
	SigninMethod     string `json:"signinMethod"`
}

// IsRefreshToken returns true if the token is a refresh token
func (c Claims) IsRefreshToken() bool {
	return c.RefreshTokenType == "refresh-token"
}

func (c *Client) ParseJwtToken(token string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		switch token.Method.Alg() {
		case jwt.SigningMethodES256.Alg():
			return jwt.ParseECPublicKeyFromPEM([]byte(c.Certificate))
		case jwt.SigningMethodES512.Alg():
			return jwt.ParseECPublicKeyFromPEM([]byte(c.Certificate))
		case jwt.SigningMethodRS256.Alg():
			return jwt.ParseRSAPublicKeyFromPEM([]byte(c.Certificate))
		case jwt.SigningMethodRS512.Alg():
			return jwt.ParseRSAPublicKeyFromPEM([]byte(c.Certificate))
		default:
			return nil, fmt.Errorf("unsupported signing method: %v", token.Header["alg"])
		}
	})

	if t != nil {
		if claims, ok := t.Claims.(*Claims); ok && t.Valid {
			return claims, nil
		}
	}

	return nil, err
}
