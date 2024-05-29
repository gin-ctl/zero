package model

import (
	"fmt"
	"github.com/gin-ctl/zero/bootstrap"
	"github.com/gin-ctl/zero/package/console"
	"github.com/gin-ctl/zero/package/get"
	"github.com/spf13/cobra"
	"os"
	"strings"
	"sync"
	"text/template"
)

var (
	tableName string
)

func GenerateModelStruct() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "model",
		Short: "make model",
		Long:  "Generates a mapping structure for a table based on the database table name.",
		RunE:  GenModelStruct,
	}

	cmd.Flags().StringVarP(&tableName, "table", "t", "*", "Specify table name")

	return cmd
}

func GenModelStruct(_ *cobra.Command, _ []string) (err error) {

	if tableName == "" {
		console.Error("table name invalid.")
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		console.Error(err.Error())
		return
	}

	// get sql database.
	database := get.String(fmt.Sprintf("db.%s.database", bootstrap.DB.Config.Name()))
	// get dir.
	dir := fmt.Sprintf("%s/model/%s", pwd, database)
	if _, e := os.Stat(dir); os.IsNotExist(e) {
		errs := os.Mkdir(dir, os.ModePerm)
		if errs != nil {
			console.Error(errs.Error())
			return
		}
	}

	temp, err := template.ParseFiles(fmt.Sprintf("%s/tools/model/stub/model.stub", pwd))
	if err != nil {
		console.Error(err.Error())
		return
	}

	tables, err := GetTables(tableName)
	if err != nil {
		console.Error(err.Error())
		return
	}

	var wg sync.WaitGroup
	errChan := make(chan error, 10)

	for _, table := range tables {
		filePath := fmt.Sprintf("%s/%s.go", dir, table.TableName)
		// check table struct is existed.
		if _, errs := os.Stat(filePath); os.IsNotExist(errs) {
			wg.Add(1)
			go func(table *Table, filePath string, wg *sync.WaitGroup, errChan chan error) {
				defer wg.Done()

				columns, ers := GetColumn(table.TableName)
				if ers != nil {
					errChan <- ers
					return
				}

				table.Struct = GenerateStruct(table.TableName, columns)

				// Handling import packages.
				pkg := ""
				if strings.Contains(table.Struct, "json.RawMessage") {
					pkg += "\"encoding/json\"\n"
				}
				if strings.Contains(table.Struct, "") {
					pkg += "\t\"github.com/gin-ctl/zero/package/time\""
				}
				if pkg != "" {
					table.Import = fmt.Sprintf("import (\n  %s\n)", pkg)
				}

				newFile, ers := os.Create(filePath)
				if ers != nil {
					errChan <- ers
					return
				}
				defer func(newFile *os.File) {
					e := newFile.Close()
					if e != nil {
						errChan <- e
					}
				}(newFile)

				err = temp.Execute(newFile, table)
				if err != nil {
					errChan <- err
					return
				}
			}(table, filePath, &wg, errChan)
		}
	}

	go func() {
		wg.Wait()
		close(errChan)
	}()

	for e := range errChan {
		console.Error(e.Error())
	}
	console.Success("Done.")

	return
}
