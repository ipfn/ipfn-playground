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

package codecs

import cid "github.com/ipfs/go-cid"

// Codecs - Maps the prefix of a codec to a codec.
var Codecs = map[uint64]Codec{}

// CodecsByName - Maps the name of a codec to a codec.
var CodecsByName = map[string]Codec{}

// CodecsIDs - Maps the name of a codec to its type.
var CodecsIDs = map[string]uint64{}

// CodecToStr - Maps the numeric codec to its name.
var CodecToStr = map[uint64]string{}

// Register - Registers codec in cell codec registry and `go-cid` codec types.
func Register(name string, id uint64, codec Codec) {
	// Register codecs
	Codecs[id] = codec
	CodecsByName[name] = codec
	CodecsIDs[name] = id
	CodecToStr[id] = name
	// Content ID package registration
	cid.Codecs[name] = id
	cid.CodecToStr[id] = name
}

// CodecByPrefix - Returns codec by ID prefix.
func CodecByPrefix(id uint64) (codec Codec, ok bool) {
	codec, ok = Codecs[id]
	return
}

// CodecByName - Returns codec by name.
func CodecByName(name string) (codec Codec, ok bool) {
	codec, ok = CodecsByName[name]
	return
}
