package cmd

import (
	"github.com/charmbracelet/log"
	"github.com/jspreddy/go-algo/src"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var versionCmd = &cobra.Command{
	Use:   "gloss",
	Short: "Render the lipgloss example.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		debug := viper.GetBool("debug")
		if debug {
			log.SetLevel(log.DebugLevel)
		}

		log.Debug("all settings", "settings", viper.AllSettings())

		src.RenderLayout()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
