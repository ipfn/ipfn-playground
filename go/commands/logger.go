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
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

var (
	logger         *zap.Logger
	verboseLogs    bool
	productionLogs bool
)

func init() {
	cobra.OnInitialize(initLogger)
	RootCmd.PersistentFlags().BoolVarP(&verboseLogs, "verbose", "v", false, "verbose logs output (stdout/stderr)")
	RootCmd.PersistentFlags().BoolVar(&productionLogs, "prod", false, "production logs output (JSON formatted)")
}

func initLogger() {
	var err error
	if verboseLogs {
		logger, err = zap.NewDevelopment()
	} else if productionLogs {
		logger, err = zap.NewProduction()
	} else {
		logger = zap.NewNop()
	}
	if err != nil {
		panic(err)
	}
}
