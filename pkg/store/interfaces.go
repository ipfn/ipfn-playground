// Copyright © 2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2017 Yuxuan 'fishy' Wang. All Rights Reserved.
//
// Redistribution and use in source and binary forms, with or without
// modification, are permitted provided that the following conditions are
// met:
//
//    * Redistributions of source code must retain the above copyright
// notice, this list of conditions and the following disclaimer.
//    * Redistributions in binary form must reproduce the above
// copyright notice, this list of conditions and the following disclaimer
// in the documentation and/or other materials provided with the
// distribution.
//    * Neither the copyright holder nor the names of its
// contributors may be used to endorse or promote products derived from
// this software without specific prior written permission.
//
// THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
// "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
// LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
// A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
// OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
// SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
// LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
// DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
// THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
// (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
// OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.

package store

import (
	"context"
	"io"
)

// Reader - Store reader.
//
// Source: https://github.com/fishy/fsdb/blob/44ba08f/fsdb.go#L9
type Reader interface {
	// Read opens an entry and returns a ReadCloser.
	//
	// If the key does not exist, it should return a NoSuchKeyError.
	//
	// It should never return both nil reader and nil err.
	//
	// It's the caller's responsibility to close the ReadCloser returned.
	Read(ctx context.Context, key Key) (reader io.ReadCloser, err error)

	// Write opens an entry.
	//
	// If the key already exists, it will be overwritten.
	//
	// If data is actually a ReadCloser,
	// it's the caller's responsibility to close it after Write function returns.
	Write(ctx context.Context, key Key, data io.Reader) error

	// Delete deletes an entry.
	//
	// If the key does not exist, it should return a NoSuchKeyError.
	Delete(ctx context.Context, key Key) error
}

// Scanner - Store scanner.
//
// Source: https://github.com/fishy/fsdb/blob/44ba08f/fsdb.go#L37
type Scanner interface {
	// ScanKeys scans all the keys locally.
	//
	// This function would be heavy on IO and takes a long time. Use with caution.
	//
	// The behavior is undefined for keys changed after the scan started,
	// but it should never visit the same key twice in a single scan.
	ScanKeys(ctx context.Context, keyFunc KeyFunc, errFunc ErrFunc) error
}

// KeyFunc is used in ScanKeys function in Scanner interface.
//
// It's the callback function called for every key scanned.
//
// It should return true to continue the scan and false to abort the scan.
//
// It's OK for KeyFunc to block.
type KeyFunc func(key Key) bool

// ErrFunc is used in ScanKeys function in Scanner interface.
//
// It's the callback function called when the scan encounters an I/O error that
// is possible to be ignored.
//
// It should return true to ignore the error, or false to abort the scan.
type ErrFunc func(path string, err error) bool

// StopAll is an ErrFunc that can be used in Local.ScanKeys().
//
// It always returns false,
// means that the scan stops at the first I/O error it encounters.
func StopAll(path string, err error) bool {
	return false
}

// IgnoreAll is an ErrFunc that can be used in Local.ScanKeys().
//
// It always returns true,
// means that the scan ignores all I/O errors it could ignore.
func IgnoreAll(path string, err error) bool {
	return true
}
