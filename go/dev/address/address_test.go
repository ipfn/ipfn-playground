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

package address

import (
	"encoding/json"
	"fmt"
	. "testing"

	"github.com/ipfn/ipfn/go/keypair"
	"github.com/stretchr/testify/assert"
)

var (
	testSrc  = "beqpdfdhq87dkncb"
	testAddr = &Address{ID: 2191370559816, Extra: 13471}

	testKeyAddr = "bnx37fk4wmxur3j0puapwv"
	testKeyCID  = "zFNScYMHCVTuggbomkvLDPpoLruXKwWEu6SV2S45P9QwuLwhfMeq"
	testPrivKey = "fnprvqWuSwMXSGgZYML8TUkm7zJfMVnMVS7PpXUYDRUbfFnN6s29mRmW82FZcks64djUmYZ8t9CLaxc4dFAMvxxdiKGn9iqjp783LLv9c45Z9HpR"
)

func TestParseAddress(t *T) {
	a, err := ParseAddress(testSrc)
	assert.Equal(t, nil, err)
	assert.Equal(t, testAddr.ID, a.ID)
	assert.Equal(t, testAddr.Extra, a.Extra)
	assert.Equal(t, testAddr.String(), testSrc)
	assert.Equal(t, testSrc, a.String())
}

func TestAddressJSON(t *T) {
	b, _ := json.Marshal(testAddr)
	assert.Equal(t, `"beqpdfdhq87dkncb"`, fmt.Sprintf("%s", b))
	r := new(Address)
	err := json.Unmarshal(b, r)
	assert.Equal(t, nil, err)
	b, _ = json.Marshal(r)
	assert.Equal(t, `"beqpdfdhq87dkncb"`, fmt.Sprintf("%s", b))
}

func TestAddressJSON_CID(t *T) {
	key, err := keypair.NewKeyFromString(testPrivKey)
	assert.Equal(t, nil, err)
	cid, _ := key.CID()
	assert.Equal(t, testKeyCID, cid.String())
	a := FromCID(cid)
	assert.Equal(t, testKeyAddr, a.String())
	b, _ := json.Marshal(a)
	assert.Equal(t, `"zFNScYMHCVTuggbomkvLDPpoLruXKwWEu6SV2S45P9QwuLwhfMeq"`, fmt.Sprintf("%s", b))
	var a2 Address
	json.Unmarshal(b, &a2)
	assert.Equal(t, a.ID, a2.ID)
	assert.Equal(t, a.Extra, a2.Extra)
	assert.Equal(t, a.CID.String(), a2.CID.String())
}
