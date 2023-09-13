package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/jspreddy/go-algo/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var infoCmd = &cobra.Command{
	Use:   "info",
	Short: "Show info about config and other things.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		debug := viper.GetBool("debug")
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		log.Debug("all settings", "settings", viper.AllSettings())

		fmt.Println()
		fmt.Println(src.RenderTitleStyle("Info"))
		fmt.Println()
		fmt.Println("Resolved Configuration from viper: ")
		fmt.Println()
		fmt.Print(src.RenderMap(viper.AllSettings()))
	},
}

func init() {
	rootCmd.AddCommand(infoCmd)
}
