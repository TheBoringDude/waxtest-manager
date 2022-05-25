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

var name string

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Create a new testnet account name",
	Long: `Create a new testnet account name.
	
This will send an api request to create a new testnet with the current account.`,
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			return
		}

		var resp internal.TestnetResponse

		if err := requests.Get(fmt.Sprintf("https://faucet.waxsweden.org/create_account?%s", name), &resp); err != nil {
			log.Fatal("\n[i] There was a problem when trying to create a new testnet account. Are you connected to the internet?")
		}

		if resp.Msg != "succeeded" {
			log.Fatalf("\n[e] %s", resp.Msg)
		}

		account := internal.TestnetAccount{
			Account: resp.Account,
			Keys:    resp.Keys,
		}
		db.Push(account)

		fmt.Println("\n[i] Successfully created a new account name. You can check it with the `list` command.")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.Flags().StringVarP(&name, "name", "n", "", "wax account name (required)")
	newCmd.MarkFlagRequired("name")
}
