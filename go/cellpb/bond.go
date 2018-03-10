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

package cellpb

import cell "github.com/ipfn/go-ipfn-cell"

// NewBond - Creates new protocol buffers bond.
func NewBond(bond cell.Bond) (res *Bond) {
	switch msg := bond.(type) {
	case *BondWrapper:
		return msg.bond
	default:
		return &Bond{
			Name: bond.Name(),
			Kind: bond.Kind(),
			From: bond.From(),
			To:   bond.To(),
		}
	}
}

// NewBondWrapper - Creates a new bond from protocol buffers message.
func NewBondWrapper(bond *Bond) *BondWrapper {
	return &BondWrapper{bond: bond}
}

// BondWrapper - Protocol buffers cell wrapper.
type BondWrapper struct {
	bond *Bond
}

// Name - Returns bond name.
func (wrapper *BondWrapper) Name() string {
	return wrapper.bond.Name
}

// Kind - Returns bond kind.
func (wrapper *BondWrapper) Kind() string {
	return wrapper.bond.Kind
}

// From - Returns bond from.
func (wrapper *BondWrapper) From() string {
	return wrapper.bond.From
}

// To - Returns bond to.
func (wrapper *BondWrapper) To() string {
	return wrapper.bond.To
}
