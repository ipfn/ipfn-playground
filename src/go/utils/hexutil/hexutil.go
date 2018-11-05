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

// Package hexutil implements hex utilities.
package hexutil

import "encoding/hex"

// Encode - Encodes hex.
func Encode(src []byte) []byte {
	res := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(res, src)
	return res
}

// ToString - Encodes hex.
func ToString(src []byte) []byte {
	res := make([]byte, hex.EncodedLen(len(src)))
	hex.Encode(res, []byte(src))
	return res
}

// Decode - Decodes hex.
func Decode(src []byte) []byte {
	res := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(res, src)
	return res
}

// FromString - Decodes hex.
func FromString(src string) []byte {
	res := make([]byte, hex.DecodedLen(len(src)))
	hex.Decode(res, []byte(src))
	return res
}
