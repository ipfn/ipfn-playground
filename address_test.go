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

package pubkeyhash

import (
	"fmt"
	"log"
	. "testing"

	"github.com/stretchr/testify/assert"

	keywallet "github.com/ipfn/go-ipfn-key-wallet"
)

var (
	testPass     = []byte("mypass")
	testMnemonic = []byte("wisdom quantum bachelor solution strike harbor push electric gate subject waste worth safe key glance happy notice corn rate occur accuse answer gown census")
)

func TestAddressEthereum(t *T) {
	seed := keywallet.NewSeed(testMnemonic, testPass)

	// Start by getting an extended key instance for the master node.
	// This gives the path:
	//   m
	masterPrivKey, err := keywallet.NewMaster(seed, nil)
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
		acc, _ := keywallet.DerivePath(masterPrivKey, path)
		pub, _ := acc.ECPubKey()
		addr := AddressEthereum(*pub.ToECDSA()).String()
		assert.Equal(t, addr, expected[index])
	}
}

func TestPKHash(t *T) {
	seed := keywallet.NewSeed(testMnemonic, testPass)

	// Start by getting an extended key instance for the master node.
	// This gives the path:
	//   m
	masterPrivKey, err := keywallet.NewMaster(seed, nil)
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
		acc, _ := keywallet.DerivePath(masterPrivKey, path)
		pub, _ := acc.ECPubKey()
		addr, _ := PKHash(pub, 0x89)
		assert.Equal(t, expected[index], addr.String())
	}
}
