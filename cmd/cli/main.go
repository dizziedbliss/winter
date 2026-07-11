package cli

import "github.com/dizziedbliss/winter/internal/deployment"
import "github.com/spf13/cobra"

var deploycmd = &cobra.Command{
	Use:   "deploy",
	Short: "Builds a Docker Image of a project",
	Long: `Locates dockerfile, tests the dockerfile, and builds a docker image of the project
		
command usage: 
	winter deploy --directory <path to project>
	winter deploy -d <path to project>`,

	RunE: func(cmd *cobra.Command, args []string) error {

		verbose, err := cmd.Flags().GetBool("verbose")
		if err != nil {
			return err
		}

		dir, err := cmd.Flags().GetString("directory")
		if err != nil {
			return err
		}

		return deployment.Deploy(deployment.DeploymentOpts{
			Path:    dir,
			Verbose: verbose,
		})
	},
}

func init() {

	deploycmd.Flags().StringP("directory", "d", ".", `Mention the Project to build (e.g., ".", "/user/username/project/")`)
}

func Deploycmd() *cobra.Command {
	return deploycmd
}
