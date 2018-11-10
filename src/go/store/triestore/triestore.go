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

	"github.com/ipfn/ipfn/src/go/chain/dev/contents"
	"github.com/ipfn/ipfn/src/go/store"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/trie"
	"github.com/gogo/protobuf/proto"
	cells "github.com/ipfn/go-ipfn-cells"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
)

// Store - Execution store.
type Store interface {
	store.Bytes

	// Commit - Saves store state to database.
	Commit() (*cells.CID, error)

	// Clone - Clones store with current state.
	Clone() (Store, error)
}

// NewStore - Creates new mutable execution store.
func NewStore(state *cells.CID, triedb *trie.Database) (_ Store, err error) {
	var hash common.Hash
	if state != nil {
		hash = common.BytesToHash(state.Digest())
	}
	t, err := trie.New(hash, triedb)
	if err != nil {
		return
	}
	tb, err := t.TryGet(totalKey)
	if err != nil {
		return
	}
	total, _ := proto.DecodeVarint(tb)
	return &execStore{
		trie:   t,
		triedb: triedb,
		commit: state,
	}, nil
}

var totalKey = []byte("rootchain.total_power")

type execStore struct {
	trie   *trie.Trie
	triedb *trie.Database
	commit *cells.CID
}

func (s *execStore) ReadBytes(key store.Key) (val []byte, err error) {
	return s.trie.TryGet(key.Bytes())
}

func (s *execStore) WriteBytes(key store.Key, body []byte) (err error) {
	prev, err := s.ReadBytes(key)
	if err != nil {
		return
	}
	err = s.trie.TryUpdate(key.Bytes(), body)
	if err != nil {
		return
	}
	s.commit = nil
	s.total += value - prev
	logger.Debugw("Store Update", "key", key, "value", value, "total", s.total, "prev", prev)
	return nil
}

func (s *execStore) Total() uint64 {
	return s.total
}

func (s *execStore) Commit() (_ *cells.CID, err error) {
	err = s.trie.TryUpdate(totalKey, proto.EncodeVarint(s.total))
	if err != nil {
		return
	}
	commit, err := s.trie.Commit(nil)
	if err != nil {
		return
	}
	err = s.triedb.Commit(commit, false)
	if err != nil {
		return
	}
	s.commit = cells.NewCIDFromHash(contents.StateTrie, commit[:], contents.StateTriePrefix.MhType)
	return s.commit, nil
}

func (s *execStore) Clone() (Store, error) {
	if s.commit == nil {
		return nil, errors.New("store not committed")
	}
	t, err := trie.New(common.BytesToHash(s.commit.Digest()), s.triedb)
	if err != nil {
		return nil, err
	}
	return &execStore{
		trie:   t,
		triedb: s.triedb,
	}, nil
}
