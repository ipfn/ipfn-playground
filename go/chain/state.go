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
	"errors"
	"fmt"

	"github.com/ipfn/ipfn/go/opcode/chainops"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ipfn/ipfn/go/opcode"
)

// State - Chain state structure.
type State struct {
	// Header - State header.
	Header *Header `json:"header,omitempty"`

	// ExecOps - State execution.
	ExecOps []*opcode.BinaryCell `json:"exec_ops,omitempty"`

	// Signatures - State signatures.
	Signatures [][]byte `json:"signatures,omitempty"`

	sigOps []*opcode.BinaryCell
}

// NewState - Creates new state structure.
func NewState(index uint64, prevHash *opcode.CID, execOps []*opcode.BinaryCell) (_ *State, err error) {
	if prevHash == nil && index > 0 {
		return nil, fmt.Errorf("prev hash cannot be empty with index %d", index)
	}
	execCID, err := opcode.RootOp(execOps).CID()
	if err != nil {
		return
	}
	header, err := NewHeader(index, prevHash, execCID)
	if err != nil {
		return
	}
	return &State{
		Header:  header,
		ExecOps: execOps,
	}, nil
}

// Head - Returns head CID.
func (state *State) Head() *opcode.CID {
	return state.Header.CID
}

// Signed - Returns signed head CID.
func (state *State) Signed() *opcode.CID {
	return state.Header.Signed
}

// Prev - Returns previous CID.
func (state *State) Prev() *opcode.CID {
	return state.Header.Prev
}

// Index - Returns state index.
func (state *State) Index() uint64 {
	return state.Header.Index
}

// Next - Returns next state including given ops.
func (state *State) Next(exec []*opcode.BinaryCell) (*State, error) {
	if len(exec) == 0 {
		return nil, errors.New("cannot produce state with zero operations")
	}
	return NewState(state.Index()+1, state.Head(), exec)
}

// Sign - Signs state with given private key.
// Computes new signed header hash.
func (state *State) Sign(key *btcec.PrivateKey) (err error) {
	// BUG(crackcomm): proper fucking signature xD
	signature, err := btcec.SignCompact(btcec.S256(), key, state.Head().Bytes(), false)
	if err != nil {
		return
	}
	state.Signatures = append(state.Signatures, signature)
	state.sigOps = append(state.sigOps, chainops.Signature(signature))
	// TODO this cid is just fucking wrong
	c, err := chainops.Signed(opcode.RootOp(state.ExecOps), state.sigOps).CID()
	if err != nil {
		return
	}
	state.Header.Signed, err = opcode.SumCID(SignedPrefix, c.Bytes())
	return
}
