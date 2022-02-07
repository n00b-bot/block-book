package main

import (
	"blockchain/model"
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func balancesCmd() *cobra.Command {
	var balancesCmd = &cobra.Command{
		Use:   "balances",
		Short: "Interact with balances ",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("something error")
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	balancesCmd.AddCommand(balancesListCmd)
	return balancesCmd
}

var balancesListCmd = &cobra.Command{
	Use:   "list",
	Short: "List all balances",
	Run: func(cmd *cobra.Command, args []string) {
		state, err := model.NewStateFromDisk()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		defer state.DbFile.Close()
		fmt.Println("Account Balances")
		fmt.Println("------------------------------")
		for user, balance := range state.User {
			fmt.Println("Account " + user + ": " + model.Account(strconv.Itoa(balance)))
		}
	},
}
