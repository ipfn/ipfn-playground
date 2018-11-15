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
	"encoding/binary"
	"fmt"
	"hash"
)

// Hasher - Hasher interface.
type Hasher interface {
	hash.Hash

	// Algorithm - Hashing algorithm.
	Algorithm() Type
}

// NewHasher - Wraps hasher with multihash encoding.
func NewHasher(code Type, h hash.Hash) Hasher {
	return &hasher{Hash: h, code: code}
}

type hasher struct {
	hash.Hash
	code Type
}

// Algorithm - Hashing algorithm.
func (h *hasher) Algorithm() Type {
	return h.code
}

// Sum appends the current hash to b and returns the resulting slice.
// It does not change the underlying hash state.
func (h *hasher) Sum(b []byte) []byte {
	// digest: size
	size := h.Hash.Size()
	// multihash: header
	body := make([]byte, len(b)+HeaderSize)
	// hash: copy source b
	n := copy(body, b)
	// multihash: hash codec
	n += binary.PutUvarint(body[n:], h.code.Code())
	// multihash: digest size
	n += binary.PutUvarint(body[n:], uint64(size))
	// multihash: join user input to header and body
	return h.Hash.Sum(body[:n])
}

// Size returns the number of bytes Sum will return.
func (h *hasher) Size() int {
	// digest: size
	size := h.Hash.Size()
	// multihash: hash codec
	n := varintSize(h.code.Code())
	// multihash: digest size
	n += varintSize(uint64(size))
	// multihash: save space;
	if n > HeaderSize {
		panic(fmt.Errorf("size too big for code=%d size=%d", h.code, size))
	}
	return n + size
}

func varintSize(x uint64) int {
	i := 0
	for x >= 0x80 {
		x >>= 7
		i++
	}
	return i + 1
}
