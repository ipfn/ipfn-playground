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

package digest

import (
	"testing"

	"github.com/crackcomm/sha256-simd"
	"github.com/ipfn/ipfn/pkg/utils/hexutil"
	multihash "github.com/multiformats/go-multihash"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/sha3"
)

func TestFromHex(t *testing.T) {
	digest := FromHex("56e81f171bcc55a6ff8345e692c0f86e5b48e01b996cadc001622fb5e363b421")
	expect := Digest{0x56, 0xe8, 0x1f, 0x17, 0x1b, 0xcc, 0x55, 0xa6, 0xff, 0x83, 0x45, 0xe6, 0x92, 0xc0, 0xf8, 0x6e, 0x5b, 0x48, 0xe0, 0x1b, 0x99, 0x6c, 0xad, 0xc0, 0x01, 0x62, 0x2f, 0xb5, 0xe3, 0x63, 0xb4, 0x21}
	assert.Equal(t, digest, expect)
	// we need == operator forever
	assert.Equal(t, digest == expect, true)
}

func TestSum(t *testing.T) {
	hashed := Sum(sha256.New(), []byte("test"))
	digest := FromHex("9F86D081884C7D659A2FEAA0C55AD015A3BF4F1B2B0B822CD15D6C15B0F00A08")
	assert.Equal(t, digest, hashed)
}

func TestIsEmpty(t *testing.T) {
	assert.Equal(t, false, IsEmpty(FromHex("9F86D081884C7D659A2FEAA0C55AD015A3BF4F1B2B0B822CD15D6C15B0F00A08")))
	assert.Equal(t, true, IsEmpty(FromHex("0000000000000000000000000000000000000000000000000000000000000000")))
	assert.Equal(t, true, IsEmpty(Empty()))
	assert.Equal(t, true, Empty() == Digest{})
}

func TestEqual(t *testing.T) {
	assert.Equal(t, true, FromHex("0000000000000000000000000000000000000000000000000000000000000000") == Empty())
}

func TestHashEncoding(t *testing.T) {
	hashed := Sum(sha256.New(), []byte("test"))
	hash := HashFromDigest(Sha2_256, hashed)
	assert.Equal(t, hash.Digest(), hashed[:])
	assert.Equal(t, hash.Size(), Size)

	h2o, err := DecodeHash(hexutil.FromString("12209f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08"))
	assert.Equal(t, err, nil)
	assert.Equal(t, h2o.Algorithm(), hash.Algorithm())
	assert.Equal(t, h2o.Size(), Size)
	assert.Equal(t, h2o.Digest(), hashed[:])
	assert.Equal(t, h2o.Bytes(), HashFromDigest(Sha2_256, hashed).Bytes())
}

func TestHashFromDigest(t *testing.T) {
	buf := []byte("test")
	algo := Sha2_256
	hashed := Sum(sha256.New(), buf)
	mh, err := multihash.Sum(buf, algo.Code(), 32)
	assert.Equal(t, err, nil)

	hash := HashFromDigest(algo, hashed)
	assert.Equal(t, hash.Size(), len(hashed))
	assert.Equal(t, hash.Digest(), hashed[:])
	assert.Equal(t, hash.Bytes(), []byte(mh))

	dec, err := DecodeHash([]byte(mh))
	assert.Equal(t, err, nil)
	assert.Equal(t, dec.Size(), hash.Size())
	assert.Equal(t, dec.Digest(), hash.Digest())
	assert.Equal(t, dec.Bytes(), []byte(mh))
}

func TestMultihasher(t *testing.T) {
	buf := []byte("test")
	mh, err := multihash.Sum(buf, uint64(Sha3_256), 32)
	assert.Equal(t, err, nil)
	hasher := NewHasher(Sha3_256, sha3.New256())
	hasher.Reset()
	hasher.Write(buf)
	assert.Equal(t, hasher.Sum(nil), []byte(mh))
}
