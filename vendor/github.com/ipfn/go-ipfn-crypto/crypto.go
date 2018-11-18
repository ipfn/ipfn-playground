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

// Package crypto implements cryptographic utilities.
package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"

	"golang.org/x/crypto/pbkdf2"

	sha3 "golang.org/x/crypto/sha3"
)

// NewEntropy - Generates random entropy of a given size.
func NewEntropy(size int) (entropy []byte, err error) {
	entropy = make([]byte, size)
	_, err = io.ReadFull(rand.Reader, entropy)
	return
}

// DeriveEncKey - Derive encryption key from password using PBKDF2.
// Calls to Decrypt and Encrypt automatically use it.
func DeriveEncKey(password, salt []byte) []byte {
	hash := sha3.Sum512(password)
	return pbkdf2.Key(hash[:32], salt, 10240, 32, sha3.New512)
}

// Decrypt - Decrypts stored message using custom key derived from password.
// Encryption key is derived using DeriveEncKey function.
func Decrypt(ciphertext, nonce, password, salt []byte) (_ []byte, err error) {
	aesgcm, err := newCipherGCM(password, salt)
	if err != nil {
		return
	}
	return aesgcm.Open(nil, nonce, ciphertext, nil)
}

// Encrypt - Encrypts message for storage using custom key derived from password.
// Encryption key is derived using DeriveEncKey function, returned salt should be saved.
func Encrypt(message, password []byte) (ciphertext, nonce, salt []byte, err error) {
	salt, err = NewEntropy(32)
	if err != nil {
		return
	}
	aesgcm, err := newCipherGCM(password, salt)
	if err != nil {
		return
	}
	nonce, err = NewEntropy(12)
	if err != nil {
		return
	}
	ciphertext = aesgcm.Seal(nil, nonce, message, nil)
	return
}

func newCipherGCM(password, salt []byte) (_ cipher.AEAD, err error) {
	block, err := aes.NewCipher(DeriveEncKey(password, salt))
	if err != nil {
		return
	}
	return cipher.NewGCM(block)
}
