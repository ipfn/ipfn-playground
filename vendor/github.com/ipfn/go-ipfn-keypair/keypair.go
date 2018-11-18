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

package keypair

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcutil/hdkeychain"
	cells "github.com/ipfn/go-ipfn-cells"
)

// KeyPair - Hierarchical deterministic derived key chain.
type KeyPair struct {
	*hdkeychain.ExtendedKey
}

// CID - Returns public key as cid.
func (key *KeyPair) CID() (_ *cells.CID, err error) {
	pub, err := key.ECPubKey()
	if err != nil {
		return
	}
	return CID(pub), nil
}

// DerivePath - Derives extended path.
func (key *KeyPair) DerivePath(path string) (*KeyPair, error) {
	return derivePath(key, path)
}

// ExtendedChild - Derives extended key child by adding 0x80000000 (2^31) to path.
func (key *KeyPair) ExtendedChild(path uint32) (*KeyPair, error) {
	return key.Child(hdkeychain.HardenedKeyStart + path)
}

// Child - Derives extended key child.
func (key *KeyPair) Child(path uint32) (*KeyPair, error) {
	k, err := key.ExtendedKey.Child(path)
	if err != nil {
		return nil, err
	}
	return &KeyPair{ExtendedKey: k}, nil
}

// Serialize - Returns serialized public key bytes.
func (key *KeyPair) Serialize() []byte {
	pubkey, _ := key.ECPubKey()
	return pubkey.SerializeCompressed()
}

// Bytes - Returns public key bytes.
func (key *KeyPair) Bytes() []byte {
	pubkey, _ := key.ECPubKey()
	return PubkeyBytes(pubkey)
}

// String - Returns public key string.
func (key *KeyPair) String() (_ string) {
	if key.ExtendedKey.IsPrivate() {
		neuter, err := key.Neuter()
		if err != nil {
			return
		}
		return neuter.String()
	}
	return key.ExtendedKey.String()
}

// PrivateString - Returns private key string.
func (key *KeyPair) PrivateString() (_ string) {
	if key.ExtendedKey.IsPrivate() {
		return key.ExtendedKey.String()
	}
	return "<not-private-key>"
}

// derivePath - Derives string BIP44 path like `m/44'/0'/1'/0/0`.
func derivePath(key *KeyPair, path string) (res *KeyPair, err error) {
	if tr := strings.TrimPrefix(path, "x/"); tr != path {
		path = HashPath(tr)
	}
	elems := strings.Split(path, "/")
	if len(elems) > 0 && elems[0] == "" {
		elems = elems[1:]
	}
	if len(elems) > 0 && elems[0] == "m" {
		elems = elems[1:]
	} else {
		return nil, fmt.Errorf("invalid derivation path %q", path)
	}
	if len(elems) == 0 {
		return nil, fmt.Errorf("empty derivation path %q", path)
	}
	res = key
	for _, value := range elems {
		v, err := strconv.Atoi(strings.TrimRight(value, "'"))
		if err != nil {
			return nil, fmt.Errorf("wrong derivation path element %q in %q", value, path)
		}
		if strings.HasSuffix(value, "'") {
			res, err = res.ExtendedChild(uint32(v))
		} else {
			res, err = res.Child(uint32(v))
		}
		if err != nil {
			return nil, err
		}
	}
	return
}
