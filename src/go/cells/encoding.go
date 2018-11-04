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
	"math"

	"github.com/gogo/protobuf/proto"
)

// ToHex - Encodes binary cell to hex.
func ToHex(cell Cell) (body []byte, err error) {
	src, err := Marshal(cell)
	if err != nil {
		return
	}
	body = make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(body, src)
	return
}

// FromHex - Decodes cell from hex encoded binary cell.
func FromHex(src []byte) (_ Cell, err error) {
	n, err := hex.Decode(src, src)
	if err != nil {
		return
	}
	cell := new(BinaryCell)
	err = Unmarshal(cell, src[:n])
	if err != nil {
		return
	}
	return cell, nil
}

// FromHexString - Decodes cell from hex encoded binary cell.
func FromHexString(src string) (_ Cell, err error) {
	return FromHex([]byte(src))
}

// MarshalJSON - Marshals cell as JSON.
func MarshalJSON(cell Cell) (_ []byte, err error) {
	type jsonCell struct {
		OpCode   ID             `json:"op,omitempty"`
		Memory   []byte         `json:"value,omitempty"`
		Children []*CellPrinter `json:"ops,omitempty"`
	}
	childSize := cell.ChildrenSize()
	children := make([]*CellPrinter, childSize)
	for index := 0; index < childSize; index++ {
		children[index] = NewPrinter(cell.Child(index))
	}
	return json.Marshal(jsonCell{
		OpCode:   cell.OpCode(),
		Memory:   cell.Memory(),
		Children: children,
	})
}

// Marshal - Marshals cell as byte array.
func Marshal(cell Cell) (body []byte, err error) {
	switch v := cell.(type) {
	case *BinaryCell:
		if v.body != nil {
			return v.body, nil
		}
	}
	buff := proto.NewBuffer(nil)
	err = marshal(cell, buff)
	if err != nil {
		return
	}
	body = buff.Bytes()
	switch v := cell.(type) {
	case *BinaryCell:
		v.body = body
	}
	return
}

// Unmarshal - Unmarshals cell from byte array.
func Unmarshal(cell MutableCell, body []byte) error {
	switch v := cell.(type) {
	case *BinaryCell:
		v.body = body
	}
	return unmarshal(cell, proto.NewBuffer(body))
}

func unmarshal(cell MutableCell, buff *proto.Buffer) (err error) {
	opCode, err := buff.DecodeVarint()
	if err != nil {
		return err
	}
	cell.SetOpCode(ID(opCode))
	memory, err := buff.DecodeRawBytes(false)
	if err != nil {
		return err
	}
	cell.SetMemory(memory)
	childSize, err := buff.DecodeVarint()
	if err != nil {
		return err
	}
	if childSize >= math.MaxInt32 {
		return fmt.Errorf("children length too big %d", childSize)
	}
	children := make([]Cell, childSize)
	for index := 0; index < int(childSize); index++ {
		child := new(BinaryCell)
		if err := unmarshal(child, buff); err != nil {
			return err
		}
		children[index] = child
	}
	cell.SetChildren(children)
	return
}

func marshal(cell Cell, buff *proto.Buffer) (err error) {
	// if cell.body != nil {
	// 	buff.SetBuf(append(buff.Bytes(), cell.body...))
	// 	return
	// }
	if err := buff.EncodeVarint(uint64(cell.OpCode())); err != nil {
		return err
	}
	if err := buff.EncodeRawBytes(cell.Memory()); err != nil {
		return err
	}
	children := cell.ChildrenSize()
	if err := buff.EncodeVarint(uint64(children)); err != nil {
		return err
	}
	for index := 0; index < children; index++ {
		if err := marshal(cell.Child(index), buff); err != nil {
			return err
		}
	}
	return
}
