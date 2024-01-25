package main

import (
	"fmt"

	"github.com/ggeorgiev/instant-chess/src/chess"
	"github.com/spf13/cobra"
)

func preplayCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "preplay",
		Short: "Preplays positions with certain peaces and creates a database of them.",
		Long:  "Preplays positions with certain peaces and creates a database of them. It is a work in progress.",
		RunE:  preplay,
	}

	cmd.Flags().StringP(
		"directory",
		"d",
		".",
		"The storage directory",
	)

	cmd.Flags().StringP(
		"peaces",
		"p",
		"[]string{}",
		"Peaces to preplay. Example: --peaces=♔♖♖♘♘♗♗♕♙♙♙♙♙♙♙♙♚♜♜♞♞♝♝♛♟︎♟︎♟︎♟︎♟︎♟︎♟︎♟︎",
	)

	return cmd
}

func preplay(cmd *cobra.Command, args []string) error {
	peaces, err := cmd.Flags().GetString("peaces")
	if err != nil {
		return err
	}

	if peaces == "" {
		return fmt.Errorf("peaces are required")
	}

	return chess.Generate(peaces)
}
