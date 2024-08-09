// Copyright 2024 team@soda.zone
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/sodazone/go-steward/client"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "steward",
	Short: "Ocelloids Data Steward Agent command-line interface",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.stw.yaml)")
	rootCmd.PersistentFlags().StringP("api-key", "k", "", "Ocelloids API key")
	rootCmd.PersistentFlags().StringP("http-url", "u", client.HTTP_URL, "HTTP API base URL")
	rootCmd.PersistentFlags().StringP("cursor", "c", "", "Page cursor")
	rootCmd.PersistentFlags().Uint16P("limit", "l", 25, "Max results per page")
	rootCmd.CompletionOptions.HiddenDefaultCmd = true

	viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("http-url", rootCmd.PersistentFlags().Lookup("http-url"))
	viper.BindPFlag("limit", rootCmd.PersistentFlags().Lookup("limit"))
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".stw" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".stw")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
