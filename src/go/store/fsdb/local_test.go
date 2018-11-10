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
	"bytes"
	"compress/gzip"
	"context"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"reflect"
	"strings"
	"testing"
	"time"

	"github.com/ipfn/ipfn/src/go/store"
)

const lorem = `Lorem ipsum dolor sit amet,
consectetur adipiscing elit,
sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

Ut enim ad minim veniam,
quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.

Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur.

Excepteur sint occaecat cupidatat non proident,
sunt in culpa qui officia deserunt mollit anim id est laborum.`

func TestReadWriteDelete(t *testing.T) {
	root, err := ioutil.TempDir("", "fsdb_")
	if err != nil {
		t.Fatalf("failed to get tmp dir: %v", err)
	}
	defer os.RemoveAll(root)
	opts := NewDefaultOptions(root).SetUseGzip(false)
	db := Open(opts)

	key := store.Key("foo")
	// Empty
	testDeleteEmpty(t, db, key)
	testReadEmpty(t, db, key)
	// Write
	testWrite(t, db, key, lorem)
	testRead(t, db, key, lorem)
	testRead(t, db, key, lorem)
	// Overwrite
	content := ""
	testWrite(t, db, key, content)
	testRead(t, db, key, content)
	// Delete
	testDelete(t, db, key)
	testReadEmpty(t, db, key)
}

func TestGzip(t *testing.T) {
	root, err := ioutil.TempDir("", "fsdb_")
	if err != nil {
		t.Fatalf("failed to get tmp dir: %v", err)
	}
	defer os.RemoveAll(root)
	opts := NewDefaultOptions(root).SetUseGzip(true)
	db := Open(opts)

	key := store.Key("foo")
	// Empty
	testDeleteEmpty(t, db, key)
	testReadEmpty(t, db, key)
	// Write
	testWrite(t, db, key, lorem)
	testRead(t, db, key, lorem)
	testRead(t, db, key, lorem)
	// Overwrite
	content := ""
	testWrite(t, db, key, content)
	testRead(t, db, key, content)
	// Delete
	testDelete(t, db, key)
	testReadEmpty(t, db, key)
}

func TestChangeCompression(t *testing.T) {
	root, err := ioutil.TempDir("", "fsdb_")
	if err != nil {
		t.Fatalf("failed to get tmp dir: %v", err)
	}
	defer os.RemoveAll(root)
	gzipOpts := NewDefaultOptions(root).SetUseGzip(true)
	gzipDb := Open(gzipOpts)

	key := store.Key("foo")
	testWrite(t, gzipDb, key, lorem)
	testRead(t, gzipDb, key, lorem)

	opts := NewDefaultOptions(root).SetUseGzip(false)
	db := Open(opts)
	testRead(t, db, key, lorem)
	content := ""
	testWrite(t, db, key, content)
	testRead(t, db, key, content)

	testRead(t, gzipDb, key, content)
	testDelete(t, gzipDb, key)
	testReadEmpty(t, gzipDb, key)
}

func TestScan(t *testing.T) {
	ctx := context.Background()
	root, err := ioutil.TempDir("", "fsdb_")
	if err != nil {
		t.Fatalf("failed to get tmp dir: %v", err)
	}
	defer os.RemoveAll(root)
	opts := NewDefaultOptions(root)
	db := Open(opts)

	keys := make(map[string]bool)
	keyFunc := func(ret bool) func(key store.Key) bool {
		return func(key store.Key) bool {
			keys[string(key)] = true
			return ret
		}
	}
	err = db.ScanKeys(ctx, keyFunc(true), store.IgnoreAll)
	if err != nil {
		t.Fatalf("ScanKeys failed: %v", err)
	}
	if len(keys) != 0 {
		t.Errorf("Scan empty db got keys: %+v", keys)
	}

	expectKeys := map[string]bool{
		"foo":    true,
		"bar":    true,
		"foobar": true,
	}
	for key := range expectKeys {
		testWrite(t, db, store.Key(key), "")
	}
	if err := db.ScanKeys(ctx, keyFunc(true), store.StopAll); err != nil {
		t.Fatalf("ScanKeys failed: %v", err)
	}
	if !reflect.DeepEqual(keys, expectKeys) {
		t.Errorf("ScanKeys expected %+v, got %+v", expectKeys, keys)
	}

	keys = make(map[string]bool)
	if err := db.ScanKeys(ctx, keyFunc(false), store.StopAll); err != nil {
		t.Fatalf("ScanKeys failed: %v", err)
	}
	if len(keys) != 1 {
		t.Errorf("Scan should stop after the first key, got: %+v", keys)
	}
}

func TestScanCancel(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping test in short mode")
	}

	sleep := time.Millisecond * 100
	shorter := time.Millisecond * 50

	ctx, cancel := context.WithTimeout(context.Background(), shorter)
	defer cancel()

	root, err := ioutil.TempDir("", "fsdb_")
	if err != nil {
		t.Fatalf("failed to get tmp dir: %v", err)
	}
	defer os.RemoveAll(root)
	opts := NewDefaultOptions(root)
	db := Open(opts)

	keys := []store.Key{
		store.Key("foo"),
		store.Key("bar"),
	}
	for _, key := range keys {
		testWrite(t, db, key, "")
	}

	keyFunc := func(key store.Key) bool {
		time.Sleep(sleep)
		return true
	}
	started := time.Now()
	err = db.ScanKeys(ctx, keyFunc, store.IgnoreAll)
	elapsed := time.Now().Sub(started)
	t.Logf("ScanKeys took %v", elapsed)
	if err != context.DeadlineExceeded {
		t.Errorf("ScanKeys should return %v, got %v", context.DeadlineExceeded, err)
	}
	if elapsed > sleep*time.Duration(len(keys)) {
		t.Errorf("ScanKeys took too long: %v", elapsed)
	}
}

func BenchmarkReadWrite(b *testing.B) {
	root, err := ioutil.TempDir(".", "_fsdb_bench_test_")
	if err != nil {
		b.Fatalf("failed to get tmp dir: %v", err)
	}
	defer os.RemoveAll(root)

	ctx := context.Background()
	keySize := 12
	r := rand.New(rand.NewSource(time.Now().Unix()))

	var benchmarkSizes = map[string]int{
		"1K":   1024,
		"10K":  10 * 1024,
		"1M":   1024 * 1024,
		"10M":  10 * 1024 * 1024,
		"256M": 256 * 1024 * 1024,
	}

	var options = map[string]Options{
		"nocompression": NewDefaultOptions(root).SetUseGzip(false),
		"gzip-min":      NewDefaultOptions(root).SetUseGzip(false).SetGzipLevel(gzip.BestSpeed),
		"gzip-default":  NewDefaultOptions(root).SetUseGzip(false).SetGzipLevel(gzip.DefaultCompression),
		"gzip-max":      NewDefaultOptions(root).SetUseGzip(false).SetGzipLevel(gzip.BestCompression),
	}

	for label, size := range benchmarkSizes {
		b.Run(
			label,
			func(b *testing.B) {
				content := randomBytes(b, r, size)

				for label, opts := range options {
					b.Run(
						label,
						func(b *testing.B) {
							os.RemoveAll(root)
							db := Open(opts)
							keys := make([]store.Key, 0)
							b.Run(
								"write",
								func(b *testing.B) {
									for i := 0; i < b.N; i++ {
										key := store.Key(randomBytes(b, r, keySize))
										keys = append(keys, key)

										err := db.Write(ctx, key, bytes.NewReader(content))
										if err != nil {
											b.Fatalf("Write failed: %v", err)
										}
									}
								},
							)
							b.Run(
								"read",
								func(b *testing.B) {
									for i := 0; i < b.N; i++ {
										key := keys[r.Int31n(int32(len(keys)))]
										reader, err := db.Read(ctx, key)
										if err != nil {
											b.Fatalf("Read failed: %v", err)
										}
										reader.Close()
									}
								},
							)
						},
					)
				}
			},
		)
	}
}

func randomBytes(b *testing.B, r *rand.Rand, size int) []byte {
	b.Helper()

	reader := io.LimitReader(r, int64(size))
	content, err := ioutil.ReadAll(reader)
	if err != nil {
		b.Fatalf("Generate content failed: %v", err)
	}
	if len(content) != size {
		b.Fatalf(
			"Generate content failed, expected %d bytes, got %d",
			size,
			len(content),
		)
	}
	return content
}

func testDeleteEmpty(t *testing.T, db store.Reader, key store.Key) {
	t.Helper()
	if err := db.Delete(context.Background(), key); !store.IsNoSuchKeyError(err) {
		t.Errorf("Expected NoSuchKeyError, got: %v", err)
	}
}

func testDelete(t *testing.T, db store.Reader, key store.Key) {
	t.Helper()
	if err := db.Delete(context.Background(), key); err != nil {
		t.Errorf("Delete failed: %v", err)
	}
}

func testReadEmpty(t *testing.T, db store.Reader, key store.Key) {
	t.Helper()
	if _, err := db.Read(context.Background(), key); !store.IsNoSuchKeyError(err) {
		t.Errorf("Expected NoSuchKeyError, got: %v", err)
	}
}

func testRead(t *testing.T, db store.Reader, key store.Key, expect string) {
	t.Helper()
	reader, err := db.Read(context.Background(), key)
	if err != nil {
		t.Fatalf("Read failed: %v", err)
	}
	defer reader.Close()
	actual, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatalf("Read content failed: %v", err)
	}
	if string(actual) != expect {
		t.Errorf("Read content expected %q, got %q", expect, actual)
	}
}

func testWrite(t *testing.T, db store.Reader, key store.Key, data string) {
	t.Helper()
	err := db.Write(context.Background(), key, strings.NewReader(data))
	if err != nil {
		t.Fatalf("Write failed: %v", err)
	}
}
