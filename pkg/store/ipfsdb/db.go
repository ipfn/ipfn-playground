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

package ipfsdb

import (
	"github.com/ipfn/ipfn/pkg/trie/ethdb"

	cid "gx/ipfs/QmR8BauakNcBa3RbE4nbQu76PDiJgoQgz8AJdhJuiU4TAw/go-cid"
)

// Wrap - Wraps database with IPFS storage.
func Wrap(prefix cid.Prefix, db ethdb.Database) ethdb.Database {
	return WrapURL(prefix, db, "http://localhost:5001")
}

// WrapURL - Wraps database with IPFS storage.
func WrapURL(prefix cid.Prefix, db ethdb.Database, url string) ethdb.Database {
	client := newClient(prefix, url)
	return &wrapDB{Database: db, client: client}
}

type wrapDB struct {
	ethdb.Database
	client *wrapClient
}

func (db *wrapDB) Get(key []byte) (value []byte, err error) {
	if v, err := db.Database.Get(key); err == nil {
		return v, nil
	}
	return db.client.Get(key)
}

func (db *wrapDB) Put(key []byte, value []byte) error {
	if err := db.Database.Put(key, value); err != nil {
		return err
	}
	return db.client.Put(value)
}

func (db *wrapDB) NewBatch() ethdb.Batch {
	return &wrapBatch{Batch: db.Database.NewBatch(), client: db.client}
}
