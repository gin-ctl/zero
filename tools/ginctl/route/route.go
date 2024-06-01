package route

import (
	"fmt"
	"github.com/gin-ctl/zero/package/console"
	"github.com/gin-ctl/zero/package/helper"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"os"
	"text/template"
)

var (
	apply string
)

type Template struct {
	Route string
}

func GenerateRoute() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "route",
		Short: "make route",
		Long: `
		`,
		RunE: GenRoute,
	}

	cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")

	return cmd
}

func GenRoute(_ *cobra.Command, _ []string) (err error) {

	if apply == "" {
		console.Error("invalid apply name.")
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		console.Error(err.Error())
		return
	}
	dir := fmt.Sprintf("%s/app/http/%s/route", pwd, apply)
	err = helper.CreateDirIfNotExist(dir)
	if err != nil {
		console.Error(err.Error())
		return
	}

	temp, err := template.ParseFiles(fmt.Sprintf("%s/tools/ginctl/route/stub/route.stub", pwd))
	if err != nil {
		console.Error(err.Error())
		return
	}

	filePath := fmt.Sprintf("%s/route.go", dir)
	if _, errs := os.Stat(filePath); os.IsNotExist(errs) {
		var r Template
		r.Route = strcase.ToCamel(apply)

		newFile, ers := os.Create(filePath)
		if ers != nil {
			console.Error(ers.Error())
			return
		}

		defer newFile.Close()

		err = temp.Execute(newFile, r)
		if err != nil {
			console.Error(err.Error())
			return
		}

		insert := fmt.Sprintf(
			"func Register%sApiRoute(router *gin.Engine) {\n\t// route not found.\n\thttp.Alert404Route(router)\n\t// global middleware.\n\tRegisterGlobalMiddleware(router)\n\t// Initialize route.\n\troute.Register%sAPI(router)\n}",
			r.Route, r.Route)
		routePath := fmt.Sprintf("%s/bootstrap/route.go", pwd)
		lines, er := helper.ReadLines(routePath)
		if er != nil {
			console.Error(er.Error())
			return
		}
		modifiedLines := helper.InsertOffset(
			lines, insert, "// {{.ApiRoute}}")
		err = helper.WriteLines(routePath, modifiedLines)
		if err != nil {
			console.Error(err.Error())
			return
		}

		imports := fmt.Sprintf("\t\"github.com/gin-ctl/zero/app/http/%s/route\"", apply)
		lines, er = helper.ReadLines(routePath)
		if er != nil {
			console.Error(er.Error())
			return
		}
		modifiedLines = helper.InsertOffset(
			lines, imports, "// {{.Import}}")
		err = helper.WriteLines(routePath, modifiedLines)
		if err != nil {
			console.Error(err.Error())
			return
		}
	}
	console.Success("Done.")
	return
}
