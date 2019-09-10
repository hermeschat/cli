package cmd

import (
	"context"
	"fmt"
	"google.golang.org/grpc/metadata"
	"hermescli/api"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/amirrezaask/config"

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
		config.Init()
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
		defer cancel()
		md := metadata.Pairs("Authorization", config.Get("receiver_token"))
		ctx = metadata.NewOutgoingContext(ctx, md)
		buff, err := cli.EventBuff(ctx)
		if err != nil {
			fmt.Fprintf(os.Stdout, "error in calling event buff: %v", err)
			os.Exit(1)
		}
		for {
			e, err := buff.Recv()
			if err != nil {
				fmt.Fprintf(os.Stderr, "error in receiving event: %v", err)
				time.Sleep(time.Second * 3)
				continue
			}
			fmt.Println("event is")
			fmt.Printf("%+v\n", e)
			switch e.GetEvent().(type) {
			case *api.Event_NewMessage:
				fmt.Println("New Message recieved")
				m := e.GetNewMessage()
				fmt.Printf("%+v\n", m)
			}
		}
		<-sigs

	},
}

func init() {

	rootCmd.AddCommand(receiveCmd)
	receiveCmd.Aliases = []string{"recv"}

}
