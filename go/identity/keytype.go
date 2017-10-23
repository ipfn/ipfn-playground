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

package identity

import (
	"errors"

	crypto "github.com/libp2p/go-libp2p-crypto"
)

// KeyType - Type of a key.
type KeyType int

const (
	// RSA - RSA key type.
	RSA KeyType = crypto.RSA
	// Ed25519 - Ed25519 key type.
	Ed25519 KeyType = crypto.Ed25519
	// Secp256k1 - Secp256k1 key type.
	Secp256k1 KeyType = crypto.Secp256k1
	// Unknown - Unknown key type.
	Unknown KeyType = -1
)

var (
	// ErrUnknownKeyType - Error returned when key type is unknown.
	ErrUnknownKeyType = errors.New("unknown key type")
)

// KeyTypeFromString - Returns key type from string.
func KeyTypeFromString(typ string) (KeyType, error) {
	switch typ {
	case "RSA":
		return RSA, nil
	case "Ed25519":
		return Ed25519, nil
	case "Secp256k1":
		return Secp256k1, nil
	}
	return Unknown, ErrUnknownKeyType
}

// String - Returns string representation of a key type.
func (typ KeyType) String() string {
	switch typ {
	case RSA:
		return "RSA"
	case Ed25519:
		return "Ed25519"
	case Secp256k1:
		return "Secp256k1"
	}
	return "Unknown"
}
