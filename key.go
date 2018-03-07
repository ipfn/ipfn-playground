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

package keystore

import crypto "github.com/ipfn/go-ipfn-crypto"

// EncryptedKey - Encrypted cryptographic key.
type EncryptedKey struct {
	Salt       []byte `json:"salt,omitempty"`
	Nonce      []byte `json:"nonce,omitempty"`
	Ciphertext []byte `json:"ciphertext,omitempty"`
}

// NewEncryptedKey - Creates new encrypted key from private key and password.
func NewEncryptedKey(body, password string) (key *EncryptedKey, err error) {
	ciphertext, nonce, salt, err := crypto.Encrypt([]byte(body), []byte(password))
	if err != nil {
		return
	}
	return &EncryptedKey{
		Ciphertext: ciphertext,
		Nonce:      nonce,
		Salt:       salt,
	}, nil
}

// Decrypt - Decrypts cryptographic key using password encryption key.
func (key *EncryptedKey) Decrypt(password []byte) ([]byte, error) {
	return crypto.Decrypt(key.Ciphertext, key.Nonce, password, key.Salt)
}
