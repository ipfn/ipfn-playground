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

package x25519

import (
	"golang.org/x/crypto/curve25519"
)

// Public - Computes public key from seed.
func Public(seed *[32]byte) (pubkey [32]byte) {
	curve25519.ScalarBaseMult(&pubkey, seed)
	return
}

// Shared - Computes shared secret key.
func Shared(secret, public *[32]byte) (shared [32]byte) {
	curve25519.ScalarMult(&shared, secret, public)
	return
}
