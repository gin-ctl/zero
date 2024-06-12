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

type StubCode uint

const (
	FromStubBasic StubCode = iota + 1
	FromStubImport
	FromStubLogicFunc
	FromStubTypes
	FromStubTypeStruct
	FromStubTypeFunc
	ToLogic
)

var StubMap = map[StubCode]string{
	FromStubBasic:      "%s/api/stub/basic_logic.stub",
	FromStubImport:     "%s/api/stub/logic_import.stub",
	FromStubLogicFunc:  "%s/api/stub/logic_func.stub",
	FromStubTypes:      "%s/api/stub/types.stub",
	FromStubTypeStruct: "%s/api/stub/type_struct.stub",
	FromStubTypeFunc:   "%s/api/stub/type_func.stub",
	ToLogic:            "%s/api/stub/logic.stub",
}

// GenLogic generate apply logic.
func GenLogic(filePath string, from, to StubCode) (err error) {
	dir := filepath.Dir(filePath)

	err = helper.CreateDirIfNotExist(dir)
	if err != nil {
		return
	}

	pwd, err := os.Getwd()
	if err != nil {
		return
	}

	tmp, err := template.ParseFiles(fmt.Sprintf(StubMap[from], pwd))
	if err != nil {
		return
	}

	outFile, err := os.Create(filePath)
	if err != nil {
		return
	}
	defer outFile.Close()

	var logic Logic
	logic.Content, err = helper.GetFileContent(fmt.Sprintf(StubMap[to], pwd))
	if err != nil {
		return
	}
	err = tmp.Execute(outFile, logic)
	return
}
