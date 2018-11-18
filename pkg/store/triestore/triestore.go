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

package triestore

import (
	"errors"

	"github.com/ipfn/ipfn/pkg/cells"
	"github.com/ipfn/ipfn/pkg/chain"
	"github.com/ipfn/ipfn/pkg/digest"
	"github.com/ipfn/ipfn/pkg/store"
	"github.com/ipfn/ipfn/pkg/trie"
	"github.com/ipfn/ipfn/pkg/utils/flog"
)

// Store - Execution store.
type Store interface {
	store.Simple

	// Commit - Saves store state to database.
	Commit() (cells.CID, error)

	// Clone - Clones store with current state.
	Clone() (Store, error)
}

var logger = flog.MustGetLogger("triestore")

// NewStore - Creates new mutable execution store.
func NewStore(state cells.CID, triedb *trie.Database) (_ Store, err error) {
	// TODO: check for content id and hash algo
	var hash digest.Digest
	if state.Defined() {
		hash = digest.FromBytes(state.Digest())
	}
	t, err := trie.New(hash, triedb)
	if err != nil {
		return
	}
	return &execStore{
		trie:   t,
		triedb: triedb,
		commit: state,
	}, nil
}

type execStore struct {
	trie   *trie.Trie
	triedb *trie.Database
	commit cells.CID
}

func (s *execStore) Get(key store.Key) (val []byte, err error) {
	return s.trie.TryGet(key)
}

func (s *execStore) Set(key store.Key, body []byte) (err error) {
	err = s.trie.TryUpdate(key, body)
	if err != nil {
		return
	}
	s.commit = cells.UndefCID
	return nil
}

func (s *execStore) Commit() (_ cells.CID, err error) {
	commit, err := s.trie.Commit(nil)
	if err != nil {
		return
	}
	err = s.triedb.Commit(commit)
	if err != nil {
		return
	}
	s.commit = cells.NewCIDFromHash(chain.StateTrie, commit[:], chain.StateTriePrefix.MhType)
	return s.commit, nil
}

func (s *execStore) Clone() (Store, error) {
	if !s.commit.Defined() {
		return nil, errors.New("store not committed")
	}
	t, err := trie.New(digest.FromBytes(s.commit.Digest()), s.triedb)
	if err != nil {
		return nil, err
	}
	return &execStore{
		trie:   t,
		triedb: s.triedb,
	}, nil
}
