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

import (
	"encoding/hex"
	"math/big"

	"github.com/gogo/protobuf/proto"
	"github.com/ipfn/ipfn/go/opcode"
)

// BigInt - Creates new big int cell.
func BigInt(num *big.Int) *opcode.BinaryCell {
	return opcode.New(OpBigInt, num.Bytes())
}

// Uint64 - Creates new uint64 cell.
func Uint64(num uint64) *opcode.BinaryCell {
	return opcode.New(OpUint64, proto.EncodeVarint(num))
}

// Bytes - Creates new bytes cell.
func Bytes(bytes []byte) *opcode.BinaryCell {
	return opcode.New(OpBytes, bytes)
}

// ParseBigInt - Creates new big int cell from hex string.
func ParseBigInt(str string) (_ *opcode.BinaryCell, err error) {
	bytes, err := hex.DecodeString(str)
	if err != nil {
		return
	}
	return opcode.New(OpUint64, bytes), nil
}

// MustParseBigInt - Creates new big int cell from string.
func MustParseBigInt(str string) (c *opcode.BinaryCell) {
	c, err := ParseBigInt(str)
	if err != nil {
		panic(err)
	}
	return
}
