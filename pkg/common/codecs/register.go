// Copyright © 2017-2018 The IPFN Authors. All Rights Reserved.
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

// Package codecs registers codecs in CID codec types.
//
// Example usage:
//
// 	import _ "github.com/ipfn/ipfn/pkg/chain/dev/contents"
//
// // or
//
// 	// optionally
// 	func init() {
// 		contents.RegisterPrefixes(cid.Codecs, cid.CodecToStr)
// 	}
package codecs

import (
	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
)

const (
	// PubkeyV1 - Content ID of Public Kt Version 1. (name = "ipfn-pubkey-v1", id = 24748)
	PubkeyV1 = 0x60ac
	// BinaryCellV1 - Content ID of Binary Cell Version 1. (name = "ipfn-cell-v1", id = 28860)
	BinaryCellV1 = 0x70bc
)

//
// Reserved:
//
//   // BinaryCellV2 - Content ID of Binary Cell Version 2. (45244)
//   BinaryCellV2 = 0xb0bc
//

func init() {
	// this one should be always before
	// subsequent calls to RegisterTarget
	// are using these maps to clone later
	RegisterTarget(Codecs, CodecToStr)
	Register(map[string]uint64{
		"ipfn-pubkey-v1": PubkeyV1,
		"ipfn-cell-v1":   BinaryCellV1,
	})
	// this one is after to ensure it works
	RegisterTarget(cid.Codecs, cid.CodecToStr)
}

type target struct {
	Codecs     map[string]uint64
	CodecToStr map[uint64]string
}

var targets []target

// Codecs - Maps the name of a codec to its type.
var Codecs = make(map[string]uint64)

// CodecToStr - Maps the numeric codec to its name.
var CodecToStr = make(map[uint64]string)

// Register - Registers codecs in remote cids package.
func Register(codecs map[string]uint64) {
	for _, target := range targets {
		for name, codec := range codecs {
			target.Codecs[name] = codec
			target.CodecToStr[codec] = name
		}
	}
}

// RegisterTarget - Registers codecs in remote cids package.
func RegisterTarget(codecs map[string]uint64, codecToStr map[uint64]string) {
	targets = append(targets, target{
		Codecs:     codecs,
		CodecToStr: codecToStr,
	})
	for name, codec := range Codecs {
		codecs[name] = codec
		codecToStr[codec] = name
	}
}
