package route

import (
	"fmt"
	"github.com/gin-ctl/zero/package/console"
	"github.com/gin-ctl/zero/package/helper"
	"github.com/iancoleman/strcase"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"text/template"
)

var (
	apply string
)

type Template struct {
	Route string
}

var Cmd = &cobra.Command{
	Use:   "route",
	Short: "make route",
	Long:  ``,
	RunE:  GenRoute,
}

func init() {
	Cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
}

func GenRoute(_ *cobra.Command, _ []string) (err error) {

	if apply == "" {
		console.Error("invalid apply name.")
		return
	}
	apply = strings.ToLower(apply)

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

	filePath := fmt.Sprintf("%s/route.go", dir)
	if _, errs := os.Stat(filePath); os.IsNotExist(errs) {
		temp, ee := template.ParseFiles(fmt.Sprintf("%s/tools/ginctl/route/stub/route.stub", pwd))
		if ee != nil {
			console.Error(ee.Error())
			return
		}

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

		imports := fmt.Sprintf("%s \"%s/app/http/%s/route\"", apply, helper.GetModule(pwd), apply)
		route := fmt.Sprintf("%s/bootstrap/route.go", pwd)
		err = helper.InsertImport(route, imports, "import ", "")
		if err != nil {
			console.Error(err.Error())
			return
		}

		content, er := helper.GetFileContent(fmt.Sprintf("%s/tools/ginctl/route/stub/register_route.stub", pwd))
		if er != nil {
			console.Error(er.Error())
			return
		}
		camelApply := strcase.ToCamel(apply)
		content = fmt.Sprintf(content, camelApply, apply, camelApply)
		err = helper.AppendToFile(route, content)
		if err != nil {
			console.Error(err.Error())
			return
		}
		console.Success(fmt.Sprintf("Exist %s route done.", apply))
	}
	return
}
