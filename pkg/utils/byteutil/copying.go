// Copyright © 2018 The IPFN Developers Authors. All Rights Reserved.
// Copyright © 2014-2018 The go-ethereum Authors. All Rights Reserved.
//
// This file is part of the IPFN project.
// This file was part of the go-ethereum library.
//
// The IPFN project is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The IPFN project is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the IPFN project. If not, see <http://www.gnu.org/licenses/>.

package byteutil

// Copy returns an exact copy of the provided bytes.
func Copy(src []byte) (dst []byte) {
	if src == nil {
		return
	}
	dst = make([]byte, len(src))
	copy(dst, src)
	return
}
