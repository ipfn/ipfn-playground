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

package fsdb

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ipfn/ipfn/src/go/store"
)

func Example() {
	root, _ := ioutil.TempDir("", "fsdb_")
	defer os.RemoveAll(root)

	db := Open(NewDefaultOptions(root).SetUseGzip(true))
	key := store.Key("name")
	ctx := context.Background()

	db.Write(ctx, key, strings.NewReader("Anakin Skywalker"))
	reader, err := db.Read(ctx, key)
	if err != nil {
		// TODO: handle error
	}
	name, err := ioutil.ReadAll(reader)
	reader.Close()
	if err != nil {
		// TODO: handle error
	}
	fmt.Println(string(name))

	db.Write(ctx, key, strings.NewReader("Darth Vader"))
	reader, err = db.Read(ctx, key)
	if err != nil {
		// TODO: handle error
	}
	name, err = ioutil.ReadAll(reader)
	reader.Close()
	if err != nil {
		// TODO: handle error
	}
	fmt.Println(string(name))

	db.Delete(ctx, key)
	_, err = db.Read(ctx, key)
	if store.IsNoSuchKeyError(err) {
		fmt.Println("Joined force")
	}

	// Output:
	// Anakin Skywalker
	// Darth Vader
	// Joined force
}
