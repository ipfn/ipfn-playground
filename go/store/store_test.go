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

package store

import (
	"io/ioutil"
	. "testing"

	"github.com/stretchr/testify/assert"
)

var testBody = []byte("xpub661MyMwAqRbcF2uf1w2zGfELNWjdwasLeZ1vLUXP6dkKEdpzGm4ndyUsKUaH9ok2942o3Ke4Q3wUG9d3NLv8o4enh7g5G38ePJNU5a4mRMG")

func TestMapStore(t *T) {
	key := "test"
	store := NewMapStore()
	assert.Equal(t, store.Put(key, testBody), nil)
	value, err := store.Get(key)
	assert.Equal(t, value, testBody)
	assert.Equal(t, err, nil)
}

func TestJSONStore(t *T) {
	type testStruct struct {
		Example []byte
	}

	key := "test"
	store := NewJSONStore(NewMapStore())
	assert.Equal(t, store.Put(key, testStruct{Example: testBody}), nil)
	var value testStruct
	err := store.Get(key, &value)
	assert.Equal(t, err, nil)
	assert.Equal(t, value.Example, testBody)
}

func TestFileStore(t *T) {
	dir, err := ioutil.TempDir("", "ipfn-kvstoretest")
	assert.Equal(t, err, nil)
	key := "test"
	store := NewFileStore(dir)
	has, err := store.Has(key)
	assert.Equal(t, has, false)
	assert.Equal(t, err, nil)
	assert.Equal(t, store.Put(key, testBody), nil)
	has, err = store.Has(key)
	assert.Equal(t, has, true)
	assert.Equal(t, err, nil)
	value, err := store.Get(key)
	assert.Equal(t, value, testBody)
	assert.Equal(t, err, nil)
	keys, err := store.Keys()
	assert.Equal(t, keys, []string{"test"})
	assert.Equal(t, err, nil)
}
