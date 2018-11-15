// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2016-2018 IBM Corp. All Rights Reserved.
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

// Test for bccsp plugin interface.
package main

import (
	"hash"

	"github.com/ipfn/ipfn/pkg/crypto/bccsp"
	"github.com/ipfn/ipfn/pkg/digest"
)

type impl struct{}

// New returns a new instance of the BCCSP implementation
func New(config map[string]interface{}) (bccsp.BCCSP, error) {
	return &impl{}, nil
}

// ReadOnly returns true if this KeyStore is read only, false otherwise.
// If ReadOnly is true then StoreKey will fail.
func (csp *impl) ReadOnly() bool {
	return false
}

// StoreKey stores the key k in this KeyStore.
// If this KeyStore is read only then the method will fail.
func (csp *impl) StoreKey(k bccsp.Key) (err error) {
	return nil
}

// KeyGen generates a key using opts.
func (csp *impl) KeyGen(opts bccsp.KeyGenOpts) (k bccsp.Key, err error) {
	return nil, nil
}

// KeyDeriv derives a key from k using opts.
// The opts argument should be appropriate for the primitive used.
func (csp *impl) KeyDeriv(k bccsp.Key, opts bccsp.KeyDerivOpts) (dk bccsp.Key, err error) {
	return nil, nil
}

// KeyImport imports a key from its raw representation using opts.
// The opts argument should be appropriate for the primitive used.
func (csp *impl) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (k bccsp.Key, err error) {
	return nil, nil
}

// Key returns the key this CSP associates to
// the Subject Key Identifier ski.
func (csp *impl) Key(ski []byte) (k bccsp.Key, err error) {
	return nil, nil
}

// Hash hashes messages msg using options opts.
// If opts is nil, the default hash function will be used.
func (csp *impl) Hash(msg []byte, t digest.Type) (hash []byte, err error) {
	return nil, nil
}

// GetHash returns and instance of hash.Hash using options opts.
// If opts is nil, the default hash function will be returned.
func (csp *impl) Hasher(t digest.Type) (h hash.Hash, err error) {
	return nil, nil
}

// Sign signs digest using key k.
// The opts argument should be appropriate for the algorithm used.
//
// Note that when a signature of a hash of a larger message is needed,
// the caller is responsible for hashing the larger message and passing
// the hash (as digest).
func (csp *impl) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) (signature []byte, err error) {
	return nil, nil
}

// Verify verifies signature against key k and digest
// The opts argument should be appropriate for the algorithm used.
func (csp *impl) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (valid bool, err error) {
	return true, nil
}

// Encrypt encrypts plaintext using key k.
// The opts argument should be appropriate for the algorithm used.
func (csp *impl) Encrypt(k bccsp.Key, plaintext []byte, opts bccsp.EncrypterOpts) (ciphertext []byte, err error) {
	return nil, nil
}

// Decrypt decrypts ciphertext using key k.
// The opts argument should be appropriate for the algorithm used.
func (csp *impl) Decrypt(k bccsp.Key, ciphertext []byte, opts bccsp.DecrypterOpts) (plaintext []byte, err error) {
	return nil, nil
}
