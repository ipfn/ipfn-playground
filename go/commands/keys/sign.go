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

package keys

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	base58 "github.com/jbenet/go-base58"
	prompt "github.com/segmentio/go-prompt"

	"github.com/ethereum/go-ethereum/crypto/sha3"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
)

var (
	sigHashSize      int
	passwordAlphabet bool
)

func init() {
	RootCmd.AddCommand(SignCmd)
	SignCmd.PersistentFlags().BoolVarP(&forcePath, "force", "f", false, "Force derivation path")
	SignCmd.PersistentFlags().BoolVarP(&customSeedPwd, "custom", "u", false, "Custom seed password")
	SignCmd.PersistentFlags().BoolVar(&btcAddr, "btc", false, "BTC address format")
	SignCmd.PersistentFlags().IntVarP(&sigHashSize, "size", "s", 32, "Signature hash size")
	SignCmd.PersistentFlags().BoolVarP(&passwordAlphabet, "password", "p", false, "Password encoding alphabet")
}

// SignCmd - Key sign command.
var SignCmd = &cobra.Command{
	Use:         "sign [seed] [path] [content]",
	Short:       "Signs message using derived key",
	Annotations: map[string]string{"category": "key"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("seed argument is required")
		}
		if viper.Get(fmt.Sprintf("seeds.%s", args[0])) == nil {
			return fmt.Errorf("seed %q was not found", args[0])
		}
		return nil
	},
	Run: cmdutil.WrapCommand(HandleSignCmd),
}

// HandleSignCmd - Handles key sign command.
func HandleSignCmd(cmd *cobra.Command, args []string) (err error) {
	var path string
	if len(args) > 1 {
		path = args[1]
	} else {
		path = prompt.StringRequired("Derivation path")
	}
	acc, err := deriveKey(args[0], path)
	if err != nil {
		return
	}
	priv, err := acc.ECPrivKey()
	if err != nil {
		return
	}
	var content []byte
	if len(args) > 2 {
		content = []byte(args[2])
	} else {
		content = []byte(prompt.StringRequired("Message content"))
	}
	signature, err := priv.Sign(content)
	if err != nil {
		return
	}
	pub, err := acc.ECPubKey()
	if err != nil {
		return
	}
	if !signature.Verify(content, pub) {
		return errors.New("Cannot verify signature")
	}
	if printKey {
		if err := printAccount(acc); err != nil {
			return err
		}
	}
	sigBytes := signature.Serialize()
	logger.Debugf("Hex encoded signature: %x", sigBytes)
	sighash := sha3.Sum512(sigBytes)
	if passwordAlphabet {
		logger.Printf("Signature hash: %s", encodePass(sighash[:sigHashSize]))
	} else {
		logger.Printf("Signature hash: %x", sighash[:sigHashSize])
	}
	return
}

func encodePass(b []byte) string {
	return base58.EncodeAlphabet(b, "12=45-789_BCDEFGHJKLMNPQRSTUV#XYZabcdefghijkmnop?rstuvwxyz")
}
