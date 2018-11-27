package main

import (
	"encoding/binary"
	"fmt"
	"net/http"

	"github.com/cespare/xxhash"
	"github.com/ipfn/ipfn/pkg/digest"
	"github.com/ipfn/ipfn/pkg/trie"
	"github.com/ipfn/ipfn/pkg/trie/ethdb"

	// profiling
	_ "net/http/pprof"
)

func main() {

	db := trie.NewDatabase(ethdb.NewMemDatabase())
	state, err := trie.New(digest.Empty(), db)
	if err != nil {
		panic(err)
	}

	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			panic(err)
		}
	}()

	var index uint64
	for {
		k := hashBytes(fmt.Sprintf("test1%d", index))
		v := hashBytes(fmt.Sprintf("test2%d", index))
		err = state.TryUpdate(k[:], v[:])
		if err != nil {
			panic(err)
		}
		index++
	}
}

func hashBytes(s string) (res [binary.MaxVarintLen64]byte) {
	n := xxhash.Sum64String(s)
	binary.PutUvarint(res[:], n)
	return

}
