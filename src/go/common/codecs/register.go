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
// 	import _ "github.com/ipfn/ipfn/src/go/chain/dev/contents"
//
// // or
//
// 	// optionally
// 	func init() {
// 		contents.RegisterPrefixes(cid.Codecs, cid.CodecToStr)
// 	}
package codecs

import (
	cid "gx/ipfs/QmapdYm1b22Frv3k17fqrBYTFRxwiaVJkB299Mfn33edeB/go-cid"
)

const (
	// PubkeyHashV1 - Content ID of Sealed Cell Version 1. (24748)
	PubkeyHashV1 = 0x60ac
	// BinaryCellV1 - Content ID of Binary Cell Version 1. (28860)
	BinaryCellV1 = 0x70bc
	// ChainHeaderV0 - Content ID of Chain Header Version 1. (79278)
	ChainHeaderV0 = 0x51df0
	// ChainSignedV0 - Content ID of Chain Signed Header Version 1. (335344)
	ChainSignedV0 = 0x135ae
	// OperationTrieV0 - Content ID of Cell Trie Version 1. (26156)
	OperationTrieV0 = 0x662c
	// StateTrieV1 - Content ID of Cell Trie Version 1. (27549)
	StateTrieV1 = 0x6b9d
)

//
// Reserved:
//
//   // BinaryCellV2 - Content ID of Binary Cell Version 2. (45244)
//   BinaryCellV2 = 0xb0bc
//

// Codecs - Maps the name of a codec to its type.
var Codecs = map[string]uint64{
	"pubkey-hash-v1":    PubkeyHashV1,
	"cell-binary-v1":    BinaryCellV1,
	"chain-header-v0":   ChainHeaderV0,
	"chain-signed-v0":   ChainSignedV0,
	"operation-trie-v0": OperationTrieV0,
	"state-trie-v1":     StateTrieV1,
}

// CodecToStr - Maps the numeric codec to its name.
var CodecToStr = map[uint64]string{}

// RegisterPrefixes - Registers codecs in remote cids package.
func RegisterPrefixes(codecs map[string]uint64, codecToStr map[uint64]string) {
	for name, codec := range Codecs {
		codecs[name] = codec
		codecToStr[codec] = name
	}
}

func init() {
	for name, codec := range Codecs {
		CodecToStr[codec] = name
	}
	RegisterPrefixes(cid.Codecs, cid.CodecToStr)
}
