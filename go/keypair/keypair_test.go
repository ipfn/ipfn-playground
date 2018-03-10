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

package keypair

import (
	"crypto/rand"
	"encoding/json"
	"fmt"

	. "testing"
)

func ExampleNew() {
	keys, err := New(RSA)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Public ID: %s", keys.ID)
}

func TestNew(t *T) {
	keys, err := New(RSA)
	if err != nil {
		t.Error(err)
	}
	if len(keys.ID) == 0 {
		t.Fail()
	}
}

func TestJSON(t *T) {
	keys, err := New(RSA)
	if err != nil {
		t.Error(err)
	}
	body, err := json.Marshal(keys)
	if err != nil {
		t.Error(err)
	}
	var res *KeyPair
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Error(err)
	}
	if keys.ID.Pretty() != res.ID.Pretty() {
		t.Fail()
	}
	// TODO we can test more
}

func BenchmarkRSA_124_New(b *B) {
	for i := 0; i < b.N; i++ {
		NewCustomUnsafe(RSA, 124)
	}
}

func BenchmarkRSA_256_New(b *B) {
	for i := 0; i < b.N; i++ {
		NewCustomUnsafe(RSA, 256)
	}
}

func BenchmarkRSA_512_New(b *B) {
	for i := 0; i < b.N; i++ {
		NewCustomUnsafe(RSA, 512)
	}
}

func BenchmarkRSA_1024_New(b *B) {
	for i := 0; i < b.N; i++ {
		NewCustomUnsafe(RSA, 1024)
	}
}

func BenchmarkRSA_2048_New(b *B) {
	for i := 0; i < b.N; i++ {
		New(RSA)
	}
}

func BenchmarkEd25519_New(b *B) {
	for i := 0; i < b.N; i++ {
		New(Ed25519)
	}
}

func BenchmarkSecp256k1_New(b *B) {
	for i := 0; i < b.N; i++ {
		New(Secp256k1)
	}
}

func BenchmarkRSA_512_Sign(b *B) {
	key := NewCustomUnsafe(RSA, 512)
	for i := 0; i < b.N; i++ {
		key.Sign(randBytes2048)
	}
}

func BenchmarkRSA_1024_Sign(b *B) {
	key := NewCustomUnsafe(RSA, 1024)
	for i := 0; i < b.N; i++ {
		key.Sign(randBytes2048)
	}
}

func BenchmarkRSA_2048_Sign(b *B) {
	key, _ := New(RSA)
	for i := 0; i < b.N; i++ {
		key.Sign(randBytes2048)
	}
}

func BenchmarkEd25519_Sign(b *B) {
	key, _ := New(Ed25519)
	for i := 0; i < b.N; i++ {
		key.Sign(randBytes2048)
	}
}

func BenchmarkSecp256k1_Sign(b *B) {
	key, _ := New(Secp256k1)
	for i := 0; i < b.N; i++ {
		key.Sign(randBytes2048)
	}
}

var randBytes2048 = make([]byte, 2048)

func init() {
	_, err := rand.Read(randBytes2048)
	if err != nil {
		panic(err)
	}
}
