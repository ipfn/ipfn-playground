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
	"encoding/hex"
	"errors"
	"fmt"
	"math"
)

// HeaderSize - Maximum multihash header size.
// NOTE: It can be changed if there is an issue.
const HeaderSize = 10

// EncodedSize - Default hash byte size.
const EncodedSize = Size + HeaderSize

// Encoded - Encoded multihash digest with codec.
type Encoded [EncodedSize]byte

// DecodeHash - Decodes multihash.
func DecodeHash(input []byte) (_ Hash, err error) {
	// Handle empty hash case
	if len(input) == 0 {
		return EmptyHash(), nil
	}
	// Source code from: go-multihash
	if len(input) < 2 {
		return nil, errors.New("multihash is too short")
	}
	if len(input) > EncodedSize {
		return nil, errors.New("multihash is too long")
	}
	code, body, err := uvarint(input)
	if err != nil {
		return
	}
	length, body, err := uvarint(body)
	if err != nil {
		return
	}
	if code > math.MaxInt32 {
		return nil, errors.New("implementation supports only 32 bit integer codes")
	}
	if length > math.MaxInt32 {
		return nil, errors.New("implementation supports only 32 bit integer length")
	}
	size := len(body)
	if length != uint64(size) {
		return nil, fmt.Errorf("inconsistent multihash length %d != %d", length, size)
	}
	var result [EncodedSize]byte
	copy(result[:], input)
	return customHash{
		code:  Type(code),
		start: len(input) - size,
		size:  size,
		body:  result,
	}, nil
}

// HashFromSum - Constructs common hash interface from computed hash and algorithm ID.
// It must encode multihash body structure first.
func HashFromSum(algo Type, digest []byte) Hash {
	size := len(digest)
	var result Encoded
	// multihash: hash codec
	n := binary.PutUvarint(result[:], algo.Code())
	// multihash: digest size
	n += binary.PutUvarint(result[n:], uint64(size))
	// multihash: copy digest
	copy(result[n:], digest)
	return customHash{code: algo, size: size, start: n, body: result}
}

// HashFromDigest - Constructs common hash interface from computed hash and algorithm ID.
// It must encode multihash body structure first.
func HashFromDigest(algo Type, digest Digest) Hash {
	return HashFromSum(algo, digest[:])
}

// HashFromHex - Creates hash from hex string and algo.
func HashFromHex(algo Type, src string) Hash {
	size := hex.DecodedLen(len(src))
	var result Encoded
	// multihash: hash codec
	n := binary.PutUvarint(result[:], algo.Code())
	// multihash: digest size
	n += binary.PutUvarint(result[n:], uint64(size))
	// multihash: decode hex digest
	hex.Decode(result[n:], []byte(src))
	return customHash{code: algo, size: size, start: n, body: result}
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
