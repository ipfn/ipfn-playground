// Copyright © 2017-2018 The IPFN Developers. All Rights Reserved.
// Copyright © 2017-2018 The go-ethereum Authors. All Rights Reserved.
//
// This file is part of the IPFN Project.
// This file was part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

// +build ios linux,arm64 windows !darwin,!freebsd,!linux,!netbsd,!solaris

// This is the fallback implementation of directory watching.
// It is used on unsupported platforms.

package keystore

type watcher struct{ running bool }

func newWatcher(*accountCache) *watcher { return new(watcher) }
func (*watcher) start()                 {}
func (*watcher) close()                 {}
