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

package chainops

import (
	"github.com/btcsuite/btcd/btcec"
	base32check "github.com/ipfn/go-base32check"
	"github.com/ipfn/ipfn/go/dev/address"
	"github.com/ipfn/ipfn/go/keypair"
	"github.com/ipfn/ipfn/go/opcode"
	"github.com/ipfn/ipfn/go/opcode/synaptic"
	multihash "github.com/multiformats/go-multihash"
)

// ID - Creates new uint64 cell.
func ID(num opcode.ID) *opcode.BinaryCell {
	return opcode.New(OpID, num.Bytes())
}

// IDFromString - Creates new uint64 cell.
func IDFromString(body string) *opcode.BinaryCell {
	return ID(opcode.NewIDFromString(body))
}

// ParseID - Creates new uint64 cell from string by parsing it.
func ParseID(src string) (_ *opcode.BinaryCell, err error) {
	id, err := base32check.CheckDecodeString(src)
	if err != nil {
		return
	}
	return opcode.New(synaptic.OpUint64, id), nil
}

// MustParseID - Creates new uint64 cell from string.
func MustParseID(str string) (c *opcode.BinaryCell) {
	c, err := ParseID(str)
	if err != nil {
		panic(err)
	}
	return
}

// ParseAddress - Parses short address from string.
func ParseAddress(src string) (_ *opcode.BinaryCell, err error) {
	addr, err := address.ParseAddress(src)
	if err != nil {
		return
	}
	bytes, err := addr.Marshal()
	if err != nil {
		return
	}
	return opcode.New(OpAddress, bytes), nil
}

// MustParseAddress - Parses short address or panics.
func MustParseAddress(src string) (c *opcode.BinaryCell) {
	c, err := ParseAddress(src)
	if err != nil {
		panic(err)
	}
	return
}

// CID - Creates CID binary cell.
func CID(c *opcode.CID) (_ *opcode.BinaryCell) {
	if c == nil {
		return opcode.Op(OpCID)
	}
	return opcode.New(OpCID, c.Bytes())
}

// Multihash - Creates multihash binary cell.
func Multihash(mh multihash.Multihash) *opcode.BinaryCell {
	return opcode.New(OpMultihash, []byte(mh))
}

// Sign - Signs binary cell and creates signed operation.
func Sign(op *opcode.BinaryCell, pk *btcec.PrivateKey) (_ *opcode.BinaryCell, err error) {
	body, err := op.Marshal()
	if err != nil {
		return
	}
	sig, err := btcec.SignCompact(btcec.S256(), pk, body, false)
	if err != nil {
		return
	}
	return opcode.Op(OpSigned, op, Signature(sig)), nil
}

// Signature - Creates signature binary cell.
func Signature(sig []byte) (_ *opcode.BinaryCell) {
	return opcode.New(OpSignature, sig)
}

// Signed - Creates signed binary cell.
func Signed(op *opcode.BinaryCell, signatures []*opcode.BinaryCell) *opcode.BinaryCell {
	ops := append(opcode.Ops(op), signatures...)
	return opcode.Op(OpSigned, ops...)
}

// Pubkey - Creates public key cell.
func Pubkey(pubkey *btcec.PublicKey) (c *opcode.BinaryCell) {
	return PubkeyBytes(pubkey.SerializeCompressed())
}

// PubkeyBytes - Creates public key cell.
func PubkeyBytes(pubkey []byte) (c *opcode.BinaryCell) {
	return opcode.New(OpPubkey, pubkey)
}

// Genesis - Creates genesis operation.
func Genesis() (c *opcode.BinaryCell) {
	return opcode.Op(OpGenesis)
}

// AssignPower - Creates assign power operation.
func AssignPower(nonce opcode.ID, quantity uint64, pubkey []byte) (c *opcode.BinaryCell) {
	return opcode.Op(OpAssignPower,
		// opcode.New(OpNonce, nonce.Bytes()),
		synaptic.Uint64(quantity),
		PubkeyBytes(pubkey))
}

// AssignPowerAddr - Creates assign power operation.
func AssignPowerAddr(nonce opcode.ID, quantity uint64, addr *opcode.CID) (c *opcode.BinaryCell) {
	return opcode.Op(OpAssignPower,
		// opcode.New(OpNonce, nonce.Bytes()),
		synaptic.Uint64(quantity),
		CID(addr))
}

// DelegatePower - Creates delegate power operation.
func DelegatePower(nonce opcode.ID, quantity uint64, pubkeys ...[]byte) (c *opcode.BinaryCell) {
	c = opcode.Op(OpAssignPower,
		// opcode.New(OpNonce, nonce.Bytes()),
		synaptic.Uint64(quantity))
	if len(pubkeys) > 0 {
		for _, pubkey := range pubkeys {
			c.Add(PubkeyBytes(pubkey))
		}
	}
	return
}

// PubkeyToAddr - Creates public key hash cell from public key.
func PubkeyToAddr(bytes []byte) *opcode.BinaryCell {
	c, err := opcode.SumCID(keypair.CIDPrefix, bytes)
	if err != nil {
		panic(err)
	}
	return opcode.New(OpPubkeyAddr, c.Bytes())
}
