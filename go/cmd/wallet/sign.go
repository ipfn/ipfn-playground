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
	"strings"

	"github.com/spf13/cobra"

	"golang.org/x/crypto/sha3"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
)

func init() {
	RootCmd.AddCommand(SignCmd)
	SignCmd.PersistentFlags().BoolVarP(&hashPath, "hash", "x", false, "derive hash path")
	SignCmd.PersistentFlags().StringVarP(&keyPath, "key-path", "k", "", "wallet key path (<wallet>/<x|m>/<path>)")
	SignCmd.PersistentFlags().StringVarP(&derivePath, "derive-path", "d", "", "derive BIP32 hierarchical path")
	SignCmd.PersistentFlags().StringVarP(&walletName, "wallet", "w", "default", "wallet name")
}

// SignCmd - Key sign command.
var SignCmd = &cobra.Command{
	Use:   "sign [content]",
	Short: "Sign with ECDSA",
	Example: `  $ ipfn wallet sign -w example -xd mnemonic '{"value": "0xd"}'
  $ ipfn wallet sign -d m/44'/138'/0'/0/0 '{"value": "0xd"}'
  $ ipfn wallet sign -w example -d m/44'/138'/0'/0/0 '{"value": "0xd"}'`,
	Annotations: map[string]string{"category": "wallet"},
	Args:        cobra.MinimumNArgs(1),
	Run:         cmdutil.WrapCommand(HandleSignCmd),
}

// HandleSignCmd - Handles key sign command.
func HandleSignCmd(cmd *cobra.Command, args []string) (err error) {
	acc, err := deriveWallet()
	if err != nil {
		return
	}
	priv, err := acc.ECPrivKey()
	if err != nil {
		return
	}
	hash := sha3.Sum512([]byte(strings.Join(args, " ")))
	signature, err := priv.Sign(hash[:])
	if err != nil {
		return
	}
	pubKey, err := acc.ECPubKey()
	if err != nil {
		return
	}
	if !signature.Verify(hash[:], pubKey) {
		return errors.New("Cannot verify signature")
	}
	c, err := acc.CID()
	if err != nil {
		return
	}
	sigBytes := signature.Serialize()
	logger.Print()
	logger.Printf("Address:        %s", c)
	logger.Printf("Signature hex:  %x", sigBytes)
	logger.Printf("Signature hash: %x", sha3.Sum512(sigBytes))
	return
}
