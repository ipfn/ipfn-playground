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

package cellcid

import (
	"fmt"

	cid "github.com/ipfs/go-cid"
	multihash "github.com/multiformats/go-multihash"

	cell "github.com/ipfn/ipfn/go/cell"
	codecs "github.com/ipfn/ipfn/go/codecs"
)

// MhType - CID multihash type.
const MhType = multihash.SHA3_256

// stdCodec - Standard codec used for CIDs.
// Multicodec ID of Protocol Buffers Cell Version 1. (28860)
// Definition: /ipfs/QmeX5H9x2qNdGC1R5uhyX2HuG5izxR2SGi71jSWyEQjV6Q
// Requires registered codec, current implementation:
// https://github.com/ipfn/ipfn/go/cellpb
const stdCodec = "cell-pb-v1"

// CID - Creates CID from cell using standard codec.
func CID(c cell.Cell) (*cid.Cid, error) {
	return EncodeByName(c, stdCodec)
}

// Encode - Creates CID from cell using codec by ID.
func Encode(c cell.Cell, id uint64) (*cid.Cid, error) {
	codec, ok := codecs.CodecByPrefix(id)
	if !ok {
		return nil, fmt.Errorf("codec by ID %x not found", codec)
	}
	body, err := codec.Encode(c)
	if err != nil {
		return nil, err
	}
	return Encoded(body, id)
}

// EncodeByName - Creates CID from cell using codec by name.
func EncodeByName(c cell.Cell, name string) (*cid.Cid, error) {
	codec, ok := codecs.CodecByName(name)
	if !ok {
		return nil, fmt.Errorf("codec by name %s not found", name)
	}
	body, err := codec.Encode(c)
	if err != nil {
		return nil, err
	}
	return EncodedByName(body, name)
}

// Encoded - Creates CID from encoded body and codec ID.
func Encoded(body []byte, codec uint64) (*cid.Cid, error) {
	prefix := cid.Prefix{
		Version:  1,
		Codec:    codec,
		MhType:   MhType,
		MhLength: -1, // default length
	}
	return prefix.Sum(body)
}

// EncodedByName - Creates CID from encoded body and codec ID.
func EncodedByName(body []byte, name string) (*cid.Cid, error) {
	codec, ok := codecs.CodecsIDs[name]
	if !ok {
		return nil, fmt.Errorf("codec by name %s not found", name)
	}
	return Encoded(body, codec)
}
