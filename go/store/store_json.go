// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"encoding/json"
)

// NewJSONStore - Creates a new JSON marshaled key-value storage.
func NewJSONStore(storage RawStore) EncodedStore {
	return &jsonStore{RawStore: storage}
}

// jsonStore - JSON key-store wrapper.
type jsonStore struct {
	RawStore
}

// Get - Gets unmarshaled value by key.
func (store *jsonStore) Get(name string, value interface{}) (err error) {
	body, err := store.RawStore.Get(name)
	if err != nil {
		return
	}
	return json.Unmarshal(body, value)
}

// Put - Puts marshaled key.
func (store *jsonStore) Put(name string, value interface{}) error {
	marshaled, err := json.Marshal(value)
	if err != nil {
		return err
	}
	return store.RawStore.Put(name, marshaled)
}
