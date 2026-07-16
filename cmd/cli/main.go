package cli

import (
	"github.com/spf13/cobra"

	"github.com/dizziedbliss/winter/internal/project"
	"github.com/dizziedbliss/winter/internal/ui/bubbletea"
)

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

		bubble := bubbletea.New()
		proj := project.NewProject(
			dir,
			verbose,
			bubble,
		)

		return bubble.Run(func() error {
				return proj.Deploy()
			})
	},
}

func init() {

	deploycmd.Flags().StringP("directory", "d", ".", `Mention the Project to build (e.g., ".", "/user/username/project/")`)
}

func Deploycmd() *cobra.Command {
	return deploycmd
}
