package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

type SendCmdOptions struct {
	connector string
	message   string
}

func init() {
	o := &SendCmdOptions{}
	RootCmd.AddCommand(sendCmd)
	sendCmd.Flags().StringVar(&o.connector, "connector", "", "The connector id to send the message to")
	sendCmd.Flags().StringVar(&o.connector, "message", "", "Message content")
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
