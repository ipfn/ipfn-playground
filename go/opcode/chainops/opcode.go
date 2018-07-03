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
	"github.com/ipfn/ipfn/go/opcode"
)

const (
	// ChainOpOffset - Offset of chain operation code.
	ChainOpOffset opcode.ID = iota + 0x3c

	// OpMultihash - Multihash native operation code.
	OpMultihash

	// OpID - ID operation.
	OpID

	// OpCID - Contentn ID native operation code.
	OpCID

	// OpHeader - Header operation.
	OpHeader

	// OpGenesis - Genesis operation.
	OpGenesis

	// OpClaim - Address claim operation.
	OpClaim

	// OpAssignPower - Allocation of power operation.
	OpAssignPower

	// OpDelegatePower - Investment of power operation.
	OpDelegatePower

	// OpPubkey - Public key operation.
	OpPubkey

	// OpPubkeyAddr - Public key hash operation.
	OpPubkeyAddr

	// OpSignature - Signature operation.
	OpSignature

	// OpSigned - Signed operation.
	OpSigned

	// OpAddress - Address native operation code.
	OpAddress

	// OpTransfer - Transfer of an asset.
	OpTransfer
)

func init() {
	opcode.Register(OpID, "id")
	opcode.Register(OpHeader, "header")
	opcode.Register(OpGenesis, "genesis")
	opcode.Register(OpAssignPower, "assign_power")
	opcode.Register(OpDelegatePower, "delegate_power")
	opcode.Register(OpSignature, "signature")
	opcode.Register(OpPubkey, "pubkey")
	opcode.Register(OpPubkeyAddr, "pubkey_addr")
	opcode.Register(OpSigned, "signed")
	opcode.Register(OpAddress, "address")
	opcode.Register(OpMultihash, "multihash")
	opcode.Register(OpCID, "cid")
	opcode.Register(OpClaim, "claim")
	opcode.Register(OpTransfer, "transfer")
}
