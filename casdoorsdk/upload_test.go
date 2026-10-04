// Copyright 2026 The Casdoor Authors. All Rights Reserved.
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
	"archive/zip"
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// newTestXlsx builds a minimal Excel (.xlsx) file, whose first sheet has the given rows.
func newTestXlsx(t *testing.T, rows [][]string) []byte {
	t.Helper()

	escape := func(s string) string {
		var b bytes.Buffer
		if err := xml.EscapeText(&b, []byte(s)); err != nil {
			t.Fatalf("Failed to escape the cell: %v", err)
		}
		return b.String()
	}

	var sharedStrings, sheetRows strings.Builder
	count := 0
	for i, row := range rows {
		fmt.Fprintf(&sheetRows, `<row r="%d">`, i+1)
		for j, cell := range row {
			fmt.Fprintf(&sheetRows, `<c r="%c%d" t="s"><v>%d</v></c>`, 'A'+j, i+1, count)
			fmt.Fprintf(&sharedStrings, `<si><t>%s</t></si>`, escape(cell))
			count++
		}
		sheetRows.WriteString(`</row>`)
	}

	files := []struct{ name, content string }{
		{"[Content_Types].xml", `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Types xmlns="http://schemas.openxmlformats.org/package/2006/content-types">
<Default Extension="rels" ContentType="application/vnd.openxmlformats-package.relationships+xml"/>
<Default Extension="xml" ContentType="application/xml"/>
<Override PartName="/xl/workbook.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml"/>
<Override PartName="/xl/worksheets/sheet1.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"/>
<Override PartName="/xl/sharedStrings.xml" ContentType="application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"/>
</Types>`},
		{"_rels/.rels", `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
<Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument" Target="xl/workbook.xml"/>
</Relationships>`},
		{"xl/workbook.xml", `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<workbook xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" xmlns:r="http://schemas.openxmlformats.org/officeDocument/2006/relationships">
<sheets><sheet name="Sheet1" sheetId="1" r:id="rId1"/></sheets>
</workbook>`},
		{"xl/_rels/workbook.xml.rels", `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<Relationships xmlns="http://schemas.openxmlformats.org/package/2006/relationships">
<Relationship Id="rId1" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet" Target="worksheets/sheet1.xml"/>
<Relationship Id="rId2" Type="http://schemas.openxmlformats.org/officeDocument/2006/relationships/sharedStrings" Target="sharedStrings.xml"/>
</Relationships>`},
		{"xl/worksheets/sheet1.xml", `<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<worksheet xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main"><sheetData>` + sheetRows.String() + `</sheetData></worksheet>`},
		{"xl/sharedStrings.xml", fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<sst xmlns="http://schemas.openxmlformats.org/spreadsheetml/2006/main" count="%d" uniqueCount="%d">%s</sst>`, count, count, sharedStrings.String())},
	}

	var buf bytes.Buffer
	w := zip.NewWriter(&buf)
	for _, file := range files {
		f, err := w.Create(file.name)
		if err != nil {
			t.Fatalf("Failed to create %s in the xlsx file: %v", file.name, err)
		}
		if _, err = f.Write([]byte(file.content)); err != nil {
			t.Fatalf("Failed to write %s in the xlsx file: %v", file.name, err)
		}
	}
	if err := w.Close(); err != nil {
		t.Fatalf("Failed to close the xlsx file: %v", err)
	}

	return buf.Bytes()
}

// TestUploadRequest checks the requests sent by the Upload*() functions against a fake
// server, and that a failed import is returned as an error.
func TestUploadRequest(t *testing.T) {
	var gotPath, gotAuth, gotFileName string
	var gotFile []byte
	response := `{"status":"ok","msg":""}`

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		gotAuth = r.Header.Get("Authorization")

		mediaType, params, err := mime.ParseMediaType(r.Header.Get("Content-Type"))
		if err != nil || mediaType != "multipart/form-data" {
			t.Errorf("Unexpected content type: %q, %v", r.Header.Get("Content-Type"), err)
			return
		}

		part, err := multipart.NewReader(r.Body, params["boundary"]).NextPart()
		if err != nil {
			t.Errorf("Failed to read the multipart body: %v", err)
			return
		}
		gotFileName = part.FormName()
		if gotFile, err = io.ReadAll(part); err != nil {
			t.Errorf("Failed to read the file: %v", err)
		}

		_, _ = w.Write([]byte(response))
	}))
	defer server.Close()

	c := NewClient(server.URL, "clientId", "clientSecret", "", "casbin", "app")
	userClient := c.WithAccessToken("userToken")

	fileBytes := []byte("xlsx content")
	tests := []struct {
		action string
		client *Client
		call   func(*Client, []byte) (bool, error)
		auth   string
	}{
		{"upload-users", userClient, (*Client).UploadUsers, "Bearer userToken"},
		{"upload-groups", c, (*Client).UploadGroups, ""},
		{"upload-roles", c, (*Client).UploadRoles, ""},
		{"upload-permissions", c, (*Client).UploadPermissions, ""},
	}
	for _, tt := range tests {
		t.Run(tt.action, func(t *testing.T) {
			response = `{"status":"ok","msg":""}`
			ok, err := tt.call(tt.client, fileBytes)
			if err != nil {
				t.Fatalf("Failed to upload: %v", err)
			}
			if !ok {
				t.Fatalf("Expected the upload to succeed")
			}

			if gotPath != "/api/"+tt.action {
				t.Fatalf("Unexpected path: %s != /api/%s", gotPath, tt.action)
			}
			if gotFileName != "file" || !bytes.Equal(gotFile, fileBytes) {
				t.Fatalf("Unexpected file: the form field is %q and the content is %q", gotFileName, gotFile)
			}
			if tt.auth != "" && gotAuth != tt.auth {
				t.Fatalf("Unexpected Authorization header: %s != %s", gotAuth, tt.auth)
			}
			if tt.auth == "" && !strings.HasPrefix(gotAuth, "Basic ") {
				t.Fatalf("Expected the application's Basic Auth, got %q", gotAuth)
			}

			response = `{"status":"error","msg":"Failed to import"}`
			ok, err = tt.call(tt.client, fileBytes)
			if err == nil || err.Error() != "Failed to import" || ok {
				t.Fatalf("Expected the error from Casdoor, got %v and %v", ok, err)
			}
		})
	}
}

func TestUploadGroups(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	names := []string{getRandomName("UploadGroup"), getRandomName("UploadGroup")}
	file := newTestXlsx(t, [][]string{
		{"Owner", "Name", "CreatedTime", "DisplayName", "Type", "IsEnabled"},
		{TestCasdoorOrganization, names[0], GetCurrentTime(), "Uploaded 1", "Virtual", "1"},
		{TestCasdoorOrganization, names[1], GetCurrentTime(), "Uploaded 2", "Virtual", "1"},
	})

	ok, err := UploadGroups(file)
	if err != nil || !ok {
		t.Fatalf("Failed to upload groups: %v, %v", ok, err)
	}
	for _, name := range names {
		t.Cleanup(func() {
			if _, err := DeleteGroup(&Group{Owner: TestCasdoorOrganization, Name: name}); err != nil {
				t.Errorf("Failed to delete group: %v", err)
			}
		})
	}

	for i, name := range names {
		group, err := GetGroup(name)
		if err != nil || group == nil {
			t.Fatalf("Failed to get the uploaded group %s: %v", name, err)
		}
		if want := fmt.Sprintf("Uploaded %d", i+1); group.DisplayName != want {
			t.Fatalf("Unexpected display name: %s != %s", group.DisplayName, want)
		}
		if !group.IsEnabled {
			t.Fatalf("Expected the group %s to be enabled", name)
		}
	}

	// the groups that already exist are skipped, so there is nothing to import
	ok, err = UploadGroups(file)
	if err == nil || ok {
		t.Fatalf("Expected an error when all the groups already exist, got %v, %v", ok, err)
	}

	// the groups that already exist are skipped, and the new ones are still imported
	newName := getRandomName("UploadGroup")
	file = newTestXlsx(t, [][]string{
		{"Owner", "Name", "CreatedTime", "DisplayName", "Type", "IsEnabled"},
		{TestCasdoorOrganization, names[0], GetCurrentTime(), "Changed", "Virtual", "1"},
		{TestCasdoorOrganization, newName, GetCurrentTime(), "Uploaded 3", "Virtual", "1"},
	})
	ok, err = UploadGroups(file)
	if err != nil || !ok {
		t.Fatalf("Failed to upload a new group along with an existing one: %v, %v", ok, err)
	}
	t.Cleanup(func() {
		if _, err := DeleteGroup(&Group{Owner: TestCasdoorOrganization, Name: newName}); err != nil {
			t.Errorf("Failed to delete group: %v", err)
		}
	})

	group, err := GetGroup(newName)
	if err != nil || group == nil {
		t.Fatalf("Failed to get the uploaded group %s: %v", newName, err)
	}
	if group, err = GetGroup(names[0]); err != nil || group == nil || group.DisplayName != "Uploaded 1" {
		t.Fatalf("Expected the existing group %s to be kept as is, got %+v, %v", names[0], group, err)
	}
}

func TestUploadRoles(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("UploadRole")
	file := newTestXlsx(t, [][]string{
		{"Owner", "Name", "CreatedTime", "DisplayName", "Users", "IsEnabled"},
		{TestCasdoorOrganization, name, GetCurrentTime(), "Uploaded role", `["casbin/admin"]`, "1"},
	})

	ok, err := UploadRoles(file)
	if err != nil || !ok {
		t.Fatalf("Failed to upload roles: %v, %v", ok, err)
	}
	t.Cleanup(func() {
		if _, err := DeleteRole(&Role{Owner: TestCasdoorOrganization, Name: name}); err != nil {
			t.Errorf("Failed to delete role: %v", err)
		}
	})

	role, err := GetRole(name)
	if err != nil || role == nil {
		t.Fatalf("Failed to get the uploaded role: %v", err)
	}
	if role.DisplayName != "Uploaded role" || len(role.Users) != 1 || role.Users[0] != "casbin/admin" || !role.IsEnabled {
		t.Fatalf("The uploaded role is not as expected: %+v", role)
	}

	ok, err = UploadRoles(file)
	if err == nil || ok {
		t.Fatalf("Expected an error when the role already exists, got %v, %v", ok, err)
	}
}

func TestUploadPermissions(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	name := getRandomName("UploadPermission")
	file := newTestXlsx(t, [][]string{
		{"Owner", "Name", "CreatedTime", "DisplayName", "Users", "ResourceType", "Resources", "Actions", "Effect", "IsEnabled", "State"},
		{TestCasdoorOrganization, name, GetCurrentTime(), "Uploaded permission", `["casbin/admin"]`, "Application", `["app-casibase"]`, `["Read","Write"]`, "Allow", "1", "Approved"},
		// the rows without an owner or a name are ignored
		{"", "", GetCurrentTime(), "Invalid", "", "", "", "", "", "", ""},
	})

	ok, err := UploadPermissions(file)
	if err != nil || !ok {
		t.Fatalf("Failed to upload permissions: %v, %v", ok, err)
	}
	t.Cleanup(func() {
		if _, err := DeletePermission(&Permission{Owner: TestCasdoorOrganization, Name: name}); err != nil {
			t.Errorf("Failed to delete permission: %v", err)
		}
	})

	permission, err := GetPermission(name)
	if err != nil || permission == nil {
		t.Fatalf("Failed to get the uploaded permission: %v", err)
	}
	if permission.DisplayName != "Uploaded permission" || permission.Effect != "Allow" || !permission.IsEnabled {
		t.Fatalf("The uploaded permission is not as expected: %+v", permission)
	}
	if len(permission.Actions) != 2 || permission.Actions[0] != "Read" || len(permission.Resources) != 1 || permission.Resources[0] != "app-casibase" {
		t.Fatalf("The lists of the uploaded permission are not as expected: %+v", permission)
	}

	ok, err = UploadPermissions(file)
	if err == nil || ok {
		t.Fatalf("Expected an error when the permission already exists, got %v, %v", ok, err)
	}
}

func TestUploadUsers(t *testing.T) {
	InitConfig(TestCasdoorEndpoint, TestClientId, TestClientSecret, TestJwtPublicKey, TestCasdoorOrganization, TestCasdoorApplication)

	// the upload-users API acts on behalf of an admin user, not of the application
	token, err := GetOAuthTokenByPassword("admin", "123")
	if err != nil {
		t.Fatalf("Failed to get the token of the admin: %v", err)
	}
	userClient := WithAccessToken(token.AccessToken)

	names := []string{getRandomName("UploadUser"), getRandomName("UploadUser")}
	file := newTestXlsx(t, [][]string{
		{"Owner", "Name", "CreatedTime", "DisplayName", "Email", "IsAdmin", "Groups"},
		{TestCasdoorOrganization, names[0], GetCurrentTime(), "Uploaded 1", names[0] + "@example.com", "1", `["casbin/group1"]`},
		{TestCasdoorOrganization, names[1], GetCurrentTime(), "Uploaded 2", names[1] + "@example.com", "", ""},
		// Excel keeps blank rows in a sheet, they are skipped
		{"", "", "", "", "", "", ""},
	})

	// the application's credentials are not enough
	if ok, err := UploadUsers(file); err == nil || ok {
		t.Fatalf("Expected the upload as the application to be rejected, got %v, %v", ok, err)
	}

	ok, err := userClient.UploadUsers(file)
	if err != nil || !ok {
		t.Fatalf("Failed to upload users: %v, %v", ok, err)
	}
	for _, name := range names {
		t.Cleanup(func() {
			if _, err := DeleteUser(&User{Owner: TestCasdoorOrganization, Name: name}); err != nil {
				t.Errorf("Failed to delete user: %v", err)
			}
		})
	}

	for i, name := range names {
		user, err := GetUser(name)
		if err != nil || user == nil {
			t.Fatalf("Failed to get the uploaded user %s: %v", name, err)
		}
		if want := fmt.Sprintf("Uploaded %d", i+1); user.DisplayName != want || user.Email != name+"@example.com" {
			t.Fatalf("The uploaded user is not as expected: %+v", user)
		}
		if user.IsAdmin != (i == 0) {
			t.Fatalf("Unexpected IsAdmin of the user %s: %v", name, user.IsAdmin)
		}
	}

	// the users that already exist make the import fail
	if ok, err = userClient.UploadUsers(file); err == nil || ok {
		t.Fatalf("Expected an error when the users already exist, got %v, %v", ok, err)
	}
}
