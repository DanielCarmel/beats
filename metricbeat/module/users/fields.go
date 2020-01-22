// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package users

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "users", asset.ModuleFieldsPri, AssetUsers); err != nil {
		panic(err)
	}
}

// AssetUsers returns asset data.
// This is the base64 encoded gzipped contents of ../metricbeat/module/users.
func AssetUsers() string {
	return "eJx8jzsOwyAQRHtOMXLvC2yRLgchYRwhY4MAK/HtI5t8HITyOnbFm9keI1fBkhiTArLNjoJuf3cKiHTUiYILs1aAYbpGG7L1s+CkAJS/mLxZHBUwWDqTZF/1mPXEr34jr4GCW/RLeE0azl/NUWXnwX+GLdtG3fpNM6lQeev4YwU+9BT2W4+UJiPXu4+m2v3J3TgXYQlVzwAAAP//TARqFw=="
}
