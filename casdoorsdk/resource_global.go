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

func GetResource(id string) (*Resource, error) {
	return globalClient.GetResource(id)
}

func GetResourceEx(owner, name string) (*Resource, error) {
	return globalClient.GetResourceEx(owner, name)
}

func GetResources(owner, user, field, value, sortField, sortOrder string) ([]*Resource, error) {
	return globalClient.GetResources(owner, user, field, value, sortField, sortOrder)
}

func GetPaginationResources(owner, user, field, value string, pageSize, page int, sortField, sortOrder string) ([]*Resource, error) {
	return globalClient.GetPaginationResources(owner, user, field, value, pageSize, page, sortField, sortOrder)
}

func UploadResource(user string, tag string, parent string, fullFilePath string, fileBytes []byte) (string, string, error) {
	return globalClient.UploadResource(user, tag, parent, fullFilePath, fileBytes)
}

func UploadResourceEx(user string, tag string, parent string, fullFilePath string, fileBytes []byte, createdTime string, description string) (string, string, error) {
	return globalClient.UploadResourceEx(user, tag, parent, fullFilePath, fileBytes, createdTime, description)
}

func DeleteResource(resource *Resource) (bool, error) {
	return globalClient.DeleteResource(resource)
}
