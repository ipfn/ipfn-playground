// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2013-2014 The btcsuite developers. All Rights Reserved.
//
// Use of this source code is governed by an ISC license.
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package base32i

import (
	"errors"
	"hash/crc32"

	// Only for proto.EncodeVarint
	"github.com/gogo/protobuf/proto"
)

// ErrChecksum indicates that the checksum of a check-encoded string does not verify against
// the checksum.
var ErrChecksum = errors.New("checksum error")

// ErrInvalidFormat indicates that the check-encoded string has an invalid format.
var ErrInvalidFormat = errors.New("invalid format: checksum bytes missing")

const cSize = 1

func checksum(input []byte) (cksum byte) {
	return proto.EncodeVarint(uint64(crc32.ChecksumIEEE(input)))[0]
}

// CheckEncode prepends and appends a four byte checksum.
func CheckEncode(input []byte) []byte {
	return Encode(checkBuffer(input))
}

// CheckEncodeToString is CheckEncode to string.
func CheckEncodeToString(input []byte) string {
	return EncodeToString(checkBuffer(input))
}

// CheckDecodeString decodes a string that was encoded with CheckEncode and verifies the checksum.
func CheckDecodeString(input string) (result []byte, err error) {
	decoded, err := DecodeString(input)
	if err != nil {
		return
	}
	if len(decoded) < 1 {
		err = ErrInvalidFormat
		return
	}
	cksum := decoded[len(decoded)-1]
	result = decoded[:len(decoded)-1]
	if checksum(result) != cksum {
		err = ErrChecksum
		return
	}
	return
}

func checkBuffer(input []byte) []byte {
	return append(input, checksum(input))
}
