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
	"bytes"
	"fmt"
	"math"
)

// Family - Hash digest family multihash ID.
// Identifiers used for families are 256bit IDs.
type Family uint64

const (
	// FamilySha1 - SHA1 hashing algorithm.
	FamilySha1 Family = 0x11
	// FamilySha2 - SHA2 hashing algorithm.
	FamilySha2 Family = 0x12
	// FamilySha3 - SHA3 hashing algorithm.
	FamilySha3 Family = 0x16
	// FamilyKeccak - Keccak hashing algorithm.
	FamilyKeccak Family = 0x1B
	// FamilyShake - Shake hashing algorithm.
	FamilyShake Family = 0x19
	// FamilyBlake2b - Blake2B hashing algorithm.
	FamilyBlake2b Family = 0xb201
	// FamilyBlake2s - Blake2S hashing algorithm.
	FamilyBlake2s Family = 0xb241
	// FamilyDoubleSha2 - Double SHA2 hashing algorithm.
	FamilyDoubleSha2 Family = 0x56
	// FamilyMurmur3 - MURMUR3 hashing algorithm.
	FamilyMurmur3 Family = 0x22
	// FamilyUnknown - Unknown hashing algorithm family.
	FamilyUnknown Family = math.MaxUint64
)

func (family Family) String() string {
	switch family {
	case FamilySha1:
		return "sha1"
	case FamilySha2:
		return "sha2"
	case FamilySha3:
		return "sha3"
	case FamilyKeccak:
		return "keccak"
	case FamilyShake:
		return "shake"
	case FamilyBlake2b:
		return "blake2b"
	case FamilyBlake2s:
		return "blake2s"
	case FamilyDoubleSha2:
		return "doublesha2"
	case FamilyMurmur3:
		return "murmur3"
	default:
		return "unknown"
	}
}

// MarshalText - Unmarshals JSON string.
func (family Family) MarshalText() ([]byte, error) {
	return []byte(family.String()), nil
}

// MarshalJSON - Unmarshals JSON string.
func (family Family) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", family)), nil
}

// UnmarshalJSON - Unmarshals JSON string.
func (family *Family) UnmarshalJSON(body []byte) (_ error) {
	return family.UnmarshalText(bytes.Trim(body, `"`))
}

// UnmarshalText - Unmarshals text.
func (family *Family) UnmarshalText(body []byte) (_ error) {
	switch string(bytes.ToLower(body)) {
	case "sha1":
		*family = FamilySha1
		return
	case "sha2":
		*family = FamilySha2
		return
	case "sha3":
		*family = FamilySha3
		return
	case "keccak":
		*family = FamilyKeccak
		return
	case "shake":
		*family = FamilyShake
		return
	case "blake2b":
		*family = FamilyBlake2b
		return
	case "blake2s":
		*family = FamilyBlake2s
		return
	case "doublesha2":
		*family = FamilyDoubleSha2
		return
	case "murmur3":
		*family = FamilyMurmur3
		return
	default:
		return fmt.Errorf("Unknown hash family %q", body)
	}
}
