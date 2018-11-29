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
	"encoding/base32"
)

// Base32Alphabet - Rootchain address encoding alphabet.
const Base32Alphabet = "0pzqy9x8bf2tvrwds3jn54khce6mua7l"

// Encoding - Rootchain address encoder.
var Encoding = base32.NewEncoding(Base32Alphabet).WithPadding(base32.NoPadding)

// Decode - Encodes rootchain address bytes.
func Decode(src []byte) (body []byte, err error) {
	body = make([]byte, Encoding.DecodedLen(len(src)))
	_, err = Encoding.Decode(body, src)
	return
}

// DecodeString - Encodes rootchain address bytes.
func DecodeString(src string) (body []byte, err error) {
	body, err = Encoding.DecodeString(src)
	return
}

// Encode - Encodes rootchain address bytes.
func Encode(src []byte) []byte {
	buf := make([]byte, Encoding.EncodedLen(len(src)))
	Encoding.Encode(buf, src)
	return buf
}

// EncodeToString - Encodes rootchain address bytes to string.
func EncodeToString(src []byte) string {
	return Encoding.EncodeToString(src)
}
