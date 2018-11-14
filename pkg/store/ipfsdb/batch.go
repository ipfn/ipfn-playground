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

package ipfsdb

import (
	"github.com/ipfn/ipfn/pkg/trie/ethdb"
)

type wrapBatch struct {
	ethdb.Batch
	client *wrapClient
}

func (batch *wrapBatch) Put(key, value []byte) error {
	if err := batch.Batch.Put(key, value); err != nil {
		return err
	}
	return batch.client.Put(value)
}
