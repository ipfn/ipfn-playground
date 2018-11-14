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

package cells

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
	cidCell    = New(0xc1d, cidMem)
	simpleCell = New(0, []byte("test"),
		Op(0x01,
			New(0x10, []byte("0x0023")),
			New(0x10, []byte("0x0325"))))
)

func TestBinaryCell(t *T) {
	t.Parallel()

	b, err := Marshal(simpleCell)
	assert.Equal(t, nil, err)
	de := new(BinaryCell)
	err = Unmarshal(de, b)
	assert.Equal(t, nil, err)
	b2, _ := Marshal(de)
	assert.Equal(t, b, b2)
	assert.Equal(t, simpleCell.OpCode(), de.OpCode())

	id, err := de.Checksum()
	assert.Equal(t, nil, err)
	assert.Equal(t, ID(0x1663e9dbd3c404b2), id)

	cid := de.CID()
	assert.Equal(t, nil, err)
	assert.Equal(t, "zFuncktLhQmWU8P24RYpuMGWgY5RVX9QpTcKmfGf2TdDQiaxF1aA", cid.String())

	jb1, _ := json.Marshal(simpleCell)
	jb2, _ := json.Marshal(de)
	assert.Equal(t, jb1, jb2)
}

func TestBinaryCell_Size(t *T) {
	t.Parallel()

	c := Op(0x1b1,
		New(0x1b2, proto.EncodeVarint(xxhash.Sum64String("acc1"))),
		New(0x1b3, proto.EncodeVarint(1e6)),
		New(0x1c1, proto.EncodeVarint(xxhash.Sum64String("v1"))),
		New(0x1c2, proto.EncodeVarint(xxhash.Sum64String("r2"))),
		New(0x1c3, proto.EncodeVarint(xxhash.Sum64String("s3"))))
	b, _ := Marshal(c)
	assert.Equal(t, 64, len(b))
	assert.Equal(t, 130, len(fmt.Sprintf("0x%x", b)))
	b, _ = json.Marshal(c)
	// logger.Printf("%s", b)
	assert.Equal(t, 158, len(b))
}

func TestBinaryCell_Size2(t *T) {
	t.Parallel()

	c := Root(Ops(Op(1), Op(1), Op(1), Op(1)))
	b, _ := Marshal(c)
	assert.Equal(t, 15, len(b))
	c = Root(Ops(Op(1), Op(1), Op(1)))
	b, _ = Marshal(c)
	assert.Equal(t, 12, len(b))
	c = Root(Ops(Op(1), Op(1)))
	b, _ = Marshal(c)
	assert.Equal(t, 9, len(b))
	c = Root(Ops(Op(1)))
	b, _ = Marshal(c)
	assert.Equal(t, 6, len(b))
	c = Op(1, Op(1))
	b, _ = Marshal(c)
	assert.Equal(t, 6, len(b))
	c = Op(1)
	b, _ = Marshal(c)
	assert.Equal(t, 3, len(b))
}

var benchCell = New(0, []byte("test"),
	Op(0x01,
		New(0x10, cidMem),
		New(0x10, cidMem),
		New(0x10, cidMem),
		New(0x10, cidMem),
		New(0x10, cidMem),
		New(0x10, cidMem)))

func BenchmarkBinaryCell_Marshal(b *B) {
	for i := 0; i < b.N; i++ {
		_, _ = Marshal(benchCell)
	}
}
