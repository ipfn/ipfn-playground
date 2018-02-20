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

// Kind - Synaptic type.
// This enum is internal representation only.
type Kind int8

const (
	// String - Synaptic string. Soul address: `/synaptic/string`.
	String Kind = iota
	// Number - Synaptic number. Soul address: `/synaptic/number`.
	Number Kind = iota
	// Bool - Synaptic bool. Soul address: `/synaptic/bool`.
	Bool Kind = iota
)

// KindFromString - Returns synaptic kind from string.
func KindFromString(kind string) (_ Kind, _ bool) {
	switch kind {
	case "string":
		return String, true
	case "number":
		return Number, true
	case "bool":
		return Bool, true
	default:
		return
	}
}

// Soul - Returns address of the cell soul.
func (kind Kind) Soul() string {
	switch kind {
	case String:
		return "/synaptic/string"
	case Number:
		return "/synaptic/number"
	case Bool:
		return "/synaptic/bool"
	default:
		panic("unknown kind")
	}
}
