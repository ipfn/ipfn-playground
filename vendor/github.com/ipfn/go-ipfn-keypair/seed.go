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

	bip39 "github.com/ipfn/go-bip39"
	crypto "github.com/ipfn/go-ipfn-crypto"
)

// EncryptedSeed - Encrypted master key or seed.
type EncryptedSeed struct {
	// Name - Name of a seed.
	Name string `json:"name"`

	// SeedType - Type of a seed.
	SeedType `json:"seed_type"`

	// EncryptedKey - Encrypted master key or seed.
	EncryptedKey crypto.SealedBox `json:"encrypted_key"`
}

// SeedType - Type of a master key.
type SeedType int

const (
	// Mnemonic - Represents mnemonic seed.
	Mnemonic SeedType = 0
	// Base58Key - Represents base58 encoded seed.
	Base58Key SeedType = 1
)

// GetSeedType - Gets type of a seed.
func GetSeedType(seed []byte) SeedType {
	if bip39.IsMnemonicValid(string(seed)) {
		return Mnemonic
	}
	return Base58Key
}

// NewSeed - Creates a new key seed.
// Returns human readable mnemonic.
func NewSeed() ([]byte, error) {
	// generate entropy
	entropy, err := crypto.NewEntropy(32)
	if err != nil {
		return nil, fmt.Errorf("failed to generate entropy: %v", err)
	}
	// convert entropy to mnemonic
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return nil, fmt.Errorf("failed to create mnemonic: %v", err)
	}
	return []byte(mnemonic), nil
}

// EncryptSeed - Encrypts and creates new encrypted seed structure.
func EncryptSeed(name string, seed, password []byte) (key EncryptedSeed, err error) {
	if name == "" {
		err = errors.New("name cannot be empty")
		return
	}
	if len(seed) == 0 {
		err = errors.New("seed cannot be empty")
		return
	}
	key.Name = name
	key.SeedType = GetSeedType(seed)
	key.EncryptedKey, err = crypto.SealBox(seed, password)
	return
}

// Decrypt - Decrypts sealed box with password.
func (key EncryptedSeed) Decrypt(password []byte) ([]byte, error) {
	return key.EncryptedKey.Decrypt(password)
}

// ToKeyPair - Decrypts sealed box with password and creates a new key chain.
func (key EncryptedSeed) ToKeyPair(password []byte) (_ *KeyPair, err error) {
	seed, err := key.Decrypt(password)
	if err != nil {
		return
	}
	if key.SeedType == Mnemonic {
		return NewMaster(bip39.NewSeed(seed, password))
	}
	return NewKeyFromString(string(seed))
}
