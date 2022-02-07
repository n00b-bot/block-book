package main

import (
	"blockchain/model"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func txCmd() *cobra.Command {
	var txCmd = &cobra.Command{
		Use:   "tx",
		Short: "Interact with tx",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return fmt.Errorf("tx command wrong")
		},
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	txCmd.AddCommand(txAddCmd())
	return txCmd
}
func txAddCmd() *cobra.Command {
	txAddCmd := &cobra.Command{
		Use:   "add",
		Short: "Add tx ",
		Run: func(cmd *cobra.Command, args []string) {
			from, _ := cmd.Flags().GetString("from")
			to, _ := cmd.Flags().GetString("to")
			value, _ := cmd.Flags().GetInt("value")
			tx := model.Tx{
				From:  model.Account(from),
				To:    model.Account(to),
				Value: value,
				Data:  "",
			}
			state, err := model.NewStateFromDisk()
			if err != nil {
				log.Println(err)
				os.Exit(1)
			}
			defer state.DbFile.Close()
			if err := state.Add(tx); err != nil {
				log.Println(err)
				os.Exit(1)
			}
			if _, err = state.Save(); err != nil {
				log.Println(err)
				os.Exit(1)
			}
			fmt.Println("tx complete")

		},
	}
	txAddCmd.Flags().String("from", "", "From what account to send tokens")
	txAddCmd.MarkFlagRequired("from")

	txAddCmd.Flags().String("to", "", "To what account to send tokens")
	txAddCmd.MarkFlagRequired("to")

	txAddCmd.Flags().Int("value", 0, "How many tokens to send")
	txAddCmd.MarkFlagRequired("value")
	return txAddCmd
}
