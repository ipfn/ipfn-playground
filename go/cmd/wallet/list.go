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
	"github.com/ipfn/ipfn/go/wallet"
)

func init() {
	RootCmd.AddCommand(ListCmd)
	// ListCmd.PersistentFlags().BoolVarP(&printKey, "print-key", "p", false, "Prints public keys")
}

// ListCmd - Key list command.
var ListCmd = &cobra.Command{
	Use:         "list",
	Short:       "List master keys",
	Annotations: map[string]string{"category": "key"},
	Run:         cmdutil.WrapCommand(HandleListCmd),
}

// HandleListCmd - Handles key list command.
func HandleListCmd(cmd *cobra.Command, args []string) (err error) {
	names, err := wallet.NewDefault().KeyNames()
	if err != nil {
		return
	}

	for _, name := range names {
		fmt.Println(name)
	}

	return
}
