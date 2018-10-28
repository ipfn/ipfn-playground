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

package core

import (
	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	"github.com/ipfn/ipfn/go/cmd/wallet"
	"github.com/spf13/cobra"
)

// RegisterCommands - Registers core commands.
func RegisterCommands(root *cobra.Command) {
	root.AddCommand(InitCmd)
}

// InitCmd - Init command.
var InitCmd = &cobra.Command{
	Use:         "init [wallet]",
	Short:       "Initialize config",
	Annotations: map[string]string{"category": "init"},
	Run:         cmdutil.WrapCommand(HandleInitCmd),
}

// HandleInitCmd - Handles init command.
func HandleInitCmd(cmd *cobra.Command, args []string) (err error) {
	if wallet.CheckCreateArgs(cmd, args) == nil {
		err = wallet.HandleCreateCmd(cmd, args)
		if err != nil {
			return
		}
	} else {
		logger.Print("Wallet already exists...")
	}

	return
}
