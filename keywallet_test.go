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
	"fmt"
	"log"
	. "testing"

	"github.com/stretchr/testify/assert"

	"github.com/btcsuite/btcd/chaincfg"

	bip39 "github.com/ipfn/go-bip39"
	pubkeyhash "github.com/ipfn/go-ipfn-pubkey-hash"
)

var (
	testPass     = []byte("mypass")
	testMnemonic = []byte("wisdom quantum bachelor solution strike harbor push electric gate subject waste worth safe key glance happy notice corn rate occur accuse answer gown census")
)

func TestMnemonicRecover(t *T) {
	seed := bip39.NewSeed(testMnemonic, testPass)

	masterKey, _ := NewMaster(seed)
	publicKey, _ := masterKey.Neuter()

	assert.Equal(t, masterKey.String(), "xprv9s21ZrQH143K2YqBuuVyuXHbpUu9Y89VHL6KY67mYJDLMqVqjDkY6BAPUDbVCY16UCs67Goco9cPpBgTXaAQSfhnjJtryzNomPhJevqKwCR")
	assert.Equal(t, publicKey.String(), "xpub661MyMwAqRbcF2uf1w2zGfELNWjdwasLeZ1vLUXP6dkKEdpzGm4ndyUsKUaH9ok2942o3Ke4Q3wUG9d3NLv8o4enh7g5G38ePJNU5a4mRMG")
}

func TestDerivePath(t *T) {
	seed := bip39.NewSeed(testMnemonic, testPass)

	// Start by getting an extended key instance for the master node.
	// This gives the path:
	//   m
	masterPrivKey, err := NewMaster(seed)
	if err != nil {
		log.Fatal(err)
	}

	expected := []string{
		"1Nbfs1gN6FWuNxjoFmVwV4C37ZmMjfDDTa",
		"1JyCNadqv2GfVv6RaJba7vpCFuZMC2Px2J",
		"16iuX6EHAhCLMTubHPrqaSJQLrNf2jSSHt",
	}

	for index := uint32(0); index < 3; index++ {
		path := fmt.Sprintf("m/44'/0'/123'/123/%d", index)
		acc, _ := DerivePath(masterPrivKey, path)
		addr, _ := acc.Address(&chaincfg.MainNetParams)
		assert.Equal(t, addr.String(), expected[index])
	}

	for index := uint32(0); index < 3; index++ {
		// the `m/` prefix is required always
		path := fmt.Sprintf("44/0/123/123/%d", index)
		_, err := DerivePath(masterPrivKey, path)
		assert.Error(t, err)
	}
}

func TestDeriveEth(t *T) {
	seed := bip39.NewSeed(testMnemonic, testPass)

	// Start by getting an extended key instance for the master node.
	// This gives the path:
	//   m
	masterPrivKey, err := NewMaster(seed)
	if err != nil {
		log.Fatal(err)
	}

	expected := []string{
		"0x72D6EC7a17C58693E6d098d714b77F94CC20C2Ba",
		"0xE92c4BaD8C6d52b9E2759e3824f08E624a7c64dA",
		"0x3A734aEb1E149618c7B7e230D9f78862F1dDEfAC",
	}

	for index := uint32(0); index < 3; index++ {
		path := fmt.Sprintf("m/44'/60'/7'/1/%d", index)
		acc, _ := DerivePath(masterPrivKey, path)
		pub, _ := acc.ECPubKey()
		addr := pubkeyhash.AddressEthereum(*pub.ToECDSA())
		assert.Equal(t, addr.String(), expected[index])
	}
}

func TestPKHash(t *T) {
	seed := bip39.NewSeed(testMnemonic, testPass)

	// Start by getting an extended key instance for the master node.
	// This gives the path:
	//   m
	masterPrivKey, err := NewMaster(seed)
	if err != nil {
		log.Fatal(err)
	}

	expected := []string{
		"xHRa3RzLLaq4UKy8ZrDKNX9KuvePy5np8U",
		"xE9xoRVkuPZ9ry9W52P5pfDckm5PkTs7Bi",
		"xES96MKuK3GPFaW2LyWgqd4Enk6BkhUsY8",
	}

	for index := uint32(0); index < 3; index++ {
		path := fmt.Sprintf("m/44'/2'/7'/1/%d", index)
		acc, _ := DerivePath(masterPrivKey, path)
		pub, _ := acc.ECPubKey()
		addr, _ := pubkeyhash.PKHash(pub, 0x89)
		assert.Equal(t, expected[index], addr.String())
	}
}
