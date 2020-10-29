/*
Copyright © 2020 MasseElch <info@masseelch.de>

*/
package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"strings"

	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-api-skeleton",
	Short: "Skeleton to use for a fresh rest api in go",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Database configuration
	rootCmd.PersistentFlags().String("database.host", "localhost", "database host")
	rootCmd.PersistentFlags().String("database.port", "3306", "database host")
	rootCmd.PersistentFlags().String("database.table", "todo_ent_cobra", "database table")
	rootCmd.PersistentFlags().String("database.user", "todo_ent_cobra", "database user")
	rootCmd.PersistentFlags().String("database.password", "", "database password")

	// Bind all yet defined flags. Every flag after this line will not be passed to viper.
	if err := viper.BindPFlags(rootCmd.PersistentFlags()); err != nil {
		panic(err)
	}

	// Allow setting flags by config file.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "config.yaml", "/path/to/config.yaml")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)

		if err := viper.ReadInConfig(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}
