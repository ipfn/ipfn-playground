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

package pubkeyhash

import (
	"crypto/ecdsa"

	"github.com/btcsuite/btcd/btcec"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// PKHash - Creates bitcoin address style pubkey hash.
func PKHash(pub *btcec.PublicKey, netID byte) (*btcutil.AddressPubKeyHash, error) {
	params := chaincfg.Params{PubKeyHashAddrID: netID}
	pkHash := btcutil.Hash160(pub.SerializeCompressed())
	return btcutil.NewAddressPubKeyHash(pkHash, &params)
}

// AddressEthereum - Creates new CID for ECDSA Public Key and codec prefix.
func AddressEthereum(pub ecdsa.PublicKey) common.Address {
	return crypto.PubkeyToAddress(pub)
}
