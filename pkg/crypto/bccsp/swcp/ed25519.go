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
	"crypto/rand"
	"errors"
	"fmt"

	"github.com/ipfn/ipfn/pkg/crypto/bccsp"
	"github.com/ipfn/ipfn/pkg/digest"
	"golang.org/x/crypto/ed25519"
)

type ed25519KeyGenerator struct{}

func (kg *ed25519KeyGenerator) KeyGen(opts bccsp.KeyGenOpts) (bccsp.Key, error) {
	publicKey, privateKey, err := ed25519.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate ed25519 key: [%s]", err)
	}

	return &ed25519PrivateKey{
		privKey: privateKey,
		pubKey:  &ed25519PublicKey{publicKey},
	}, nil
}

type ed25519PrivateKey struct {
	privKey ed25519.PrivateKey
	pubKey  *ed25519PublicKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *ed25519PrivateKey) Bytes() ([]byte, error) {
	return nil, errors.New("not supported")
}

// SKI returns the subject key identifier of this key.
func (k *ed25519PrivateKey) SKI() []byte {
	return k.pubKey.SKI()
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k *ed25519PrivateKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k *ed25519PrivateKey) Private() bool {
	return true
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k *ed25519PrivateKey) PublicKey() (bccsp.Key, error) {
	return k.pubKey, nil
}

type ed25519PublicKey struct {
	pubKey ed25519.PublicKey
}

// Bytes converts this key to its byte representation,
// if this operation is allowed.
func (k *ed25519PublicKey) Bytes() (raw []byte, err error) {
	return k.pubKey, nil
}

// SKI returns the subject key identifier of this key.
func (k *ed25519PublicKey) SKI() []byte {
	return digest.SumSha256Bytes(k.pubKey)
}

// Symmetric returns true if this key is a symmetric key,
// false if this key is asymmetric
func (k *ed25519PublicKey) Symmetric() bool {
	return false
}

// Private returns true if this key is a private key,
// false otherwise.
func (k *ed25519PublicKey) Private() bool {
	return false
}

// PublicKey returns the corresponding public key part of an asymmetric public/private key pair.
// This method returns an error in symmetric key schemes.
func (k *ed25519PublicKey) PublicKey() (bccsp.Key, error) {
	return k, nil
}

type ed25519Signer struct{}

func (s *ed25519Signer) Sign(k bccsp.Key, digest []byte, opts bccsp.SignerOpts) ([]byte, error) {
	return ed25519.Sign(k.(*ed25519PrivateKey).privKey, digest), nil
}

type ed25519PrivateKeyVerifier struct{}

func (v *ed25519PrivateKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return ed25519.Verify(k.(*ed25519PrivateKey).pubKey.pubKey, digest, signature), nil
}

type ed25519PublicKeyKeyVerifier struct{}

func (v *ed25519PublicKeyKeyVerifier) Verify(k bccsp.Key, signature, digest []byte, opts bccsp.SignerOpts) (bool, error) {
	return ed25519.Verify(k.(*ed25519PublicKey).pubKey, digest, signature), nil
}
