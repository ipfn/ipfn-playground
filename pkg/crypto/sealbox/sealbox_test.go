// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2016-2018 The go-ethereum Authors. All Rights Reserved.
//
// This file is part of the IPFN project.
// This file was part of the go-ethereum library.
//
// The IPFN project is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The IPFN project is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package sealbox

// Implements #TST-crypto-sealed

import (
	"encoding/json"
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
		var box SealedBox
		err = json.Unmarshal(keyjson, &box)
		if err != nil {
			return
		}
		if _, err := box.Decrypt(password + "bad"); err == nil {
			t.Errorf("test %d: json key decrypted with bad password", i)
		}
		// Decrypt with the correct password
		body, err := box.Decrypt(password)
		if err != nil {
			t.Fatalf("test %d: json key failed to decrypt: %v", i, err)
		}
		// Recrypt with a new password and start over
		password += "new data appended"
		box, err = Encrypt(body, []byte(password), veryLightScryptN, veryLightScryptP)
		if err != nil {
			t.Errorf("test %d: failed to encrypt %v", i, err)
		}
		keyjson, err = json.Marshal(box)
		if err != nil {
			t.Errorf("test %d: failed to marshal json %v", i, err)
		}
	}
}
