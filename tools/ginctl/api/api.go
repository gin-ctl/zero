package api

import "github.com/spf13/cobra"

var (
	apply string
	model string
	api   string
)

func GenerateApi() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model",
		Short: "make model",
		Long: `
		`,
		RunE: GenApi,
	}

	cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	cmd.Flags().StringVarP(&model, "model", "m", "", "Specify model name")
	cmd.Flags().StringVarP(&api, "api", "i", "", "Specify api name")

	return cmd
}

func GenApi(_ *cobra.Command, _ []string) (err error) {
	err = GenBasicLogic("")
	return
}
