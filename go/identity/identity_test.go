// Copyright © 2017 The IPFN Authors. All Rights Reserved.
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

package identity

import (
	"encoding/json"
	"fmt"

	. "testing"
)

func ExampleNewSafe() {
	keys, err := NewSafe(RSA, 2048)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Public ID: %s", keys.ID)
}

func TestNew(t *T) {
	keys, err := NewSafe(RSA, 2048)
	if err != nil {
		t.Error(err)
	}
	if len(keys.ID) == 0 {
		t.Fail()
	}
}

func TestJSON(t *T) {
	keys, err := NewSafe(RSA, 2048)
	if err != nil {
		t.Error(err)
	}
	body, err := json.Marshal(keys)
	if err != nil {
		t.Error(err)
	}
	var res Identity
	err = json.Unmarshal(body, &res)
	if err != nil {
		t.Error(err)
	}
	if keys.ID.Pretty() != res.ID.Pretty() {
		t.Fail()
	}
	// TODO we can test more
}
