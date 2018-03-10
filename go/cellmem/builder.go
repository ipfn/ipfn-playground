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

// TODO: bonds
import "github.com/ipfn/go-ipfn-cell"

// Builder - In-memory cell builder.
type Builder struct {
	cell *Cell
}

// NewBuilder - Creates a new Synaptic cell.
func NewBuilder() *Builder {
	return &Builder{cell: new(Cell)}
}

// Cell - Returns result of the builder.
func (builder *Builder) Cell() *Cell {
	return builder.cell
}

// SetMemory - Sets memory of the cell.
func (builder *Builder) SetMemory(memory cell.Memory) {
	builder.cell.cid = nil
	builder.cell.memory = memory
}

// SetMemoryBytes - Sets memory bytes of the cell.
func (builder *Builder) SetMemoryBytes(bytes []byte) {
	builder.cell.cid = nil
	builder.cell.memory = NewBytes(bytes)
}

// SetName - Sets name of the cell.
func (builder *Builder) SetName(name string) {
	builder.cell.cid = nil
	builder.cell.name = name
}

// SetSoul - Sets soul of the cell.
func (builder *Builder) SetSoul(soul string) {
	builder.cell.cid = nil
	builder.cell.soul = soul
}

// SetBonds - Sets bonds of the cell.
func (builder *Builder) SetBonds(bonds []cell.Bond) {
	builder.cell.cid = nil
	builder.cell.bonds = bonds
}

// SetBody - Sets body of the cell.
func (builder *Builder) SetBody(body []cell.Cell) {
	builder.cell.cid = nil
	builder.cell.body = body
}
