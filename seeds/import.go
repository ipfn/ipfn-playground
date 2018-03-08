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

package seeds

import (
	"errors"
	"fmt"

	prompt "github.com/segmentio/go-prompt"
	"github.com/spf13/cobra"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-key-wallet"

	"github.com/crackcomm/viperkeys"
)

func init() {
	RootCmd.AddCommand(ImportCmd)
}

// ImportCmd - Seed import command.
var ImportCmd = &cobra.Command{
	Use:         "import [name]",
	Short:       "Imports existing seed",
	Annotations: map[string]string{"category": "seed"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) >= 1 && viperkeys.Default.Has(args[0]) {
			return fmt.Errorf("seed %q already exists", args[0])
		}
		return nil
	},
	Run: cmdutil.WrapCommand(HandleImportCmd),
}

// HandleImportCmd - Handles seed import command.
func HandleImportCmd(cmd *cobra.Command, args []string) (err error) {
	mnemonic := cmdutil.PromptConfirmed("mnemonic seed", keywallet.IsMnemonicValid)
	password := prompt.PasswordMasked("Seed password")
	if password == "" {
		return errors.New("failed to get password")
	}
	// Ask for *unique* name
	name := args[0]
	if viperkeys.Default.Has(name) || name == "" {
		name = cmdutil.PromptConfirmed("seed name", func(name string) bool {
			return !viperkeys.Default.Has(name)
		})
	}
	return viperkeys.Default.CreateKey(name, mnemonic, password)
}
