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
	desc      string
	curd      bool
)

var Cmd = &cobra.Command{
	Use:   "api",
	Short: "make api",
	Long: `Create a basic service for an application.
example: api -a test,Generate the base business code for the test application.`,
	RunE: GenApi,
}

func init() {
	Cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	Cmd.AddCommand(apiCmd)
}

func GenApi(_ *cobra.Command, _ []string) (err error) {

	if apply == "" {
		console.Error("invalid apply name.")
		return
	}

	// generate basic logic.
	filePath := fmt.Sprintf("app/http/%s/logic/logic.go", strings.ToLower(apply))
	err = GenLogic(filePath, FromStubBasic, ToLogic, nil)
	if err != nil {
		console.Error(err.Error())
		return
	}

	return
}
