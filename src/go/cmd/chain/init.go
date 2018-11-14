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

package chain

import (
	"github.com/ipfn/ipfn/src/go/chain/dev/contents"
	"github.com/spf13/cobra"

	"github.com/ethereum/go-ethereum/trie"
	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	wallet "github.com/ipfn/go-ipfn-wallet"
	"github.com/ipfn/ipfn/src/go/chain/dev/genesis"
	"github.com/ipfn/ipfn/src/go/trie/ethdb"
	ipfsdb "github.com/rootchain/go-ipfs-db"
)

var assignPaths []string

func init() {
	RootCmd.AddCommand(InitCmd)
	InitCmd.PersistentFlags().StringSliceVarP(&assignPaths, "assign", "a", nil, "key path or address and power in key:power:delegated format")
}

// InitCmd - Config get command.
var InitCmd = &cobra.Command{
	Use:   "init [config]",
	Short: "Initializes a chain",
	Long: `Initializes a new chain.

See wallet usage for more information on key derivation path.`,
	Example: `  $ rcx chain init -n mychain -a wallet:1e6:1e6 -a default/x/test:1e6:0
  $ rcx chain init -a zFNScYMGz4wQocWbvHVqS1HcbzNzJB5JK3eAkzF9krbSLZiV8cNr:1`,
	Annotations: map[string]string{"category": "chain"},
	Run:         cmdutil.WrapCommand(HandleInitCmd),
}

// HandleInitCmd - Handles chain init command.
func HandleInitCmd(cmd *cobra.Command, args []string) (err error) {
	store := ipfsdb.Wrap(contents.StateTriePrefix, ethdb.NewMemDatabase())
	config := &genesis.Config{
		Wallet:   wallet.NewDefault(),
		Database: trie.NewDatabase(store),
	}

	// create chain for default wallet by default
	if len(assignPaths) == 0 {
		assignPaths = []string{"default:1e6:1e6"}
	}

	for _, path := range assignPaths {
		power, err := genesis.ParsePowerString(path)
		if err != nil {
			return err
		}
		config.Assign(power)
	}

	for _, key := range config.WalletKeys() {
		_, err = wallet.PromptUnlock(config.Wallet, key)
		if err != nil {
			return
		}
	}

	state, err := genesis.Init(config)
	if err != nil {
		return
	}

	logger.PrintJSON(state)
	return
}
