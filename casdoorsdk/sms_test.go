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
	"testing"
)

func TestSms(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	sms := &smsForm{
		Content:   "casdoor",
		Receivers: []string{"+8613854673829", "+441932567890"},
	}
	err := SendSms(sms.Content, sms.Receivers...)
	if err != nil {
		t.Fatalf("Failed to send sms: %v", err)
	}

}

func TestSmsByProvider(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	sms := &smsForm{
		Content:   "casdoor",
		Receivers: []string{"+8613854673829", "+441932567890"},
	}
	err := SendSmsByProvider(sms.Content, "provider_casbin_sms", sms.Receivers...)
	if err != nil {
		t.Fatalf("Failed to send sms: %v", err)
	}
}
