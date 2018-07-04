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
	// OpOffset - Offset of chain operation code.
	OpOffset opcode.ID = iota + 0x3c

	// OpMultihash - Multihash native operation code.
	OpMultihash = OpOffset + 1

	// OpID - ID operation.
	OpID = OpOffset + 2

	// OpCID - Contentn ID native operation code.
	OpCID = OpOffset + 3

	// OpHeader - Header operation.
	OpHeader = OpOffset + 4

	// OpGenesis - Genesis operation.
	OpGenesis = OpOffset + 5

	// OpClaim - Address claim operation.
	OpClaim = OpOffset + 6

	// OpAssignPower - Allocation of power operation.
	OpAssignPower = OpOffset + 7

	// OpDelegatePower - Investment of power operation.
	OpDelegatePower = OpOffset + 8

	// OpPubkey - Public key operation.
	OpPubkey = OpOffset + 9

	// OpPubkeyAddr - Public key hash operation.
	OpPubkeyAddr = OpOffset + 10

	// OpSignature - Signature operation.
	OpSignature = OpOffset + 11

	// OpSigned - Signed operation.
	OpSigned = OpOffset + 12

	// OpAddress - Address native operation code.
	OpAddress = OpOffset + 13

	// OpTransfer - Transfer of an asset.
	OpTransfer = OpOffset + 14

	// OpNonce - Nonce op (noop).
	OpNonce = OpOffset + 15

	// OpRoot - Offset of chain operation code.
	OpRoot opcode.ID = 0
)

func init() {
	opcode.Register(OpID, "id")
	opcode.Register(OpRoot, "root")
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
	opcode.Register(OpNonce, "nonce")
}
