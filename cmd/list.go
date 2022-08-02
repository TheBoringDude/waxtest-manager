/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/TheBoringDude/waxtest-manager/cmd/internal"
	"github.com/olekukonko/tablewriter"
	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List your created wax wallets",
	Long:  `List the wallets that you have created with their keys`,
	Run: func(cmd *cobra.Command, args []string) {
		list := db.List()
		if len(list) == 0 {
			fmt.Println("No accounts yet~~")
			return
		}

		data := [][]string{}

		for _, v := range list {
			var r internal.TestnetAccount

			if err := internal.InterfaceToStruct(v, &r); err != nil {
				log.Fatal(err)
			}

			data = append(data,
				[]string{
					r.Account,
					fmt.Sprintf("(priv) %s", r.Keys.ActiveKey.Private),
					fmt.Sprintf("(priv) %s", r.Keys.OwnerKey.Private),
				},
				[]string{
					r.Account,
					fmt.Sprintf("(pub) %s", r.Keys.ActiveKey.Public),
					fmt.Sprintf("(pub) %s", r.Keys.OwnerKey.Public),
				})
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"Account", "Active Keys", "Owner Keys"})
		table.SetFooter([]string{"", "Total Accounts", strconv.Itoa(len(list))})
		table.SetAutoMergeCells(true)
		table.SetRowLine(true)
		table.AppendBulk(data)

		fmt.Println("")
		table.Render()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
