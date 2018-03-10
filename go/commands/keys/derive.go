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
	"encoding/hex"
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	"github.com/ipfn/ipfn/go/pubkeyhash"
)

var (
	forcePath     bool
	customSeedPwd bool
	keyAddrID     string
)

func init() {
	RootCmd.AddCommand(DeriveCmd)
	DeriveCmd.PersistentFlags().BoolVarP(&forcePath, "force", "f", false, "Force derivation path")
	DeriveCmd.PersistentFlags().BoolVarP(&customSeedPwd, "custom", "u", false, "Use ustom seed derivation password")
	DeriveCmd.PersistentFlags().StringVarP(&keyAddrID, "addr", "a", "0x0", "Custom pubkeyhash address network ID")
}

// DeriveCmd - Key derive command.
var DeriveCmd = &cobra.Command{
	Use:         "derive [seed] [path]",
	Short:       "Derives key from seed",
	Annotations: map[string]string{"category": "key"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("seed argument is required")
		}
		if viper.Get(fmt.Sprintf("seeds.%s", args[0])) == nil {
			return fmt.Errorf("seed %q was not found", args[0])
		}
		if len(args) < 2 {
			return errors.New("path argument is required")
		}
		return nil
	},
	Run: cmdutil.WrapCommand(HandleDeriveCmd),
}

// HandleDeriveCmd - Handles key derive command.
func HandleDeriveCmd(cmd *cobra.Command, args []string) (err error) {
	acc, err := deriveKey(args[0], args[1])
	if err != nil {
		return
	}
	pubkey, err := acc.ECPubKey()
	if err != nil {
		return
	}
	id, err := hexToByte(keyAddrID)
	if err != nil {
		return
	}
	addr, err := pubkeyhash.PKHash(pubkey, id)
	if err != nil {
		return
	}
	logger.Print(addr.String())
	return
}

// hexToByte - Converts hex string to byte.
func hexToByte(input string) (_ byte, err error) {
	input = strings.TrimLeft(input, "0x")
	arr, err := hex.DecodeString(input)
	if err != nil {
		return
	}
	if len(arr) != 1 {
		return 0, nil
	}
	return arr[0], nil
}
