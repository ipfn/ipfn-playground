// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2014-2018 The go-ethereum Authors. All Rights Reserved.
//
// This file is part of the IPFN project.
// This file was part of the go-ethereum library.
//
// The IPFN project is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The IPFN project is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package sealbox

import (
	"crypto/aes"
	"encoding/hex"

	"github.com/ipfn/ipfn/pkg/crypto/entropy"
	"github.com/ipfn/ipfn/pkg/digest"
	"golang.org/x/crypto/scrypt"
)

// Encrypt - Encrypts a box using the specified scrypt parameters.
func Encrypt(body, pwd []byte, scryptN, scryptP int) (_ SealedBox, err error) {
	salt, err := entropy.New(32)
	if err != nil {
		return
	}
	derivedKey, err := scrypt.Key(pwd, salt, scryptN, scryptR, scryptP, scryptDKLen)
	if err != nil {
		return
	}
	iv, err := entropy.New(aes.BlockSize)
	if err != nil {
		return
	}
	cipherText, err := aesCTRXOR(derivedKey[:16], body, iv)
	if err != nil {
		return
	}
	return SealedBox{
		Version: version,
		Crypto: Crypto{
			Cipher:     "aes-128-ctr",
			CipherText: hex.EncodeToString(cipherText),
			CipherParams: CipherParams{
				IV: hex.EncodeToString(iv),
			},
			KDF: keyHeaderKDF,
			KDFParams: KDFParams{
				N:     scryptN,
				R:     scryptR,
				P:     scryptP,
				DKLen: scryptDKLen,
				Salt:  hex.EncodeToString(salt),
			},
			MAC: hex.EncodeToString(digest.SumKeccak256Bytes(derivedKey[16:32], cipherText)),
		},
	}, nil
}
