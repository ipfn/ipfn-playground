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
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/cespare/xxhash"
	"github.com/gogo/protobuf/proto"

	. "testing"

	"github.com/stretchr/testify/assert"
)

var (
	cidMem, _  = hex.DecodeString("01acc1011b208b332db5843f3c465414c524e409029a94a3d4644054f4a54d267833e818cc3c")
	cidCell    = &BinaryCell{OpCode: 0xc1d, Memory: cidMem}
	simpleCell = &BinaryCell{
		Memory: []byte("test"),
		Children: []*BinaryCell{
			{OpCode: 0x01, Children: []*BinaryCell{
				{OpCode: 0x10, Memory: []byte("0x0023")},
				{OpCode: 0x10, Memory: []byte("0x0325")},
			}},
		},
	}
)

func TestBinaryCell(t *T) {
	b, err := simpleCell.Marshal()
	assert.Equal(t, nil, err)
	unc := new(BinaryCell)
	err = unc.Unmarshal(b)
	assert.Equal(t, nil, err)
	b2, _ := unc.Marshal()
	assert.Equal(t, b, b2)
	assert.Equal(t, simpleCell.OpCode, unc.OpCode)

	id, err := unc.Checksum()
	assert.Equal(t, nil, err)
	assert.Equal(t, ID(0x1663e9dbd3c404b2), id)

	cid, err := unc.CID()
	assert.Equal(t, nil, err)
	assert.Equal(t, "zFuncm69CFEp3YTS4vgNGpXJHUKZVgo6Qzr6syTNjZrrKFugu4K1", cid.String())

	jb1, _ := json.Marshal(simpleCell)
	jb2, _ := json.Marshal(unc)
	assert.Equal(t, jb1, jb2)
}

func TestBinaryCell_Size(t *T) {
	c := &BinaryCell{
		OpCode: 0x1b1,
		Children: []*BinaryCell{
			{OpCode: 0x1b2, Memory: proto.EncodeVarint(xxhash.Sum64String("acc1"))},
			{OpCode: 0x1b3, Memory: proto.EncodeVarint(1e6)},
			{OpCode: 0x1c1, Memory: proto.EncodeVarint(xxhash.Sum64String("v1"))},
			{OpCode: 0x1c2, Memory: proto.EncodeVarint(xxhash.Sum64String("r2"))},
			{OpCode: 0x1c3, Memory: proto.EncodeVarint(xxhash.Sum64String("s3"))},
		},
	}
	b, _ := c.Marshal()
	assert.Equal(t, 64, len(b))
	assert.Equal(t, 130, len(fmt.Sprintf("0x%x", b)))
	b, _ = json.Marshal(c)
	// logger.Printf("%s", b)
	assert.Equal(t, 142, len(b))
}

func TestBinaryCell_Size2(t *T) {
	c := &BinaryCell{Children: []*BinaryCell{{OpCode: 1}, {OpCode: 1}, {OpCode: 1}, {OpCode: 1}}}
	b, _ := c.Marshal()
	assert.Equal(t, 15, len(b))
	c = &BinaryCell{Children: []*BinaryCell{{OpCode: 1}, {OpCode: 1}, {OpCode: 1}}}
	b, _ = c.Marshal()
	assert.Equal(t, 12, len(b))
	c = &BinaryCell{Children: []*BinaryCell{{OpCode: 1}, {OpCode: 1}}}
	b, _ = c.Marshal()
	assert.Equal(t, 9, len(b))
	c = &BinaryCell{Children: []*BinaryCell{{OpCode: 1}}}
	b, _ = c.Marshal()
	assert.Equal(t, 6, len(b))
	c = &BinaryCell{OpCode: 1, Children: []*BinaryCell{{OpCode: 1}}}
	b, _ = c.Marshal()
	assert.Equal(t, 6, len(b))
	c = &BinaryCell{OpCode: 1}
	b, _ = c.Marshal()
	assert.Equal(t, 3, len(b))
}

var benchCell = &BinaryCell{
	Memory: []byte("test"),
	Children: []*BinaryCell{
		{OpCode: 0x01, Children: []*BinaryCell{
			{OpCode: 0x10, Memory: cidMem},
			{OpCode: 0x10, Memory: cidMem},
			{OpCode: 0x10, Memory: cidMem},
			{OpCode: 0x10, Memory: cidMem},
			{OpCode: 0x10, Memory: cidMem},
			{OpCode: 0x10, Memory: cidMem},
		}},
	},
}

func BenchmarkBinaryCell_Marshal(b *B) {
	for i := 0; i < b.N; i++ {
		_, _ = benchCell.Marshal()
	}
}
