package main

import (
	"fmt"
	"os"
)

// TODO: init log

func main() {
	// params.SetAddressPrefixes()

	// rootCmd := cmd.NewRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}
}
