package cmd

import (
	"cli-go/utils"
	"fmt"

	"github.com/spf13/cobra"
)

// whoisCmd represents the whois command
var whoisCmd = &cobra.Command{
	Use:   "whois",
	Short: "Whois domain",
	Long: `Whois domain.`,
	Run: func(cmd *cobra.Command, args []string) {
		host := cmd.Flag("host").Value.String()

		if host == "" {
			fmt.Println("Insira o host a consultado no whois.")
			return
		} else {
			fmt.Println(utils.GetWhois(host))			
		}
	},
}

func init() {
	rootCmd.AddCommand(whoisCmd)
 
	whoisCmd.PersistentFlags().String("host", "", "host no whois")

}
