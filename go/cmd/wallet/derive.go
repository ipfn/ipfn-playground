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

package wallet

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	"github.com/ipfn/ipfn/go/dev/address"
	"github.com/ipfn/ipfn/go/keypair"
	"github.com/ipfn/ipfn/go/wallet"
)

var (
	hashPath    bool
	printInfo   bool
	keyPath     string
	derivePath  string
	accountName string
	walletName  string
)

func init() {
	RootCmd.AddCommand(DeriveCmd)
	DeriveCmd.PersistentFlags().BoolVarP(&hashPath, "hash", "x", false, "derive hash path")
	DeriveCmd.PersistentFlags().BoolVarP(&printInfo, "print", "p", true, "prints derivation info")
	DeriveCmd.PersistentFlags().StringVarP(&keyPath, "key-path", "k", "", "wallet key path (<wallet>/<x|m>/<path>)")
	DeriveCmd.PersistentFlags().StringVarP(&derivePath, "derive-path", "d", "", "derive BIP32 hierarchical path")
	DeriveCmd.PersistentFlags().StringVarP(&walletName, "wallet", "w", "default", "wallet name")
}

// DeriveCmd - Key derive command.
var DeriveCmd = &cobra.Command{
	Use:   "derive",
	Short: "Derive from seed",
	Long: `Derives key from seed and path or mnemonic name.

Path is defined as: "/<purpose>'/<coin_type>'/<account>'/<change>/<address_index>".

Mnemonic can be used for path by using --hash or -x flag.`,
	Example: `  $ ipfn wallet derive -xd memo
  $ ipfn wallet derive -k wallet/x/memo
  $ ipfn wallet derive -w wallet -xd memo
  $ ipfn wallet derive -d m/44'/138'/0'/0/0`,
	Annotations: map[string]string{"category": "wallet"},
	Args:        CheckDeriveArgs,
	Run:         cmdutil.WrapCommand(HandleDeriveCmd),
}

// HandleDeriveCmd - Handles key derive command.
func HandleDeriveCmd(cmd *cobra.Command, args []string) (err error) {
	acc, err := deriveWallet()
	if err != nil {
		return
	}
	if printInfo {
		neuter, _ := acc.Neuter()
		c, _ := acc.CID()
		a := address.FromCID(c)
		logger.Print()
		logger.Printf("Short:       %s", a)
		logger.Printf("Address:     %s", c)
		logger.Printf("Public key:  %s", neuter)
		logger.Printf("Private key: %s", acc.PrivateString())
	}
	return
}

// CheckDeriveArgs - Checks derivation path.
func CheckDeriveArgs(cmd *cobra.Command, args []string) (err error) {
	if hashPath && strings.Contains(derivePath, "/") {
		return fmt.Errorf("hash path %q cannot contain \"/\"", derivePath)
	}
	return nil
}

func deriveWallet() (_ *keypair.KeyPair, err error) {
	var path *wallet.KeyPath
	if keyPath != "" {
		path, err = wallet.ParseKeyPath(keyPath)
	} else if derivePath != "" {
		path = wallet.NewKeyPath(walletName, derivePath, hashPath)
	} else {
		err = errors.New("derivation path cannot be empty")
	}
	if err != nil {
		return
	}
	return wallet.PromptDeriveKey(path)
}
