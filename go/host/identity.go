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

package host

import (
	"github.com/btcsuite/btcd/btcec"
	"github.com/ipfn/ipfn/go/keypair"
	libp2p "github.com/libp2p/go-libp2p"
	crypto "github.com/libp2p/go-libp2p-crypto"
)

// Identity - Creates new identity from private key.
func Identity(privkey *btcec.PrivateKey) libp2p.Option {
	return libp2p.Identity((*crypto.Secp256k1PrivateKey)(privkey))
}

// KeyPair - Creates new identity from keypair.
// Panics if given keypair is not private key.
func KeyPair(keys *keypair.KeyPair) libp2p.Option {
	privkey, err := keys.ECPrivKey()
	if err != nil {
		panic(err)
	}
	return Identity(privkey)
}
