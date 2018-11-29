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

package swcp

import (
	"crypto/rand"
	"crypto/sha512"
	"io"
	"testing"

	"gx/ipfs/QmW7VUmSvhvSGbYbdsh7uRjhGmsYkc9fL8aJ5CorxxrU5N/go-crypto/ed25519"

	"github.com/agl/ed25519/edwards25519"
	"github.com/ipfn/ipfn/pkg/crypto/entropy"
)

func BenchmarkED25519Sign_32(b *testing.B) {
	seed, _ := entropy.New(ed25519.SeedSize)
	priv := ed25519.NewKeyFromSeed(seed)
	msg, _ := entropy.New(32)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = ed25519.Sign(priv, msg)
	}
}

func BenchmarkED25519Verify_32(b *testing.B) {
	seed, _ := entropy.New(ed25519.SeedSize)
	priv := ed25519.NewKeyFromSeed(seed)
	var pubkey [ed25519.PublicKeySize]byte
	copy(pubkey[:], priv[32:])
	msg, _ := entropy.New(32)
	sig := ed25519.Sign(priv, msg)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if !ed25519.Verify(pubkey[:], msg, sig) {
			b.Fatal("invalid")
		}
	}
}

func generateKey() (_ [32]byte, _ [32]byte, err error) {
	var seed [32]byte
	_, err = io.ReadFull(rand.Reader, seed[:])
	if err != nil {
		return
	}
	return newKeyFromSeed(seed[:])
}

func newKeyFromSeed(seed []byte) (_ [32]byte, _ [32]byte, err error) {
	resultingKey := ed25519.NewKeyFromSeed(seed)
	var privateKey [32]byte
	var publicKey [32]byte
	copy(privateKey[:], resultingKey[:32])
	copy(publicKey[:], resultingKey[32:])
	return publicKey, privateKey, nil
}

func computePublic(privateKey [32]byte) (publicKey [32]byte) {
	digest := sha512.Sum512(privateKey[:])
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64
	var A edwards25519.ExtendedGroupElement
	var hBytes [32]byte
	copy(hBytes[:], digest[:])
	edwards25519.GeScalarMultBase(&A, &hBytes)
	A.ToBytes(&publicKey)
	return
}
