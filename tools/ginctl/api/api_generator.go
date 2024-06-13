package api

import (
	"bytes"
	"fmt"
	"github.com/gin-ctl/zero/package/helper"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Logic struct {
	Content string
}

type Body struct {
	LowerModel string
	Apply      string
}

type Operation struct {
	Opt string
}

type StubCode uint

const (
	FromStubBasic StubCode = iota
	FromStubImport
	FromStubLogicFunc
	FromStubTypes
	FromStubTypeStruct
	FromStubTypeFunc
	ToLogic
)

var StubMap = map[StubCode]string{
	FromStubBasic:      "%s/tools/ginctl/api/stub/basic_logic.stub",
	FromStubImport:     "%s/tools/ginctl/api/stub/logic_import.stub",
	FromStubLogicFunc:  "%s/tools/ginctl/api/stub/logic_func.stub",
	FromStubTypes:      "%s/tools/ginctl/api/stub/types.stub",
	FromStubTypeStruct: "%s/tools/ginctl/api/stub/type_struct.stub",
	FromStubTypeFunc:   "%s/tools/ginctl/api/stub/type_func.stub",
	ToLogic:            "%s/tools/ginctl/api/stub/logic.stub",
}

// GenLogic generate apply logic.
func GenLogic(filePath string, from, to StubCode, body *Body) (err error) {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}

	dir := fmt.Sprintf("%s/%s", pwd, strings.TrimLeft(filepath.Dir(filePath), "/"))
	err = helper.CreateDirIfNotExist(dir)
	if err != nil {
		return
	}

	filePath = fmt.Sprintf("%s/%s", pwd, strings.TrimLeft(filePath, "/"))
	if helper.PathExists(filePath) {
		return
	}

	tmp, err := template.ParseFiles(fmt.Sprintf(StubMap[to], pwd))
	if err != nil {
		return
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		return
	}

	var logic Logic
	logic.Content, err = helper.GetFileContent(fmt.Sprintf(StubMap[from], pwd))
	if err != nil {
		return
	}
	err = tmp.Execute(outFile, logic)
	if err != nil {
		return
	}

	err = outFile.Close()
	if err != nil {
		return
	}

	if body != nil {
		tmp, err = template.ParseFiles(filePath)
		if err != nil {
			return
		}

		var output bytes.Buffer
		err = tmp.Execute(&output, body)
		if err != nil {
			return
		}

		err = os.WriteFile(filePath, output.Bytes(), os.ModePerm)
		return
	}

	return
}

func ExecuteContent(filePath string, body *Body) (err error) {
	pwd, err := os.Getwd()
	if err != nil {
		return
	}

	filePath = fmt.Sprintf("%s/%s", pwd, strings.TrimLeft(filePath, "/"))
	tmp, err := template.ParseFiles(filePath)
	if err != nil {
		return
	}

	var output bytes.Buffer
	err = tmp.Execute(&output, body)
	if err != nil {
		return
	}

	err = os.WriteFile(filePath, output.Bytes(), os.ModePerm)

	return
}
