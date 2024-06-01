package main

import (
	"github.com/gin-ctl/zero/tools/ginctl/apply"
	"github.com/gin-ctl/zero/tools/ginctl/grpc"
	"github.com/gin-ctl/zero/tools/ginctl/middleware"
	"github.com/gin-ctl/zero/tools/ginctl/model"
	"github.com/gin-ctl/zero/tools/ginctl/route"
	"github.com/spf13/cobra"
)

func main() {

	// This is a basic CLI application.
	var rootCmd = &cobra.Command{
		Use:   "ginctl",
		Short: "gin ctl",
		Long:  `This is a basic CLI application.`,
	}

	rootCmd.AddCommand(
		apply.GenerateHttpApply(),
		grpc.GenerateGrpcApply(),
		model.GenerateModelStruct(),
		middleware.GenerateMiddleware(),
		route.GenerateRoute(),
	)

	// Execute command.
	cobra.CheckErr(rootCmd.Execute())

}
