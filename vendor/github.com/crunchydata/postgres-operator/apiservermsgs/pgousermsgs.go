package apiservermsgs

/*
Copyright 2019 Crunchy Data Solutions, Inc.
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

import ()

// ShowPgouserRequest ...
type ShowPgouserRequest struct {
	Namespace     string
	AllFlag       bool
	ClientVersion string
	PgouserName   []string
}

type PgouserInfo struct {
	Username  string
	Role      []string
	Namespace []string
}

// ShowPgouserResponse ...
type ShowPgouserResponse struct {
	UserInfo []PgouserInfo
	Status
}

// CreatePgouserRequest ...
type CreatePgouserRequest struct {
	PgouserName       string
	PgouserPassword   string
	PgouserRoles      string
	AllNamespaces     bool
	PgouserNamespaces string
	Namespace         string
	ClientVersion     string
}

// CreatePgouserResponse ...
type CreatePgouserResponse struct {
	Status
}

// UpdatePgouserRequest ...
type UpdatePgouserRequest struct {
	Name              string
	PgouserRoles      string
	PgouserNamespaces string
	AllNamespaces     bool
	PgouserPassword   string
	PgouserName       string
	Namespace         string
	ClientVersion     string
}

// ApplyPgouserResponse ...
type UpdatePgouserResponse struct {
	Status
}

// DeletePgouserRequest ...
type DeletePgouserRequest struct {
	PgouserName   []string
	Namespace     string
	AllFlag       bool
	ClientVersion string
}

// DeletePgouserResponse ...
type DeletePgouserResponse struct {
	Results []string
	Status
}
