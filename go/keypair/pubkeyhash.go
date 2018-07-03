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

	"github.com/btcsuite/btcd/btcec"
	"github.com/ipfn/ipfn/go/cids"
	"github.com/ipfn/ipfn/go/opcode"
	cid "github.com/ipfs/go-cid"
	multihash "github.com/multiformats/go-multihash"
)

// CIDPrefix - Key CID prefix.
var CIDPrefix = cid.Prefix{
	Version:  1,
	Codec:    cids.PubkeyHash,
	MhType:   multihash.KECCAK_256,
	MhLength: 32,
}

// CID - Creates CID from public key.
func CID(pub *btcec.PublicKey) (c *opcode.CID) {
	pubBytes := PubkeyBytes(pub)
	c, _ = opcode.SumCID(CIDPrefix, pubBytes[1:])
	return
}

// PubkeyBytes - Gets public key bytes.
func PubkeyBytes(pub *btcec.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(btcec.S256(), pub.X, pub.Y)
}
