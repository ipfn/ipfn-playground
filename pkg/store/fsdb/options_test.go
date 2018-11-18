// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2017 Yuxuan 'fishy' Wang. All Rights Reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the copyright holder nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package fsdb

import (
	"crypto/sha512"
	"encoding/hex"
	"hash"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/ipfn/ipfn/pkg/store"
)

func TestOptions(t *testing.T) {
	root := "/foobar"
	opts := NewDefaultOptions(root)

	t.Run(
		"dirs",
		func(t *testing.T) {
			var expect, actual string

			opts.SetDataDir("data2")
			expect = "/foobar" + PathSeparator + "data2" + PathSeparator
			actual = opts.GetRootDataDir()
			if expect != actual {
				t.Errorf("data dir expected %q, got %q", expect, actual)
			}

			opts.SetDataDir("data2" + PathSeparator)
			expect = "/foobar" + PathSeparator + "data2" + PathSeparator
			actual = opts.GetRootDataDir()
			if expect != actual {
				t.Errorf("data dir expected %q, got %q", expect, actual)
			}

			opts.SetTempDir("tmp")
			expect = "/foobar" + PathSeparator + "tmp" + PathSeparator
			actual = opts.GetRootTempDir()
			if expect != actual {
				t.Errorf("data dir expected %q, got %q", expect, actual)
			}

			opts.SetTempDir("tmp" + PathSeparator)
			expect = "/foobar" + PathSeparator + "tmp" + PathSeparator
			actual = opts.GetRootTempDir()
			if expect != actual {
				t.Errorf("data dir expected %q, got %q", expect, actual)
			}
		},
	)

	t.Run(
		"key-hash",
		func(t *testing.T) {
			key := store.Key("key")
			data := "data"
			opts.SetDataDir(data)
			var expect, actual string

			expect = strings.Join(
				[]string{
					root,
					data,
					"6c",
					"b1",
					"b0",
					"e50d74419e2244eaa7328235f71b48c7e1c33b23f6f9517d14",
					"",
				},
				PathSeparator,
			)
			actual = opts.GetDirForKey(key)
			if actual != expect {
				t.Errorf("hash dir for key %q expected %q, got %q", key, expect, actual)
			}

			opts.SetDirLevel(sha512.Size224 + 10)
			expect = strings.Join(
				[]string{
					root,
					data,
					"6c",
					"b1",
					"b0",
					"e5",
					"0d",
					"74",
					"41",
					"9e",
					"22",
					"44",
					"ea",
					"a7",
					"32",
					"82",
					"35",
					"f7",
					"1b",
					"48",
					"c7",
					"e1",
					"c3",
					"3b",
					"23",
					"f6",
					"f9",
					"51",
					"7d",
					"14",
					"",
				},
				PathSeparator,
			)
			actual = opts.GetDirForKey(key)
			if actual != expect {
				t.Errorf("hash dir for key %v expected %q, got %q", key, expect, actual)
			}

			opts.SetDirLevel(sha512.Size224)
			actual = opts.GetDirForKey(key)
			if actual != expect {
				t.Errorf("hash dir for key %v expected %q, got %q", key, expect, actual)
			}
		},
	)

	t.Run(
		"hash-reentrant",
		func(t *testing.T) {
			if testing.Short() {
				t.Skip("skipping test in short mode")
			}

			calcHash := func(h hash.Hash, key store.Key, sleep time.Duration) string {
				h.Write(key)
				time.Sleep(sleep)
				return hex.EncodeToString(h.Sum([]byte{}))
			}

			keys := []store.Key{
				store.Key("foo"),
				store.Key("bar"),
				store.Key("key"),
			}
			expect := make([]string, len(keys))
			for i, key := range keys {
				expect[i] = calcHash(sha512.New512_224(), key, 0)
			}

			opts.SetHashFunc(sha512.New512_224)
			var wg sync.WaitGroup
			wg.Add(len(keys))
			sleep := time.Millisecond * 100
			for i, key := range keys {
				go func(key store.Key, expect string) {
					defer wg.Done()

					actual := calcHash(opts.GetHashFunc()(), key, sleep)
					if actual != expect {
						t.Errorf("hash for %v expected %q, got %q", key, expect, actual)
					}
				}(key, expect[i])
			}
			wg.Wait()
		},
	)
}
