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

package keypair

import (
	"errors"
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
)

// New - Creates a new master key.
func New() (_ *KeyPair, err error) {
	seed, err := NewSeed()
	if err != nil {
		return
	}
	return NewMaster(seed)
}

// NewMaster - Creates a new master key from seed.
func NewMaster(seed []byte) (*KeyPair, error) {
	return NewCustomMaster(seed, keyParams)
}

// NewKeyFromString - Creates a new master key from string.
func NewKeyFromString(v string) (*KeyPair, error) {
	key, err := hdkeychain.NewKeyFromString(v)
	if err != nil {
		return nil, err
	}
	return &KeyPair{ExtendedKey: key}, nil
}

// NewCustomMaster - Creates a new master key from seed.
func NewCustomMaster(seed []byte, net *chaincfg.Params) (*KeyPair, error) {
	if net == nil {
		return nil, errors.New("cannot generate with empty chain params")
	}
	key, err := hdkeychain.NewMaster(seed, net)
	if err != nil {
		return nil, err
	}
	return &KeyPair{ExtendedKey: key}, nil
}

// HashPath - Creates custom derivation path from bytes of crc32 hash.
func HashPath(path string) string {
	// crc32 checksum of path
	h := crc32.NewIEEE()
	h.Write([]byte(path))
	s := h.Sum(nil)
	// create derivation path string
	r := []string{"m", "44'", "43153'", "120'", "0"}
	// iterate over crc32 bytes
	for n, v := range s {
		if n < 2 {
			r = append(r, fmt.Sprintf("%d'", v))
		} else {
			r = append(r, strconv.Itoa(int(v)))
		}
	}
	return strings.Join(r, "/")
}
