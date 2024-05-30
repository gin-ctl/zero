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

func GenerateMiddleware() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "middleware",
		Short: "make middleware",
		Long: `Generate middleware.
Enter --table = * or -t = * to generate all table mapping structures. Multiple tables are separated by ",".
		`,
		RunE: GenMiddleware,
	}

	cmd.Flags().StringVarP(&middleware, "name", "n", "", "Specify middleware name")
	cmd.Flags().StringVarP(&apply, "apply", "a", "", "Specify apply name")
	cmd.Flags().BoolVarP(&global, "global", "g", false, "Specify global middleware")

	return cmd
}

func GenMiddleware(_ *cobra.Command, _ []string) (err error) {

	if middleware == "" {
		console.Error("invalid middleware name.")
		return
	}

	if apply != "" && global {
		console.Error("application middleware and global middleware cannot be specified at the same time.")
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		console.Error(err.Error())
		return
	}

	temp, err := template.ParseFiles(fmt.Sprintf("%s/tools/ginctl/middleware/stub/middleware.stub", pwd))
	if err != nil {
		console.Error(err.Error())
		return
	}

	var ware Ware
	lower := strings.ToLower(middleware)
	dir, pkg := "", ""
	if global {
		pkg = "middlewares"
		dir = fmt.Sprintf("%s/app/middleware", pwd)
	} else {
		if apply == "" {
			console.Error("the application name cannot be empty.")
			return
		}
		pkg = "middleware"
		dir = fmt.Sprintf("%s/app/http/%s/middleware", pwd, strings.ToLower(apply))
	}

	err = helper.CreateDirIfNotExist(dir)
	if err != nil {
		console.Error(err.Error())
		return
	}
	// check middleware is existed.
	filePath := fmt.Sprintf("%s/%s.go", dir, lower)
	if _, err = os.Stat(filePath); os.IsNotExist(err) {

		ware.Middleware = strcase.ToCamel(middleware)
		ware.Pkg = pkg

		// create middleware file.
		newFile, ers := os.Create(filePath)
		if ers != nil {
			console.Error(ers.Error())
			return
		}

		defer func(newFile *os.File) {
			e := newFile.Close()
			if e != nil {
				console.Error(e.Error())
			}
		}(newFile)

		err = temp.Execute(newFile, ware)
		if err != nil {
			console.Error(err.Error())
			return
		}

		// insert offset.
		if global {
			routePath := fmt.Sprintf("%s/bootstrap/route.go", pwd)
			lines, errs := helper.ReadLines(routePath)
			if errs != nil {
				console.Error(errs.Error())
				return
			}
			modifiedLines := helper.InsertOffset(
				lines, fmt.Sprintf("\t\t%s.%s(),", pkg, ware.Middleware), "// {{.GlobalMiddleware}}")
			err = helper.WriteLines(routePath, modifiedLines)
			if err != nil {
				console.Error(err.Error())
				return
			}
		} else {
			// TODO apply middleware.
		}
		console.Success("Done.")
	}

	return
}
