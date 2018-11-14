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

package bccsp

import (
	"bytes"
	"fmt"
	"math"
)

// HashFamily - Hash family ID.
// Identifiers are 256bit IDs.
type HashFamily uint64

const (
	// Sha1Family - SHA1 hashing algorithm.
	Sha1Family HashFamily = 0x11
	// Sha2Family - SHA2 hashing algorithm.
	Sha2Family HashFamily = 0x12
	// Sha3Family - SHA3 hashing algorithm.
	Sha3Family HashFamily = 0x16
	// KeccakFamily - Keccak hashing algorithm.
	KeccakFamily HashFamily = 0x1B
	// ShakeFamily - Shake hashing algorithm.
	ShakeFamily HashFamily = 0x19
	// Blake2bFamily - Blake2B hashing algorithm.
	Blake2bFamily HashFamily = 0xb201
	// Blake2sFamily - Blake2S hashing algorithm.
	Blake2sFamily HashFamily = 0xb241
	// DoubleSha2Family - Double SHA2 hashing algorithm.
	DoubleSha2Family HashFamily = 0x56
	// Murmur3Family - MURMUR3 hashing algorithm.
	Murmur3Family HashFamily = 0x22
	// UnknownFamily - Unknown hashing algorithm family.
	UnknownFamily HashFamily = math.MaxUint64
)

func (family HashFamily) String() string {
	switch family {
	case Sha1Family:
		return "sha1"
	case Sha2Family:
		return "sha2"
	case Sha3Family:
		return "sha3"
	case KeccakFamily:
		return "keccak"
	case ShakeFamily:
		return "shake"
	case Blake2bFamily:
		return "blake2b"
	case Blake2sFamily:
		return "blake2s"
	case DoubleSha2Family:
		return "doublesha2"
	case Murmur3Family:
		return "murmur3"
	default:
		return "unknown"
	}
}

// MarshalText - Unmarshals JSON string.
func (family HashFamily) MarshalText() ([]byte, error) {
	return []byte(family.String()), nil
}

// MarshalJSON - Unmarshals JSON string.
func (family HashFamily) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", family)), nil
}

// UnmarshalJSON - Unmarshals JSON string.
func (family *HashFamily) UnmarshalJSON(body []byte) (_ error) {
	return family.UnmarshalText(bytes.Trim(body, `"`))
}

// UnmarshalText - Unmarshals text.
func (family *HashFamily) UnmarshalText(body []byte) (_ error) {
	switch string(bytes.ToLower(body)) {
	case "sha1":
		*family = Sha1Family
		return
	case "sha2":
		*family = Sha2Family
		return
	case "sha3":
		*family = Sha3Family
		return
	case "keccak":
		*family = KeccakFamily
		return
	case "shake":
		*family = ShakeFamily
		return
	case "blake2b":
		*family = Blake2bFamily
		return
	case "blake2s":
		*family = Blake2sFamily
		return
	case "doublesha2":
		*family = DoubleSha2Family
		return
	case "murmur3":
		*family = Murmur3Family
		return
	default:
		return fmt.Errorf("Unknown hash family %q", body)
	}
}
