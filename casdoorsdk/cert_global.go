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

func GetGlobalCerts() ([]*Cert, error) {
	return globalClient.GetGlobalCerts()
}

func GetCerts() ([]*Cert, error) {
	return globalClient.GetCerts()
}

func UpdateCert(cert *Cert) (bool, error) {
	return globalClient.UpdateCert(cert)
}

func AddCert(cert *Cert) (bool, error) {
	return globalClient.AddCert(cert)
}

func DeleteCert(cert *Cert) (bool, error) {
	return globalClient.DeleteCert(cert)
}
