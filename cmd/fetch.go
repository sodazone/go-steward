// Copyright 2024 team@soda.zone
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"fmt"
	"os"

	"github.com/sodazone/go-steward/client"
	"github.com/spf13/cobra"
)

func croak(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

// fetchCmd represents the fetch command
var fetchCmd = &cobra.Command{
	Use:       "fetch [assets or chains]",
	Short:     "Streams either assets or chains data to stdout",
	ValidArgs: []string{"assets", "chains"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, err := cmd.Root().PersistentFlags().GetString("api-key")
		croak(err)

		if apiKey == "" {
			apiKey = PUB_KEY
		}
		httpUrl, err := cmd.Root().PersistentFlags().GetString("http-url")
		croak(err)

		var oc = client.NewOcelloidsClient(apiKey, httpUrl)

		if len(args) > 0 && args[0] == "chains" {
			err = oc.FetchChains()
		} else {
			err = oc.FetchAssets()
		}

		croak(err)
	},
}

func init() {
	rootCmd.AddCommand(fetchCmd)
}
