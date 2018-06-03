// Copyright © 2017-2018 Łukasz Kurowski. All Rights Reserved.
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

// Package viperkeys implements viper key-value store.
package viperkeys

import (
	"encoding/hex"
	"fmt"

	"github.com/spf13/viper"

	"github.com/ipfn/ipfn/go/keystore"
)

// Storage - Viper-config based key-value viperkeys.
type Storage struct {
}

// Default - Default store viperkeys.
var Default = keystore.New(New())

// New - Creates new viperkeys.
func New() *Storage {
	return new(Storage)
}

// Get - Gets marshaled encrypted key.
func (store *Storage) Get(name string) (_ *keystore.EncryptedKey, err error) {
	data := viper.GetStringMapString(fmt.Sprintf("seeds.%s", name))
	ciphertext, err := hex.DecodeString(data["ciphertext"])
	if err != nil {
		return
	}
	if len(ciphertext) == 0 {
		return
	}
	salt, err := hex.DecodeString(data["salt"])
	if err != nil {
		return
	}
	nonce, err := hex.DecodeString(data["nonce"])
	if err != nil {
		return
	}
	return &keystore.EncryptedKey{
		Salt:       salt,
		Nonce:      nonce,
		Ciphertext: ciphertext,
	}, nil
}

// Put - Puts marshaled encrypted key.
func (store *Storage) Put(name string, key *keystore.EncryptedKey) error {
	viper.Set(fmt.Sprintf("seeds.%s", name), map[string]interface{}{
		"salt":       fmt.Sprintf("%x", key.Salt),
		"nonce":      fmt.Sprintf("%x", key.Nonce),
		"ciphertext": fmt.Sprintf("%x", key.Ciphertext),
	})
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config: %v", err)
	}
	return nil
}
