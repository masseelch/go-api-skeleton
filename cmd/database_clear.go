/*
Copyright © 2020 MasseElch <info@masseelch.de>

*/
package cmd

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masseelch/go-api-skeleton/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// clearCmd represents the clear command
var dropCmd = &cobra.Command{
	Use: "clear",
	// todo
	Run: func(cmd *cobra.Command, args []string) {
		// The user has provide the --force flag.
		ok, err := cmd.Flags().GetBool("force")
		if err != nil {
			panic(err) // todo
		}

		if ok {
			db, err := sql.Open("mysql", util.MysqlDSN())
			if err != nil {
				panic(err) // todo
			}
			defer db.Close()

			// Get a list of all tables.
			rs, err := db.Query(
				"select concat('drop table if exists `', table_name, '`;') from information_schema.tables where table_schema = ?",
				viper.GetString("database.name"),
			)
			if err != nil {
				panic(err) // todo
			}
			defer rs.Close()

			var s string
			stmts := make([]string, 0)
			for rs.Next() {
				if err := rs.Scan(&s); err != nil {
					panic(err) // todo
				}

				stmts = append(stmts, s)
			}

			// Disable foreign key checks, delete the tables, enable foreign key checks.
			if _, err := db.Exec("set foreign_key_checks = 0"); err != nil {
				panic(err) // todo
			}

			for _, stmt :=  range stmts {
				if _, err := db.Exec(stmt); err != nil {
					panic(err) // todo
				}
			}

			if _, err := db.Exec("set foreign_key_checks = 1"); err != nil {
				panic(err) // todo
			}
		} else {
			fmt.Println("Please add the --force flag to drop the database")
		}
	},
}

func init() {
	databaseCmd.AddCommand(dropCmd)

	dropCmd.Flags().Bool("force", false, "Drop the database")
}
