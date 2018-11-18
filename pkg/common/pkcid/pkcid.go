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

// Package pkcid implements helpers for generating public-key CID.
package pkcid

import (
	"crypto/ecdsa"
	"crypto/elliptic"

	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	mh "gx/ipfs/QmerPMzPk1mJVowm8KgmoknWa4yCYvvugMPsgWmDNUvDLW/go-multihash"

	"github.com/ipfn/ipfn/pkg/cells"
	"github.com/ipfn/ipfn/pkg/common/codecs"
)

// PrefixV1 - CID prefix of IPFN public key.
//
// It can be any key like ECDSA or RSA.
// Security directs knowledge of its kind.
var PrefixV1 = cid.Prefix{
	Version:  1,
	Codec:    codecs.PubkeyHashV1,
	MhType:   mh.SHA2_256,
	MhLength: 32,
}

// CID - Creates CID from public key.
//
// Resulting CID has codec ID of `codecs.PubkeyHashV1`.
func CID(pub *ecdsa.PublicKey) (c *cells.CID) {
	return BytesToCID(PubkeyBytes(pub))
}

// BytesToCID - Creates CID from public key.
//
// Follows ethereum pattern and strips one byte.
func BytesToCID(pubBytes []byte) (c *cells.CID) {
	c, err := cells.SumCID(PrefixV1, pubBytes[1:])
	if err != nil {
		panic(err)
	}
	return
}

// PubkeyBytes - Marshals public key to bytes.
func PubkeyBytes(pub *ecdsa.PublicKey) []byte {
	if pub == nil || pub.X == nil || pub.Y == nil {
		return nil
	}
	return elliptic.Marshal(pub.Curve, pub.X, pub.Y)
}
