package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// sendCmd represents the send command
var sendCmd = &cobra.Command{
	Use:   "send",
	Short: "hermes-cli send mode",
	Long: `in send mode you can send messages
	usage:
		hermes-cli send [receiver] [body]`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			fmt.Fprintf(os.Stdout, "send needs exactly two arguments")
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// sendCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// sendCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
