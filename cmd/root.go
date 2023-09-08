/*
Copyright © 2023 Sai Jonnala <jspreddy@users.noreply.github.com>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/jspreddy/go-algo/src"
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-algo",
	Short: "Entry point to run my implemented algorithms.",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		debug := false

		src.PrettyPrint(src.Sqrt, 2, debug)
		src.PrettyPrint(src.Sqrt, 4, debug)
		src.PrettyPrint(src.Sqrt, 144, debug)

		input := []int{1, 2, 4, 1, 5, 1, 6, 76, 87, 35, 54, 7878, 8}
		fmt.Println(input)
		fmt.Println(src.FirstPositiveInteger(input))
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-algo.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
