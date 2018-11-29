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

package x25519

import (
	"fmt"
	"testing"

	"github.com/ipfn/ipfn/pkg/digest"

	"github.com/stretchr/testify/assert"
)

func TestX25519(t *testing.T) {
	var sk1 [32]byte = digest.FromHex("1b4f0e9851971998e732078544c96b36c3d01cedf7caa332359d6f1d83567014") // sha256.Sum256([]byte("test1"))
	var sk2 [32]byte = digest.FromHex("60303ae22b998861bce3b28f33eec1be758a213c86c93c076dbe9f558c11c752") // sha256.Sum256([]byte("test2"))
	pk1 := Public(&sk1)
	pk2 := Public(&sk2)
	assert.Equal(t, "4652486ebc271520d844e5bdda9ac243c05dcbe7bc9b93807073a32177a6f73d", fmt.Sprintf("%x", pk1))
	assert.Equal(t, "ffbc7ba2e4c43be03f8a7f020d0651f582ad1901c254eebb4ec2ecb73148e50d", fmt.Sprintf("%x", pk2))
	shk1 := Shared(&sk1, &pk2)
	shk2 := Shared(&sk2, &pk1)
	assert.Equal(t, "42dedd506f22f8bbe71c2dbfc31e50e2db53861a6f55a2cc77e07e4e271f9807", fmt.Sprintf("%x", shk1))
	assert.Equal(t, "42dedd506f22f8bbe71c2dbfc31e50e2db53861a6f55a2cc77e07e4e271f9807", fmt.Sprintf("%x", shk2))
	assert.Equal(t, "5b5bb68b4ec37023cf216e71391e8411547d41545c6efc4f9d746c96f2a43751", fmt.Sprintf("%x", Public(&shk1)))
}
