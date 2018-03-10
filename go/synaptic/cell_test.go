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

	cell "github.com/ipfn/ipfn/go/cell"
)

func TestInterface(t *T) {
	var c cell.Cell
	c = NewCell(String, []byte("test"))
	assert.Equal(t, c.Soul(), "/synaptic/string")
	bytes, err := c.Memory().Bytes()
	assert.Equal(t, err, nil)
	assert.Equal(t, bytes, []byte("test"))
}
