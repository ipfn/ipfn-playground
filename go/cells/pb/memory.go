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
	"github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes"
	any "github.com/golang/protobuf/ptypes/any"
	"github.com/golang/protobuf/ptypes/duration"
	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/golang/protobuf/ptypes/timestamp"
)

// PackMemory - Turns interface{} into a serialized `google.protobuf.Any` structure.
// TODO(crackcomm): it only works on primitive types and nil
func PackMemory(value interface{}) (*any.Any, error) {
	if value == nil {
		return &any.Any{TypeUrl: "∅"}, nil
	}
	res, ok, err := primitiveToProtoBuf(value)
	if ok {
		return res, err
	}
	msg, ok := value.(proto.Message)
	if ok {
		return intoAny(msg)
	}
	return nil, nil
}

// UnpackMemory - Unpacks serialized protocol buffers message.
func UnpackMemory(a *any.Any) (msg proto.Message, err error) {
	switch a.TypeUrl {
	case "∅":
		return &structpb.Value{Kind: &structpb.Value_NullValue{NullValue: structpb.NullValue_NULL_VALUE}}, nil
	case "⚛":
		msg = new(Cell)
	case "⚖":
		msg = new(structpb.Value)
	case "⌚":
		msg = new(duration.Duration)
	case "📅":
		msg = new(timestamp.Timestamp)
	default:
		// BUG(crackcomm):
		// this implies a.TypeUrl is a protobuf...WRONG!
		// memory could be anything...

		// Cell memory field is not like any other any field.
		// It already has some packed typing system (see above).
		// Packed memory field could come from node.js implementation.
		// Or even worse from a browser or MongoDB - or whatever in that case,
		// and it could be a serialized JSON document or even some extra codec.
		// If we would not have a TypeUrl field we would use multicodec,
		// but in this case we are using it as a multiaddr of codecs.
		// Simple examples would be: `/json`, `/pb`, `/msgp` etc.
		//
		// This "complicated" to read cell memory is a possibility,
		// and is treated as exception. It is not implemented in the core.
		// It is a way of allowing a freedom of expression and innovation.
		//
		// Some codecs can use unicode prefixes, otherwise compatibility with:
		// https://github.com/multiformats/multicodec/pull/63
		// https://github.com/multiformats/multicodec/issues/64
		// If both above will get merged prefixes will be multicodec compatible.

		// TODO(crackcomm):
		// this part of app we are unpacking memory...
		// so we will either read type url (protobuf, json or else)
		// or we will return error - no codec available.
		// for now its ok

		var res ptypes.DynamicAny
		err = ptypes.UnmarshalAny(a, &res)
		if err != nil {
			return
		}
		return res.Message, nil
	}
	err = proto.Unmarshal(a.Value, msg)
	return
}

// intoAny - Converts protocol buffers message into serialized Any struct.
// Any message type URL is internally mapped using `TypeURL` function.
func intoAny(msg proto.Message) (*any.Any, error) {
	value, err := proto.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return &any.Any{TypeUrl: TypeURL(msg), Value: value}, nil
}

// TODO(crackcomm): unimplemented cases:
// case complex64:
// case complex128:
// case uintptr:
// case uint64:
// case uint:
func primitiveToProtoBuf(value interface{}) (*any.Any, bool, error) {
	switch val := value.(type) {
	case bool:
		res, err := intoAny(&structpb.Value{Kind: &structpb.Value_BoolValue{BoolValue: val}})
		return res, true, err
	case string:
		res, err := intoAny(&structpb.Value{Kind: &structpb.Value_StringValue{StringValue: val}})
		return res, true, err
	case int:
		return numberValue(float64(val))
	case int16:
		return numberValue(float64(val))
	case int32:
		return numberValue(float64(val))
	case int64:
		return numberValue(float64(val))
	case int8:
		return numberValue(float64(val))
	case uint8:
		return numberValue(float64(val))
	case uint16:
		return numberValue(float64(val))
	case uint32:
		return numberValue(float64(val))
	case float32:
		return numberValue(float64(val))
	case float64:
		return numberValue(float64(val))
	default:
		return nil, false, nil
	}
}

func numberValue(n float64) (*any.Any, bool, error) {
	res, err := intoAny(&structpb.Value{Kind: &structpb.Value_NumberValue{NumberValue: n}})
	return res, true, err
}
