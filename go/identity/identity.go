// Copyright © 2017 The IPFN Authors. All Rights Reserved.
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

// Package identity implements cryptographic identity utils.
package identity

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"io/ioutil"

	crypto "github.com/libp2p/go-libp2p-crypto"
	peer "github.com/libp2p/go-libp2p-peer"
)

// Identity - Node key pair.
type Identity struct {
	ID   peer.ID
	Pub  crypto.PubKey
	Priv crypto.PrivKey
}

var (
	// ErrInvalidKeyID - Error returned during deserialization of key with invalid ID.
	ErrInvalidKeyID = errors.New("invalid key ID")
)

// New - Same as NewSafe but panics on error.
func New(typ KeyType, size int) *Identity {
	id, err := NewSafe(typ, size)
	if err != nil {
		panic(err)
	}
	return id
}

// NewSafe - Generates a random identity with given size.
func NewSafe(typ KeyType, size int) (_ *Identity, err error) {
	priv, pub, err := crypto.GenerateKeyPairWithReader(int(typ), size, rand.Reader)
	if err != nil {
		return
	}
	id, err := peer.IDFromPublicKey(pub)
	if err != nil {
		return
	}
	return &Identity{
		ID:   id,
		Pub:  pub,
		Priv: priv,
	}, nil
}

// ReadFromFile - Reads keypair from file.
func ReadFromFile(filename string) (_ *Identity, err error) {
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return
	}
	id := new(Identity)
	err = json.Unmarshal(body, id)
	if err != nil {
		return
	}
	return id, nil
}

// UnmarshalJSON - Marshals keypair to JSON.
func (id *Identity) UnmarshalJSON(body []byte) (err error) {
	msg := struct {
		ID   string `json:"id,omitempty"`
		Pub  []byte `json:"pubKey,omitempty"`
		Priv []byte `json:"privKey,omitempty"`
	}{}
	err = json.Unmarshal(body, &msg)
	if err != nil {
		return
	}
	if len(msg.Pub) != 0 {
		id.Pub, err = crypto.UnmarshalPublicKey(msg.Pub)
		if err != nil {
			return
		}
		id.ID, err = peer.IDFromPublicKey(id.Pub)
		if err != nil {
			return
		}
		if len(msg.ID) != 0 && msg.ID != id.ID.Pretty() {
			return ErrInvalidKeyID
		}
	}
	if len(msg.Priv) != 0 {
		id.Priv, err = crypto.UnmarshalPrivateKey(msg.Priv)
		if err != nil {
			return
		}
	}
	return
}

// MarshalJSON - Marshals keypair to JSON.
func (id *Identity) MarshalJSON() (_ []byte, err error) {
	pubKey, err := id.Pub.Bytes()
	if err != nil {
		return
	}
	privKey, err := id.Priv.Bytes()
	if err != nil {
		return
	}
	return json.MarshalIndent(struct {
		ID   string `json:"id,omitempty"`
		Pub  []byte `json:"pubKey,omitempty"`
		Priv []byte `json:"privKey,omitempty"`
	}{
		ID:   id.ID.Pretty(),
		Pub:  pubKey,
		Priv: privKey,
	}, "", "  ")
}
