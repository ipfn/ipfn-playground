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

package config

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
)

func init() {
	RootCmd.AddCommand(SetCmd)
}

// SetCmd - Config set command.
var SetCmd = &cobra.Command{
	Use:         "set [key] [value]",
	Short:       "Sets config value",
	Annotations: map[string]string{"category": "config"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("key and value argument is required")
		}
		if len(args) < 2 {
			return errors.New("value argument is required")
		}
		return nil
	},
	Run: cmdutil.WrapCommand(HandleSetCmd),
}

// HandleSetCmd - Handles config set command.
func HandleSetCmd(cmd *cobra.Command, args []string) (err error) {
	var (
		key   = args[0]
		value = args[1]
	)

	logger.Printf("Key %q value:", key)
	logger.Line()
	logger.Printf("%s", value)

	return setConfig(key, value)
}

func setConfig(key string, value interface{}) error {
	viper.Set(key, value)
	if err := viper.WriteConfig(); err != nil {
		return fmt.Errorf("failed to write config: %v", err)
	}
	return nil
}
