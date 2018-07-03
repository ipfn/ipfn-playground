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
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// NewFileStore - Creates a new file-based key-value store.
func NewFileStore(dir string) RawStore {
	return &fileStorage{
		dir: dir,
	}
}

// NewFileJSONStore - Creates a new file-based JSON-encoded key-value store.
func NewFileJSONStore(dir string) EncodedStore {
	return NewJSONStore(NewFileStore(dir))
}

type fileStorage struct {
	dir string
}

func (store *fileStorage) Has(key string) (bool, error) {
	_, err := os.Stat(store.keyPath(key))
	if os.IsNotExist(err) {
		return false, nil
	}
	return err == nil, err

}

func (store *fileStorage) Keys() (keys []string, err error) {
	keys, err = filepath.Glob(store.keyPath("*"))
	if err != nil {
		return
	}
	for i := 0; i < len(keys); i++ {
		keys[i] = strings.TrimSuffix(filepath.Base(keys[i]), ".ipfn.json")
	}
	return
}

func (store *fileStorage) Get(key string) ([]byte, error) {
	return ioutil.ReadFile(store.keyPath(key))
}

func (store *fileStorage) Put(key string, value []byte) error {
	return ioutil.WriteFile(store.keyPath(key), value, 0666)
}

func (store *fileStorage) keyPath(key string) string {
	return filepath.Join(store.dir, fmt.Sprintf("%s.ipfn.json", key))
}
