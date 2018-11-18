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

package chain

import (
	"github.com/ipfn/ipfn/pkg/common/codecs"
)

const (
	// ChainHeader - Content ID of Chain Header Version 1. (79278)
	ChainHeader = 0x51df0
	// ChainSigned - Content ID of Chain Signed Header Version 1. (335344)
	ChainSigned = 0x135ae
	// OperationTrie - Content ID of Cell Trie Version 1. (26156)
	OperationTrie = 0x662c
	// StateTrie - Content ID of Cell Trie Version 1. (27549)
	StateTrie = 0x6b9d
)

// Codecs - Maps the name of a codec to its type.
var Codecs = map[string]uint64{
	"chain-header":   ChainHeader,
	"chain-signed":   ChainSigned,
	"operation-trie": OperationTrie,
	"state-trie":     StateTrie,
}

func init() {
	codecs.Register(Codecs)
}
