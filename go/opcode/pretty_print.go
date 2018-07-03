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

package opcode

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"
	cid "github.com/ipfs/go-cid"
)

func prettyPrint(cell *BinaryCell) (_ []byte) {
	buff := bytes.NewBuffer(nil)
	buff.WriteByte('"')
	writeStringScript(cell, buff)
	buff.WriteByte('"')
	return buff.Bytes()
}

func writeStringScript(cell *BinaryCell, buff *bytes.Buffer) {
	buff.WriteString(strings.ToUpper(fmt.Sprintf("OP_%s", cell.OpCode)))
	if len(cell.Memory) > 0 {
		buff.WriteByte(' ')
		if cell.OpCode == 31 || cell.OpCode == 62 { // uint64 or id
			i, _ := proto.DecodeVarint(cell.Memory)
			buff.WriteString(fmt.Sprintf("%d", i))
		} else if cell.OpCode == 63 || cell.OpCode == 70 { // cid or pubkey addr
			c, _ := cid.Cast(cell.Memory)
			buff.WriteString(c.String())
		} else {
			buff.WriteString(fmt.Sprintf("0x%x", cell.Memory))
		}
	}
	if len(cell.Children) > 0 {
		buff.WriteString(" [ ")
		for _, child := range cell.Children {
			writeStringScript(child, buff)
			buff.WriteByte(' ')
		}
		buff.WriteString("]")
	}
}
