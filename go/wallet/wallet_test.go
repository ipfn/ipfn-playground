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

package wallet

import (
	"fmt"
	"log"
	. "testing"

	"github.com/stretchr/testify/assert"

	bip39 "github.com/ipfn/go-bip39"
	"github.com/ipfn/ipfn/go/keypair"
	"github.com/ipfn/ipfn/go/store"
)

var (
	testPass = []byte("mypass")
	testSeed = []byte("wisdom quantum bachelor solution strike harbor push electric gate subject waste worth safe key glance happy notice corn rate occur accuse answer gown census")

	testPubKey  = "fnpubdSC7bxYyo3zTtqFboAAYKAti2pWY5hecgcndjjNaxBdfn8ZHtTK2wMaiHG9j4semnwJTTmGvBm9TToMeY5t9bHYT9gz2MzrzHxynANGWWk1"
	testPrivKey = "fnprvqWuSwMXSGgZYML8TUkm7zJfMVnMVS7PpXUYDRUbfFnN6s29mRmW82FZcks64djUmYZ8t9CLaxc4dFAMvxxdiKGn9iqjp783LLv9c45Z9HpR"
)

func TestWalletStore(t *T) {
	wallet := New(store.NewJSONStore(store.NewMapStore()))

	err := wallet.ImportKey("test", testSeed, testPass)
	assert.Equal(t, err, nil)

	has, err := wallet.KeyExists("test")
	assert.Equal(t, has, true)
	assert.Equal(t, err, nil)

	key, err := wallet.MasterKey("test", testPass)
	assert.Equal(t, key.PrivateString(), testPrivKey)
	assert.Equal(t, err, nil)
}

func TestWalletCreate(t *T) {
	wallet := New(store.NewJSONStore(store.NewMapStore()))

	seed, err := wallet.CreateSeed("test", testPass)
	assert.Equal(t, err, nil)

	k, err := wallet.ExportKey("test")
	assert.Equal(t, k.Name, "test")
	assert.Equal(t, k.SeedType, keypair.Mnemonic)
	assert.Equal(t, err, nil)
	res, err := k.Decrypt(testPass)
	assert.Equal(t, err, nil)
	assert.Equal(t, res, seed)

	key, err := wallet.MasterKey("test", testPass)
	assert.Equal(t, key != nil, true)
	assert.Equal(t, err, nil)

	names, err := wallet.KeyNames()
	assert.Equal(t, names, []string{"test"})
	assert.Equal(t, err, nil)
}

func TestSeedRecover(t *T) {
	seed := bip39.NewSeed(testSeed, testPass)

	masterKey, _ := keypair.NewMaster(seed)
	publicKey, _ := masterKey.Neuter()

	assert.Equal(t, masterKey.PrivateString(), testPrivKey)
	assert.Equal(t, publicKey.String(), testPubKey)
}

func TestDerive(t *T) {
	seed := bip39.NewSeed(testSeed, testPass)

	// Start by getting an extended key instance for the master node.
	// This gives the path:
	//   m
	masterPrivKey, err := keypair.NewMaster(seed)
	if err != nil {
		log.Fatal(err)
	}

	expected := []struct {
		addr string
		pub  string
		priv string
	}{
		{
			addr: "zFNScYMGyjcr82aXUxFD4rdS34AY4vAgpZMdmBjC2YLFaytJkyfP",
			pub:  "fnpubdbsYn2EKDhRnujd238maZUqm4yoXLT2fNLvUW8GMYHT8sL3LkSeUCQq2kokzujZ2gDGjBHNMKyY5WvzKdt3yseZXwpYwnTMLMRuSJwobbcc",
			priv: "fnprvqgat7RCmhKzsNEVsijNAEccQXweUgrmsDCg4BsVRqtBZxDdpHkqZHJowEQf34jDVX99cieC6gvzRthk7euuKXNZ9qmjUajp16XtHk4x7BTo",
		},
		{
			addr: "zFNScYMH2yUE5MTaRM6712E6gP65LgFGXkQ5B7ftMhB33tEVGyv1",
			pub:  "fnpubdbsYn2EKDhRnwruhib2BRww9ZTsnkbncduEM9NYxo31JZSMw3ji9aYHcGud7rE4JwBqi27gP7xHDdF98Lnz9tYqDoLFjsYSe8DvjcwwLBeF",
			priv: "fnprvqgat7RCmhKzsQMnZQBcm75ho2Rik71XpUkyvq7n36djjeKxQb3uEfSGWkXJWm2kcmcapqoyVU6xn5N1C6oZKYJsPWkvhJemYCq4KfMXCcjc",
		},
		{
			addr: "zFNScYMHBo7TSinu4SaG2L49Hy1W4AHVgFyUKwJ4qVRCigSzhHMu",
			pub:  "fnpubdbsYn2EKDhRnzK6xmVYoew94wqkLPphtkG8fPu7rHjisM9Q5ghXp7ATNuZf1qqVAamh3YuKWKpDkgoSsCjxxf2QNUDg1Th8HtE1zxMbCwCT",
			priv: "fnprvqgat7RCmhKzsSoypT69PL4uiQobHkET6b7tF5eLvbLTJS2zZE1iuC4SHP86XeRYPw8SqgqimBNYZj1VHZy2kCEKyCo5S44QivATiPRwruha",
		},
	}

	for index := uint32(0); index < 3; index++ {
		path := fmt.Sprintf("m/44'/60'/7'/1/%d", index)
		acc, _ := masterPrivKey.DerivePath(path)
		c, _ := acc.CID()
		assert.Equal(t, c.String(), expected[index].addr)
		assert.Equal(t, acc.String(), expected[index].pub)
		assert.Equal(t, acc.PrivateString(), expected[index].priv)
	}
}
