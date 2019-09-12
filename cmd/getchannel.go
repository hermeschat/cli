/*
Copyright © 2019 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/amirrezaask/config"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"hermescli/api"
)

// getchannelCmd represents the getchannel command
var getchannelCmd = &cobra.Command{
	Use:   "getchannel",
	Short: "gets channel",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		config.Init()
		if len(args) < 1 {
			fmt.Fprintf(os.Stderr, "send needs a channelID")
			os.Exit(1)
		}
		channelID := args[0]
		con, err := grpc.Dial(fmt.Sprintf("%s:%s", config.Get("host"), config.Get("port")), grpc.WithInsecure())
		if err != nil {
			fmt.Fprintf(os.Stderr, "error in grpc dial: %v", err)
			os.Exit(1)
		}
		now := time.Now()
		now = now.Add(-time.Second)
		fmt.Printf("timestamp is %v\n", now.Unix())
		cli := api.NewHermesClient(con)
		ctx, cancel := context.WithCancel(context.Background())
		md := metadata.Pairs("Authorization", config.Get("sender_token"))
		ctx = metadata.NewOutgoingContext(ctx, md)
		defer cancel()
		ch, err := cli.GetChannel(ctx, &api.GetChannelRequest{
			Id:        channelID,
			Timestamp: fmt.Sprint(now.Unix()),
		})
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Printf("%+v", ch)
	},
}

func init() {
	rootCmd.AddCommand(getchannelCmd)
	getchannelCmd.Aliases = []string{"ch"}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getchannelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getchannelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
