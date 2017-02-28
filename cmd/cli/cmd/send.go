package cmd

import (
	"github.com/spf13/cobra"
	"fmt"
)

func init() {
	RootCmd.AddCommand(sendCmd)
}

var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "Send a raw message to the specified goctopus connector",
	Long:  `Send a raw message to the specified goctopus connector`,
	Run: func(cmd *cobra.Command, args []string) {
		//TODO
		fmt.Println("Message sent")
	},
}