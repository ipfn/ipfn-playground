// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package sealbox

import (
	"io/ioutil"
	"testing"
)

const (
	veryLightScryptN = 2
	veryLightScryptP = 1
)

// Tests that a json key file can be decrypted and encrypted in multiple rounds.
func TestKeyEncryptDecrypt(t *testing.T) {
	keyjson, err := ioutil.ReadFile("testdata/very-light-scrypt.json")
	if err != nil {
		t.Fatal(err)
	}
	password := ""

	// Do a few rounds of decryption and encryption
	for i := 0; i < 3; i++ {
		// Try a bad password first
		if _, err := DecryptJSON(keyjson, password+"bad"); err == nil {
			t.Errorf("test %d: json key decrypted with bad password", i)
		}
		// Decrypt with the correct password
		body, err := DecryptJSON(keyjson, password)
		if err != nil {
			t.Fatalf("test %d: json key failed to decrypt: %v", i, err)
		}
		// TODO(crackcomm): ...
		// if !bytes.Equal(key.Address, address) {
		// 	t.Errorf("test %d: key address mismatch: have %x, want %x", i, key.Address, address)
		// }
		// Recrypt with a new password and start over
		password += "new data appended"
		if keyjson, err = EncryptJSON(body, []byte(password), veryLightScryptN, veryLightScryptP); err != nil {
			t.Errorf("test %d: failed to recrypt key %v", i, err)
		}
	}
}
