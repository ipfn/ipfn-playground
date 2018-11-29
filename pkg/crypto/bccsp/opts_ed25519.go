// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2016-2018 IBM Corp. All Rights Reserved.
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

// ED25519KeyGenOpts contains options for ed25519 key generation.
type ED25519KeyGenOpts struct {
	Temporary bool
}

// Algorithm returns the key generation algorithm identifier (to be used).
func (opts *ED25519KeyGenOpts) Algorithm() string {
	return ED25519
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *ED25519KeyGenOpts) Ephemeral() bool {
	return opts.Temporary
}

// ED25519ReRandKeyOpts contains options for ed25519 key re-randomization.
type ED25519ReRandKeyOpts struct {
	Temporary bool
	Expansion []byte
}

// Algorithm returns the key derivation algorithm identifier (to be used).
func (opts *ED25519ReRandKeyOpts) Algorithm() string {
	return ED25519ReRand
}

// Ephemeral returns true if the key to generate has to be ephemeral,
// false otherwise.
func (opts *ED25519ReRandKeyOpts) Ephemeral() bool {
	return opts.Temporary
}

// ExpansionValue returns the re-randomization factor
func (opts *ED25519ReRandKeyOpts) ExpansionValue() []byte {
	return opts.Expansion
}
