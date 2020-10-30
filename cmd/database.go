/*
Copyright © 2020 MasseElch <info@masseelch.de>

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// databaseCmd represents the database command
var databaseCmd = &cobra.Command{
	Use:   "database",
	Short: "Command to manipulate the database",
	Aliases: []string{"db"},
}

func init() {
	rootCmd.AddCommand(databaseCmd)
}
