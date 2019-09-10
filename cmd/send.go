package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/amirrezaask/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"hermescli/api"

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
		config.Init()
		if len(args) < 2 {
			fmt.Fprintf(os.Stderr, "send needs exactly two arguments")
			os.Exit(1)
		}
		receiverID := args[0]
		msgBody := args[1]
		con, err := grpc.Dial(fmt.Sprintf("%s:%s", config.Get("host"), config.Get("port")), grpc.WithInsecure())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error in grpc dial: %v", err)
			os.Exit(1)
		}

		cli := api.NewHermesClient(con)
		ctx, cancel := context.WithCancel(context.Background())
		md := metadata.Pairs("Authorization", config.Get("sender_token"))
		ctx = metadata.NewOutgoingContext(ctx, md)
		defer cancel()
		buff, err := cli.EventBuff(ctx)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error in calling event buff: %v", err)
			os.Exit(1)
		}

		err = buff.Send(&api.Event{Event: &api.Event_NewMessage{&api.Message{
			To:   receiverID,
			Body: msgBody,
		}}})
		if err != nil {
			fmt.Fprintf(os.Stderr, "error in sending message:%v", err)
			os.Exit(1)

		}
		time.Sleep(time.Hour * 2)
		fmt.Println("message sent")
	},
}

func init() {
	rootCmd.AddCommand(sendCmd)

}
