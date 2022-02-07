package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var ttbCommand = &cobra.Command{
		Use:   "tbb",
		Short: "Blockchain cli",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}
	ttbCommand.AddCommand(versionCommand)
	ttbCommand.AddCommand(balancesCmd())
	ttbCommand.AddCommand(txCmd())
	err := ttbCommand.Execute()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
