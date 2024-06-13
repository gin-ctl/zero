package api

import (
	"fmt"
	"github.com/gin-ctl/zero/package/console"
	"github.com/spf13/cobra"
	"strings"
)

var (
	apply     string
	model     string
	operation string
	curd      bool
)

func GenerateApi() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "api",
		Short: "make api",
		Long:  ``,
		RunE:  GenApi,
	}

	cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	cmd.Flags().StringVarP(&model, "model", "m", "", "Specify model name")

	cmd.AddCommand(apiCmd)
	apiCmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	apiCmd.Flags().StringVarP(&model, "model", "m", "", "Specify model name")
	apiCmd.Flags().StringVarP(&operation, "operation", "o", "", "Specify operation name")
	apiCmd.Flags().BoolVarP(&curd, "curd", "c", false, "Specifies whether you need to generate add, delete, update and get operations for the module")

	return cmd
}

func GenApi(cmd *cobra.Command, args []string) (err error) {

	if apply == "" {
		console.Error("invalid apply name.")
		return
	}

	if model == "" {
		console.Error("invalid model name.")
		return
	}

	// generate basic logic.
	filePath := fmt.Sprintf("app/http/%s/logic/logic.go", strings.ToLower(apply))
	err = GenLogic(filePath, FromStubBasic, ToLogic, nil)
	if err != nil {
		console.Error(err.Error())
		return
	}

	// execute subcommand.
	childCmd, _, err := cmd.Find(args)
	if err != nil {
		return err
	}
	if childCmd != nil && childCmd != cmd {
		return childCmd.Execute()
	}
	console.Success("Done.")
	return
}
