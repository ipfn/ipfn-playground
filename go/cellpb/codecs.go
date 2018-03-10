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

import (
	codecs "github.com/ipfn/go-ipfn-cell-codecs"
)

// HeaderPath - Name of Protocol Buffers cell codec.
const HeaderPath = "/cell-pb"

// CodecID - Multicodec header ID of Protocol Buffers Cell Version 1.
// Definition: /ipfs/QmeX5H9x2qNdGC1R5uhyX2HuG5izxR2SGi71jSWyEQjV6Q
const CodecID = 0x70bc // (28860)

func init() {
	codecs.Register(HeaderPath, CodecID, new(Codec))
}
