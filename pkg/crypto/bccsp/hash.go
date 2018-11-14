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

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math"

	"github.com/ipfn/ipfn/pkg/utils/hexutil"
)

// Hash - Common hash interface.
type Hash interface {
	// Algorithm - Hash algorithm.
	Algorithm() HashType
	// Bytes - Encoded multihash.
	Bytes() []byte
	// Digest - Hash digest bytes.
	Digest() []byte
	// String - Hash hex string.
	String() string
	// Size - Hash digest size.
	Size() int
}

// DefaultHashSize - Default hash byte size.
const DefaultHashSize = 32

// DecodeHash - Decodes multihash.
func DecodeHash(input []byte) (_ Hash, err error) {
	// Source code from: go-multihash
	if len(input) < 2 {
		return nil, errors.New("multihash is too short")
	}
	code, body, err := uvarint(input)
	if err != nil {
		return
	}
	length, body, err := uvarint(body)
	if err != nil {
		return
	}
	if length > math.MaxInt32 {
		return nil, errors.New("digest too long, supporting only <= 2^31-1")
	}
	size := len(body)
	if length != uint64(size) {
		return nil, fmt.Errorf("inconsistent multihash length %d != %d", length, size)
	}
	return customHash{
		code:  HashType(code),
		start: len(input) - size,
		size:  size,
		body:  input,
	}, nil
}

// MustDecodeHash - Decodes multihash.
func MustDecodeHash(body []byte) Hash {
	hash, err := DecodeHash(body)
	if err != nil {
		panic(fmt.Sprintf("cannot decode hash: %v", err))
	}
	return hash
}

// HashDigest - Constructs common hash interface from computed hash and algorithm ID.
// It must encode multihash body structure first.
func HashDigest(algo HashType, digest []byte) Hash {
	size := len(digest)
	body := make([]byte, 2*binary.MaxVarintLen64+size)
	n := binary.PutUvarint(body, algo.Code())
	n += binary.PutUvarint(body[n:], uint64(size))
	for index := 0; index < size; index++ {
		body[index+n] = digest[index]
	}
	return customHash{code: algo, size: size, start: n, body: body[:size+n]}
}

// HashFromHex - Creates hash from hex string and algo.
func HashFromHex(algo HashType, hex string) Hash {
	return HashDigest(algo, hexutil.FromString(hex))
}

// IsHashEmpty - Returns true if hash is empty.
// Hash can be still empty without valid type.
func IsHashEmpty(h Hash) bool {
	return h.Size() == 0 || h.Algorithm() == UnknownHash
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
// NOTE: This hash is considered invalid and its
// only being purpose is to test hash validity.
func EmptyHash() Hash {
	return customHash{code: UnknownHash, body: []byte{}}
}

type customHash struct {
	code  HashType
	start int
	size  int
	body  []byte
}

// Algorithm - Hash algorithm.
func (hash customHash) Algorithm() HashType {
	return hash.code
}

// Digest - Hash digest bytes.
func (hash customHash) Digest() []byte {
	return hash.body[hash.start:]
}

// Bytes - Multihash bytes.
func (hash customHash) Bytes() []byte {
	return hash.body
}

// String - Hash hex string.
func (hash customHash) String() string {
	return hexutil.ToString(hash.body)
}

// Size - Hash bytes size.
func (hash customHash) Size() int {
	return hash.size
}

func uvarint(body []byte) (uint64, []byte, error) {
	n, c := binary.Uvarint(body)
	if c == 0 {
		return n, body, errors.New("uvarint: buffer too small")
	} else if c < 0 {
		return n, body[-c:], errors.New("uvarint: varint too big (max 64bit)")
	} else {
		return n, body[c:], nil
	}
}
