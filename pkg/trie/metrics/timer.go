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

import "time"

// Timers capture the duration and rate of events.
type Timer interface {
	Count() int64
	Max() int64
	Mean() float64
	Min() int64
	Percentile(float64) float64
	Percentiles([]float64) []float64
	Rate1() float64
	Rate5() float64
	Rate15() float64
	RateMean() float64
	Snapshot() Timer
	StdDev() float64
	Stop()
	Sum() int64
	Time(func())
	Update(time.Duration)
	UpdateSince(time.Time)
	Variance() float64
}

func NewRegisteredResettingTimer(name string, r Registry) Timer {
	return NilTimer{}
}

// NilTimer is a no-op Timer.
type NilTimer struct {
}

// Count is a no-op.
func (NilTimer) Count() int64 { return 0 }

// Max is a no-op.
func (NilTimer) Max() int64 { return 0 }

// Mean is a no-op.
func (NilTimer) Mean() float64 { return 0.0 }

// Min is a no-op.
func (NilTimer) Min() int64 { return 0 }

// Percentile is a no-op.
func (NilTimer) Percentile(p float64) float64 { return 0.0 }

// Percentiles is a no-op.
func (NilTimer) Percentiles(ps []float64) []float64 {
	return make([]float64, len(ps))
}

// Rate1 is a no-op.
func (NilTimer) Rate1() float64 { return 0.0 }

// Rate5 is a no-op.
func (NilTimer) Rate5() float64 { return 0.0 }

// Rate15 is a no-op.
func (NilTimer) Rate15() float64 { return 0.0 }

// RateMean is a no-op.
func (NilTimer) RateMean() float64 { return 0.0 }

// Snapshot is a no-op.
func (NilTimer) Snapshot() Timer { return NilTimer{} }

// StdDev is a no-op.
func (NilTimer) StdDev() float64 { return 0.0 }

// Stop is a no-op.
func (NilTimer) Stop() {}

// Sum is a no-op.
func (NilTimer) Sum() int64 { return 0 }

// Time is a no-op.
func (NilTimer) Time(func()) {}

// Update is a no-op.
func (NilTimer) Update(time.Duration) {}

// UpdateSince is a no-op.
func (NilTimer) UpdateSince(time.Time) {}

// Variance is a no-op.
func (NilTimer) Variance() float64 { return 0.0 }
