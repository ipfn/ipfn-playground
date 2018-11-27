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

package hashutil

import (
	"testing"

	"github.com/cespare/xxhash"
)

func BenchmarkXXHash64(b *testing.B) {
	input := "http://cyan4973.github.io/xxHash/"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = xxhash.Sum64String(input)
	}
}

func BenchmarkFNV64(b *testing.B) {
	input := "http://cyan4973.github.io/xxHash/"
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = fnv64aString(input)
	}
}

const (
	// offset64 FNVa offset basis. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	offset64 = 14695981039346656037
	// prime64 FNVa prime value. See https://en.wikipedia.org/wiki/Fowler–Noll–Vo_hash_function#FNV-1a_hash
	prime64 = 1099511628211
)

func fnv64aString(key string) uint64 {
	var hash uint64 = offset64
	for i := 0; i < len(key); i++ {
		hash ^= uint64(key[i])
		hash *= prime64
	}

	return hash
}
