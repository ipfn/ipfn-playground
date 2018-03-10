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

// Package keypair implements cryptographic keypair utils.
package keypair

import (
	"encoding/json"
	"fmt"

	crypto "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
)

// KeyPair - Cryptographic key pair, public and private.
type KeyPair struct {
	peer.ID
	crypto.PrivKey
}

// New - Generates a random identity with given size.
func New(typ KeyType) (_ *KeyPair, err error) {
	switch typ {
	case RSA:
		return NewCustom(typ, 2048)
	default:
		return NewCustom(typ, -1)
	}
}

// NewCustom - Generates a random identity with given size.
func NewCustom(typ KeyType, size int) (_ *KeyPair, err error) {
	priv, pub, err := crypto.GenerateKeyPair(int(typ), size)
	if err != nil {
		return
	}
	id, err := peer.IDFromPublicKey(pub)
	if err != nil {
		return
	}
	return &KeyPair{
		ID:      id,
		PrivKey: priv,
	}, nil
}

// UnmarshalJSON - Marshals keypair to JSON.
func (keys *KeyPair) UnmarshalJSON(body []byte) (err error) {
	msg := struct {
		ID      string `json:"id,omitempty"`
		PrivKey []byte `json:"private_key,omitempty"`
	}{}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return
	}
	if len(msg.PrivKey) != 0 {
		keys.PrivKey, err = crypto.UnmarshalPrivateKey(msg.PrivKey)
		if err != nil {
			return
		}
	}
	if len(msg.ID) != 0 && keys.PrivKey != nil {
		keys.ID, err = peer.IDFromPublicKey(keys.PrivKey.GetPublic())
		if err != nil {
			return
		}
		if len(msg.ID) != 0 && msg.ID != keys.ID.Pretty() {
			return fmt.Errorf("invalid keypair ID %q", msg.ID)
		}
	}
	return
}

// MarshalJSON - Marshals keypair to JSON.
func (keys *KeyPair) MarshalJSON() (_ []byte, err error) {
	privKey, err := keys.PrivKey.Bytes()
	if err != nil {
		return
	}
	return json.MarshalIndent(struct {
		ID      string `json:"id,omitempty"`
		PrivKey []byte `json:"private_key,omitempty"`
	}{
		ID:      keys.ID.Pretty(),
		PrivKey: privKey,
	}, "", "  ")
}
