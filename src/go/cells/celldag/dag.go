// Copyright © 2017-2018 The IPFN Authors. All Rights Reserved.
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

package celldag

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"strconv"

	blocks "gx/ipfs/QmTRCUvZLiir12Qr6MV3HKfKMHX8Nf1Vddn6t2g5nsQSb9/go-block-format"
	ipld "gx/ipfs/QmWi2BYBL5gJ3CiAiQchg6rn1A8iBsrWy51EYxvHVjFvLb/go-ipld-format"
	cid "gx/ipfs/QmapdYm1b22Frv3k17fqrBYTFRxwiaVJkB299Mfn33edeB/go-cid"

	"github.com/ipfn/ipfn/src/go/cells"
	"github.com/ipfs/go-ipfs/core/coredag"
)

const formatName = "cell-binary"

func init() {
	// TODO(crackcomm):
	//   import contents "github.com/rootchain/go-rootchain/dev/contents"
	//   ipld.DefaultBlockDecoder.Register(contents.BinaryCell, decodeCellBlock)
	coredag.DefaultInputEncParsers.AddParser(formatName, formatName, parseCellDag)
}

func decodeCellBlock(block blocks.Block) (_ ipld.Node, err error) {
	return decodeCell(block.RawData())
}

func parseCellDag(r io.Reader, mhType uint64, mhLen int) (nodes []ipld.Node, err error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return
	}
	cell, err := decodeCell(data)
	if err != nil {
		return
	}
	return []ipld.Node{cell}, nil
}

func decodeCell(body []byte) (_ ipld.Node, err error) {
	cell, err := cells.UnmarshalBinary(body)
	if err != nil {
		return
	}
	return &dagCell{Cell: cell, body: body}, nil
}

type dagCell struct {
	cells.Cell
	body []byte
}

// func (c *dagCell) MarshalJSON() (_ []byte, err error) {
// 	return cells.NewPrinter(c.BinaryCell).MarshalJSON()
// }

func (c *dagCell) RawData() []byte {
	return c.body
}

func (c *dagCell) Cid() *cid.Cid {
	return c.Cell.CID().Cid
}

func (c *dagCell) String() string {
	return "<cell>"
}

func (c *dagCell) Loggable() map[string]interface{} {
	return nil
}

// ResolveLink is a helper function that calls resolve and asserts the
// output is a link
func (c *dagCell) ResolveLink(path []string) (_ *ipld.Link, _ []string, err error) {
	return
}

// Copy returns a deep copy of this node
func (c *dagCell) Copy() (_ ipld.Node) {
	return
}

// Links is a helper function that returns all links within this object
func (c *dagCell) Links() (_ []*ipld.Link) {
	return
}

// TODO: not sure if stat deserves to stay
func (c *dagCell) Stat() (_ *ipld.NodeStat, err error) {
	return
}

// Size returns the size in bytes of the serialized object
func (c *dagCell) Size() (_ uint64, err error) {
	return
}

// Resolve resolves a path through this node, stopping at any link boundary
// and returning the object found as well as the remaining path to traverse
func (c *dagCell) Resolve(path []string) (_ interface{}, _ []string, err error) {
	if len(path) == 0 {
		return
	}
	if len(path) < 2 {
		err = errors.New("invalid path")
		return
	}
	switch path[0] {
	case "op":
		num, err := strconv.ParseUint(path[1], 10, 32)
		if err != nil {
			return nil, nil, err
		}
		if c.ChildrenSize() > int(num) {
			child := c.Child(int(num))
			return &dagCell{Cell: child}, path[2:], nil
		}
	}
	err = fmt.Errorf("unrecognized path:%v", path)
	return
}

// Tree lists all paths within the object under 'path', and up to the given depth.
// To list the entire object (similar to `find .`) pass "" and -1
func (c *dagCell) Tree(path string, depth int) (_ []string) {
	return
}
