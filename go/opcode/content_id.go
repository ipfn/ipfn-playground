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

package opcode

import (
	"fmt"

	cid "github.com/ipfs/go-cid"
)

// CID - Content ID wrapper.
type CID struct {
	*cid.Cid
}

// SumCID - Sums content id and wraps.
func SumCID(prefix cid.Prefix, body []byte) (_ *CID, err error) {
	c, err := prefix.Sum(body)
	if err != nil {
		return
	}
	return WrapCID(c), nil
}

// WrapCID - Wraps content id.
func WrapCID(c *cid.Cid) *CID {
	return &CID{Cid: c}
}

// DecodeCID - Parses CID.
func DecodeCID(v string) (_ *CID, err error) {
	c, err := cid.Decode(v)
	if err != nil {
		return
	}
	return &CID{Cid: c}, nil
}

// UnmarshalJSON - Parses the JSON string representation of a cid.
func (c *CID) UnmarshalJSON(b []byte) (err error) {
	if len(b) < 2 {
		return fmt.Errorf("invalid cid json blob")
	}
	b = b[1 : len(b)-1]
	c.Cid, err = cid.Decode(string(b))
	return
}

// MarshalJSON - Marshals the cid as string.
func (c *CID) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("%q", c)), nil
}
