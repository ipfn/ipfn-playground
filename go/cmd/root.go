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

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/ipfn/go-ipfn-cmd-util/logger"

	"github.com/ipfn/ipfn/go/cmd/account"
	"github.com/ipfn/ipfn/go/cmd/chain"
	"github.com/ipfn/ipfn/go/cmd/config"
	"github.com/ipfn/ipfn/go/cmd/core"
	"github.com/ipfn/ipfn/go/cmd/exp"
	"github.com/ipfn/ipfn/go/cmd/wallet"
)

func init() {
	RootCmd.AddCommand(exp.RootCmd)
	core.RegisterCommands(RootCmd)
	RootCmd.AddCommand(chain.RootCmd)
	RootCmd.AddCommand(account.RootCmd)
	RootCmd.AddCommand(wallet.RootCmd)
	RootCmd.AddCommand(config.RootCmd)
	RootCmd.PersistentFlags().BoolVarP(&logger.Verbose, "verbose", "v", false, "verbose logs output (stdout/stderr)")
}

var cfgFile string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use: "ipfn commands",
	Short: `IPFN – Interplanetary Functions

https://github.com/ipfn/go-ipfn`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.EnableCommandSorting = false
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	RootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "~/.ipfn.json", "config file")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	cfgFile, _ = homedir.Expand(cfgFile)
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".ipfn" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".ipfn")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		logger.Debugf("Using config file: %q", viper.ConfigFileUsed())
	}
}
