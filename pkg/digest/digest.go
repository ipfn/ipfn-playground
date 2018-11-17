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
	"encoding/hex"
	"hash"
	"io"
)

// Size - Default digest size.
const Size = 32

// Digest - Hash digest.
// Common type for map keys.
type Digest [Size]byte

// emptyDigest - used for comparison in IsEmpty.
var emptyDigest = Digest{0}

// Sum - Sums hash digest using provided hasher.
func Sum(h hash.Hash, data ...[]byte) (digest Digest) {
	if r, ok := h.(io.Reader); ok {
		h.Reset()
		for _, body := range data {
			h.Write(body)
		}
		r.Read(digest[:])
		return
	}
	return FromBytes(SumBytes(h, data...))
}

// SumBytes - Sums hash digest using provided hasher.
func SumBytes(h hash.Hash, data ...[]byte) (digest []byte) {
	h.Reset()
	for _, body := range data {
		h.Write(body)
	}
	return h.Sum(nil)
}

// FromHex - Creates hash digest from parsed hex hash.
func FromHex(src string) (digest Digest) {
	hex.Decode(digest[:], []byte(src))
	return
}

// FromBytes - Creates hash digest from bytes.
func FromBytes(src []byte) (digest Digest) {
	copy(digest[:], src)
	return
}

// IsEmpty - Returns true if hash digest is empty.
func IsEmpty(h Digest) bool {
	return h == emptyDigest
}

// Empty - Returns empty hash digest.
func Empty() Digest {
	return emptyDigest
}

// Compare - Checks two digests for equality.
func Compare(a, b Digest) int {
	return bytes.Compare(a[:], b[:])
}

// Bytes - Returns digest bytes.
func (digest Digest) Bytes() []byte {
	return digest[:]
}
