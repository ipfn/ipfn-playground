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

package cells

import (
	. "testing"

	"github.com/stretchr/testify/assert"
)

func TestComputeID(t *T) {
	id, err := cidCell.Checksum()
	assert.Equal(t, nil, err)
	assert.Equal(t, ID(0x13fb93194e573bd2), id)
}

func TestParseID(t *T) {
	var (
		nID = ID(216547582402989)
		src = "b4hp7dnrwncce5"
		num = "216547582402989"
	)

	assert.Equal(t, num, nID.String())
	assert.Equal(t, src, nID.Encode())

	id, err := ParseID(src)
	assert.Equal(t, nil, err)
	assert.Equal(t, nID, id)
}
