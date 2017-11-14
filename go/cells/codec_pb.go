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

package cells

import (
	"github.com/gogo/protobuf/proto"
	"github.com/ipfn/ipfn/go/cells/pb"
)

type codecProtoBuf struct {
}

// Encode - Encodes a cell.
func (codec *codecProtoBuf) Encode(cell Cell) (_ []byte, err error) {
	switch msg := cell.(type) {
	case *cellProtoBuf:
		return proto.Marshal(msg.Cell)
	default:
		panic("not implemented")
	}
}

// Decode - Decodes a cell.
// TODO(crackcomm): decode memory
func (codec *codecProtoBuf) Decode(data []byte) (_ Cell, err error) {
	cell := new(pb.Cell)
	err = proto.Unmarshal(data, cell)
	if err != nil {
		return
	}
	return &cellProtoBuf{Cell: cell}, nil
}

type cellProtoBuf struct {
	*pb.Cell

	// interface slice
	body []Cell
}

func newProtoBufCell(c Cell) (_ *pb.Cell, err error) {
	cellBody := c.Body()
	body := make([]*pb.Cell, len(cellBody))
	for n, child := range cellBody {
		body[n], err = newProtoBufCell(child)
		if err != nil {
			return
		}
	}
	return &pb.Cell{
		Name:   c.Name(),
		Soul:   c.Soul(),
		Body:   body,
		Bonds:  c.Bonds(),
		Memory: c.Memory().([]byte),
	}, nil
}

func newCellFromProtoBuf(c *pb.Cell) *cellProtoBuf {
	body := make([]Cell, len(c.Body))
	for n, child := range c.Body {
		body[n] = newCellFromProtoBuf(child)
	}
	return &cellProtoBuf{Cell: c, body: body}
}

// Name - Returns name of the cell.
func (cell *cellProtoBuf) Name() string {
	return cell.GetName()
}

// Soul - Returns name of the cell soul.
func (cell *cellProtoBuf) Soul() string {
	return cell.GetSoul()
}

// Body - Returns body of the cell.
func (cell *cellProtoBuf) Body() []Cell {
	return cell.body
}

// Bonds - Returns bonds of the cell.
func (cell *cellProtoBuf) Bonds() []string {
	return cell.GetBonds()
}

// Memory - Returns memory of the cell.
// BUG(crackcomm): this returns byte array
func (cell *cellProtoBuf) Memory() interface{} {
	return cell.GetMemory()
}
