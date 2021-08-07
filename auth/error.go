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

import "fmt"

type CasdoorError struct {
	Reason string
}

func (e CasdoorError) Error() string {
	return e.Reason
}

func TokenExpiredError(expireTime int64) CasdoorError {
	return CasdoorError{fmt.Sprintf("Token expired at: %d", expireTime)}
}

func TokenInvalidError() CasdoorError {
	return CasdoorError{"This token is invalid. "}
}
