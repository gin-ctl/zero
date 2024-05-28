package model

import (
	"github.com/gin-ctl/zero/package/console"
	"github.com/spf13/cobra"
)

func GenerateModelStruct() *cobra.Command {
	var (
		tableName string
	)

	cmd := &cobra.Command{
		Use:   "model",
		Short: "make model",
		Long:  "Specifies the standard restful APIs for table generation",
		RunE:  GenModelStruct,
	}

	cmd.Flags().StringVarP(&tableName, "table", "t", "", "Specify table name")

	return cmd
}

func GenModelStruct(cmd *cobra.Command, args []string) (err error) {

	columns, err := GetColumn("policies")
	if err != nil {
		console.Error(err.Error())
		return
	}

	str := GenerateStruct(columns)
	console.Success(str)

	return
}
