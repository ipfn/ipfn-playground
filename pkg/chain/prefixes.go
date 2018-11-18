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

package chain

import (
	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
	mh "gx/ipfs/QmerPMzPk1mJVowm8KgmoknWa4yCYvvugMPsgWmDNUvDLW/go-multihash"
)

var (
	// HeaderPrefix - Header CID prefix.
	HeaderPrefix = cid.Prefix{
		Version:  1,
		Codec:    ChainHeader,
		MhType:   mh.SHA2_256,
		MhLength: 32,
	}

	// SignedPrefix - Signed header CID prefix.
	SignedPrefix = cid.Prefix{
		Version:  1,
		Codec:    ChainSigned,
		MhType:   mh.SHA2_256,
		MhLength: 32,
	}

	// OperationTriePrefix - Operation trie CID prefix.
	OperationTriePrefix = cid.Prefix{
		Version:  1,
		Codec:    OperationTrie,
		MhType:   mh.SHA2_256,
		MhLength: 32,
	}

	// StateTriePrefix - State trie CID prefix.
	StateTriePrefix = cid.Prefix{
		Version:  1,
		Codec:    StateTrie,
		MhType:   mh.SHA2_256,
		MhLength: 32,
	}
)
