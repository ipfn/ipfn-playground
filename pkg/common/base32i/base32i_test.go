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
	"encoding/base64"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testTxt = []byte("test")

func TestBase32ToString(t *testing.T) {
	encoded := EncodeToString(testTxt)
	assert.Equal(t, encoded, "w3jhxa0")
	back, _ := DecodeString(encoded)
	assert.Equal(t, testTxt, back)
}

func TestBase32Encode(t *testing.T) {
	encoded := Encode(testTxt)
	assert.Equal(t, encoded, []byte("w3jhxa0"))
	back, _ := Decode(encoded)
	assert.Equal(t, testTxt, back)
}

func TestBase32Decode_cppcompat(t *testing.T) {
	decoded, err := base64.StdEncoding.DecodeString("YW55IGNhcm5hbCBwbGVhc3VyZQ==")
	assert.NoError(t, err)
	encoded := EncodeToString(decoded)
	assert.Equal(t, "v9h8jbqqv9exuctvypcxcetpwr6hyeb", encoded)
}
