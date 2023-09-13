/*
Copyright © 2023 Sai Jonnala <jspreddy@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string = "$HOME/.go-algo/config.yaml"
	rootCmd        = &cobra.Command{
		Use:   "go-algo",
		Short: "Entry point to run my implemented algorithms.",
		Long:  ``,
		// Run: func(cmd *cobra.Command, args []string) {
		// 	debug := viper.GetBool("debug")
		// 	if debug {
		// 		log.SetLevel(log.DebugLevel)
		// 	}
		// },
	}
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", cfgFile, "config file")

	rootCmd.PersistentFlags().Bool("debug", false, "output debug logs")
	viper.BindPFlag("debug", rootCmd.PersistentFlags().Lookup("debug"))
}

func initConfig() {
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.go-algo")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// Config file not found; ignore error if desired
		} else {
			panic(fmt.Errorf("fatal error config file: %w", err))
		}
	} else {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
