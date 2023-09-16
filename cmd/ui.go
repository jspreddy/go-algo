package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/jspreddy/go-algo/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var uiCmd = &cobra.Command{
	Use:   "ui",
	Short: "Run the ui",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		debug := viper.GetBool("debug")
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		fmt.Println()
		fmt.Println(src.RenderTitleStyle("Running the UI"))
		fmt.Println()

		src.UI()
	},
}

func init() {
	rootCmd.AddCommand(uiCmd)
}
