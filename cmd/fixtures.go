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

// migrateCmd represents the migrate command
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
			panic(err)
		}

		// if ok, _ := cmd.Flags().GetBool("dump-sql"); ok {
		// 	if err := c.Schema.WriteTo(ctx, os.Stdout); err != nil {
		// 		panic(err)
		// 	}
		// } else if ok, _ := cmd.Flags().GetBool("force"); ok {
		// 	if err := c.Schema.Create(
		// 		ctx,
		// 		migrate.WithDropIndex(true),
		// 		migrate.WithDropColumn(true),
		// 	); err != nil {
		// 		panic(err)
		// 	}
		// } else {
		// 	fmt.Println("Either --dump-sql or --force must be given")
		// }
	},
}

func init() {
	rootCmd.AddCommand(fixturesCmd)

	// migrateCmd.Flags().Bool("dump-sql", false, "Dump the sql to be executed")
	// migrateCmd.Flags().Bool("force", false, "Execute the migration")
}
