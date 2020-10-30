/*
Copyright © 2020 MasseElch <info@masseelch.de>

*/
package cmd

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/masseelch/go-api-skeleton/ent"
	"github.com/masseelch/go-api-skeleton/fixtures"
	"github.com/masseelch/go-api-skeleton/util"
	"github.com/spf13/cobra"
)

// fixturesCmd represents the fixtures command
var fixturesCmd = &cobra.Command{
	Use: "fixtures",
	// todo
	Run: func(cmd *cobra.Command, args []string) {
		c, err := ent.Open("mysql", util.MysqlDSN())
		if err != nil {
			panic(err)
		}
		defer c.Close()

		if err := fixtures.Load(c); err != nil {
			panic(err) // todo
		}
	},
}

func init() {
	databaseCmd.AddCommand(fixturesCmd)
}
