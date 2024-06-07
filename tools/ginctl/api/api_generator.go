package api

import (
	"fmt"
	"github.com/gin-ctl/zero/package/helper"
	"os"
	"path/filepath"
	"text/template"
)

type Logic struct {
	Content string
}

// GenBasicLogic generate apply basic logic.
func GenBasicLogic(filePath string) (err error) {
	dir := filepath.Dir(filePath)

	err = helper.CreateDirIfNotExist(dir)
	if err != nil {
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		return
	}

	tmp, err := template.ParseFiles(fmt.Sprintf("%s/api/stub/basic_logic.stub", pwd))
	if err != nil {
		return
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer outFile.Close()

	var logic Logic
	logic.Content, err = helper.GetFileContent(fmt.Sprintf("%s/api/stub/logic.stub", pwd))
	if err != nil {
		return
	}
	err = tmp.Execute(outFile, logic)
	return
}
