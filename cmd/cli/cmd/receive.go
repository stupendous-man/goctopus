package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(recCmd)
}

var recCmd = &cobra.Command{
	Use:   "receive",
	Short: "Synchronously poll for a message from specified connector",
	Long:  `Synchronously poll for a message from specified connector`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO
		fmt.Println("Waiting for message...")
	},
}