package model

import (
	"errors"
	"fmt"
	"github.com/gin-ctl/zero/package/database"
	"github.com/gin-ctl/zero/package/get"
	"github.com/gin-ctl/zero/package/helper"
	"strings"
)

type Table struct {
	TableName string `gorm:"column:TABLE_NAME"`
	CamelCase string `gorm:"-"`
	Struct    string `gorm:"-"`
	Import    string `gorm:"-"`
}

type Column struct {
	Name                   string `gorm:"column:COLUMN_NAME"`
	Type                   string `gorm:"column:COLUMN_TYPE"`
	IsNullAble             string `gorm:"column:IS_NULLABLE"`
	CharacterMaximumLength *int   `gorm:"column:CHARACTER_MAXIMUM_LENGTH"`
	Extra                  string `gorm:"column:EXTRA"`
	Comment                string `gorm:"column:COLUMN_COMMENT"`
}

func GetTables(args string) (tables []*Table, err error) {
	// get tables.
	if args == `*` {
		err = database.DB.Raw("SELECT table_name FROM information_schema.tables WHERE table_schema=?;",
			get.String(fmt.Sprintf("db.%s.database", database.DB.Config.Name()))).
			Scan(&tables).Error
		if err == nil {
			for i, table := range tables {
				table.CamelCase = helper.CamelCase(table.TableName)
				tables[i] = table
			}
		}
	} else {
		for _, name := range strings.Split(args, ",") {
			// check table is existed.
			exist := database.DB.Migrator().HasTable(name)
			if !exist {
				err = errors.New(fmt.Sprintf("`%s` not found.", name))
				return
			}
			tables = append(tables, &Table{
				TableName: name,
				CamelCase: helper.CamelCase(name),
			})
		}
	}
	return
}

func GetColumn(tableName string) (columns []*Column, err error) {

	// get table columns.
	err = database.DB.Raw("SELECT COLUMN_NAME,COLUMN_TYPE,IS_NULLABLE,CHARACTER_MAXIMUM_LENGTH,EXTRA,COLUMN_COMMENT FROM INFORMATION_SCHEMA.COLUMNS WHERE TABLE_NAME = ? AND TABLE_SCHEMA = ?;",
		tableName, get.String(fmt.Sprintf("db.%s.database", database.DB.Config.Name()))).
		Scan(&columns).Error

	return
}

func GenerateStruct(tableName string, columns []*Column) string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprintf("type %s struct {\n", helper.CamelCase(tableName)))
	for _, col := range columns {
		fieldName := helper.CamelCase(col.Name)
		goType := mapSQLTypeToGoType(col.Type)
		jsonTag := fmt.Sprintf("json:\"%s\"", col.Name)
		gormTag := fmt.Sprintf("gorm:\"column:%s", col.Name)
		commitTag := ""
		if col.Comment != "" {
			commitTag = fmt.Sprintf("\t// %s", col.Comment)
		}

		if strings.ToLower(col.Extra) == "auto_increment" {
			gormTag = fmt.Sprintf("%s;primaryKey;autoIncrement\"", gormTag)
		} else {
			gormTag += "\""
		}

		//validateTag := "omitempty"
		//if col.IsNullAble == "NO" {
		//	validateTag = "required"
		//}
		//switch goType {
		//case "string":
		//	if col.CharacterMaximumLength != nil {
		//		validateTag = fmt.Sprintf("%s,max=%d", validateTag, *col.CharacterMaximumLength)
		//	}
		//case "int", "int8", "int16", "int32", "int64", "uint", "uint8", "uint16", "uint32", "uint64", "float32", "float64":
		//	validateTag = fmt.Sprintf("%s,numeric", validateTag)
		//case "time.Time":
		//	validateTag = "omitempty,datetime"
		//}
		//
		//if validateTag != "" {
		//	validateTag = fmt.Sprintf("validate:\"%s\"", validateTag)
		//}

		//builder.WriteString(fmt.Sprintf("    %s %s `%s %s %s`\n", fieldName, goType, jsonTag, gormTag, validateTag))
		builder.WriteString(fmt.Sprintf("    %s %s `%s %s`%s \n", fieldName, goType, jsonTag, gormTag, commitTag))
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
	case "json":
		return "json.RawMessage"
	case "date", "time", "datetime", "timestamp":
		return "time.Time"
	default:
		return "string"
	}
}
