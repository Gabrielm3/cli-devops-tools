package cmd

import (
	"cli-go/utils"
	"strings"

	"github.com/spf13/cobra"
)

// checkportCmd represents the checkport command
var checkportCmd = &cobra.Command{
	Use:   "checkport",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		host, _ := cmd.Flags().GetString("h")
		if host == "" {
			cmd.Println("Host a ser validado")
			return
		}

		ports, _ := cmd.Flags().GetString("p")
		if ports == "" {
			cmd.Println("Porta / lista de portas a serem validadas.")
			return
		}
		portList := strings.Split(ports, ",")
		utils.CheckPort(host, portList)
	},
}

func init() {
	rootCmd.AddCommand(checkportCmd)

	checkportCmd.PersistentFlags().String("h", "", "Host a ser validado")
	checkportCmd.PersistentFlags().String("p", "", "Lista de portas separadas por virgula. Ex: 80, 443, 8080")
}
