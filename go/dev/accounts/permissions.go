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

package accounts

import "github.com/ipfn/ipfn/go/dev/address"

// Identity - Identity of account client.
type Identity struct {
	// Threshold - Key threshold.
	Threshold uint `json:"threshold,omitempty"`

	// Keys - Identity keys.
	Keys []*Key `json:"keys,omitempty"`
}

// Key - Identity key identifier.
type Key struct {
	// Weight - Identity key weight.
	Weight uint `json:"weight,omitempty"`

	// Address - Identity key address.
	Address *address.Address `json:"address,omitempty"`

	// PublicKey - Optional public key bytes.
	PublicKey []byte `json:"public_key,omitempty"`
}

// Permissions - Account permissions.
type Permissions struct {
	// Action - Action name.
	Action string `json:"action,omitempty"`

	// Identity - Identity allowed to perform the action.
	*Identity
}
