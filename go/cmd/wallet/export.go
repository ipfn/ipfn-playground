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

	prompt "github.com/segmentio/go-prompt"
	"github.com/spf13/cobra"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	"github.com/ipfn/ipfn/go/keypair"
	"github.com/ipfn/ipfn/go/wallet"
)

func init() {
	RootCmd.AddCommand(ExportCmd)
}

// ExportCmd - Seed export command.
var ExportCmd = &cobra.Command{
	Use:         "export [name]",
	Short:       "Export master key",
	Annotations: map[string]string{"category": "wallet"},
	Args:        checkExistsArgs,
	Run:         cmdutil.WrapCommand(HandleExportCmd),
}

// HandleExportCmd - Handles seed export command.
func HandleExportCmd(cmd *cobra.Command, args []string) (err error) {
	name := cmdutil.ArgDefault(args, 0, "default")
	w := wallet.NewDefault()
	has, err := w.KeyExists(name)
	if err != nil {
		return fmt.Errorf("failed to read keystore: %v", err)
	}
	if !has {
		return fmt.Errorf("seed %q was not found", name)
	}
	password := prompt.PasswordMasked("Encryption password")
	if password == "" {
		return errors.New("failed to get encryption password")
	}
	key, err := w.ExportKey(name)
	if err != nil {
		return
	}
	seed, err := key.Decrypt([]byte(password))
	if key.SeedType == keypair.Mnemonic {
		logger.Printf("Mnemonic: %s", seed)
	} else {
		logger.Printf("Private key: %s", seed)
	}
	return
}

func checkExistsArgs(cmd *cobra.Command, args []string) (err error) {
	name := cmdutil.ArgDefault(args, 0, "default")
	has, err := wallet.NewDefault().KeyExists(name)
	if err != nil {
		return err
	}
	if !has {
		return fmt.Errorf("%q wallet does not exist", name)
	}
	return nil
}
