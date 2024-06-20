package main

import (
	"github.com/gin-ctl/zero/tools/ginctl/api"
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
		model.Cmd,
		api.Cmd,
		middleware.Cmd,
		route.Cmd,
		apply.Cmd,
		grpc.Cmd,
	)

	// Execute command.
	cobra.CheckErr(rootCmd.Execute())

}
