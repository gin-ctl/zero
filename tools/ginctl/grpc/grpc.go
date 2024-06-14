package grpc

import "github.com/spf13/cobra"

var Cmd = &cobra.Command{
	Use:   "grpc",
	Short: "make grpc apply",
	Long:  `Generate an grpc application.`,
	RunE:  GenGrpcApply,
}

func GenGrpcApply(cmd *cobra.Command, args []string) (err error) {

	return
}
