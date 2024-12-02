package main

import (
	"fmt"
	"os"

	"github.com/TopDev113/torram-relayer/cmd/cmd"
	"github.com/spf13/cobra"
)

func main() {
	// Create the root command
	rootCmd := &cobra.Command{
		Use:   "relayer",
		Short: "Relayer for Torram and Bitcoin networks",
	}

	// Add submitter command
	rootCmd.AddCommand(cmd.GetSubmitterCmd())

	// Execute the root command
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
