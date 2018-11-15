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

package pkcs11

import (
	"encoding/asn1"
	"fmt"
	"hash"

	"github.com/ipfn/ipfn/pkg/crypto/bccsp"
	"github.com/minio/sha256-simd"
	"golang.org/x/crypto/sha3"
)

type config struct {
	ellipticCurve asn1.ObjectIdentifier
	hashFunction  func() hash.Hash
	aesBitLength  int
	rsaBitLength  int
}

func (conf *config) setSecurityLevel(securityLevel int, hashFamily digest.Family) (err error) {
	switch hashFamily {
	case digest.FamilySha2:
		err = conf.setSecurityLevelSHA2(securityLevel)
	case digest.FamilySha3:
		err = conf.setSecurityLevelSHA3(securityLevel)
	default:
		err = fmt.Errorf("Hash Family not supported [%s]", hashFamily)
	}
	return
}

func (conf *config) setSecurityLevelSHA2(level int) (err error) {
	switch level {
	case 256:
		conf.ellipticCurve = oidNamedCurveP256
		conf.hashFunction = sha256.New
		conf.rsaBitLength = 2048
		conf.aesBitLength = 32
	default:
		err = fmt.Errorf("Security level not supported [%d]", level)
	}
	return
}

func (conf *config) setSecurityLevelSHA3(level int) (err error) {
	switch level {
	case 256:
		conf.ellipticCurve = oidNamedCurveP256
		conf.hashFunction = sha3.New256
		conf.rsaBitLength = 2048
		conf.aesBitLength = 32
	case 384:
		conf.ellipticCurve = oidNamedCurveP384
		conf.hashFunction = sha3.New384
		conf.rsaBitLength = 3072
		conf.aesBitLength = 32
	default:
		err = fmt.Errorf("Security level not supported [%d]", level)
	}
	return
}

// PKCS11Opts contains options for the P11Factory
type PKCS11Opts struct {
	// Default algorithms when not specified (Deprecated?)
	SecLevel   int              `mapstructure:"security" json:"security"`
	HashFamily digest.Family `mapstructure:"hash" json:"hash"`

	// Keystore options
	Ephemeral     bool               `mapstructure:"tempkeys,omitempty" json:"tempkeys,omitempty"`
	FileKeystore  *FileKeystoreOpts  `mapstructure:"filekeystore,omitempty" json:"filekeystore,omitempty"`
	DummyKeystore *DummyKeystoreOpts `mapstructure:"dummykeystore,omitempty" json:"dummykeystore,omitempty"`

	// PKCS11 options
	Library    string `mapstructure:"library" json:"library"`
	Label      string `mapstructure:"label" json:"label"`
	Pin        string `mapstructure:"pin" json:"pin"`
	SoftVerify bool   `mapstructure:"softwareverify,omitempty" json:"softwareverify,omitempty"`
	Immutable  bool   `mapstructure:"immutable,omitempty" json:"immutable,omitempty"`
}

// FileKeystoreOpts currently only ECDSA operations go to PKCS11, need a keystore still
// Pluggable Keystores, could add JKS, P12, etc..
type FileKeystoreOpts struct {
	KeyStorePath string `mapstructure:"keystore" json:"keystore" yaml:"KeyStore"`
}

// DummyKeystoreOpts is placeholder for testing purposes
type DummyKeystoreOpts struct{}
