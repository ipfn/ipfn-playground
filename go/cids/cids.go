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

package cids

import (
	cid "github.com/ipfs/go-cid"
)

const (
	// PubkeyHash - Content ID of Sealed Cell Version 1. (24748)
	PubkeyHash = 0x60ac
	// BinaryCell - Content ID of Binary Cell Version 1. (28860)
	BinaryCell = 0x70bc
	// ChainHeader - Content ID of Chain Header Version 1. (79278)
	ChainHeader = 0x51df0
	// ChainSigned - Content ID of Chain Signed Header Version 1. (335344)
	ChainSigned = 0x135ae
	// ChainStateTrie - Content ID of Chain State Trie Version 1. (27549)
	ChainStateTrie = 0x6b9d
)

// Codecs - Maps the name of a codec to its type.
var Codecs = map[string]uint64{
	"pubkey-hash":      PubkeyHash,
	"cell-binary":      BinaryCell,
	"chain-header":     ChainHeader,
	"chain-signed":     ChainSigned,
	"chain-state-trie": ChainStateTrie,
}

// CodecToStr - Maps the numeric codec to its name.
var CodecToStr = map[uint64]string{}

// Register - Registers codecs in remote cids package.
func Register(codecs map[string]uint64, codecToStr map[uint64]string) {
	for name, codec := range Codecs {
		codecs[name] = codec
		codecToStr[codec] = name
	}
}

// Register codecs in `go-cid` package to inject IPFN codec types into IPFS.
// `import _ "github.com/ipfn/ipfn/go/cids"`.
func init() {
	for name, codec := range Codecs {
		CodecToStr[codec] = name
	}
	Register(cid.Codecs, cid.CodecToStr)
}
