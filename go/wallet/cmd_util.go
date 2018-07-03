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

package wallet

import (
	"errors"
	"fmt"

	"github.com/ipfn/ipfn/go/keypair"
	prompt "github.com/segmentio/go-prompt"
)

// PromptDeriveKey - Derives key from wallet and path prompting for password in console.
// If `hash` is set to true `path` is hash-pathed.
func PromptDeriveKey(path *KeyPath) (_ *keypair.KeyPair, err error) {
	w := NewDefault()
	has, err := w.KeyExists(path.SeedName)
	if err != nil {
		return nil, err
	}
	if !has {
		return nil, fmt.Errorf("%q wallet does not exist", path.SeedName)
	}
	password, err := PromptWalletPassword(path.SeedName)
	if err != nil {
		return
	}
	return w.DeriveKey(path, []byte(password))
}

// PromptWalletPassword - Prompts for wallet password after checking for its existence.
func PromptWalletPassword(name string) (_ []byte, err error) {
	// Check if key exists in default wallet
	has, err := NewDefault().KeyExists(name)
	if err != nil {
		return
	}
	if !has {
		return nil, fmt.Errorf("%q wallet does not exist", name)
	}

	// Get wallet password
	password := prompt.PasswordMasked(fmt.Sprintf("Wallet %q password", name))
	if password == "" {
		return nil, errors.New("failed to get decryption password")
	}
	return []byte(password), nil
}
