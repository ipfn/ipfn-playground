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

package keys

import (
	"errors"

	prompt "github.com/segmentio/go-prompt"
	"github.com/spf13/cobra"

	"github.com/crackcomm/viperkeys"

	"github.com/ipfn/go-ipfn-cmd-util/logger"
	keywallet "github.com/ipfn/ipfn/go/keywallet"
)

// RootCmd - Root key RootCmd.
var RootCmd = &cobra.Command{
	Use:         "key",
	Short:       "Keys commands",
	Annotations: map[string]string{"category": "key"},
}

// uses global `forcePath` variable
func deriveKey(name, path string) (_ *keywallet.ExtendedKey, err error) {
	wallet := keywallet.New(viperkeys.Default)
	password := prompt.PasswordMasked("Decryption password")
	if password == "" {
		return nil, errors.New("failed to get decryption password")
	}
	if forcePath {
		path = keywallet.HashPath(path)
		logger.Printf("Derive path: /%s", path)
	}
	return wallet.Derive(name, path, []byte(password))
}
