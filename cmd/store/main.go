package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	root := &cobra.Command{
		Use:   "store",
		Short: "cloth store api cli",
	}
	root.AddCommand(
		createUser(),
		createIndexes(),
		server(),
		populate(),
		sendSaleEmail(),
		sendTransportEmail(),
	)
	err := root.Execute()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
