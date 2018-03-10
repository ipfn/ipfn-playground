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

package cell

import cid "github.com/ipfs/go-cid"

// Cell - Core cell interface.
type Cell interface {
	// CID - Returns content ID of the cell.
	CID() (*cid.Cid, error)

	// Name - Returns name of the cell.
	Name() string

	// Soul - Returns address of the cell soul.
	Soul() string

	// Bonds - Returns bonds of the cell.
	Bonds() []Bond

	// Body - Returns body of the cell.
	Body() []Cell

	// Memory - Returns memory of the cell.
	Memory() Memory
}
