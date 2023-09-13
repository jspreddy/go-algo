package cmd

import (
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/jspreddy/go-algo/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var sqrtCmd = &cobra.Command{
	Use:   "sqrt",
	Short: "Run the sqrt algo",
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

		src.PrettyPrint(src.Sqrt, 2, debug)
		src.PrettyPrint(src.Sqrt, 4, debug)
		src.PrettyPrint(src.Sqrt, 144, debug)
	},
}

func init() {
	rootCmd.AddCommand(sqrtCmd)
}
