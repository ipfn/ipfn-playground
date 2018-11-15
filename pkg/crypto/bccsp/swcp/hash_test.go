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
	"crypto/sha256"
	"errors"
	"testing"

	"github.com/ipfn/ipfn/pkg/crypto/bccsp"
	"github.com/ipfn/ipfn/pkg/crypto/bccsp/swcp/mocks"
	"github.com/ipfn/ipfn/pkg/digest"
	"github.com/stretchr/testify/assert"
)

func TestHash(t *testing.T) {
	t.Parallel()

	expectetMsg := []byte{1, 2, 3, 4}
	expectedOpts := digest.Sha2_256
	expectetValue := []byte{1, 2, 3, 4, 5}
	expectedErr := errors.New("Expected Error")

	hashers := make(map[digest.Type]bccsp.Hasher)
	hashers[expectedOpts] = &mocks.Hasher{
		MsgArg:  expectetMsg,
		OptsArg: expectedOpts,
		Value:   expectetValue,
		Err:     nil,
	}
	csp := CSP{hashers: hashers}
	value, err := csp.Hash(expectetMsg, expectedOpts)
	assert.Equal(t, expectetValue, value)
	assert.Nil(t, err)

	hashers = make(map[digest.Type]bccsp.Hasher)
	hashers[expectedOpts] = &mocks.Hasher{
		MsgArg:  expectetMsg,
		OptsArg: expectedOpts,
		Value:   nil,
		Err:     expectedErr,
	}
	csp = CSP{hashers: hashers}
	value, err = csp.Hash(expectetMsg, expectedOpts)
	assert.Nil(t, value)
	assert.Contains(t, err.Error(), expectedErr.Error())
}

func TestHasher(t *testing.T) {
	t.Parallel()

	expectedOpts := digest.Sha2_256
	expectetValue := sha256.New()
	expectedErr := errors.New("Expected Error")

	hashers := make(map[digest.Type]bccsp.Hasher)
	hashers[expectedOpts] = &mocks.Hasher{
		OptsArg:   expectedOpts,
		ValueHash: expectetValue,
		Err:       nil,
	}
	csp := CSP{hashers: hashers}
	value, err := csp.Hasher(expectedOpts)
	assert.Equal(t, expectetValue, value)
	assert.Nil(t, err)

	hashers = make(map[digest.Type]bccsp.Hasher)
	hashers[expectedOpts] = &mocks.Hasher{
		OptsArg:   expectedOpts,
		ValueHash: expectetValue,
		Err:       expectedErr,
	}
	csp = CSP{hashers: hashers}
	value, err = csp.Hasher(expectedOpts)
	assert.Nil(t, value)
	assert.Contains(t, err.Error(), expectedErr.Error())
}

func TestHasherSha256(t *testing.T) {
	t.Parallel()

	hasher := &hasher{algo: digest.Sha2_256, impl: sha256.New}

	msg := []byte("Hello World")
	out, err := hasher.Hash(msg, digest.Sha2_256)
	assert.NoError(t, err)
	h := sha256.New()
	h.Write(msg)
	out2 := h.Sum(nil)
	assert.Equal(t, out, out2)

	hf, err := hasher.Hasher(digest.Sha2_256)
	assert.NoError(t, err)
	assert.Equal(t, hf, sha256.New())
}
