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
	base32check "github.com/ipfn/go-base32check"
)

// PKHash - Creates bitcoin address style pubkey hash.
func PKHash(pub *btcec.PublicKey, netID byte) (*btcutil.AddressPubKeyHash, error) {
	params := chaincfg.Params{PubKeyHashAddrID: netID}
	pkHash := btcutil.Hash160(pub.SerializeCompressed())
	return btcutil.NewAddressPubKeyHash(pkHash, &params)
}

// Base32PKHash - Creates base32 encoded pubkey hash.
func Base32PKHash(pub *btcec.PublicKey, netID byte) (_ []byte, err error) {
	addr, err := PKHash(pub, netID)
	if err != nil {
		return
	}
	return base32check.CheckEncode(addr.ScriptAddress(), 0xe8), nil
}

// Base32PKHashString - Creates base32 encoded pubkey hash.
func Base32PKHashString(pub *btcec.PublicKey, netID byte) (_ string, err error) {
	addr, err := PKHash(pub, netID)
	if err != nil {
		return
	}
	return base32check.CheckEncodeToString(addr.ScriptAddress(), 0xe8), nil
}

// AddressEthereum - Creates new CID for ECDSA Public Key and codec prefix.
func AddressEthereum(pub ecdsa.PublicKey) common.Address {
	return crypto.PubkeyToAddress(pub)
}
