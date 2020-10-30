/*
Copyright © 2020 MasseElch <info@masseelch.de>

*/
package cmd

import (
	"context"
	"fmt"
	"github.com/masseelch/go-api-skeleton/ent"
	"github.com/masseelch/go-api-skeleton/ent/migrate"
	"github.com/masseelch/go-api-skeleton/util"
	"github.com/spf13/cobra"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// migrateCmd represents the migrate command
var migrateCmd = &cobra.Command{
	Use:   "migrate",
	// todo
	Run: func(cmd *cobra.Command, args []string) {
		c, err := ent.Open("mysql", util.MysqlDSN())
		if err != nil {
			panic(err)
		}
		defer c.Close()

		ctx := context.Background()

		if ok, _ := cmd.Flags().GetBool("dump-sql"); ok {
			if err := c.Schema.WriteTo(ctx, os.Stdout); err != nil {
				panic(err)
			}
		} else if ok, _ := cmd.Flags().GetBool("force"); ok {
			if err := c.Schema.Create(
				ctx,
				migrate.WithDropIndex(true),
				migrate.WithDropColumn(true),
			); err != nil {
				panic(err)
			}
		} else {
			fmt.Println("Either --dump-sql or --force must be given")
		}
	},
}

func init() {
	databaseCmd.AddCommand(migrateCmd)

	migrateCmd.Flags().Bool("dump-sql", false, "Dump the sql to be executed")
	migrateCmd.Flags().Bool("force", false, "Execute the migration")
}
