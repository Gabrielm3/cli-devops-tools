package cmd

import (
	"cli-go/utils"

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
		encodeStr, _ := cmd.Flags().GetString("e")
		decodeStr, _ := cmd.Flags().GetString("d")

		if encodeStr != "" {
			encode := utils.EncodeString(encodeStr)
			cmd.Println(encode)
		}

		if decodeStr != "" {
			decode := utils.DecodeString(decodeStr)
			cmd.Println(decode)
		}
	},
}

func init() {
	rootCmd.AddCommand(base64Cmd)

	base64Cmd.PersistentFlags().String("e", "", "Encode string")
	base64Cmd.PersistentFlags().String("d", "", "Decode string")
}
