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

// Key represents a cryptographic key
type Key interface {
	// SKI returns the subject key identifier of this key.
	SKI() []byte
	// Bytes converts this key to its byte representation.
	Bytes() ([]byte, error)
	// Symmetric returns true if this key is a symmetric key.
	Symmetric() bool
	// Private returns true if this key is a private key.
	Private() bool
	// PublicKey returns the corresponding public key.
	PublicKey() (Key, error)
}

// KeyGenOpts contains options for key-generation with a CSP.
type KeyGenOpts interface {
	// Algorithm returns the key generation algorithm identifier (to be used).
	Algorithm() string
	// Ephemeral returns true if the key to generate has to be ephemeral.
	Ephemeral() bool
}

// KeyDerivOpts contains options for key-derivation with a CSP.
type KeyDerivOpts interface {
	// Algorithm returns the key derivation algorithm identifier (to be used).
	Algorithm() string
	// Ephemeral returns true if the key to derived has to be ephemeral.
	Ephemeral() bool
}

// KeyImportOpts contains options for importing the raw material of a key with a CSP.
type KeyImportOpts interface {
	// Algorithm returns the key importation algorithm identifier (to be used).
	Algorithm() string
	// Ephemeral returns true if the key generated has to be ephemeral.
	Ephemeral() bool
}
