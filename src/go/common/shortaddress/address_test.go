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

package shortaddress

import (
	"crypto/x509"
	"encoding/json"
	"fmt"
	. "testing"

	"github.com/ipfn/ipfn/src/go/common/pkcid"
	"github.com/ipfn/ipfn/src/go/utils/hexutil"
	"github.com/stretchr/testify/assert"
)

var (
	testSrc  = "beqpdfdhq87dkncb"
	testAddr = &Address{ID: 2191370559816, Extra: 13471}

	testKeyAddr = "bnx37fk4wmxur3j0puapwv"
	testKeyCID  = "zFNScYMHCVTuggbomkvLDPpoLruXKwWEu6SV2S45P9QwuLwhfMeq"
	testPrivKey = "3077020101042080deb3b165f87db1bbd2e5a5eaa33d001efecf37b8af18e1e99489bd7c09c41da00a06082a8648ce3d030107a1440342000470a2ab334e6ba0f9cae349f027f9edc76f89a916a2cfa47bc9dbf5b3582b69fe5d187328f0c862969deccfb282906adb71ade1908fca3da55494570c0a75f320"
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
	pk, err := x509.ParseECPrivateKey(hexutil.FromString(testPrivKey))
	assert.Equal(t, nil, err)
	cid := pkcid.CID(&pk.PublicKey)
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
