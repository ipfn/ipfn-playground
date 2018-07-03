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
	"fmt"

	"github.com/spf13/cobra"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	"github.com/ipfn/ipfn/go/wallet"
)

var (
	newPrintSeed bool
)

func init() {
	RootCmd.AddCommand(CreateCmd)
	CreateCmd.PersistentFlags().BoolVarP(&newPrintSeed, "print", "p", false, "Print mnemonic")
}

// CreateCmd - Seed new command.
var CreateCmd = &cobra.Command{
	Use:         "create [name]",
	Short:       "Create new seed",
	Annotations: map[string]string{"category": "wallet"},
	Args:        CheckCreateArgs,
	Run:         cmdutil.WrapCommand(HandleCreateCmd),
}

// HandleCreateCmd - Handles seed new command.
func HandleCreateCmd(cmd *cobra.Command, args []string) (err error) {
	name := cmdutil.ArgDefault(args, 0, "default")
	password := cmdutil.PromptPasswordRepeated(fmt.Sprintf("%q wallet password", walletName))
	seed, err := wallet.NewDefault().CreateSeed(name, []byte(password))
	if err != nil {
		return
	}
	if newPrintSeed {
		logger.Printf("Mnemonic: %s", seed)
	}
	return
}

// CheckCreateArgs - Checks if wallet already exists.
func CheckCreateArgs(cmd *cobra.Command, args []string) (err error) {
	name := cmdutil.ArgDefault(args, 0, "default")
	has, err := wallet.NewDefault().KeyExists(name)
	if err != nil {
		return err
	}
	if has {
		return fmt.Errorf("%q wallet already exists", name)
	}
	return nil
}
