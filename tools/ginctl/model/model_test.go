package model

import "testing"

func TestGetGotype(t *testing.T) {
	columnType := []string{
		"int unsigned",
		"enum('READ','WRITE','DELETE')",
		"timestamp",
	}

	for _, s := range columnType {
		gotype := mapSQLTypeToGoType(s)
		t.Log(gotype)
	}
}
