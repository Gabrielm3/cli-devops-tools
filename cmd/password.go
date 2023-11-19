package cmd

import (
	"cli-go/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// passwordCmd represents the password command
var passwordCmd = &cobra.Command{
	Use:   "password",
	Short: "password",
	Long: `Generate Password`,
	Run: func(cmd *cobra.Command, args []string) {
		tamanho, _ := cmd.Flags().GetInt("t")
		fmt.Println(utils.GeneratePassword(tamanho))
	},
}

func init() {
	rootCmd.AddCommand(passwordCmd)
	
	passwordCmd.PersistentFlags().Int("t", 16, "Tamanho da senha gerada.")
}