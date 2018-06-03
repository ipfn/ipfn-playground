// The MIT License (MIT)
//
// Copyright © 2014 Tyler Smith.
// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Package bip39 implements BIP0039 spec for mnemonic seeds.
package bip39

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"golang.org/x/crypto/pbkdf2"
)

// Some bitwise operands for working with big.Ints
var (
	Last11BitsMask          = big.NewInt(2047)
	RightShift11BitsDivider = big.NewInt(2048)
	BigOne                  = big.NewInt(1)
	BigTwo                  = big.NewInt(2)
)

// NewEntropy will create random entropy bytes
// so long as the requested size bitSize is an appropriate size.
func NewEntropy(bitSize int) ([]byte, error) {
	err := validateEntropyBitSize(bitSize)
	if err != nil {
		return nil, err
	}

	entropy := make([]byte, bitSize/8)
	_, err = rand.Read(entropy)
	return entropy, err
}

// NewDefaultEntropy will create random entropy bytes of size 32 bits.
func NewDefaultEntropy() ([]byte, error) {
	return NewEntropy(32)
}

// NewSeedWithErrorChecking creates a hashed seed output given the mnemonic string and a password.
// An error is returned if the mnemonic is not convertible to a byte array.
func NewSeedWithErrorChecking(mnemonic, password string) ([]byte, error) {
	_, err := MnemonicToByteArray(mnemonic)
	if err != nil {
		return nil, err
	}
	return NewSeed([]byte(mnemonic), []byte(password)), nil
}

// NewSeed creates a hashed seed output given a provided string and password.
// No checking is performed to validate that the string provided is a valid mnemonic.
func NewSeed(mnemonic, salt []byte) []byte {
	salt = append([]byte("mnemonic"), salt[:]...)
	return NewCustomSeed(mnemonic, salt, 2048, 64)
}

// NewSeedFromString creates a new hashed seed from string. See NewSeed.
func NewSeedFromString(mnemonic, salt string) []byte {
	return NewSeed([]byte(mnemonic), []byte(salt))
}

// NewCustomSeed  Creates a new key generation seed with given size
// from mnemonic, salt and amount of derivation iterations to perform.
// NewSeed creates a hashed seed output given a provided string and password.
// No checking is performed to validate that the string provided is a valid mnemonic.
func NewCustomSeed(mnemonic, salt []byte, iter, size int) []byte {
	return pbkdf2.Key(mnemonic, salt, iter, size, sha512.New)
}

// NewMnemonic will return a string consisting of the mnemonic words for
// the given entropy.
// If the provide entropy is invalid, an error will be returned.
func NewMnemonic(entropy []byte) (string, error) {
	// Compute some lengths for convenience
	entropyBitLength := len(entropy) * 8
	checksumBitLength := entropyBitLength / 32
	sentenceLength := (entropyBitLength + checksumBitLength) / 11

	err := validateEntropyBitSize(entropyBitLength)
	if err != nil {
		return "", err
	}

	// Add checksum to entropy
	entropy = addChecksum(entropy)

	// Break entropy up into sentenceLength chunks of 11 bits
	// For each word AND mask the rightmost 11 bits and find the word at that index
	// Then bitshift entropy 11 bits right and repeat
	// Add to the last empty slot so we can work with LSBs instead of MSB

	// Entropy as an int so we can bitmask without worrying about bytes slices
	entropyInt := new(big.Int).SetBytes(entropy)

	// Slice to hold words in
	words := make([]string, sentenceLength)

	// Throw away big int for AND masking
	word := big.NewInt(0)

	for i := sentenceLength - 1; i >= 0; i-- {
		// Get 11 right most bits and bitshift 11 to the right for next time
		word.And(entropyInt, Last11BitsMask)
		entropyInt.Div(entropyInt, RightShift11BitsDivider)

		// Get the bytes representing the 11 bits as a 2 byte slice
		wordBytes := padByteSlice(word.Bytes(), 2)

		// Convert bytes to an index and add that word to the list
		words[i] = WordList[binary.BigEndian.Uint16(wordBytes)]
	}

	return strings.Join(words, " "), nil
}

// MnemonicToByteArray takes a mnemonic string and turns it into a byte array
// suitable for creating another mnemonic.
// An error is returned if the mnemonic is invalid.
// FIXME
// This does not work for all values in
// the test vectors.  Namely
// Vectors 0, 4, and 8.
// This is not really important because BIP39 doesnt really define a conversion
// from string to bytes.
func MnemonicToByteArray(mnemonic string) ([]byte, error) {
	if IsMnemonicValid(mnemonic) == false {
		return nil, fmt.Errorf("Invalid mnemonic")
	}
	mnemonicSlice := strings.Split(mnemonic, " ")

	bitSize := len(mnemonicSlice) * 11
	err := validateEntropyWithChecksumBitSize(bitSize)
	if err != nil {
		return nil, err
	}
	checksumSize := bitSize % 32

	b := big.NewInt(0)
	modulo := big.NewInt(2048)
	for _, v := range mnemonicSlice {
		index, found := ReverseWordMap[v]
		if found == false {
			return nil, fmt.Errorf("Word `%v` not found in reverse map", v)
		}
		add := big.NewInt(int64(index))
		b = b.Mul(b, modulo)
		b = b.Add(b, add)
	}
	hex := b.Bytes()
	checksumModulo := big.NewInt(0).Exp(big.NewInt(2), big.NewInt(int64(checksumSize)), nil)
	entropy, _ := big.NewInt(0).DivMod(b, checksumModulo, big.NewInt(0))

	entropyHex := entropy.Bytes()

	byteSize := bitSize/8 + 1
	if len(hex) != byteSize {
		tmp := make([]byte, byteSize)
		diff := byteSize - len(hex)
		for i := 0; i < len(hex); i++ {
			tmp[i+diff] = hex[i]
		}
		hex = tmp
	}

	validationHex := addChecksum(entropyHex)
	if len(validationHex) != byteSize {
		tmp2 := make([]byte, byteSize)
		diff2 := byteSize - len(validationHex)
		for i := 0; i < len(validationHex); i++ {
			tmp2[i+diff2] = validationHex[i]
		}
		validationHex = tmp2
	}

	if len(hex) != len(validationHex) {
		panic("[]byte len mismatch - it shouldn't happen")
	}
	for i := range validationHex {
		if hex[i] != validationHex[i] {
			return nil, fmt.Errorf("Invalid byte at position %v", i)
		}
	}
	return hex, nil
}

// IsMnemonicValid attempts to verify that the provided mnemonic is valid.
// Validity is determined by both the number of words being appropriate,
// and that all the words in the mnemonic are present in the word list.
func IsMnemonicValid(mnemonic string) bool {
	// Create a list of all the words in the mnemonic sentence
	words := strings.Fields(mnemonic)

	//Get num of words
	numOfWords := len(words)

	// The number of words should be 12, 15, 18, 21 or 24
	if numOfWords%3 != 0 || numOfWords < 12 || numOfWords > 24 {
		return false
	}

	// Check if all words belong in the wordlist
	for i := 0; i < numOfWords; i++ {
		if !contains(WordList, words[i]) {
			return false
		}
	}

	return true
}

// Appends to data the first (len(data) / 32)bits of the result of sha256(data)
// Currently only supports data up to 32 bytes
func addChecksum(data []byte) []byte {
	// Get first byte of sha256
	hasher := sha256.New()
	hasher.Write(data)
	hash := hasher.Sum(nil)
	firstChecksumByte := hash[0]

	// len() is in bytes so we divide by 4
	checksumBitLength := uint(len(data) / 4)

	// For each bit of check sum we want we shift the data one the left
	// and then set the (new) right most bit equal to checksum bit at that index
	// staring from the left
	dataBigInt := new(big.Int).SetBytes(data)
	for i := uint(0); i < checksumBitLength; i++ {
		// Bitshift 1 left
		dataBigInt.Mul(dataBigInt, BigTwo)

		// Set rightmost bit if leftmost checksum bit is set
		if uint8(firstChecksumByte&(1<<(7-i))) > 0 {
			dataBigInt.Or(dataBigInt, BigOne)
		}
	}

	return dataBigInt.Bytes()
}

func contains(list []string, value string) bool {
	for _, v := range list {
		if v == value {
			return true
		}
	}
	return false
}

func padByteSlice(slice []byte, length int) []byte {
	res := make([]byte, length-len(slice))
	return append(res, slice...)
}

func validateEntropyBitSize(size int) error {
	if (size%32) != 0 || size < 128 || size > 256 {
		return errors.New("Entropy length must be [128, 256] and a multiple of 32")
	}
	return nil
}

func validateEntropyWithChecksumBitSize(size int) error {
	if (size != 128+4) && (size != 160+5) && (size != 192+6) && (size != 224+7) && (size != 256+8) {
		return fmt.Errorf("Wrong entropy + checksum size - expected %v, got %v", int((size-size%32)+(size-size%32)/32), size)
	}
	return nil
}
