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

package exp

import (
	"github.com/cespare/xxhash"
	"github.com/spf13/cobra"

	cmdutil "github.com/ipfn/go-ipfn-cmd-util"
	"github.com/ipfn/go-ipfn-cmd-util/logger"
)

func init() {
	RootCmd.AddCommand(GetCmd)
}

// GetCmd - Config get command.
var GetCmd = &cobra.Command{
	Use:         "get [key]",
	Short:       "Gets exp value",
	Annotations: map[string]string{"category": "exp"},
	Run:         cmdutil.WrapCommand(HandleGetCmd),
}

// HandleGetCmd - Handles exp get command.
func HandleGetCmd(cmd *cobra.Command, args []string) (err error) {

	logger.Printf("0x%x", xxhash.Sum64([]byte(args[0])))

	return
}
