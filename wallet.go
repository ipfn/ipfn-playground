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

package keywallet

import (
	keystore "github.com/ipfn/go-ipfn-key-store"
)

// Wallet - Storage based key wallet.
type Wallet struct {
	*keystore.Store
}

// NewWallet - Creates new key wallet using storage.
func NewWallet(store *keystore.Store) *Wallet {
	return &Wallet{Store: store}
}

// Derive - Finds master key by name, decrypts using password and derives path.
// Path should be BIP32 compatible otherwise operation will fail.
func (wallet *Wallet) Derive(name, path string, password []byte) (_ *ExtendedKey, err error) {
	key, err := wallet.Decrypt(name, password)
	if err != nil {
		return
	}
	master, err := NewMaster(NewSeed(key, []byte(password)))
	if err != nil {
		return
	}
	return DerivePath(master, path)
}
