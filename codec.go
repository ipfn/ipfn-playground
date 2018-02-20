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

package cellpb

import (
	proto "github.com/gogo/protobuf/proto"

	cell "github.com/ipfn/go-ipfn-cell"
	cellcid "github.com/ipfn/go-ipfn-cell-cid"
)

// Encode - Encodes a cell.
func Encode(cell cell.Cell) ([]byte, error) {
	switch msg := cell.(type) {
	case *CellWrapper:
		return proto.Marshal(msg.Cell)
	default:
		c, err := NewCell(cell)
		if err != nil {
			return nil, err
		}
		return proto.Marshal(c)
	}
}

// Decode - Decodes a cell.
func Decode(data []byte) (_ cell.Cell, err error) {
	cell := new(Cell)
	err = proto.Unmarshal(data, cell)
	if err != nil {
		return
	}
	wrapped := NewCellWrapper(cell)
	// Calculate it here, so we don't have to encode it again
	wrapped.cid, err = cellcid.Encoded(data, CodecID)
	if err != nil {
		return
	}
	return wrapped, nil
}

// Codec - Protocol buffers cells codec.
type Codec struct{}

// Encode - Encodes a cell.
func (codec *Codec) Encode(cell cell.Cell) (_ []byte, err error) {
	return Encode(cell)
}

// Decode - Decodes a cell.
func (codec *Codec) Decode(data []byte) (_ cell.Cell, err error) {
	return Decode(data)
}
