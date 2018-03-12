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

package keywallet

import (
	"fmt"
	"hash/crc32"
	"strconv"
	"strings"
)

// HashPath - Creates custom derivation path from bytes of crc32 hash.
func HashPath(path string) string {
	// crc32 checksum of path
	h := crc32.NewIEEE()
	h.Write([]byte(path))
	s := h.Sum(nil)
	// create derivation path string
	r := []string{"m", "112'"}
	// iterate over crc32 bytes
	for n, v := range s {
		if n < 2 {
			r = append(r, fmt.Sprintf("%d'", v))
		} else {
			r = append(r, strconv.Itoa(int(v)))
		}
	}
	return strings.Join(r, "/")
}
