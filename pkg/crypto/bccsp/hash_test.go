// Copyright © 2018 The IPFN Developers. All Rights Reserved.
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

package bccsp

import (
	"encoding/binary"
	"fmt"
	"log"
	"testing"

	"github.com/ipfn/ipfn/pkg/utils/hashutil"
	"github.com/ipfn/ipfn/pkg/utils/hexutil"
	multihash "github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
)

func TestHashEncoding(t *testing.T) {
	hashed := hashutil.SumSha256([]byte("test"))
	hash := HashDigest(Sha2_256, hashed)

	start := make([]byte, 2*binary.MaxVarintLen64)
	spot := start
	n := binary.PutUvarint(spot, hash.Algorithm().Code())
	spot = start[n:]
	n += binary.PutUvarint(spot, uint64(hash.Size()))

	fmt.Printf("Length: %d\n", n)
	res := append(start[:n], hash.Digest()...)

	// padding tests
	assert.Equal(t, res[n:], hash.Digest())

	log.Printf("hash: %x", res)

	h2o, err := DecodeHash(hexutil.FromString("12209f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"))
	assert.Equal(t, err, nil)
	assert.Equal(t, h2o.Algorithm(), hash.Algorithm())
	assert.Equal(t, h2o.Size(), len(hashed))
	assert.Equal(t, h2o.Digest(), hashed)
	assert.Equal(t, h2o.Bytes(), HashDigest(Sha2_256, hashed).Bytes())
}

func TestHashDigest(t *testing.T) {
	buf := []byte("test")
	algo := Sha2_256
	hashed := hashutil.SumSha256(buf)
	mh, err := multihash.Sum(buf, algo.Code(), 32)
	assert.Equal(t, err, nil)

	hash := HashDigest(algo, hashed)
	assert.Equal(t, hash.Size(), len(hashed))
	assert.Equal(t, hash.Digest(), hashed)
	assert.Equal(t, hash.Bytes(), []byte(mh))

	dec, err := DecodeHash([]byte(mh))
	assert.Equal(t, err, nil)
	assert.Equal(t, dec.Size(), hash.Size())
	assert.Equal(t, dec.Digest(), hash.Digest())
	assert.Equal(t, dec.Bytes(), []byte(mh))
}
