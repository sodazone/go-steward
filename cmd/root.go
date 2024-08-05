// Copyright 2024 team@soda.zone
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var apiKey string
var httpUrl string

const PUB_KEY = "eyJhbGciOiJFZERTQSIsImtpZCI6IklSU1FYWXNUc0pQTm9kTTJsNURrbkJsWkJNTms2SUNvc0xBRi16dlVYX289In0.ewogICJpc3MiOiAiZGV2LWFwaS5vY2VsbG9pZHMubmV0IiwKICAianRpIjogIjAxMDAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwIiwKICAic3ViIjogInB1YmxpY0BvY2VsbG9pZHMiCn0K.bjjQYsdIN9Fx34S9Of5QSKxb8_aOtwURInOGSSc_DxrdZcnYWi-5nnZsh1v5rYWuRWNzLstX0h1ICSH_oAugAQ"

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "steward [assets or chains]",
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
	rootCmd.PersistentFlags().StringVarP(&apiKey, "api-key", "k", "", "Ocelloids API key")
	rootCmd.PersistentFlags().StringVarP(&httpUrl, "http-url", "u", "https://dev-api.ocelloids.net", "HTTP API base URL")

	viper.BindPFlag("api-key", rootCmd.PersistentFlags().Lookup("api-key"))
	viper.BindPFlag("http-url", rootCmd.PersistentFlags().Lookup("http-url"))
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
