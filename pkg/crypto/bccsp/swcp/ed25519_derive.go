// Copyright © 2018 The IPFN Developers. All Rights Reserved.
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

package swcp

import (
	"errors"
	"io"

	"github.com/ipfn/ipfn/pkg/crypto/bccsp"
	"github.com/minio/sha256-simd"
	"golang.org/x/crypto/ed25519"
	"golang.org/x/crypto/hkdf"
)

type ed25519PrivateKeyKeyDeriver struct{}

func (kd *ed25519PrivateKeyKeyDeriver) KeyDeriv(k bccsp.Key, opts bccsp.KeyDerivOpts) (_ bccsp.Key, err error) {
	// Validate opts
	if opts == nil {
		return nil, errors.New("derive opts cant be nil")
	}

	var (
		sk     = k.(*ed25519PrivateKey)
		op     = opts.(*bccsp.ED25519ReRandKeyOpts)
		r      = hkdf.New(sha256.New, sk.privKey.Seed(), op.Expansion, []byte("ad9ba3560bdcd0894f887ea27774ac98"))
		seed   = make([]byte, ed25519.PrivateKeySize)
		pubkey = make([]byte, ed25519.PublicKeySize)
	)

	_, err = io.ReadFull(r, seed)
	if err != nil {
		return nil, err
	}

	privateKey := ed25519.NewKeyFromSeed(seed)
	copy(pubkey, privateKey[32:])

	return &ed25519PrivateKey{
		privKey: privateKey,
		pubKey:  &ed25519PublicKey{pubkey},
	}, nil
}
