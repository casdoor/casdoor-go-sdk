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
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	modelName := getRandomName("enforceModel")

	affected, err := AddModel(&Model{Owner: "casbin", Name: modelName, DisplayName: modelName, ModelText: `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && r.obj == p.obj && r.act == p.act`})

	if err != nil {
		t.Fatalf("Failed to add model: %v", err.Error())
	}
	if !affected {
		t.Fatalf("Failed to add model")
	}

	adapterName := getRandomName("enforceAdapter")
	affected, err = AddAdapter(&Adapter{Owner: "casbin", Name: adapterName, Table: adapterName + "_policy", UseSameDb: true})
	if err != nil {
		t.Fatalf("Failed to add adapter: %v", err.Error())
	}
	if !affected {
		t.Fatalf("Failed to add adapter")
	}

	enforcerId := getRandomName("enforceEnforcer")
	enforcer := Enforcer{Owner: "casbin", Name: enforcerId, DisplayName: enforcerId, Model: "casbin/" + modelName, Adapter: "casbin/" + adapterName}
	affected, err = AddEnforcer(&enforcer)
	if err != nil {
		t.Fatalf("Failed to add enforcer: %v", err.Error())
	}
	if !affected {
		t.Fatalf("Failed to add enforcer")
	}

	affected, err = AddPolicy(&enforcer, &CasbinRule{Ptype: "p", V0: "alice", V1: "data1", V2: "read"})
	if err != nil {
		t.Fatalf("Failed to add policy: %v", err.Error())
	}
	if !affected {
		t.Fatalf("Failed to add policy")
	}

	affected, err = AddPolicy(&enforcer, &CasbinRule{Ptype: "p", V0: "bob", V1: "data2", V2: "write"})
	if err != nil {
		t.Fatalf("Failed to add policy: %v", err.Error())
	}
	if !affected {
		t.Fatalf("Failed to add policy")
	}

	req1 := CasbinRequest{"alice", "data1", "read"}
	res, err := Enforce("", "", "", "casbin/"+enforcerId, "", req1)
	if err != nil {
		t.Fatalf("Failed to enforce: %v", err.Error())
	}
	if !res {
		t.Fatalf("Enforce fail")
	}

	req2 := CasbinRequest{"bob", "data2", "write"}
	res, err = Enforce("", "", "", "casbin/"+enforcerId, "", req2)
	if err != nil {
		t.Fatalf("Failed to enforce: %v", err.Error())
	}
	if !res {
		t.Fatalf("Enforce fail")
	}

	reqFail := CasbinRequest{"alice", "data1", "write"}
	res, err = Enforce("", "", "", "casbin/"+enforcerId, "", reqFail)
	if err != nil {
		t.Fatalf("Failed to enforce: %v", err.Error())
	}

	if res {
		t.Fatalf("Enforce test fail")
	}

	resBatch, err := BatchEnforce("", "", "", "casbin/"+enforcerId, "", [][]interface{}{req1, reqFail})
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
