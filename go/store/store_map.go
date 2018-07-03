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

// NewMapStore - Creates a new map-based key-value store.
func NewMapStore() RawStore {
	return &mapStorage{
		inner: make(map[string][]byte),
	}
}

type mapStorage struct {
	inner map[string][]byte
}

func (store *mapStorage) Has(key string) (bool, error) {
	return store.inner[key] != nil, nil
}

func (store *mapStorage) Keys() (keys []string, _ error) {
	for key := range store.inner {
		keys = append(keys, key)
	}
	return
}

func (store *mapStorage) Get(key string) ([]byte, error) {
	return store.inner[key], nil
}

func (store *mapStorage) Put(key string, value []byte) (_ error) {
	store.inner[key] = value
	return
}
