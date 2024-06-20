package middleware

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
	middleware string
	apply      string
	global     bool
)

type Ware struct {
	Middleware string
	Pkg        string
}

var Cmd = &cobra.Command{
	Use:   "middleware",
	Short: "make middleware",
	Long: `Generate middleware.
Example: middleware --apply demo --name auth.
Example: middleware --apply demo --name auth --global true.
`,
	RunE: GenMiddleware,
}

func init() {
	Cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	Cmd.Flags().StringVarP(&middleware, "name", "n", "", "Specify middleware name")
	Cmd.Flags().BoolVarP(&global, "global", "g", false, "Specify middleware name")
}

func GenMiddleware(_ *cobra.Command, _ []string) (err error) {

	if apply == "" {
		console.Error("the application name cannot be empty.")
		return
	}

	if middleware == "" {
		console.Error("invalid middleware name.")
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		console.Error(err.Error())
		return
	}

	var ware Ware
	lower := strings.ToLower(middleware)
	apply = strings.ToLower(apply)
	dir, pkg := "", ""
	if global {
		pkg = "middlewares"
		dir = fmt.Sprintf("%s/middleware", pwd)
	} else {
		pkg = "middleware"
		dir = fmt.Sprintf("%s/app/http/%s/middleware", pwd, apply)
	}

	err = helper.CreateDirIfNotExist(dir)
	if err != nil {
		console.Error(err.Error())
		return
	}
	// check middleware is existed.
	filePath := fmt.Sprintf("%s/%s.go", dir, lower)
	if _, err = os.Stat(filePath); os.IsNotExist(err) {

		temp, errs := template.ParseFiles(fmt.Sprintf("%s/tools/ginctl/middleware/stub/middleware.stub", pwd))
		if errs != nil {
			console.Error(errs.Error())
			return
		}

		ware.Middleware = strcase.ToCamel(middleware)
		ware.Pkg = pkg

		// create middleware file.
		newFile, ers := os.Create(filePath)
		if ers != nil {
			console.Error(ers.Error())
			return
		}
		defer newFile.Close()

		err = temp.Execute(newFile, ware)
		if err != nil {
			console.Error(err.Error())
			return
		}

		// insert offset.
		routePath, imports, t := "", "", ""
		module := helper.GetModule(pwd)
		if global {
			routePath = fmt.Sprintf("%s/bootstrap/route.go", pwd)
			imports = fmt.Sprintf("\"%s/middleware\"", module)
			t = "\t"
		} else {
			routePath = fmt.Sprintf("%s/app/http/%s/route/route.go", pwd, apply)
			imports = fmt.Sprintf("\"%s/app/http/%s/middleware\"", module, apply)
		}
		err = helper.InsertImport(routePath, imports, "import ", "")
		if err != nil {
			console.Error(err.Error())
			return
		}
		err = helper.InsertImport(routePath, fmt.Sprintf("\t%s.%s(),", pkg, ware.Middleware), "r.Use", t)
		if err != nil {
			console.Error(err.Error())
			return
		}
		console.Success(fmt.Sprintf("Create middleware of %s done.", lower))
	} else {
		console.Success(fmt.Sprintf("Middleware %s is existed.", lower))
	}

	return
}
