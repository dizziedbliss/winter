package deployment

import (
	"fmt"

	"github.com/spf13/cobra"
)

// deployCmd represents the deploy command

func Deploy() *cobra.Command {
	return &cobra.Command{
		Use:   "deploy",
		Short: "A brief description of your command",
		Long: `A longer description that spans multiple lines and likely contains examples
		and usage of using your command. For example:

		Cobra is a CLI library for Go that empowers applications.
		This application is a tool to generate the needed files
		to quickly create a Cobra application.`,
		
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("deploy called")
		},
	}
}
