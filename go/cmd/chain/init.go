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
	"fmt"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec"
	"github.com/spf13/cobra"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
	"github.com/ipfn/ipfn/go/chain"
	"github.com/ipfn/ipfn/go/opcode"
	"github.com/ipfn/ipfn/go/opcode/chainops"
	"github.com/ipfn/ipfn/go/wallet"
)

var (
	keyPaths   []string
	addrPowers []string
)

func init() {
	RootCmd.AddCommand(InitCmd)
	InitCmd.PersistentFlags().StringSliceVarP(&keyPaths, "key", "k", nil, "key path and power in key:power:delegated format")
	InitCmd.PersistentFlags().StringSliceVarP(&addrPowers, "addr", "a", nil, "address and power in addr:power format")
}

// InitCmd - Config get command.
var InitCmd = &cobra.Command{
	Use:   "init [config]",
	Short: "Initializes a chain",
	Long: `Initializes a new chain.

See wallet usage for more information on key derivation path.`,
	Example: `  $ ipfn chain init -n mychain -k wallet:1e6:1e6 -k default/x/test:1e6:0
  $ ipfn chain init -p zFNScYMGz4wQocWbvHVqS1HcbzNzJB5JK3eAkzF9krbSLZiV8cNr:1`,
	Annotations: map[string]string{"category": "chain"},
	Run:         cmdutil.WrapCommand(HandleInitCmd),
}

// HandleInitCmd - Handles chain init command.
func HandleInitCmd(cmd *cobra.Command, args []string) (err error) {
	w := wallet.NewDefault()

	var (
		privKeys []*btcec.PrivateKey

		assignOps   []*opcode.BinaryCell
		delegateOps []*opcode.BinaryCell
	)

	passwords := make(map[string][]byte)

	var nonce opcode.ID
	for _, keyPath := range keyPaths {

		split := strings.Split(keyPath, ":")
		if len(split) != 3 {
			return fmt.Errorf("invalid key:power:delegated format: %q", keyPath)
		}
		power, err := strconv.ParseFloat(split[1], 64)
		if err != nil {
			return err
		}
		dpower, err := strconv.ParseFloat(split[2], 64)
		if err != nil {
			return err
		}
		path, err := wallet.ParseKeyPath(split[0])
		if err != nil {
			return err
		}
		password, ok := passwords[path.SeedName]
		if !ok {
			password, err = wallet.PromptWalletPassword(path.SeedName)
			if err != nil {
				return err
			}
			passwords[path.SeedName] = password
		}

		key, err := w.DeriveKey(path, password)
		if err != nil {
			return fmt.Errorf("wallet %s: %v", path.SeedName, err)
		}

		privkey, err := key.ECPrivKey()
		if err != nil {
			return err
		}

		if dpower == -1 {
			dpower = power
		}

		if dpower != 0 {
			op := chainops.DelegatePower(nonce, uint64(dpower))
			delegateOp, err := chainops.Sign(op, privkey)
			if err != nil {
				return err
			}
			privKeys = append(privKeys, privkey)
			delegateOps = append(delegateOps, delegateOp)
		}

		assignOps = append(assignOps, chainops.AssignPower(nonce, uint64(power), key.Serialize()))

		// TODO(crackcomm):
		nonce++
	}

	for _, addrPower := range addrPowers {
		split := strings.Split(addrPower, ":")
		if len(split) != 2 {
			return fmt.Errorf("invalid addr:power format: %q", addrPower)
		}
		power, err := strconv.ParseFloat(split[1], 64)
		if err != nil {
			return err
		}
		c, err := opcode.DecodeCID(split[0])
		if err != nil {
			return err
		}
		assignOps = append(assignOps, chainops.AssignPowerAddr(nonce, uint64(power), c))
		// TODO(crackcomm):
		nonce++
	}

	ops := []*opcode.BinaryCell{chainops.Genesis()}
	for _, op := range assignOps {
		ops = append(ops, op)
	}
	for _, op := range delegateOps {
		ops = append(ops, op)
	}

	state, err := chain.NewState(0, nil, ops) //, claimOp))
	if err != nil {
		logger.Error(err)
	}

	for _, key := range privKeys {
		if err := state.Sign(key); err != nil {
			return err
		}
	}

	logger.PrintJSON(state) //.Head().String())
	return
}
