// Copyright © 2018 The IPFN Developers Authors. All Rights Reserved.
// Copyright © 2014-2018 The go-ethereum Authors. All Rights Reserved.
//
// This file is part of the IPFN project.
// This file was part of the go-ethereum library.
//
// The IPFN project is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The IPFN project is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package metrics

// Registry
type Registry interface {
}

// Counters hold an int64 value that can be incremented and decremented.
type Counter interface {
	Clear()
	Count() int64
	Dec(int64)
	Inc(int64)
	Snapshot() Counter
}

// GetOrRegisterCounter returns an existing Counter or constructs and registers
// a new StandardCounter.
func GetOrRegisterCounter(name string, r Registry) Counter {
	return NilCounter{}
}

// NewCounter constructs a new StandardCounter.
func NewCounter() Counter {
	return NilCounter{}
}

// NewRegisteredCounter constructs and registers a new StandardCounter.
func NewRegisteredCounter(name string, r Registry) Counter {
	return NilCounter{}
}

// NilCounter is a no-op Counter.
type NilCounter struct{}

// Clear is a no-op.
func (NilCounter) Clear() {}

// Count is a no-op.
func (NilCounter) Count() int64 { return 0 }

// Dec is a no-op.
func (NilCounter) Dec(i int64) {}

// Inc is a no-op.
func (NilCounter) Inc(i int64) {}

// Snapshot is a no-op.
func (NilCounter) Snapshot() Counter { return NilCounter{} }
