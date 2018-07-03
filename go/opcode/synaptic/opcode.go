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

package synaptic

import "github.com/ipfn/ipfn/go/opcode"

// OpOffset - Synaptic opcode offset.
const OpOffset opcode.ID = 0x1c

const (
	// OpBytes - Synaptic byte array opcode.
	OpBytes = OpOffset + 1
	// OpInt64 - Synaptic uint64 opcode.
	OpInt64 = OpOffset + 2
	// OpUint64 - Synaptic uint64 opcode.
	OpUint64 = OpOffset + 3
	// OpBigInt - Synaptic big int opcode.
	OpBigInt = OpOffset + 4
	// OpString - Synaptic string opcode.
	OpString = OpBytes
)

func init() {
	opcode.Register(OpBytes, "bytes")
	opcode.Register(OpInt64, "int64")
	opcode.Register(OpUint64, "uint64")
	opcode.Register(OpBigInt, "bigint")
	opcode.Register(OpString, "string")
}
