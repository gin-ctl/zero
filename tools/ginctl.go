package main

import (
	"fmt"
	"github.com/gin-ctl/zero/bootstrap"
	"github.com/gin-ctl/zero/package/console"
	"github.com/gin-ctl/zero/package/get"
	"github.com/gin-ctl/zero/tools/api"
	"github.com/gin-ctl/zero/tools/model"
	"github.com/spf13/cobra"
	"os"
)

func main() {

	pwd, err := os.Getwd()
	if err != nil {
		console.Error(err.Error())
		return
	}
	if _, err = os.Stat(fmt.Sprintf("%s/config/env.yaml", pwd)); os.IsNotExist(err) {
		console.Error("config/env.yaml not found.")
		return
	}
	// Load configuration file.
	get.NewViper("env.yaml", fmt.Sprintf("%s/config", pwd))
	// Start basic services.
	bootstrap.SetupLogger()
	bootstrap.SetupDB()

	// This is a basic CLI application.
	var rootCmd = &cobra.Command{
		Use:   "ginctl",
		Short: "gin ctl",
		Long:  `This is a basic CLI application.`,
	}

	rootCmd.AddCommand(
		model.GenerateModelStruct(),
		api.GenerateSourceApi(),
	)

	// Execute command.
	cobra.CheckErr(rootCmd.Execute())

}
