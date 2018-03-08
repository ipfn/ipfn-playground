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

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
)

var (
	keySignSize int
	keyAlphabet bool
)

func init() {
	RootCmd.AddCommand(SignCmd)
	SignCmd.PersistentFlags().BoolVarP(&forcePath, "force", "f", false, "Force derivation path")
	SignCmd.PersistentFlags().BoolVarP(&customSeedPwd, "custom", "u", false, "Custom seed password")
	SignCmd.PersistentFlags().BoolVarP(&keyAlphabet, "encoding", "e", false, "Custom encoding alphabet")
	SignCmd.PersistentFlags().IntVarP(&keySignSize, "size", "s", 32, "Signature hash size")
	SignCmd.PersistentFlags().StringVarP(&keyAddrID, "addr", "a", "0x00", "Custom address pubkey-hash address ID")
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
	logger.Print(encodeB58(signature.Serialize())[:keySignSize])
	return
}

func encodeB58(b []byte) string {
	if keyAlphabet {
		return base58.EncodeAlphabet(b, "12=45-789_BCDEFGHJKLMNPQRSTUV#XYZabcdefghijkmnop?rstuvwxyz")
	}
	return base58.Encode(b)
}
