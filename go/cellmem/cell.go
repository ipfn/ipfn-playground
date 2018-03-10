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

package cellmem

import (
	cid "github.com/ipfs/go-cid"

	cell "github.com/ipfn/go-ipfn-cell"
	cellcid "github.com/ipfn/go-ipfn-cell-cid"
)

// Cell - In-memory cell.
type Cell struct {
	cid    *cid.Cid
	name   string
	soul   string
	memory cell.Memory
	bonds  []cell.Bond
	body   []cell.Cell
}

// NewCell - Creates a new Synaptic cell.
func NewCell(soul string, value []byte) *Cell {
	return &Cell{soul: soul, memory: NewBytes(value)}
}

// CID - Returns content ID of the cell.
func (cell *Cell) CID() (*cid.Cid, error) {
	if cell.cid != nil {
		return cell.cid, nil
	}
	return cellcid.CID(cell)
}

// Memory - Returns memory of the cell.
func (cell *Cell) Memory() cell.Memory {
	return cell.memory
}

// Name - Returns name of the cell.
func (cell *Cell) Name() string {
	return cell.name
}

// Soul - Returns soul of the cell.
func (cell *Cell) Soul() string {
	return cell.soul
}

// Bonds - Returns bonds of the cell.
func (cell *Cell) Bonds() []cell.Bond {
	return cell.bonds
}

// Body - Returns body of the cell.
func (cell *Cell) Body() []cell.Cell {
	return cell.body
}
