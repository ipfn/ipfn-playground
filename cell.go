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
	cell "github.com/ipfn/go-ipfn-cell"
	cellcid "github.com/ipfn/go-ipfn-cell-cid"
	cid "github.com/ipfs/go-cid"
)

// NewCell - Creates new protocol buffers cell.
func NewCell(c cell.Cell) (res *Cell, err error) {
	switch msg := c.(type) {
	case *CellWrapper:
		return msg.Cell, nil
	}
	memory, err := c.Memory().Bytes()
	if err != nil {
		return
	}
	res = &Cell{
		Name:  c.Name(),
		Soul:  c.Soul(),
		Value: memory,
	}
	bonds := c.Bonds()
	if l := len(bonds); l > 0 {
		res.Bonds = make([]*Bond, l)
		for n, child := range bonds {
			res.Bonds[n] = NewBond(child)
		}
	}
	body := c.Body()
	if l := len(body); l > 0 {
		res.Body = make([]*Cell, l)
		for n, child := range body {
			res.Body[n], err = NewCell(child)
			if err != nil {
				return
			}
		}
	}
	return
}

// NewCellWrapper - Creates a new cell from protocol buffers message.
func NewCellWrapper(c *Cell) (w *CellWrapper) {
	w = &CellWrapper{Cell: c}
	// Convert bonds
	if l := len(c.Bonds); l > 0 {
		w.bonds = make([]cell.Bond, l)
		for n, child := range c.Bonds {
			w.bonds[n] = NewBondWrapper(child)
		}
	}
	// Convert body
	if l := len(c.Body); l > 0 {
		w.body = make([]cell.Cell, l)
		for n, child := range c.Body {
			w.body[n] = NewCellWrapper(child)
		}
	}
	return
}

// CellWrapper - Protocol buffers cell wrapper.
type CellWrapper struct {
	Cell   *Cell
	cid    *cid.Cid
	body   []cell.Cell
	bonds  []cell.Bond
	memory cell.Memory
}

// CID - Returns content ID of the cell.
func (wrapper *CellWrapper) CID() (*cid.Cid, error) {
	if wrapper.cid != nil {
		return wrapper.cid, nil
	}
	return cellcid.CID(wrapper)
}

// Name - Returns name of the cell.
func (wrapper *CellWrapper) Name() string {
	return wrapper.Cell.Name
}

// Soul - Returns name of the cell soul.
func (wrapper *CellWrapper) Soul() string {
	return wrapper.Cell.Soul
}

// Body - Returns body of the cell.
func (wrapper *CellWrapper) Body() []cell.Cell {
	return wrapper.body
}

// Bonds - Returns bonds of the cell.
func (wrapper *CellWrapper) Bonds() []cell.Bond {
	return wrapper.bonds
}

// Memory - Returns memory of the cell.
func (wrapper *CellWrapper) Memory() cell.Memory {
	return wrapper
}

// Bytes - Returns memory of the cell.
func (wrapper *CellWrapper) Bytes() ([]byte, error) {
	return wrapper.Cell.Value, nil
}
