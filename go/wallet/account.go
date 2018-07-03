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

import "github.com/ipfn/ipfn/go/dev/address"

// Account - Contains a public account address not always private keys.
// It only describes local account configuration storage such as
// paths to our own account owner or active keys.
type Account struct {
	// Name - Account name.
	Name string `json:"name,omitempty"`

	// OwnerKeys - Owner keys.
	OwnerKeys []*AccountKeyPath `json:"master_keys,omitempty"`

	// ClientKeys - Client key with user permissions.
	ClientKeys []*KeyPath `json:"client_keys,omitempty"`
}

// AccountKeyPath - Account key structure.
type AccountKeyPath struct {
	// Weight - Key weight.
	Weight uint `json:"weight,omitempty"`
	// KeyPath - Key path.
	KeyPath *KeyPath `json:"key_path,omitempty"`
	// Address - Key address.
	Address *address.Address `json:"address,omitempty"`
}

// String - Returns account address.
func (acc *Account) String() string {
	return acc.Name
}
