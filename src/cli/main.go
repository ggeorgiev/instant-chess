package main

import (
	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:     "instant-chess",
		Short:   "Instant Chess is a chess engine written in Go.",
		Long:    "Instant Chess is a chess engine written in Go. It is a work in progress.",
		Version: "0.0.1",
	}

	rootCmd.AddCommand(preplayCommand())

	rootCmd.Execute()
}
