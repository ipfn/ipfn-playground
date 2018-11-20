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

	"github.com/ipfn/ipfn/pkg/crypto/bccsp"
	"github.com/ipfn/ipfn/pkg/digest"
	"golang.org/x/crypto/ed25519"
)

type ed25519PrivateKeyKeyDeriver struct{}

func (kd *ed25519PrivateKeyKeyDeriver) KeyDeriv(k bccsp.Key, opts bccsp.KeyDerivOpts) (bccsp.Key, error) {
	// Validate opts
	if opts == nil {
		return nil, errors.New("derive opts cant be nil")
	}

	pk := k.(*ed25519PrivateKey)
	op := opts.(*bccsp.ED25519ReRandKeyOpts)

	pkhash := digest.SumSha256(pk.privKey.Seed())
	seed := digest.SumSha256(pkhash, op.Expansion)[:ed25519.SeedSize]

	privateKey := ed25519.NewKeyFromSeed(seed)
	publicKey := make([]byte, ed25519.PublicKeySize)
	copy(publicKey, privateKey[32:])

	return &ed25519PrivateKey{
		privKey: privateKey,
		pubKey:  &ed25519PublicKey{publicKey},
	}, nil
}
