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
	"fmt"
	"math"

	"github.com/gogo/protobuf/proto"
)

// BinaryCell - Binary representation of cell.
type BinaryCell struct {
	// OpCode - Native op code identifier.
	OpCode ID `json:"op,omitempty"`

	// Memory - Encoded cell memory value.
	Memory []byte `json:"memory,omitempty"`

	// Children - Operation input body.
	Children []*BinaryCell `json:"children,omitempty"`

	cid  *CID
	body []byte
}

// CID - Computes marshaled cell cid.
func (cell *BinaryCell) CID() (_ *CID, err error) {
	if cell.cid != nil {
		return cell.cid, nil
	}
	cell.cid, err = cell.computeCID()
	if err != nil {
		return
	}
	return cell.cid, nil
}

// Checksum - Computes marshalled xxhash64 of cell content id.
func (cell *BinaryCell) Checksum() (_ ID, err error) {
	hash, err := cell.CID()
	if err != nil {
		return
	}
	return NewID(hash.Bytes()), nil
}

// Add - Appends new children operation.
func (cell *BinaryCell) Add(child *BinaryCell) {
	cell.Children = append(cell.Children, child)
}

// MarshalJSON - Marshals cell as JSON.
func (cell *BinaryCell) MarshalJSON() (_ []byte, err error) {
	return prettyPrint(cell), nil
}

// Marshal - Marshals cell as byte array.
func (cell *BinaryCell) Marshal() (_ []byte, err error) {
	if cell.body != nil {
		return cell.body, nil
	}
	buff := proto.NewBuffer(nil)
	err = cell.marshal(buff)
	if err != nil {
		return
	}
	cell.body = buff.Bytes()
	return cell.body, nil
}

// Unmarshal - Unmarshals cell from byte array.
func (cell *BinaryCell) Unmarshal(body []byte) (err error) {
	return cell.unmarshal(proto.NewBuffer(body))
}

func (cell *BinaryCell) unmarshal(buff *proto.Buffer) (err error) {
	opCode, err := buff.DecodeVarint()
	if err != nil {
		return err
	}
	cell.OpCode = ID(opCode)
	if cell.Memory, err = buff.DecodeRawBytes(false); err != nil {
		return err
	}
	children, err := buff.DecodeVarint()
	if err != nil {
		return err
	}
	if children >= math.MaxInt32 {
		return fmt.Errorf("children length too big %d", children)
	}
	cell.Children = make([]*BinaryCell, children)
	for index := 0; index < int(children); index++ {
		child := new(BinaryCell)
		if err := child.unmarshal(buff); err != nil {
			return err
		}
		cell.Children[index] = child
	}
	return
}

func (cell *BinaryCell) marshal(buff *proto.Buffer) (err error) {
	if err := buff.EncodeVarint(uint64(cell.OpCode)); err != nil {
		return err
	}
	if err := buff.EncodeRawBytes(cell.Memory); err != nil {
		return err
	}
	if err := buff.EncodeVarint(uint64(len(cell.Children))); err != nil {
		return err
	}
	for _, child := range cell.Children {
		if err := child.marshal(buff); err != nil {
			return err
		}
	}
	return
}

func (cell *BinaryCell) computeCID() (_ *CID, err error) {
	body, err := cell.Marshal()
	if err != nil {
		return
	}
	return SumCID(CellPrefix, body)
}
