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

package cells

// Cell - Core cell interface.
type Cell interface {
	// Name - Returns name of the cell.
	Name() string

	// Soul - Returns name of the cell soul.
	Soul() string

	// Body - Returns body of the cell.
	Body() []Cell

	// Bonds - Returns bonds of the cell.
	Bonds() []string

	// Memory - Returns memory of the cell.
	Memory() interface{}
}
