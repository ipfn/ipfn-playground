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

// Package keystore implements cryptographic key store.
package keystore

// RawStorage - Key-value raw storage interface.
type RawStorage interface {
	// Get - Gets raw encrypted key.
	Get(string) ([]byte, error)
	// Put - Puts raw encrypted key.
	Put(string, []byte) error
}

// Storage - Key-value storage interface.
type Storage interface {
	// Get - Gets encrypted key.
	Get(string) (*EncryptedKey, error)
	// Put - Puts encrypted key.
	Put(string, *EncryptedKey) error
}
