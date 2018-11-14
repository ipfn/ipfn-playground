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

package bccsp

import "github.com/ipfn/ipfn/pkg/utils/hexutil"

// Hash - Common hash interface.
type Hash interface {
	// Algorithm - Hash algorithm.
	Algorithm() HashType
	// Bytes - Hash bytes.
	Bytes() []byte
	// String - Hash hex string.
	String() string
	// Size - Hash bytes size.
	Size() int
}

// DefaultSize - Default hash byte size.
const DefaultSize = 32

// NewHash - Constructs common hash interface from computed hash and algorithm ID.
func NewHash(algo HashType, body []byte) Hash {
	return customHash{algo: algo, body: body}
}

// FromHexString - Creates hash from hex string.
func FromHexString(algo HashType, hex string) Hash {
	return NewHash(algo, hexutil.FromString(hex))
}

type customHash struct {
	algo HashType
	body []byte
}

// Algorithm - Hash algorithm.
func (hash customHash) Algorithm() HashType {
	return hash.algo
}

// Bytes - Hash bytes.
func (hash customHash) Bytes() []byte {
	return []byte(hash.body)
}

// String - Hash hex string.
func (hash customHash) String() string {
	return hexutil.ToString(hash.body)
}

// Size - Hash bytes size.
func (hash customHash) Size() int {
	return len(hash.body)
}
