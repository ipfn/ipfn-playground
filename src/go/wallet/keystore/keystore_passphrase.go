// Copyright 2014 The go-ethereum Authors
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

/*

This key store behaves as KeyStorePlain with the difference that
the private key is encrypted and on disk uses another JSON encoding.

The crypto is documented at https://github.com/ethereum/wiki/wiki/Web3-Secret-Storage-Definition

*/

package keystore

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/ipfn/ipfn/src/go/crypto"
	"github.com/ipfn/ipfn/src/go/crypto/sealbox"
	"github.com/ipfn/ipfn/src/go/utils/fsutil"
	"github.com/ipfn/ipfn/src/go/utils/mathutil"
	"github.com/ipfn/ipfn/src/go/wallet/accounts"
)

var templateError = `An error was encountered when saving and verifying the keystore file.
This indicates that the keystore is corrupted.
The corrupted file is stored at:
  %v

Please file a ticket at: https://github.com/ipfn/ipfn/issues

The error was: %q`

type keyStorePassphrase struct {
	keysDirPath string
	scryptN     int
	scryptP     int
	// skipKeyFileVerification disables the security-feature which does
	// reads and decrypts any newly created keyfiles. This should be 'false' in all
	// cases except tests -- setting this to 'true' is not recommended.
	skipKeyFileVerification bool
}

func (ks keyStorePassphrase) GetKey(addr accounts.Address, filename string, pwd string) (*Key, error) {
	// Load the key from the keystore and decrypt its contents
	keyjson, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	key, err := DecryptKey(keyjson, pwd)
	if err != nil {
		return nil, err
	}
	// Make sure we're really operating on the requested key (no swap attacks)
	if key.Address != addr {
		return nil, fmt.Errorf("key content mismatch: have account %x, want %x", key.Address, addr)
	}
	return key, nil
}

// StoreKey generates a key, encrypts with 'auth' and stores in the given directory
func StoreKey(dir, auth string, scryptN, scryptP int) (accounts.Address, error) {
	_, a, err := storeNewKey(&keyStorePassphrase{dir, scryptN, scryptP, false}, rand.Reader, auth)
	return a.Address, err
}

func (ks keyStorePassphrase) StoreKey(filename string, key *Key, auth string) error {
	keyjson, err := EncryptKey(key, auth, ks.scryptN, ks.scryptP)
	if err != nil {
		return err
	}
	// Write into temporary file
	tmpName, err := fsutil.WriteTemporaryKeyFile(filename, keyjson)
	if err != nil {
		return err
	}
	if !ks.skipKeyFileVerification {
		// Verify that we can decrypt the file with the given password.
		_, err = ks.GetKey(key.Address, tmpName, auth)
		if err != nil {
			return fmt.Errorf(templateError, tmpName, err)
		}
	}
	return os.Rename(tmpName, filename)
}

func (ks keyStorePassphrase) JoinPath(filename string) string {
	if filepath.IsAbs(filename) {
		return filename
	}
	return filepath.Join(ks.keysDirPath, filename)
}

// EncryptKey encrypts a key using the specified scrypt parameters into a json
// blob that can be decrypted later on.
func EncryptKey(key *Key, pwd string, scryptN, scryptP int) ([]byte, error) {
	body := mathutil.PaddedBigBytes(key.PrivateKey.D, 32)
	box, err := sealbox.Encrypt(body, []byte(pwd), scryptN, scryptP)
	if err != nil {
		return nil, err
	}
	return json.Marshal(box)
}

// DecryptKey decrypts a key from a json blob, returning the private key itself.
func DecryptKey(body []byte, pwd string) (*Key, error) {
	keyBody, err := sealbox.DecryptJSON(body, pwd)
	if err != nil {
		return nil, err
	}
	key := crypto.ToECDSAUnsafe(keyBody)
	return &Key{
		Address:    accounts.PubkeyToAddress(key.PublicKey),
		PrivateKey: crypto.ToECDSAUnsafe(keyBody),
	}, nil
}
