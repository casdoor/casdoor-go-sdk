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
// See the License for the specific language governing records and
// limitations under the License.

package casdoorsdk

func GetSessions() ([]*Session, error) {
	return globalClient.GetSessions()
}

func GetPaginationSessions(p int, pageSize int, queryMap map[string]string) ([]*Session, int, error) {
	return globalClient.GetPaginationSessions(p, pageSize, queryMap)
}

func GetSession(name string, application string) (*Session, error) {
	return globalClient.GetSession(name, application)
}

func UpdateSession(session *Session) (bool, error) {
	return globalClient.UpdateSession(session)
}

func UpdateSessionForColumns(session *Session, columns []string) (bool, error) {
	return globalClient.UpdateSessionForColumns(session, columns)
}

func AddSession(session *Session) (bool, error) {
	return globalClient.AddSession(session)
}

func DeleteSession(session *Session) (bool, error) {
	return globalClient.DeleteSession(session)
}
