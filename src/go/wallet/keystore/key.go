// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2017-2018 The go-ethereum Authors. All Rights Reserved.
//
// This file is part of the IPFN Project.
// This file was part of the go-ethereum library.
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

package keystore

import (
	"crypto/ecdsa"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/btcsuite/btcd/btcec"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ipfn/ipfn/src/go/wallet/accounts"
)

// Key - Account key.
type Key struct {
	// to simplify lookups we also store the address
	Address accounts.Address
	// we only store privkey as pubkey/address can be derived from it
	// privkey in this struct is always in plaintext
	PrivateKey *ecdsa.PrivateKey
}

type keyStore interface {
	// Loads and decrypts the key from disk.
	GetKey(addr accounts.Address, filename string, auth string) (*Key, error)
	// Writes and encrypts the key.
	StoreKey(filename string, k *Key, auth string) error
	// Joins filename with the key directory unless it is already absolute.
	JoinPath(filename string) string
}

type keyJSON struct {
	Address    string `json:"address"`
	PrivateKey string `json:"privatekey"`
}

func (k *Key) MarshalJSON() ([]byte, error) {
	return json.Marshal(keyJSON{
		Address:    hex.EncodeToString(k.Address[:]),
		PrivateKey: hex.EncodeToString(crypto.FromECDSA(k.PrivateKey)),
	})
}

func (k *Key) UnmarshalJSON(body []byte) (err error) {
	keyJSON := new(keyJSON)
	err = json.Unmarshal(body, &keyJSON)
	if err != nil {
		return
	}
	addr, err := hex.DecodeString(keyJSON.Address)
	if err != nil {
		return
	}
	k.Address = accounts.BytesToAddress(addr)
	k.PrivateKey, err = crypto.HexToECDSA(keyJSON.PrivateKey)
	return
}

func newKeyFromECDSA(privateKeyECDSA *ecdsa.PrivateKey) *Key {
	return &Key{
		Address:    accounts.PubkeyToAddress(privateKeyECDSA.PublicKey),
		PrivateKey: privateKeyECDSA,
	}
}

func newKey(rand io.Reader) (*Key, error) {
	privateKeyECDSA, err := ecdsa.GenerateKey(btcec.S256(), rand)
	if err != nil {
		return nil, err
	}
	return newKeyFromECDSA(privateKeyECDSA), nil
}

func storeNewKey(ks keyStore, rand io.Reader, auth string) (*Key, accounts.Account, error) {
	key, err := newKey(rand)
	if err != nil {
		return nil, accounts.Account{}, err
	}
	a := accounts.Account{Address: key.Address, URL: accounts.URL{Scheme: KeyStoreScheme, Path: ks.JoinPath(keyFileName(key.Address))}}
	if err := ks.StoreKey(a.URL.Path, key, auth); err != nil {
		zeroKey(key.PrivateKey)
		return nil, a, err
	}
	return key, a, err
}

// keyFileName implements the naming convention for keyfiles:
// UTC--<created_at UTC ISO8601>-<address hex>
func keyFileName(keyAddr accounts.Address) string {
	ts := time.Now().UTC()
	return fmt.Sprintf("UTC--%s--%s", toISO8601(ts), hex.EncodeToString(keyAddr[:]))
}

func toISO8601(t time.Time) string {
	var tz string
	name, offset := t.Zone()
	if name == "UTC" {
		tz = "Z"
	} else {
		tz = fmt.Sprintf("%03d00", offset/3600)
	}
	return fmt.Sprintf("%04d-%02d-%02dT%02d-%02d-%02d.%09d%s", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second(), t.Nanosecond(), tz)
}
