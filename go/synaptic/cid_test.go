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

package synaptic

import (
	. "testing"

	"github.com/stretchr/testify/assert"

	cellcid "github.com/ipfn/ipfn/go/cellcid"

	// Register `cell-pb-v1` codec
	_ "github.com/ipfn/ipfn/go/cellpb"
)

func TestCellCID(t *T) {
	c := NewCell(String, []byte("test"))
	cid, err := cellcid.CID(c)
	assert.Equal(t, err, nil)
	assert.Equal(t, cid.String(), "zFunckyav7J6aWMDYJMMTcQXh1hJowx3GD8RwvZjmGBoXg5HsCgM")
}
