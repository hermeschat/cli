package cmd

import (
	"context"
	"fmt"
	"hermescli/api"
	"hermescli/config"
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/cobra"
	"google.golang.org/grpc"
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
		con, err := grpc.Dial(fmt.Sprintf("%s:%s", config.Get("host"), config.Get("port")), grpc.WithInsecure())
		if err != nil {
			fmt.Fprintf(os.Stdout, "error in grpc dial: %v", err)
			os.Exit(1)
		}
		fmt.Println("Waiting for any message")
		cli := api.NewHermesClient(con)
		ctx, cancel := context.WithCancel(context.Background())
		buff, err := cli.EventBuff(ctx)
		if err != nil {
			fmt.Fprintf(os.Stdout, "error in calling event buff: %v", err)
			os.Exit(1)
		}
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
