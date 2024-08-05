/*
Copyright © 2024 team@soda.zone

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
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
