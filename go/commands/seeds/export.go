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
	"github.com/ipfn/go-ipfn-cmd-util/logger"

	"github.com/crackcomm/viperkeys"
)

func init() {
	RootCmd.AddCommand(ExportCmd)
}

// ExportCmd - Seed export command.
var ExportCmd = &cobra.Command{
	Use:         "export [name]",
	Short:       "Prints seed mnemonic",
	Annotations: map[string]string{"category": "seed"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("name argument is required")
		}
		has, err := viperkeys.Default.Has(args[0])
		if err != nil {
			return fmt.Errorf("failed to read keystore: %v", err)
		}
		if !has {
			return fmt.Errorf("seed %q was not found", args[0])
		}
		return nil
	},
	Run: cmdutil.WrapCommand(HandleExportCmd),
}

// HandleExportCmd - Handles seed export command.
func HandleExportCmd(cmd *cobra.Command, args []string) (err error) {
	password := prompt.PasswordMasked("Encryption password")
	if password == "" {
		return errors.New("failed to get encryption password")
	}
	mnemonic, err := viperkeys.Default.Decrypt(args[0], []byte(password))
	if err != nil {
		return
	}
	logger.Printf("Mnemonic: %s", mnemonic)
	return
}
