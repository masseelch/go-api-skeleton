/*
Copyright © 2020 MasseElch <info@masseelch.de>

*/
package cmd

import (
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/go-sql-driver/mysql"
	"github.com/masseelch/go-api-skeleton/auth"
	"github.com/masseelch/go-api-skeleton/ent"
	"github.com/masseelch/go-api-skeleton/handler"
	"github.com/masseelch/go-api-skeleton/util"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"log"
	"net/http"
)

// migrateCmd represents the migrate command
var serveCmd = &cobra.Command{
	Use: "serve",
	// todo
	Run: func(cmd *cobra.Command, args []string) {
		p, err := cmd.Flags().GetInt("port")
		if err != nil {
			panic(err)
		}

		// Get a database connection for ent.
		c, err := ent.Open("mysql", util.MysqlDSN())
		if err != nil {
			panic(err)
		}
		defer c.Close()

		// Create a validator.
		v := util.Validator()

		// Create a logger.
		l := logrus.New()

		r := chi.NewRouter()
		r.Use(middleware.Logger)// todo - replace with logrus

		r.Route("/auth", func(r chi.Router) {
			r.Post("/token", auth.LoginHandler(c, v ,l))
		})

		r.Group(func(r chi.Router) {
			r.Use(auth.Middleware(c, l))

			r.Mount("/users", handler.NewJobHandler(c, v, l))
		})

		log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", p), r))
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	serveCmd.Flags().IntP("port", "p", 8000, "Port to listen on")
}
