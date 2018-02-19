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
	"crypto/sha512"

	"golang.org/x/crypto/pbkdf2"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	bip39 "github.com/ipfn/go-bip39"
)

var (
	// DefaultEntropyLength - Default entropy bit length.
	DefaultEntropyLength uint8 = hdkeychain.RecommendedSeedLen
)

// NewEntropy - Generates new entropy with default bit size if size is zero.
func NewEntropy(size uint8) ([]byte, error) {
	if size == 0 {
		size = DefaultEntropyLength
	}
	return hdkeychain.GenerateSeed(size)
}

// NewMnemonic - Creates a mnemonic phrase from bytes.
func NewMnemonic(entropy []byte) (string, error) {
	return bip39.NewMnemonic(entropy)
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
func NewMaster(seed []byte, net *chaincfg.Params) (*hdkeychain.ExtendedKey, error) {
	if net == nil {
		net = &chaincfg.MainNetParams
	}
	return hdkeychain.NewMaster(seed, net)
}
