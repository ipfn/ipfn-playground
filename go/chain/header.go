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
	"fmt"
	"time"

	"github.com/ipfn/ipfn/go/cids"
	"github.com/ipfn/ipfn/go/opcode"
	"github.com/ipfn/ipfn/go/opcode/chainops"
	"github.com/ipfn/ipfn/go/opcode/synaptic"
	cid "github.com/ipfs/go-cid"
	multihash "github.com/multiformats/go-multihash"
)

// Header - State header structure.
type Header struct {
	// Index - State index.
	Index uint64 `json:"index,omitempty"`

	// Time - State time.
	Time time.Time `json:"timestamp,omitempty"`

	// CID - Content ID of this header.
	// Computed from previous and execution hashes.
	CID *opcode.CID `json:"head_hash,omitempty"`

	// Prev - Previous state hash.
	// Can be empty only on zero index.
	Prev *opcode.CID `json:"prev_hash,omitempty"`

	// Exec - State execution hash.
	Exec *opcode.CID `json:"exec_hash,omitempty"`

	// State - State trie hash.
	State *opcode.CID `json:"state_hash,omitempty"`

	// Signed - Signed hash.
	Signed *opcode.CID `json:"signed_hash,omitempty"`
}

// HeaderPrefix - Header CID prefix.
var HeaderPrefix = cid.Prefix{
	Version:  1,
	Codec:    cids.ChainHeader,
	MhType:   multihash.KECCAK_256,
	MhLength: 32,
}

// SignedPrefix - Signed header CID prefix.
var SignedPrefix = cid.Prefix{
	Version:  1,
	Codec:    cids.ChainSigned,
	MhType:   multihash.KECCAK_256,
	MhLength: 32,
}

// StateTriePrefix - State trie CID prefix.
var StateTriePrefix = cid.Prefix{
	Version:  1,
	Codec:    cids.ChainStateTrie,
	MhType:   multihash.KECCAK_256,
	MhLength: 32,
}

// NewHeader - Creates new state header structure.
func NewHeader(index uint64, prevHash *opcode.CID, execCID *opcode.CID) (hdr *Header, err error) {
	if prevHash == nil && index > 0 {
		return nil, fmt.Errorf("prev hash cannot be empty with index %d", index)
	}
	hdr = &Header{
		Index: index,
		Time:  time.Now(),
		Exec:  execCID,
		Prev:  prevHash,
	}

	hdr.CID, err = NewHeaderCID(hdr)
	if err != nil {
		return nil, err
	}
	// BUG(crackcomm): fucking state trie hash?!
	hdr.State, err = opcode.SumCID(StateTriePrefix, hdr.CID.Bytes())
	if err != nil {
		return
	}
	return
}

// NewHeaderCID - Computes header cid.
func NewHeaderCID(hdr *Header) (_ *opcode.CID, err error) {
	// BUG(crackcomm): fucking state in header?!
	body, err := hdr.BinaryCell().Marshal()
	if err != nil {
		return
	}
	return opcode.SumCID(HeaderPrefix, body)
}

// BinaryCell - Creates header binary cell.
func (hdr *Header) BinaryCell() *opcode.BinaryCell {
	return opcode.Op(chainops.OpHeader,
		synaptic.Uint64(hdr.Index),
		synaptic.Uint64(uint64(hdr.Time.Unix())),
		chainops.CID(hdr.Prev),
		chainops.CID(hdr.Exec),
	)
}
