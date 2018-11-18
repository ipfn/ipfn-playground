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
	"errors"
	"fmt"

	"github.com/cespare/xxhash"
	"github.com/gogo/protobuf/proto"
	"github.com/ipfn/go-base32check"
)

// ID - Cell of an ID registered on-chain.
type ID uint64

// NewID - Creates new cell ID from bytes.
func NewID(body []byte) ID {
	return ID(xxhash.Sum64(body))
}

// NewIDFromString - Creates new cell ID from string.
func NewIDFromString(body string) ID {
	return ID(xxhash.Sum64String(body))
}

// ParseID - Parses cell ID from a string.
func ParseID(body string) (_ ID, err error) {
	if len(body) <= 1 {
		err = errors.New("address too short")
		return
	}
	if body[0] != 'b' {
		err = fmt.Errorf("invalid codec %x", body[0])
		return
	}
	// remove 'b' byte
	body = body[1:]
	s, err := base32check.CheckDecodeString(body)
	if err != nil {
		return
	}
	return DecodeID(s), nil
}

// DecodeID - Decodes cell ID from byte array.
func DecodeID(body []byte) ID {
	id, _ := proto.DecodeVarint(body)
	return ID(id)
}

// Bytes - Returns cell ID in bytes format.
func (id ID) Bytes() []byte {
	return proto.EncodeVarint(uint64(id))
}

// Encode - Returns cell ID in base32 format.
func (id ID) Encode() string {
	body := base32check.CheckEncode(id.Bytes())
	return string(append([]byte{'b'}, body...))
}

// String - Returns cell ID in string format.
func (id ID) String() string {
	name, ok := registry[id]
	if !ok {
		return fmt.Sprintf("%d", id)
	}
	return fmt.Sprintf("%s", name)
}

// MarshalJSON - Marshals ID as JSON string.
func (id ID) MarshalJSON() (_ []byte, err error) {
	name, ok := registry[id]
	if !ok {
		return []byte(fmt.Sprintf("%d", id)), nil
	}
	return []byte(fmt.Sprintf("%q", name)), nil
}
