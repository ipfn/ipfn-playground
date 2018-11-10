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
	"net/http"

	cid "gx/ipfs/QmapdYm1b22Frv3k17fqrBYTFRxwiaVJkB299Mfn33edeB/go-cid"

	"github.com/ipfn/ipfn/src/go/utils/flog"
	shell "github.com/ipfs/go-ipfs-shell"
)

var logger = flog.MustGetLogger("ipfsdb")

type wrapClient struct {
	*shell.Shell
	cid.Prefix
}

func newClient(prefix cid.Prefix, url string) *wrapClient {
	return &wrapClient{
		Shell:  shell.NewShellWithClient(url, http.DefaultClient),
		Prefix: prefix,
	}
}

func (client *wrapClient) Put(value []byte) (err error) {
	if len(value) == 0 {
		return
	}
	cid, err := client.BlockPut(value, cid.CodecToStr[client.Prefix.Codec], mh.Codes[client.Prefix.MhType], client.Prefix.MhLength)
	if err != nil {
		logger.Debugw("IPFS BlockPut", "err", err)
		return
	}
	logger.Debugw("IPFS BlockPut", "cid", cid)
	return
}

func (client *wrapClient) Get(key []byte) (value []byte, err error) {
	mhash, _ := mh.Encode(key, client.Prefix.MhType)
	c := cid.NewCidV1(client.Prefix.Codec, mhash).String()
	logger.Debugw("IPFS BlockGet", "cid", c)
	value, err = client.BlockGet(c)
	return
}
