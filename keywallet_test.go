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

package keywallet

import (
	"log"
	. "testing"

	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/stretchr/testify/assert"
)

var (
	testPass     = []byte("mypass")
	testMnemonic = []byte("wisdom quantum bachelor solution strike harbor push electric gate subject waste worth safe key glance happy notice corn rate occur accuse answer gown census")
)

func TestMnemonicRecover(t *T) {
	seed := NewSeed(testMnemonic, testPass)

	masterKey, _ := NewMaster(seed, nil)
	publicKey, _ := masterKey.Neuter()

	assert.Equal(t, masterKey.String(), "xprv9s21ZrQH143K2YqBuuVyuXHbpUu9Y89VHL6KY67mYJDLMqVqjDkY6BAPUDbVCY16UCs67Goco9cPpBgTXaAQSfhnjJtryzNomPhJevqKwCR")
	assert.Equal(t, publicKey.String(), "xpub661MyMwAqRbcF2uf1w2zGfELNWjdwasLeZ1vLUXP6dkKEdpzGm4ndyUsKUaH9ok2942o3Ke4Q3wUG9d3NLv8o4enh7g5G38ePJNU5a4mRMG")
}

func TestHDKeyChain(t *T) {
	seed := NewSeed(testMnemonic, testPass)

	// Start by getting an extended key instance for the master node.
	// This gives the path:
	//   m
	masterPrivKey, err := NewMaster(seed, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Purpose = bip44
	// /m/44
	fourtyFour, err := masterPrivKey.Child(hdkeychain.HardenedKeyStart + 44)
	if err != nil {
		log.Fatal(err)
	}

	// Cointype = bitcoin
	// /m/44/0
	key, err := fourtyFour.Child(hdkeychain.HardenedKeyStart + 0)
	if err != nil {
		log.Fatal(err)
	}

	// testNetParams.PubKeyHashAddrID = byte(index)
	// Account = 1
	// /m/44/0/1
	account, err := key.Child(hdkeychain.HardenedKeyStart + 1)
	if err != nil {
		log.Fatal(err)
	}

	// Change(0) = external
	// /m/44/0/1/0
	external, err := account.Child(0)
	if err != nil {
		log.Fatal(err)
	}

	expected := []string{
		"1CnMSLqtNdwHYpiFu7rjWjdZ1SsGDPtuZT",
		"18VdTmf9c8qS19AyL7btb6s7sc5ZsJcuNb",
		"1CCxdxNkEUEjc8Aa54oKkUDAFtiMHqhy1v",
	}

	for index := uint32(0); index < 3; index++ {
		addr := indexKey(external, index)
		assert.Equal(t, addr, expected[index])
	}
}

func indexKey(key *hdkeychain.ExtendedKey, index uint32) string {
	// /m/44/0/1/0/{index}
	acc, err := key.Child(index)
	if err != nil {
		log.Fatal(err)
	}

	addr, _ := acc.Address(&parachaincfg.MainNetParamsms)
	return addr.String()
}
