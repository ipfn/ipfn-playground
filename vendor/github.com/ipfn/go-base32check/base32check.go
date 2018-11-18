// Copyright (c) 2018 The IPFN Developers
// Copyright (c) 2013-2018 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package base32check

import (
	"errors"
	"hash/crc32"

	"github.com/gogo/protobuf/proto"
)

// ErrChecksum indicates that the checksum of a check-encoded string does not verify against
// the checksum.
var ErrChecksum = errors.New("checksum error")

// ErrInvalidFormat indicates that the check-encoded string has an invalid format.
var ErrInvalidFormat = errors.New("invalid format: checksum bytes missing")

const cSize = 1

// checksum: first four bytes of sha256^2
func checksum(input []byte) (cksum byte) {
	// log.Printf("cksum: %x (-1=%x)", body, body[0])
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
