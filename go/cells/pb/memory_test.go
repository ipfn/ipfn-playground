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

package pb

import (
	"fmt"
	math "math"

	. "testing"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/any"
)

func TestIntoMemory(t *T) {
	marshal(t, true)
	marshal(t, false)
	marshal(t, "test")
	marshal(t, int(314159))
	marshal(t, int16(math.MaxInt16))
	marshal(t, int32(314159))
	marshal(t, int64(314159))
	marshal(t, int8(math.MaxInt8))
	marshal(t, uint8(math.MaxUint8))
	marshal(t, uint16(math.MaxUint16))
	marshal(t, uint32(math.MaxUint32))
	marshal(t, float32(math.Pi))
	marshal(t, float64(math.Pi))
	marshal(t, &Cell{Memory: &any.Any{Value: []byte{'0'}}})
	marshal(t, ptypes.TimestampNow())
	marshal(t, &ptypes.Empty{})
	marshal(t, nil)
}

// TODO(crackcomm): test marshal/unmarshal
func marshal(t *T, value interface{}) {
	a, err := PackMemory(value)
	if err != nil {
		t.Error(err)
	}
	mem, err := UnpackMemory(a)
	if err != nil {
		t.Errorf("value: %#v error: %q", value, err)
	}
	fmt.Printf("message: %#v\n", mem)
}
