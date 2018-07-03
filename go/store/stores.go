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

// Package store implements simple key-value store and state db.
package store

// BlockStore - Block store interface.
type BlockStore interface {
	// Has - Checks if has key.
	Has([]byte) (bool, error)
	// Get - Gets raw value.
	Get([]byte) ([]byte, error)
	// Put - Puts raw value.
	Put([]byte, []byte) error
}

// RawStore - Key-value raw store interface.
type RawStore interface {
	// Keys - Returns all keys.
	Keys() ([]string, error)
	// Has - Checks if has key.
	Has(string) (bool, error)
	// Get - Gets raw value.
	Get(string) ([]byte, error)
	// Put - Puts raw value.
	Put(string, []byte) error
}

// EncodedStore - Key-value store interface.
type EncodedStore interface {
	// Keys - Returns all keys.
	Keys() ([]string, error)
	// Has - Checks if has key.
	Has(string) (bool, error)
	// Get - Gets raw encrypted key.
	Get(string, interface{}) error
	// Put - Puts raw encrypted key.
	Put(string, interface{}) error
}
