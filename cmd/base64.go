package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// base64Cmd represents the base64 command
var base64Cmd = &cobra.Command{
	Use:   "base64",
	Short: "Efetua o encode e decode de strings",
	Long: `Efetua o encode e decode de strings para base64.
	Exemplo de uso:
		encode: ./main base64 --e "string"
		decode: ./main base64 --d "string"
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("base64 called")
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)
	
}
