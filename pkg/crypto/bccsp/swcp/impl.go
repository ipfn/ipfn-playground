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

package swcp

import (
	"hash"
	"reflect"

	"github.com/pkg/errors"

	"github.com/ipfn/ipfn/pkg/crypto/bccsp"
	"github.com/ipfn/ipfn/pkg/utils/flog"
)

var (
	logger = flog.MustGetLogger("bccsp_sw")
)

// CSP provides a generic implementation of the BCCSP interface based
// on wrappers. It can be customized by providing implementations for the
// following algorithm-based wrappers: KeyGenerator, KeyDeriver, KeyImporter,
// Encryptor, Decryptor, Signer, Verifier, Hasher. Each wrapper is bound to a
// goland type representing either an option or a key.
type CSP struct {
	ks bccsp.KeyStore

	keyGenerators map[reflect.Type]bccsp.KeyGenerator
	keyDerivers   map[reflect.Type]bccsp.KeyDeriver
	keyImporters  map[reflect.Type]bccsp.KeyImporter
	encryptors    map[reflect.Type]bccsp.Encryptor
	decryptors    map[reflect.Type]bccsp.Decryptor
	signers       map[reflect.Type]bccsp.Signer
	verifiers     map[reflect.Type]bccsp.Verifier
	hashers       map[bccsp.HashType]bccsp.Hasher
}

// New - Creates new software implemented BCCSP.
func New(keyStore bccsp.KeyStore) (*CSP, error) {
	if keyStore == nil {
		return nil, errors.Errorf("Invalid bccsp.KeyStore instance. It must be different from nil")
	}

	encryptors := make(map[reflect.Type]bccsp.Encryptor)
	decryptors := make(map[reflect.Type]bccsp.Decryptor)
	signers := make(map[reflect.Type]bccsp.Signer)
	verifiers := make(map[reflect.Type]bccsp.Verifier)
	keyGenerators := make(map[reflect.Type]bccsp.KeyGenerator)
	keyDerivers := make(map[reflect.Type]bccsp.KeyDeriver)
	keyImporters := make(map[reflect.Type]bccsp.KeyImporter)
	hashers := make(map[bccsp.HashType]bccsp.Hasher)

	csp := &CSP{keyStore,
		keyGenerators, keyDerivers, keyImporters, encryptors,
		decryptors, signers, verifiers, hashers}

	return csp, nil
}

// ReadOnly returns true if this KeyStore is read only, false otherwise.
// If ReadOnly is true then StoreKey will fail.
func (csp *CSP) ReadOnly() bool {
	return csp.ks.ReadOnly()
}

// StoreKey stores the key k in this KeyStore.
// If this KeyStore is read only then the method will fail.
func (csp *CSP) StoreKey(k bccsp.Key) (err error) {
	return csp.ks.StoreKey(k)
}

// KeyGen generates a key using opts.
func (csp *CSP) KeyGen(opts bccsp.KeyGenOpts) (k bccsp.Key, err error) {
	// Validate arguments
	if opts == nil {
		return nil, errors.New("Invalid Opts parameter. It must not be nil.")
	}

	keyGenerator, found := csp.keyGenerators[reflect.TypeOf(opts)]
	if !found {
		return nil, errors.Errorf("Unsupported 'KeyGenOpts' provided [%v]", opts)
	}

	k, err = keyGenerator.KeyGen(opts)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed generating key with opts [%v]", opts)
	}

	// If the key is not Ephemeral, store it.
	if !opts.Ephemeral() {
		// Store the key
		err = csp.ks.StoreKey(k)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed storing key [%s]", opts.Algorithm())
		}
	}

	return k, nil
}

// KeyDeriv derives a key from k using opts.
// The opts argument should be appropriate for the primitive used.
func (csp *CSP) KeyDeriv(k bccsp.Key, opts bccsp.KeyDerivOpts) (dk bccsp.Key, err error) {
	// Validate arguments
	if k == nil {
		return nil, errors.New("Invalid Key. It must not be nil.")
	}
	if opts == nil {
		return nil, errors.New("Invalid opts. It must not be nil.")
	}

	keyDeriver, found := csp.keyDerivers[reflect.TypeOf(k)]
	if !found {
		return nil, errors.Errorf("Unsupported 'Key' provided [%v]", k)
	}

	k, err = keyDeriver.KeyDeriv(k, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed deriving key with opts [%v]", opts)
	}

	// If the key is not Ephemeral, store it.
	if !opts.Ephemeral() {
		// Store the key
		err = csp.ks.StoreKey(k)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed storing key [%s]", opts.Algorithm())
		}
	}

	return k, nil
}

// KeyImport imports a key from its raw representation using opts.
// The opts argument should be appropriate for the primitive used.
func (csp *CSP) KeyImport(raw interface{}, opts bccsp.KeyImportOpts) (k bccsp.Key, err error) {
	// Validate arguments
	if raw == nil {
		return nil, errors.New("Invalid raw. It must not be nil.")
	}
	if opts == nil {
		return nil, errors.New("Invalid opts. It must not be nil.")
	}

	keyImporter, found := csp.keyImporters[reflect.TypeOf(opts)]
	if !found {
		return nil, errors.Errorf("Unsupported 'KeyImportOpts' provided [%v]", opts)
	}

	k, err = keyImporter.KeyImport(raw, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed importing key with opts [%v]", opts)
	}

	// If the key is not Ephemeral, store it.
	if !opts.Ephemeral() {
		// Store the key
		err = csp.ks.StoreKey(k)
		if err != nil {
			return nil, errors.Wrapf(err, "Failed storing imported key with opts [%v]", opts)
		}
	}

	return
}

// Key returns the key this CSP associates to
// the Subject Key Identifier ski.
func (csp *CSP) Key(ski []byte) (k bccsp.Key, err error) {
	k, err = csp.ks.Key(ski)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed getting key for SKI [%v]", ski)
	}

	return
}

// Hash hashes messages msg using options opts.
func (csp *CSP) Hash(msg []byte, hashType bccsp.HashType) (digest []byte, err error) {
	hasher, found := csp.hashers[hashType]
	if !found {
		return nil, errors.Errorf("Unsupported hash type [%v]", hashType)
	}

	digest, err = hasher.Hash(msg, hashType)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed hashing with hash type [%v]", hashType)
	}

	return
}

// Hasher returns and instance of hash.Hash using options opts.
// If opts is nil then the default hash function is returned.
func (csp *CSP) Hasher(hashType bccsp.HashType) (h hash.Hash, err error) {
	hasher, found := csp.hashers[hashType]
	if !found {
		return nil, errors.Errorf("Unsupported hash type [%v]", hashType)
	}

	h, err = hasher.Hasher(hashType)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed getting hash function with hash type [%v]", hashType)
	}

	return
}

// Sign signs digest using key k.
// The opts argument should be appropriate for the primitive used.
//
// Note that when a signature of a hash of a larger message is needed,
// the caller is responsible for hashing the larger message and passing
// the hash (as digest).
func (csp *CSP) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) (signature []byte, err error) {
	// Validate arguments
	if k == nil {
		return nil, errors.New("Invalid Key. It must not be nil.")
	}
	if len(digest) == 0 {
		return nil, errors.New("Invalid digest. Cannot be empty.")
	}

	keyType := reflect.TypeOf(k)
	signer, found := csp.signers[keyType]
	if !found {
		return nil, errors.Errorf("Unsupported 'SignKey' provided [%s]", keyType)
	}

	signature, err = signer.Sign(k, digest, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed signing with opts [%v]", opts)
	}

	return
}

// Verify verifies signature against key k and digest
func (csp *CSP) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (valid bool, err error) {
	// Validate arguments
	if k == nil {
		return false, errors.New("Invalid Key. It must not be nil.")
	}
	if len(signature) == 0 {
		return false, errors.New("Invalid signature. Cannot be empty.")
	}
	if len(digest) == 0 {
		return false, errors.New("Invalid digest. Cannot be empty.")
	}

	verifier, found := csp.verifiers[reflect.TypeOf(k)]
	if !found {
		return false, errors.Errorf("Unsupported 'VerifyKey' provided [%v]", k)
	}

	valid, err = verifier.Verify(k, signature, digest, opts)
	if err != nil {
		return false, errors.Wrapf(err, "Failed verifing with opts [%v]", opts)
	}

	return
}

// Encrypt encrypts plaintext using key k.
// The opts argument should be appropriate for the primitive used.
func (csp *CSP) Encrypt(k bccsp.Key, plaintext []byte, opts bccsp.EncrypterOpts) ([]byte, error) {
	// Validate arguments
	if k == nil {
		return nil, errors.New("Invalid Key. It must not be nil.")
	}

	encryptor, found := csp.encryptors[reflect.TypeOf(k)]
	if !found {
		return nil, errors.Errorf("Unsupported 'EncryptKey' provided [%v]", k)
	}

	return encryptor.Encrypt(k, plaintext, opts)
}

// Decrypt decrypts ciphertext using key k.
// The opts argument should be appropriate for the primitive used.
func (csp *CSP) Decrypt(k bccsp.Key, ciphertext []byte, opts bccsp.DecrypterOpts) (plaintext []byte, err error) {
	// Validate arguments
	if k == nil {
		return nil, errors.New("Invalid Key. It must not be nil.")
	}

	decryptor, found := csp.decryptors[reflect.TypeOf(k)]
	if !found {
		return nil, errors.Errorf("Unsupported 'DecryptKey' provided [%v]", k)
	}

	plaintext, err = decryptor.Decrypt(k, ciphertext, opts)
	if err != nil {
		return nil, errors.Wrapf(err, "Failed decrypting with opts [%v]", opts)
	}

	return
}

// AddWrapper binds the passed type to the passed wrapper.
// Notice that that wrapper must be an instance of one of the following interfaces:
// KeyGenerator, KeyDeriver, KeyImporter, Encryptor, Decryptor, Signer, Verifier.
func (csp *CSP) AddWrapper(t reflect.Type, w interface{}) error {
	if t == nil {
		return errors.Errorf("type cannot be nil")
	}
	if w == nil {
		return errors.Errorf("wrapper cannot be nil")
	}
	switch dt := w.(type) {
	case bccsp.KeyGenerator:
		csp.keyGenerators[t] = dt
	case bccsp.KeyImporter:
		csp.keyImporters[t] = dt
	case bccsp.KeyDeriver:
		csp.keyDerivers[t] = dt
	case bccsp.Encryptor:
		csp.encryptors[t] = dt
	case bccsp.Decryptor:
		csp.decryptors[t] = dt
	case bccsp.Signer:
		csp.signers[t] = dt
	case bccsp.Verifier:
		csp.verifiers[t] = dt
	default:
		return errors.Errorf("wrapper type not valid, must be on of: KeyGenerator, KeyDeriver, KeyImporter, Encryptor, Decryptor, Signer, Verifier, Hasher")
	}
	return nil
}

// AddHasher binds the passed type to the passed wrapper.
func (csp *CSP) AddHasher(t bccsp.HashType, hasher bccsp.Hasher) error {
	if csp.hashers[t] != nil {
		return errors.Errorf("hasher for type %s already implemented", t)
	}
	csp.hashers[t] = hasher
	return nil
}
