/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/TheBoringDude/waxtest-manager/cmd/internal"
	"github.com/spf13/cobra"
)

var getTokenName string

// getTokenCmd represents the getToken command
var getTokenCmd = &cobra.Command{
	Use:   "getToken",
	Short: "Get free WAX token for your account",
	Long: `Get free WAX token for your account
	
*NOTE: You are limited to 1000 free tokens per 24 hours`,
	Run: func(cmd *cobra.Command, args []string) {
		if getTokenName == "" {
			return
		}

		var resp internal.GetTokenResponse

		if err := requests.Get(fmt.Sprintf("https://faucet.waxsweden.org/get_token?%s", getTokenName), &resp); err != nil {
			log.Fatal("\n[i] There was a problem when trying to get free tokens for the account. Are you sure you are connected to the internet?")

		}

		if resp.Msg != "succeeded" {
			log.Fatalf("\n[e] %s", resp.Msg)
		}

		fmt.Printf("\n[i] Successfully got free tokens for the wallet. \n\tPlease check if it was successfull in this website https://wax-test.bloks.io/account/%s", getTokenName)

	},
}

func init() {
	rootCmd.AddCommand(getTokenCmd)

	getTokenCmd.Flags().StringVarP(&getTokenName, "name", "n", "", "wax account wallet (required)")
	getTokenCmd.MarkFlagRequired("name")
}
