package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
)

// receiveCmd represents the receive command
var receiveCmd = &cobra.Command{
	Use:   "receiver",
	Short: "hermes-cli receiver mode",
	Long: `in receiver mode you can receive messages
	usage:
		hermes-cli receive`,
	Run: func(cmd *cobra.Command, args []string) {

		sigs := make(chan os.Signal)
		signal.Notify(sigs, syscall.SIGTERM)
		fmt.Println("Waiting for any message")
		<-sigs
	},
}

func init() {

	rootCmd.AddCommand(receiveCmd)
	receiveCmd.Aliases = []string{"recv"}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// receiveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// receiveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
