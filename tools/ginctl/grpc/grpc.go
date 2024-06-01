package grpc

import "github.com/spf13/cobra"

func GenerateGrpcApply() *cobra.Command {

	cmd := &cobra.Command{
		Use:   "grpc",
		Short: "make grpc apply",
		Long:  `Generate an grpc application.`,
		RunE:  GenGrpcApply,
	}

	//cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	//cmd.Flags().BoolVarP(&withModel, "model", "m", false, "Specify apply name")
	//cmd.Flags().BoolVarP(&withMiddleware, "middleware", "w", false, "Specify apply name")

	return cmd
}

func GenGrpcApply(cmd *cobra.Command, args []string) (err error) {

	return
}
