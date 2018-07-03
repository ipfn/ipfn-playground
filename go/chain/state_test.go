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

package chain

import (
	. "testing"

	"github.com/ipfn/go-base32check"

	"github.com/ipfn/ipfn/go/opcode"
	"github.com/ipfn/ipfn/go/opcode/chainops"
	"github.com/ipfn/ipfn/go/opcode/synaptic"
	"github.com/stretchr/testify/assert"
)

var (
	allocEnc   = "0000zsc00fysmullctmwh69d6mz0qpht0y0p7070ss7s0"
	genesisEnc = "0000ysb00pps00jfphellshkad52l4ky0xqwk0b0qupupppa00"
)

func TestBinaryCell(t *T) {

	genesisOp := &opcode.BinaryCell{OpCode: chainops.OpGenesis}

	// signedGenesis := signedOp(genesisOp, key)

	allocOp := &opcode.BinaryCell{
		OpCode: chainops.OpAssignPower,
		Children: []*opcode.BinaryCell{
			chainops.MustParseAddress("b7dlu9ahtazhar30psm4sqlc"),
			synaptic.Uint64(1e6),
		},
	}

	var head string
	state, err := NewState(0, nil, opcode.Ops(genesisOp, allocOp))
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(0), state.Index())
	// assert.Equal(t, "zFSec2XVAw1qbBFm7rFFV81U8UwGqBLuV7SGze8vPyQbziy7zbku", state.Head().String())
	// assert.Equal(t, "zFSec2XV6dbiDRvMX7aKz3eEkyxuWPjdiwFvVbe5MXomnT5ZdwT1", state.Header.Exec.String())
	head = state.Head().String()

	body, err := opcode.RootOp(state.ExecOps).Marshal()
	assert.Equal(t, genesisEnc, base32check.EncodeToString(body))

	state, err = state.Next(opcode.Ops(allocOp))
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(1), state.Index())
	assert.Equal(t, head, state.Prev().String())
	// assert.Equal(t, "zFSec2XVBdAsns1xHLkvMhcJnvtK4tchNG8DBhiXcZd4kUqQbGVa", state.Head().String())
	// assert.Equal(t, "zFSec2XVGN5saHLfhnwm3s5TRXTPNGGcbggyfXSYLP97nxn6HStJ", state.Header.Exec.String())
	head = state.Head().String()

	body, err = opcode.RootOp(state.ExecOps).Marshal()
	assert.Equal(t, allocEnc, base32check.EncodeToString(body))

	state, err = state.Next(opcode.Ops(allocOp))
	assert.Equal(t, nil, err)
	assert.Equal(t, uint64(2), state.Index())
	assert.Equal(t, head, state.Prev().String())
	// assert.Equal(t, "zFSec2XV3e6uoSd8sTcBM717xMweVCBQFBoMQc4Qy3JGtYMtu34E", state.Head().String())
	// assert.Equal(t, "zFSec2XVGN5saHLfhnwm3s5TRXTPNGGcbggyfXSYLP97nxn6HStJ", state.Header.Exec.String())

	body, err = opcode.RootOp(state.ExecOps).Marshal()
	assert.Equal(t, allocEnc, base32check.EncodeToString(body))

}
