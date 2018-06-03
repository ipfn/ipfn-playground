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
	"fmt"

	prompt "github.com/segmentio/go-prompt"
	"github.com/spf13/cobra"

	bip39 "github.com/ipfn/go-bip39"
	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	crypto "github.com/ipfn/ipfn/go/crypto"

	"github.com/crackcomm/viperkeys"
)

var (
	createName string
	createSize int
)

func init() {
	RootCmd.AddCommand(CreateCmd)
	CreateCmd.PersistentFlags().StringVarP(&createName, "name", "n", "", "Name of the seed")
	CreateCmd.PersistentFlags().IntVarP(&createSize, "size", "s", 32, "Size of seed")
}

// CreateCmd - Seed create command.
var CreateCmd = &cobra.Command{
	Use:         "create [name]",
	Short:       "Generates random seed",
	Annotations: map[string]string{"category": "seed"},
	Args: func(cmd *cobra.Command, args []string) error {
		if createName == "" && len(args) == 1 {
			createName = args[0]
		}
		return nil
	},
	Run: cmdutil.WrapCommand(HandleCreateCmd),
}

// HandleCreateCmd - Handles seed create command.
func HandleCreateCmd(cmd *cobra.Command, args []string) (err error) {
	// ask for password with confirmation
	password := cmdutil.PromptPasswordRepeated("seed password")
	// generate entropy
	entropy, err := crypto.NewEntropy(createSize)
	if err != nil {
		return fmt.Errorf("failed to generate entropy: %v", err)
	}
	// convert entropy to mnemonic
	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return fmt.Errorf("failed to create mnemonic: %v", err)
	}
	// print mnemonic
	logger.Printf("Mnemonic: %s", mnemonic)
	// return if we don't want to save
	if !prompt.Confirm("Do you want to save the seed?") {
		return
	}
	// Ask for *unique* name
	var has bool
	if createName != "" {
		has, err = viperkeys.Default.Has(createName)
		if err != nil {
			return fmt.Errorf("failed to read keystore: %v", err)
		}
	}
	if createName == "" || has {
		createName = cmdutil.PromptConfirmed("seed name", func(name string) bool {
			has, _ := viperkeys.Default.Has(name)
			return !has
		})
	}
	return viperkeys.Default.CreateKey(createName, mnemonic, password)
}
