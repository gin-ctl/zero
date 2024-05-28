package model

import (
	"errors"
	"fmt"
	"github.com/gin-ctl/zero/package/bootstrap"
	"strings"
	"unicode"
)

type Column struct {
	Name                   string `gorm:"COLUMN_NAME"`
	Type                   string `gorm:"COLUMN_TYPE"`
	IsNullAble             string `gorm:"IS_NULLABLE"`
	CharacterMaximumLength *int   `gorm:"CHARACTER_MAXIMUM_LENGTH"`
	Extra                  string `gorm:"EXTRA"`
	Comment                string `json:"COLUMN_COMMENT"`
}

func GetColumn(tableName string) (columns []*Column, err error) {

	// check table is exist.
	exist := bootstrap.DB.Migrator().HasTable(tableName)
	if !exist {
		err = errors.New(fmt.Sprintf("`%s` not found.", tableName))
		return
	}

	// get table columns.
	query := fmt.Sprintf("SELECT COLUMN_NAME,COLUMN_TYPE,IS_NULLABLE,CHARACTER_MAXIMUM_LENGTH,EXTRA,COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = '%s';", tableName)
	err = bootstrap.DB.Raw(query).Scan(&columns).Error
	if err != nil {
		return
	}

	return
}

func GenerateStruct(columns []*Column) string {
	var builder strings.Builder
	builder.WriteString("type YourStruct struct {\n")
	for _, col := range columns {
		fieldName := camelCase(col.Name)
		goType := mapSQLTypeToGoType(col.Type)
		jsonTag := fmt.Sprintf("json:\"%s\"", fieldName)
		gormTag := fmt.Sprintf("gorm:\"column:%s\"", col.Name)
		validateTag := "omitempty"

		if col.IsNullAble == "YES" {
			validateTag = "required"
		}

		switch goType {
		case "string":
			if col.CharacterMaximumLength != nil {
				validateTag = fmt.Sprintf("%s,max=%d", validateTag, *col.CharacterMaximumLength)
			}
		case "int", "int8", "int16", "int32", "int64", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
			validateTag = fmt.Sprintf("%s,numeric", validateTag)
		case "time.Time":
			validateTag = fmt.Sprintf("%s,datetime", validateTag)
		}

		if validateTag != "" {
			validateTag = fmt.Sprintf("validate:\"%s\"", validateTag)
		}

		builder.WriteString(fmt.Sprintf("    %s %s `%s %s %s`\n", fieldName, goType, jsonTag, gormTag, validateTag))
	}
	builder.WriteString("}\n")
	return builder.String()
}

func mapSQLTypeToGoType(sqlType string) string {
	switch strings.ToLower(sqlType) {
	case "year":
		return "int"
	case "tinyint":
		return "int8"
	case "tinyint unsigned":
		return "uint8"
	case "smallint":
		return "int16"
	case "smallint unsigned":
		return "uint16"
	case "mediumint":
		return "int32"
	case "mediumint unsigned":
		return "uint32"
	case "int", "integer":
		return "int32"
	case "int unsigned":
		return "uint32"
	case "bigint":
		return "int64"
	case "bigint unsigned":
		return "uint64"
	case "float":
		return "float32"
	case "double", "real", "decimal", "numeric":
		return "float64"
	case "bit", "binary", "varbinary", "tinyblob", "blob", "mediumblob", "longblob":
		return "[]byte"
	case "varchar", "char", "text", "tinytext", "mediumtext", "longtext", "enum", "set":
		return "string"
	case "date", "time", "datetime", "timestamp":
		return "time.Time"
	default:
		return "string"
	}
}

func camelCase(input string) string {
	var output []rune
	toUpper := true

	for _, r := range input {
		if r == '_' {
			toUpper = true
			continue
		}
		if toUpper {
			output = append(output, unicode.ToUpper(r))
			toUpper = false
		} else {
			output = append(output, unicode.ToLower(r))
		}
	}

	return string(output)
}
