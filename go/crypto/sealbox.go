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

package crypto

import (
	"errors"
)

// SealedBox - Encrypted data.
type SealedBox struct {
	Salt       []byte `json:"salt,omitempty"`
	Nonce      []byte `json:"nonce,omitempty"`
	Ciphertext []byte `json:"ciphertext,omitempty"`
}

// SealBox - Creates new sealbox from data and password.
func SealBox(body, password []byte) (key SealedBox, err error) {
	ciphertext, nonce, salt, err := Encrypt(body, password)
	if err != nil {
		return
	}
	return SealedBox{
		Salt:       salt,
		Nonce:      nonce,
		Ciphertext: ciphertext,
	}, nil
}

// Decrypt - Decrypts data using password encryption key.
func (key SealedBox) Decrypt(password []byte) ([]byte, error) {
	if len(key.Ciphertext) == 0 {
		return nil, errors.New("cannot decrypt empty SealedBox")
	}
	return Decrypt(key.Ciphertext, key.Nonce, password, key.Salt)
}
