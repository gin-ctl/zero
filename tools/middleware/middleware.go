package middleware

import "github.com/spf13/cobra"

var (
	middleware string
)

func GenerateMiddleware() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "middleware",
		Short: "make middleware",
		Long: `Generates a mapping structure for a table based on the database table name.
Enter --table = * or -t = * to generate all table mapping structures. Multiple tables are separated by ",".
		`,
		RunE: GenMiddleware,
	}

	cmd.Flags().StringVarP(&middleware, "middleware", "mw", "", "Specify table name")

	return cmd
}

func GenMiddleware(_ *cobra.Command, _ []string) (err error) {
	return
}
