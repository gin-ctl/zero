package apply

import "github.com/spf13/cobra"

var (
	apply          string
	withModel      bool
	withMiddleware bool
)

func GenerateHttpApply() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "apply",
		Short: "make apply",
		Long:  `Generate an http api application.`,
		RunE:  GenHttpApply,
	}

	cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	cmd.Flags().BoolVarP(&withModel, "model", "m", false, "Specify apply name")
	cmd.Flags().BoolVarP(&withMiddleware, "middleware", "w", false, "Specify apply name")

	return cmd
}

func GenHttpApply(cmd *cobra.Command, args []string) (err error) {

	return
}
