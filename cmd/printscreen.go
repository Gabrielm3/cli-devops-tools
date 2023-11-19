package cmd

import (
	"cli-go/utils"

	"github.com/spf13/cobra"
)

// printscreenCmd represents the printscreen command
var printscreenCmd = &cobra.Command{
	Use:   "printscreen",
	Short: "Screenshot jpeg websites",
	Long: `Screenshot jpeg websites.`,
	Run: func(cmd *cobra.Command, args []string) {
		site := cmd.Flag("site").Value.String()
		utils.GetChromeScreenShot(site, 100)
	},
}

func init() {
	rootCmd.AddCommand(printscreenCmd)

	printscreenCmd.PersistentFlags().String("site", "www.google.com", "Executa o screenshot da tela do site informado.")
}
