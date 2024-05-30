package api

import "github.com/spf13/cobra"

func GenerateApi() *cobra.Command {
	var (
		path  string
		model string
	)

	cmd := &cobra.Command{
		Use:   "api",
		Short: "make api",
		Long:  "Specifies the standard restful APIs for table generation",
		RunE:  GenApi,
	}

	cmd.Flags().StringVarP(&path, "path", "p", "", "指定api文件的路径")
	cmd.Flags().StringVarP(&model, "model", "m", "", "指定模块名")

	return cmd
}

func GenApi(cmd *cobra.Command, args []string) (err error) {

	return
}
