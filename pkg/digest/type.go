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

package digest

import (
	"fmt"
)

// Type - Multihash algorithm ID.
type Type uint64

// Source of constants: https://godoc.org/github.com/multiformats/go-multihash#pkg-constants
const (
	// Sha1 - SHA1 hashing algorithm.
	Sha1 Type = 0x11
	// Sha2_256 - SHA2 256bit hashing algorithm.
	Sha2_256 Type = 0x12
	// Sha2_512 - SHA2 512bit hashing algorithm.
	Sha2_512 Type = 0x13
	// Sha3_224 - SHA3 224bit hashing algorithm.
	Sha3_224 Type = 0x17
	// Sha3_256 - SHA3 256bit hashing algorithm.
	Sha3_256 Type = 0x16
	// Sha3_384 - SHA3 384bit hashing algorithm.
	Sha3_384 Type = 0x15
	// Sha3_512 - SHA3 512bit hashing algorithm.
	Sha3_512 Type = 0x14
	// Keccak224 - Keccak 224bit hashing algorithm.
	Keccak224 Type = 0x1A
	// Keccak256 - Keccak 256bit hashing algorithm.
	Keccak256 Type = 0x1B
	// Keccak384 - Keccak 384bit hashing algorithm.
	Keccak384 Type = 0x1C
	// Keccak512 - Keccak 512bit hashing algorithm.
	Keccak512 Type = 0x1D
	// Shake128 - Shake 128bit hashing algorithm.
	Shake128 Type = 0x18
	// Shake256 - Shake 256bit hashing algorithm.
	Shake256 Type = 0x19
	// Blake2bMin - Blake2B MIN hashing algorithm.
	Blake2bMin Type = 0xb201
	// Blake2bMax - Blake2B MAX hashing algorithm.
	Blake2bMax Type = 0xb240
	// Blake2sMin - Blake2S MIN hashing algorithm.
	Blake2sMin Type = 0xb241
	// Blake2sMax - Blake2S MAX hashing algorithm.
	Blake2sMax Type = 0xb260
	// DoubleSha2_256 - Double SHA2 256bit hashing algorithm.
	DoubleSha2_256 Type = 0x56
	// Murmur3 - MURMUR3 hashing algorithm.
	Murmur3 Type = 0x22
	// UnknownType - Unknown hashing algorithm.
	UnknownType Type = 0
)

// HashNames - Multihash identifier names.
var HashNames = map[Type]string{
	Sha1:           "sha1",
	Sha2_256:       "sha2-256",
	Sha2_512:       "sha2-512",
	Sha3_224:       "sha3-224",
	Sha3_256:       "sha3-256",
	Sha3_384:       "sha3-384",
	Sha3_512:       "sha3-512",
	DoubleSha2_256: "dbl-sha2-256",
	Murmur3:        "murmur3",
	Keccak224:      "keccak-224",
	Keccak256:      "keccak-256",
	Keccak384:      "keccak-384",
	Keccak512:      "keccak-512",
	Shake128:       "shake-128",
	Shake256:       "shake-256",
}

// Types - Multihash identifier names.
var Types = map[string]Type{
	"sha1":         Sha1,
	"sha2-256":     Sha2_256,
	"sha2-512":     Sha2_512,
	"sha3-224":     Sha3_224,
	"sha3-256":     Sha3_256,
	"sha3-384":     Sha3_384,
	"sha3-512":     Sha3_512,
	"dbl-sha2-256": DoubleSha2_256,
	"murmur3":      Murmur3,
	"keccak-224":   Keccak224,
	"keccak-256":   Keccak256,
	"keccak-384":   Keccak384,
	"keccak-512":   Keccak512,
	"shake-128":    Shake128,
	"shake-256":    Shake256,
}

// NewType - Creates new hash name from string.
func NewType(name string) (Type, error) {
	if t, ok := Types[name]; ok {
		return t, nil
	}
	return UnknownType, fmt.Errorf("unknown hash identifier %q", name)
}

// Family - Returns algorithm family.
func (t Type) Family() Family {
	switch t {
	case Sha1:
		return FamilySha1
	case Sha2_256:
		return FamilySha2
	case Sha2_512:
		return FamilySha2
	case Sha3_224:
		return FamilySha3
	case Sha3_256:
		return FamilySha3
	case Sha3_384:
		return FamilySha3
	case Sha3_512:
		return FamilySha3
	case DoubleSha2_256:
		return FamilyDoubleSha2
	case Murmur3:
		return FamilyMurmur3
	case Keccak224:
		return FamilyKeccak
	case Keccak256:
		return FamilyKeccak
	case Keccak384:
		return FamilyKeccak
	case Keccak512:
		return FamilyKeccak
	case Shake128:
		return FamilyShake
	case Shake256:
		return FamilyShake
	default:
		return FamilyUnknown
	}
}

// Code - Returns algorithm multihash code.
func (t Type) Code() uint64 {
	return uint64(t)
}

// String - Returns algorithm name.
func (t Type) String() string {
	if name, ok := HashNames[t]; ok {
		return name
	}
	return "unknown"
}
