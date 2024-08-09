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
	Short:     "Prints assets or chains data to stdout",
	ValidArgs: []string{"assets", "chains"},
	Args:      cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run: func(cmd *cobra.Command, args []string) {
		apiKey, err := cmd.Root().PersistentFlags().GetString("api-key")
		croak(err)

		httpUrl, err := cmd.Root().PersistentFlags().GetString("http-url")
		croak(err)

		limit, err := cmd.Root().PersistentFlags().GetUint16("limit")
		croak(err)

		cursor, err := cmd.Root().PersistentFlags().GetString("cursor")
		croak(err)

		pagination := client.Pagination{
			Limit: limit,
		}

		if cursor != "" {
			pagination.Cursor = cursor
		}

		var oc = client.NewOcelloidsClient(apiKey, httpUrl, pagination)

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
