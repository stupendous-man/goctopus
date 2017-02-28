package cmd

import "github.com/spf13/cobra"

var RootCmd = &cobra.Command{
	Use:   "goctopus",
	Short: "A simplistic integration framework to glue messasging systems at large scale",
	Long: "\u200B" + `
  _____            _
 / ____|          | |
| |  __  ___   ___| |_ ___  _ __  _   _ ___
| | |_ |/ _ \ / __| __/ _ \| '_ \| | | / __|
| |__| | (_) | (__| || (_) | |_) | |_| \__ \
 \_____|\___/ \___|\__\___/| .__/ \__,_|___/
			   | |
			   |_|
A simplistic integration framework to glue messasging systems at large scale`,
}
