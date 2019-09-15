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

// listchannelsCmd represents the listchannels command
var listchannelsCmd = &cobra.Command{
	Use:   "listchannels",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		config.Init()
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
		ch, err := cli.ListChannels(ctx, &api.Empty{})
		if err != nil {
			fmt.Fprint(os.Stderr, err)
		}
		fmt.Printf("%+v", ch)
	},
}

func init() {
	rootCmd.AddCommand(listchannelsCmd)
	listchannelsCmd.Aliases = []string{"list"}
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listchannelsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listchannelsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
