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

package digest

import (
	"bytes"

	"github.com/ipfn/ipfn/pkg/utils/hexutil"
)

// Hash - Common hash interface.
type Hash interface {
	// Algorithm - Hash algorithm.
	Algorithm() Type
	// Bytes - Encoded multihash.
	Bytes() []byte
	// Digest - Hash digest bytes.
	Digest() []byte
	// String - Hash hex string.
	String() string
	// Size - Hash digest size.
	Size() int
}

// IsHashEmpty - Returns true if hash is empty.
// Hash can be still empty without valid type.
func IsHashEmpty(h Hash) bool {
	return h == nil || h.Size() == 0 || h.Algorithm() == UnknownType
}

// HashEqual - Checks two hashes for equality of algo and digest.
func HashEqual(a, b Hash) bool {
	return bytes.Equal(a.Bytes(), b.Bytes())
}

// HashCompare - Checks two hashes for equality of algo and digest.
func HashCompare(a, b Hash) int {
	return bytes.Compare(a.Bytes(), b.Bytes())
}

// EmptyHash - Creates empty hash structure.
// NOTE: This hash is considered invalid in most applications.
func EmptyHash() Hash {
	return customHash{}
}

type customHash struct {
	code  Type
	start int
	size  int
	body  Encoded
}

// Algorithm - Hash algorithm.
func (hash customHash) Algorithm() Type {
	return hash.code
}

// Digest - Hash digest bytes.
func (hash customHash) Digest() []byte {
	return hash.body[hash.start : hash.start+hash.size]
}

// Bytes - Multihash bytes.
func (hash customHash) Bytes() []byte {
	return hash.body[:hash.start+hash.size]
}

// String - Hash hex string.
func (hash customHash) String() string {
	return hexutil.ToString(hash.body[:])
}

// Size - Hash bytes size.
func (hash customHash) Size() int {
	return hash.size
}
