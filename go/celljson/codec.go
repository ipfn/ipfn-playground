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

// Package json implements JSON codec using cellpb package.
package json

import (
	"bytes"

	"github.com/gogo/protobuf/jsonpb"

	cell "github.com/ipfn/ipfn/go/cell"
	cellpb "github.com/ipfn/ipfn/go/cellpb"
)

// Codec - Protocol buffers cells codec.
type Codec struct{}

var (
	marshaler = &jsonpb.Marshaler{
		EmitDefaults: false,
		Indent:       "",
		OrigName:     true,
	}
	unmarshaler = new(jsonpb.Unmarshaler)
)

// Encode - Encodes a cell.
func Encode(c cell.Cell) ([]byte, error) {
	switch msg := c.(type) {
	case *cellpb.CellWrapper:
		var buf bytes.Buffer
		if err := marshaler.Marshal(&buf, msg.Cell); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	default:
		c, err := cellpb.NewCell(c)
		if err != nil {
			return nil, err
		}
		var buf bytes.Buffer
		if err := marshaler.Marshal(&buf, c); err != nil {
			return nil, err
		}
		return buf.Bytes(), nil
	}
}

// Decode - Decodes a cell.
func Decode(data []byte) (_ cell.Cell, err error) {
	c := new(cellpb.Cell)
	err = unmarshaler.Unmarshal(bytes.NewReader(data), c)
	if err != nil {
		return
	}
	return cellpb.NewCellWrapper(c), nil
}

// Encode - Encodes a cell.
func (codec *Codec) Encode(c cell.Cell) (_ []byte, err error) {
	return Encode(c)
}

// Decode - Decodes a cell.
func (codec *Codec) Decode(data []byte) (_ cell.Cell, err error) {
	return Decode(data)
}
