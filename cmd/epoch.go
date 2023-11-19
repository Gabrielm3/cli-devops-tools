package cmd

import (
	"cli-go/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// epochCmd represents the epoch command
var epochCmd = &cobra.Command{
	Use:   "epoch",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		epoch := cmd.Flag("c").Value.String()
		fmt.Println(utils.ConvertEpochToHuman(epoch))
	},
}

func init() {
	rootCmd.AddCommand(epochCmd)

	epochCmd.PersistentFlags().String("c", "", "Converter epoch para data/hora")
}
