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

package cells

import (
	. "testing"

	mh "gx/ipfs/QmPnFwZ2JXKnXgMw8CdBPxn7FWh6LLdjUjxV1fKHuJnkr8/go-multihash"
	cid "gx/ipfs/QmapdYm1b22Frv3k17fqrBYTFRxwiaVJkB299Mfn33edeB/go-cid"
)

func BenchmarkCidBytes_V0(b *B) {
	c, _ := SumCID(cid.Prefix{
		Version:  0,
		Codec:    8438,
		MhType:   mh.KECCAK_256,
		MhLength: 32,
	}, []byte("test"))
	for i := 0; i < b.N; i++ {
		_ = c.Bytes()
	}
}

func BenchmarkCidBytes_V1(b *B) {
	c, _ := SumCID(cid.Prefix{
		Version:  1,
		Codec:    8438,
		MhType:   mh.KECCAK_256,
		MhLength: 32,
	}, []byte("test"))
	for i := 0; i < b.N; i++ {
		_ = c.Bytes()
	}
}
