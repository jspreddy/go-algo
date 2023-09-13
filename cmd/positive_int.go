package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/jspreddy/go-algo/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var positiveIntCmd = &cobra.Command{
	Use:   "fpint",
	Short: "Run the first positive int algo",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		debug := viper.GetBool("debug")
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		log.Debug("all settings", "settings", viper.AllSettings())

		fmt.Println()
		fmt.Println(src.RenderTitleStyle("Sqrt"))
		fmt.Println()

		input := []int{1, 2, 4, 1, 5, 1, 6, 76, 87, 35, 54, 7878, 8}
		fmt.Println(input)
		fmt.Println(src.FirstPositiveInteger(input))
	},
}

func init() {
	rootCmd.AddCommand(positiveIntCmd)
}
