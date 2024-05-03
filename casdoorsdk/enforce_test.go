// Copyright 2024 The Casdoor Authors. All Rights Reserved.
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

import "testing"

func TestEnforce(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, "541738959670d221d59d", "66863369a64a5863827cf949bab70ed560ba24bf", TestJwtPublicKey, "built-in", "app-built-in")

	modelName := getRandomName("enforceModel")

	affected, err := AddModel(&Model{Owner: "built-in", Name: modelName, DisplayName: modelName, ModelText: `[request_definition]
r = subOwner, subName, method, urlPath, objOwner, objName

[policy_definition]
p = subOwner, subName, method, urlPath, objOwner, objName

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = (r.subOwner == p.subOwner || p.subOwner == "*") && \
    (r.subName == p.subName || p.subName == "*" || r.subName != "anonymous" && p.subName == "!anonymous") && \
    (r.method == p.method || p.method == "*") && \
    (r.urlPath == p.urlPath || p.urlPath == "*") && \
    (r.objOwner == p.objOwner || p.objOwner == "*") && \
    (r.objName == p.objName || p.objName == "*") || \
    (r.subOwner == r.objOwner && r.subName == r.objName)`})

	if err != nil {
		t.Fatalf("Failed to add model: %v", err.Error())
	}
	if !affected {
		t.Fatalf("Failed to add model")
	}

	adapterName := getRandomName("enforceAdapter")
	affected, err = AddAdapter(&Adapter{Owner: "built-in", Name: adapterName, Table: "casbin_api_rule", UseSameDb: true})
	if err != nil {
		t.Fatalf("Failed to add adapter: %v", err.Error())
	}
	if !affected {
		t.Fatalf("Failed to add adapter")
	}

	enforcerId := getRandomName("enforceEnforcer")
	affected, err = AddEnforcer(&Enforcer{Owner: "built-in", Name: enforcerId, DisplayName: enforcerId, Model: "built-in/" + modelName, Adapter: "built-in/" + adapterName})
	if err != nil {
		t.Fatalf("Failed to add enforcer: %v", err.Error())
	}
	if !affected {
		t.Fatalf("Failed to add enforcer")
	}

	var req []interface{}
	req = append(req, "*", "*", "POST", "/api/signup", "*", "*")

	res, err := Enforce("", "", "", "built-in/"+enforcerId, "", req)
	if err != nil {
		t.Fatalf("Failed to enforce: %v", err.Error())
	}
	if !res {
		t.Fatalf("Enforce fail")
	}

	reqFail := []interface{}{"*", "*", "GET", "/api/sg", "*", ""}
	res, err = Enforce("", "", "", "built-in/"+enforcerId, "", reqFail)
	if err != nil {
		t.Fatalf("Failed to enforce: %v", err.Error())
	}

	if res {
		t.Fatalf("Enforce test fail")
	}

	resBatch, err := BatchEnforce("", "", "", "built-in/"+enforcerId, "", [][]interface{}{req, reqFail})
	if err != nil {
		t.Fatalf("Failed to batchEnforce: %v", err.Error())
	}
	if !resBatch[0][0] {
		t.Fatalf("BatchEnforce test fail")
	}
	if resBatch[0][1] {
		t.Fatalf("BatchEnforce test fail")
	}
}
