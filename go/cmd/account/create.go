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

package account

import (
	"github.com/spf13/cobra"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
)

var (
	hashPath   bool
	keyPath    string
	derivePath string
	walletName string

	threshold  uint
	keyWeight  uint
	masterKeys int
	clientKeys int
)

func init() {
	RootCmd.AddCommand(CreateCmd)
	CreateCmd.PersistentFlags().BoolVarP(&hashPath, "hash-path", "x", false, "Derive hash path")
	CreateCmd.PersistentFlags().StringVarP(&keyPath, "key-path", "k", "", "Wallet key path (<wallet>/<x|m>/<path>)")
	CreateCmd.PersistentFlags().StringVarP(&derivePath, "derive-path", "d", "", "Derive BIP32 hierarchical path")
	CreateCmd.PersistentFlags().StringVarP(&walletName, "wallet", "w", "default", "Wallet name")
	CreateCmd.PersistentFlags().UintVarP(&keyWeight, "key-weight", "g", 1, "Master key weight for multisig")
	CreateCmd.PersistentFlags().UintVarP(&threshold, "threshold", "t", 1, " Threshold for multisig")
	CreateCmd.PersistentFlags().IntVarP(&masterKeys, "master-keys", "m", 1, "Master keys to generate")
	CreateCmd.PersistentFlags().IntVarP(&clientKeys, "client-keys", "n", 1, "Client keys to generate")
}

// CreateCmd - Account create command.
var CreateCmd = &cobra.Command{
	Use:     "create [account]",
	Short:   "Creates new account",
	Example: `  $ ipfn account create myacc -w mywallet -xd keyid`,
	Args: func(cmd *cobra.Command, args []string) error {
		// TODO: check if acc already exists
		// if len(args) < 1 {
		// 	return errors.New("seed or path argument is required")
		// }
		return nil
	},
	Annotations: map[string]string{"category": "account"},
	Run:         cmdutil.WrapCommand(HandleCreateCmd),
}

// HandleCreateCmd - Handles account get command.
func HandleCreateCmd(cmd *cobra.Command, args []string) (err error) {
	// accName := cmdutil.ArgDefault(args, 0, "default")

	// // Set derivation parameters
	// // Check if no flags were set
	// if derivePath == "" && keyPath == "" {
	// 	hashPath = false // just to make sure
	// 	derivePath = fmt.Sprintf("x/ipfnio.account/%s", accName)
	// }

	// var path *wallet.KeyPath
	// if keyPath != "" {
	// 	path, err = wallet.ParseKeyPath(keyPath)
	// } else {
	// 	path = wallet.NewKeyPath(walletName, derivePath, hashPath)
	// }
	// if err != nil {
	// 	return
	// }

	// // // Prompt user for password
	// // password, err := wallet.PromptWalletPassword(path.SeedName)
	// // if err != nil {
	// // 	return
	// // }

	// walletAcc := wallet.Account{Name: accName}

	// chainAcc := &accounts.Account{
	// 	Address: accounts.NewAddress([]byte("random example")),
	// 	Owner: &accounts.Identity{
	// 		Threshold: threshold,
	// 	},
	// }

	// w := wallet.NewDefault()

	// passwords := make(map[string][]byte)

	// for index := 0; index < masterKeys; index++ {
	// 	ownerKeyPath := path.Extend(fmt.Sprintf("master/%d", index))

	// 	password, ok := passwords[ownerKeyPath.SeedName]
	// 	if !ok {
	// 		password, err = wallet.PromptWalletPassword(ownerKeyPath.SeedName)
	// 		if err != nil {
	// 			return err
	// 		}
	// 		passwords[ownerKeyPath.SeedName] = password
	// 	}

	// 	key, err := w.DeriveKey(ownerKeyPath, password)
	// 	if err != nil {
	// 		return fmt.Errorf("wallet %s: %v", ownerKeyPath.SeedName, err)
	// 	}

	// 	c, err := key.Cid()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	address := accounts.FromCID(c)

	// 	walletAcc.OwnerKeys = append(walletAcc.OwnerKeys, &wallet.AccountKeyPath{
	// 		Address: address,
	// 		Weight:  keyWeight,
	// 		KeyPath: ownerKeyPath,
	// 	})
	// 	chainAcc.Owner.Keys = append(chainAcc.Owner.Keys, &accounts.Key{
	// 		Weight:  1,
	// 		Address: address,
	// 	})
	// }

	// for index := 0; index < clientKeys; index++ {
	// 	walletAcc.ClientKeys = append(walletAcc.ClientKeys, path.Extend(fmt.Sprintf("client/%d", index)))
	// }

	// b, err := json.MarshalIndent(&walletAcc, "", "  ")
	// if err != nil {
	// 	return
	// }
	// logger.Printf("%s", b)

	// b, _ = json.MarshalIndent(chainAcc, "", "  ")
	// logger.Printf("%s", b)
	// // Save account
	// acc, err := w.CreateAccount(account)
	// // if err != nil {
	// // 	return
	// // }

	// acc, err := wallet.NewDefault().DeriveKey(path, password)
	// if err != nil {
	// 	return
	// }

	// neuter, _ := acc.Neuter()
	// c, _ := acc.Cid()
	// logger.Print()
	// logger.Printf("Account: %s", accName)
	// logger.Printf("Address: %s", c)
	// logger.Printf("Pubkey:  %s", neuter)
	return
}
