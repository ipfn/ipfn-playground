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

package store

import (
	"context"
)

// Bytes - Byte reader.
type Bytes interface {
	// Get - Reads value from under key.
	//
	// If the key does not exist, it should return a NoSuchKeyError.
	Get(ctx context.Context, key Key) ([]byte, error)

	// Set - Writes value under key.
	//
	// If the key does not exist, it should return a NoSuchKeyError.
	Set(ctx context.Context, key Key, body []byte) error
}

// Simple - Byte reader.
type Simple interface {
	// Get - Reads value from under key.
	//
	// If the key does not exist, it should return a NoSuchKeyError.
	Get(key Key) ([]byte, error)

	// Set - Writes value under key.
	//
	// If the key does not exist, it should return a NoSuchKeyError.
	Set(key Key, body []byte) error
}
