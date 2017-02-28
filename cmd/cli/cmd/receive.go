package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

type ReceiveCmdOptions struct {
	connector string
}

func init() {
	o := &ReceiveCmdOptions{}
	RootCmd.AddCommand(recCmd)
	recCmd.Flags().StringVar(&o.connector, "connector", "", "The connector id to receive from")
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