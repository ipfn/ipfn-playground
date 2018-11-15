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

package bccsp

import (
	"crypto"
	"hash"

	"github.com/ipfn/ipfn/pkg/digest"
)

// BCCSP is the blockchain cryptographic service provider that offers
// the implementation of cryptographic standards and algorithms.
type BCCSP interface {
	KeyStore
	KeyGenerator
	KeyDeriver
	KeyImporter
	Hasher
	Signer
	Verifier
	Encryptor
	Decryptor
}

// KeyGenerator is a BCCSP-like interface that provides key generation algorithms
type KeyGenerator interface {
	// KeyGen generates a key using opts.
	KeyGen(opts KeyGenOpts) (k Key, err error)
}

// KeyDeriver is a BCCSP-like interface that provides key derivation algorithms
type KeyDeriver interface {
	// KeyDeriv derives a key from k using opts.
	// The opts argument should be appropriate for the primitive used.
	KeyDeriv(k Key, opts KeyDerivOpts) (dk Key, err error)
}

// KeyImporter is a BCCSP-like interface that provides key import algorithms
type KeyImporter interface {
	// KeyImport imports a key from its raw representation using opts.
	// The opts argument should be appropriate for the primitive used.
	KeyImport(raw interface{}, opts KeyImportOpts) (k Key, err error)
}

// Encryptor is a BCCSP-like interface that provides encryption algorithms
type Encryptor interface {
	// Encrypt encrypts plaintext using key k.
	// The opts argument should be appropriate for the algorithm used.
	Encrypt(k Key, plaintext []byte, opts EncrypterOpts) (ciphertext []byte, err error)
}

// Decryptor is a BCCSP-like interface that provides decryption algorithms
type Decryptor interface {
	// Decrypt decrypts ciphertext using key k.
	// The opts argument should be appropriate for the algorithm used.
	Decrypt(k Key, ciphertext []byte, opts DecrypterOpts) (plaintext []byte, err error)
}

// Signer is a BCCSP-like interface that provides signing algorithms
type Signer interface {
	// Sign signs digest using key k.
	// The opts argument should be appropriate for the algorithm used.
	//
	// Note that when a signature of a hash of a larger message is needed,
	// the caller is responsible for hashing the larger message and passing
	// the hash (as digest).
	Sign(k Key, digest []byte, opts SignerOpts) (signature []byte, err error)
}

// Verifier is a BCCSP-like interface that provides verifying algorithms
type Verifier interface {
	// Verify verifies signature against key k and digest
	// The opts argument should be appropriate for the algorithm used.
	Verify(k Key, signature, digest []byte, opts SignerOpts) (valid bool, err error)
}

// Hasher is a BCCSP-like interface that provides hash algorithms
type Hasher interface {
	// Hash hashes messages msg using options opts.
	// If opts is nil, the default hash function will be used.
	Hash(msg []byte, algo digest.Type) (hash []byte, err error)

	// Hasher returns and instance of hash.Hash using options opts.
	// If opts is nil, the default hash function will be returned.
	Hasher(algo digest.Type) (h hash.Hash, err error)
}

// SignerOpts contains options for signing with a CSP.
type SignerOpts interface {
	crypto.SignerOpts
}

// EncrypterOpts contains options for encrypting with a CSP.
type EncrypterOpts interface{}

// DecrypterOpts contains options for decrypting with a CSP.
type DecrypterOpts interface{}
