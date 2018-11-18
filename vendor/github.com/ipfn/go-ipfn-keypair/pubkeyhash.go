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

// Cryptographic public key hash utilities in Go supporting various crypto-currencies
// standards such a Bitcoin, Ethereum and custom standards.
//
// * [EIP-0055](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-55.md)
// * [BIP-0013](https://github.com/bitcoin/bips/blob/master/bip-0013.mediawiki)

package keypair

import (
	"crypto/elliptic"

	mh "gx/ipfs/QmPnFwZ2JXKnXgMw8CdBPxn7FWh6LLdjUjxV1fKHuJnkr8/go-multihash"
	cid "gx/ipfs/QmapdYm1b22Frv3k17fqrBYTFRxwiaVJkB299Mfn33edeB/go-cid"

	"github.com/btcsuite/btcd/btcec"
	cells "github.com/ipfn/go-ipfn-cells"
)

// CIDPrefix - Key CID prefix.
var CIDPrefix = cid.Prefix{
	Version:  1,
	Codec:    0x60ac,
	MhType:   mh.KECCAK_256,
	MhLength: 32,
}

// CID - Creates CID from public key.
func CID(pub *btcec.PublicKey) (c *cells.CID) {
	pubBytes := PubkeyBytes(pub)
	return BytesToCID(pubBytes)
}

// BytesToCID - Creates CID from public key.
func BytesToCID(pubBytes []byte) (c *cells.CID) {
	c, err := cells.SumCID(CIDPrefix, pubBytes[1:])
	if err != nil {
		panic(err)
	}
	return
}

// PubkeyBytes - Gets public key bytes.
func PubkeyBytes(pub *btcec.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(btcec.S256(), pub.X, pub.Y)
}
