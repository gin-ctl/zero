package apply

import "github.com/spf13/cobra"

var (
	apply          string
	withModel      bool
	withMiddleware bool
)

var Cmd = &cobra.Command{
	Use:   "apply",
	Short: "make apply",
	Long:  `Generate an http api application.`,
	RunE:  GenHttpApply,
}

func init() {
	Cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	Cmd.Flags().BoolVarP(&withModel, "model", "m", false, "Specify apply name")
	Cmd.Flags().BoolVarP(&withMiddleware, "middleware", "w", false, "Specify apply name")
}

func GenHttpApply(cmd *cobra.Command, args []string) (err error) {

	return
}
