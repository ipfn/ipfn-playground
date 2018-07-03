// Copyright © 2017 The IPFN Authors. All Rights Reserved.
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

package dag

import (
	"io"

	ipld "gx/ipfs/QmWi2BYBL5gJ3CiAiQchg6rn1A8iBsrWy51EYxvHVjFvLb/go-ipld-format"

	cid "github.com/ipfs/go-cid"
	"github.com/ipfs/go-ipfs/core/coredag"
)

const formatName = "cell-binary"

func init() {
	coredag.DefaultInputEncParsers.AddParser(formatName, formatName, cellDagParser)
}

func cellDagParser(r io.Reader, mhType uint64, mhLen int) (nodes []ipld.Node, err error) {
	return
}

type dagResolver struct {
}

// Resolve resolves a path through this node, stopping at any link boundary
// and returning the object found as well as the remaining path to traverse
func (c *dagResolver) Resolve(path []string) (_ interface{}, _ []string, err error) {
	return
}

// Tree lists all paths within the object under 'path', and up to the given depth.
// To list the entire object (similar to `find .`) pass "" and -1
func (c *dagResolver) Tree(path string, depth int) (_ []string) {
	return
}

type dagNode struct {
}

func (c *dagNode) RawData() []byte {
	return nil
}

func (c *dagNode) Cid() *cid.Cid {
	return nil
}

func (c *dagNode) String() string {
	return ""
}

func (c *dagNode) Loggable() map[string]interface{} {
	return nil
}

// ResolveLink is a helper function that calls resolve and asserts the
// output is a link
func (c *dagNode) ResolveLink(path []string) (_ *ipld.Link, _ []string, err error) {
	return
}

// Copy returns a deep copy of this node
func (c *dagNode) Copy() (_ ipld.Node) {
	return
}

// Links is a helper function that returns all links within this object
func (c *dagNode) Links() (_ []*ipld.Link) {
	return
}

// TODO: not sure if stat deserves to stay
func (c *dagNode) Stat() (_ *ipld.NodeStat, err error) {
	return
}

// Size returns the size in bytes of the serialized object
func (c *dagNode) Size() (_ uint64, err error) {
	return
}
