// Copyright © 2017 The IPFN Authors. All Rights Reserved.
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

package commands

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/spf13/cobra"
	"go.uber.org/zap"

	"github.com/ipfn/ipfn/go/identity"
)

func init() {
	RootCmd.AddCommand(idCmd)
}

var idCmd = &cobra.Command{
	Use:         "id",
	Short:       "Manage identities",
	Annotations: map[string]string{"category": "id"},
}

var (
	idGenSize int
	idGenType string
)

func init() {
	idNewCmd.PersistentFlags().IntVarP(&idGenSize, "size", "s", 2048, "Size of cryptographic key")
	idNewCmd.PersistentFlags().StringVarP(&idGenType, "type", "t", "RSA", "Identity type: RSA, Ed25519, Secp256k1")
	idCmd.AddCommand(idNewCmd)
}

var idNewCmd = &cobra.Command{
	Use:         "new [OUTPUT]",
	Short:       "Generates a new identity",
	Annotations: map[string]string{"category": "id"},
	Run: func(cmd *cobra.Command, args []string) {
		var typ identity.KeyType
		switch idGenType {
		case "RSA":
			typ = identity.RSA
		case "Ed25519":
			typ = identity.Ed25519
		case "Secp256k1":
			typ = identity.Secp256k1
		default:
			logger.Fatal("unknown key type", zap.String("type", idGenType))
		}

		logger.Info("generating key",
			zap.Int("size", idGenSize),
			zap.String("type", idGenType),
		)

		id, err := identity.NewSafe(typ, idGenSize)
		if err != nil {
			logger.Fatal("failed to generate key", zap.Error(err))
		}

		logger.Info("generated key",
			zap.String("id", id.ID.Pretty()),
		)

		body, err := json.MarshalIndent(id, "", "  ")
		if err != nil {
			logger.Fatal("failed to serialize key", zap.Error(err))
		}

		if len(args) == 0 {
			_, err = os.Stdout.Write(body)
			if err != nil {
				logger.Fatal("failed to write output", zap.Error(err))
			}
			return
		}

		err = ioutil.WriteFile(args[0], body, os.FileMode(0400))
		if err != nil {
			logger.Fatal("failed to save key",
				zap.String("file", args[0]),
				zap.Error(err),
			)
		}
		logger.Info("saved key to", zap.String("file", args[0]))
	},
}

func init() {
	idCmd.AddCommand(idVerifyCmd)
}

var idVerifyCmd = &cobra.Command{
	Use:   "verify [INPUT]",
	Short: "Verifies identity",
	Long: `Reads a serialized JSON message containing a cryptographic key and validates it.
  Currently only checks if serialized public key matches ID.`,
	Annotations: map[string]string{"category": "id"},
	Args:        cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		logger.Info("reading key",
			zap.String("input", args[0]),
		)

		keys, err := identity.ReadFromFile(args[0])
		if err != nil {
			logger.Fatal("failed to read key", zap.Error(err))
		}

		logger.Info("key unserialized successfully",
			zap.String("id", keys.ID.Pretty()),
		)
	},
}
