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

import (
	"encoding/json"

	crypto "github.com/ipfn/go-ipfn-crypto"
)
r
// JSONStorage - JSON key-store wrapper.
type JSONStorage struct {
	RawStorage
}

// NewJSONStorage - Creates a new keystore from storage.
func NewJSONStorage(storage RawStorage) (store *JSONStorage, err error) {
	return &JSONStorage{RawStorage: storage}, nil
}

// Get - Gets encrypted cryptographic key.
func (store *JSONStorage) Get(name string) (key *EncryptedKey, err error) {
	body, err := store.RawStorage.Get(name)
	if err != nil {
		return
	}
	key = new(EncryptedKey)
	err = json.Unmarshal(body, key)
	return
}

// Put - Puts encrypted cryptographic key.
func (store *JSONStorage) Put(name string, body, password []byte) (err error) {
	ciphertext, nonce, salt, err := crypto.Encrypt(body, password)
	if err != nil {
		return
	}
	marshaled, err := json.Marshal(EncryptedKey{
		Ciphertext: ciphertext,
		Nonce:      nonce,
		Salt:       salt,
	})
	if err != nil {
		return
	}
	err = store.RawStorage.Put(name, marshaled)
	return
}
