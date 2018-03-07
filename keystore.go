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

// Package keystore implements cryptographic key store.
package keystore

// Store - Cryptographic key store.
type Store struct {
	Storage
}

// New - Creates a new keystore from storage.
func New(storage Storage) (store *Store) {
	return &Store{Storage: storage}
}

// CreateKey - Creates new master key in storage and encrypts with password.
func (store *Store) CreateKey(name, mnemonic, password string) (err error) {
	key, err := NewEncryptedKey(mnemonic, password)
	if err != nil {
		return
	}
	return store.Put(name, key)
}

// Has - Naively checks if error occurs on read, returns false if yes.
func (store *Store) Has(name string) bool {
	key, err := store.Get(name)
	return key == nil && err != nil
}
