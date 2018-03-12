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
	"encoding/json"
	"errors"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
)

func init() {
	RootCmd.AddCommand(GetCmd)
}

// GetCmd - Config get command.
var GetCmd = &cobra.Command{
	Use:         "get [key]",
	Short:       "Gets config value",
	Annotations: map[string]string{"category": "config"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("key argument is required")
		}
		return nil
	},
	Run: cmdutil.WrapCommand(HandleGetCmd),
}

// HandleGetCmd - Handles config get command.
func HandleGetCmd(cmd *cobra.Command, args []string) (err error) {
	key := args[0]

	var value interface{}
	value = viper.Get(key)
	if value == nil {
		return fmt.Errorf("config value under key %q was not found", key)
	}

	switch value.(type) {
	case map[string]interface{}, []interface{}:
		value, err = json.MarshalIndent(value, "", "  ")
		if err != nil {
			return
		}
	}

	logger.Printf("Value under key %q:", key)
	logger.Line()
	logger.Printf("%s", value)

	return
}
