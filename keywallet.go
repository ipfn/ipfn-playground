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

package keywallet

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	bip39 "github.com/ipfn/go-bip39"
)

var (
	// DefaultEntropyLength - Default entropy bit length.
	DefaultEntropyLength uint8 = hdkeychain.RecommendedSeedLen
)

// NewEntropy - Generates new entropy with default bit size.
func NewEntropy() ([]byte, error) {
	return hdkeychain.GenerateSeed(DefaultEntropyLength)
}

// NewCustomEntropy - Generates new entropy with custom bit size.
func NewCustomEntropy(size uint8) ([]byte, error) {
	if size == 0 {
		size = DefaultEntropyLength
	}
	return hdkeychain.GenerateSeed(size)
}

// NewMnemonic - Converts a mnemonic phrase from bytes.
func NewMnemonic(entropy []byte) (string, error) {
	return bip39.NewMnemonic(entropy)
}

// IsMnemonicValid - Checks if a mnemonic phrase is valid.
func IsMnemonicValid(mnemonic string) bool {
	return bip39.IsMnemonicValid(mnemonic)
}

// NewSeed - Creates a new key generation seed from mnemonic and salt.
func NewSeed(mnemonic, salt []byte) []byte {
	return bip39.NewSeed(mnemonic, salt)
}

// NewCustomSeed - Creates a new key generation seed from mnemonic and salt.
func NewCustomSeed(mnemonic, salt []byte, iter, size int) []byte {
	return bip39.NewCustomSeed(mnemonic, salt, iter, size)
}

// NewMaster - Creates a new master key from seed.
func NewMaster(seed []byte) (*ExtendedKey, error) {
	return NewCustomMaster(seed, &chaincfg.MainNetParams)
}

// NewCustomMaster - Creates a new master key from seed.
func NewCustomMaster(seed []byte, net *chaincfg.Params) (*ExtendedKey, error) {
	if net == nil {
		return nil, errors.New("cannot generate with empty chain params")
	}
	key, err := hdkeychain.NewMaster(seed, net)
	if err != nil {
		return nil, err
	}
	return &ExtendedKey{ExtendedKey: key}, nil
}

// DerivePath - Derives string BIP32 path like `m/44'/0'/1'/0/0`.
func DerivePath(key *ExtendedKey, path string) (res *ExtendedKey, err error) {
	elems := strings.Split(path, "/")
	if len(elems) > 0 && elems[0] == "" {
		elems = elems[1:]
	}
	if len(elems) > 0 && elems[0] == "m" {
		elems = elems[1:]
	} else {
		return nil, fmt.Errorf("invalid derivation path %q", path)
	}
	if len(elems) == 0 {
		return nil, fmt.Errorf("empty derivation path %q", path)
	}
	res = key
	for _, value := range elems {
		v, err := strconv.Atoi(strings.TrimRight(value, "'"))
		if err != nil {
			return nil, fmt.Errorf("wrong derivation path element %q in %q", value, path)
		}
		if strings.HasSuffix(value, "'") {
			res, err = res.Derive(uint32(v))
		} else {
			res, err = res.Child(uint32(v))
		}
		if err != nil {
			return nil, err
		}
	}
	return
}

// ExtendedKey - Hierarchical deterministic wallet key derivation.
type ExtendedKey struct {
	*hdkeychain.ExtendedKey
}

// Derive - Derives extended key child by adding 0x80000000 (2^31) to path.
func (key *ExtendedKey) Derive(path uint32) (*ExtendedKey, error) {
	return key.Child(hdkeychain.HardenedKeyStart + path)
}

// Child - Derives extended key child.
func (key *ExtendedKey) Child(path uint32) (*ExtendedKey, error) {
	k, err := key.ExtendedKey.Child(path)
	if err != nil {
		return nil, err
	}
	return &ExtendedKey{ExtendedKey: k}, nil
}
