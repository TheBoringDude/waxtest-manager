/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"net/http"
	"os"

	"github.com/TheBoringDude/scuffed-go/requester"
	"github.com/TheBoringDude/waxtest-manager/cmd/internal"
	"github.com/spf13/cobra"
)

var requests = requester.NewRequester(&http.Client{})
var db = internal.DB()

var rootCmd = &cobra.Command{
	Use:   "waxtest-manager",
	Short: "Manage your wax testnet accounts.",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
