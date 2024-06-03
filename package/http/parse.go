package http

import (
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
)

const (
	PATH  = "path"
	QUERY = "query"
	FORM  = "form"
	JSON  = "json"
	XML   = "xml"
	FILE  = "file"
)

// RequestType is a generic struct that holds the parsed request data
type RequestType[T any] struct {
	data T
}

// NewRequestType is a constructor function for RequestType
func NewRequestType[T any](data T) RequestType[T] {
	return RequestType[T]{data: data}
}

// Data returns the parsed request data
func (r RequestType[T]) Data() T {
	return r.data
}

// Parse is a function that parses request parameters into the provided struct
func Parse[T any](c *gin.Context, obj *T) (err error) {

	// Handle the JSON binding first
	if err = c.ShouldBind(obj); err != nil {
		return
	}

	val := reflect.ValueOf(obj).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldType := typ.Field(i)
		tag := fieldType.Tag

		switch field.Kind() {
		case reflect.String:
			parseStringField(c, &field, tag)
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			parseIntField(c, &field, tag)
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			parseUintField(c, &field, tag)
		case reflect.Bool:
			parseBoolField(c, &field, tag)
		case reflect.Float32, reflect.Float64:
			parseFloatField(c, &field, tag)
		case reflect.Struct, reflect.Map, reflect.Slice, reflect.Array:
			parseComplexField(c, &field, tag)
		case reflect.Ptr:
			parsePtrField(c, &field, tag)
		}
	}

	return
}

// Helper function to parse string fields
func parseStringField(c *gin.Context, field *reflect.Value, tag reflect.StructTag) {
	if pathTag, ok := tag.Lookup(PATH); ok {
		field.SetString(c.Param(pathTag))
	}
	if queryTag, ok := tag.Lookup(QUERY); ok {
		field.SetString(c.Query(queryTag))
	}
}

// Helper function to parse integer fields
func parseIntField(c *gin.Context, field *reflect.Value, tag reflect.StructTag) {
	if pathTag, ok := tag.Lookup(PATH); ok {
		if va, err := strconv.ParseInt(c.Param(pathTag), 10, 64); err == nil {
			field.SetInt(va)
		}
	}
	if queryTag, ok := tag.Lookup(QUERY); ok {
		if va, err := strconv.ParseInt(c.Query(queryTag), 10, 64); err == nil {
			field.SetInt(va)
		}
	}
}

// Helper function to parse unsigned integer fields
func parseUintField(c *gin.Context, field *reflect.Value, tag reflect.StructTag) {
	if pathTag, ok := tag.Lookup(PATH); ok {
		if va, err := strconv.ParseUint(c.Param(pathTag), 10, 64); err == nil {
			field.SetUint(va)
		}
	}
	if queryTag, ok := tag.Lookup(QUERY); ok {
		if va, err := strconv.ParseUint(c.Query(queryTag), 10, 64); err == nil {
			field.SetUint(va)
		}
	}
}

// Helper function to parse boolean fields
func parseBoolField(c *gin.Context, field *reflect.Value, tag reflect.StructTag) {
	if pathTag, ok := tag.Lookup(PATH); ok {
		if va, err := strconv.ParseBool(c.Param(pathTag)); err == nil {
			field.SetBool(va)
		}
	}
	if queryTag, ok := tag.Lookup(QUERY); ok {
		if va, err := strconv.ParseBool(c.Query(queryTag)); err == nil {
			field.SetBool(va)
		}
	}
}

// Helper function to parse float fields
func parseFloatField(c *gin.Context, field *reflect.Value, tag reflect.StructTag) {
	if pathTag, ok := tag.Lookup(PATH); ok {
		if va, err := strconv.ParseFloat(c.Param(pathTag), 64); err == nil {
			field.SetFloat(va)
		}
	}
	if queryTag, ok := tag.Lookup(QUERY); ok {
		if va, err := strconv.ParseFloat(c.Query(queryTag), 64); err == nil {
			field.SetFloat(va)
		}
	}
}

// Helper function to parse complex fields (struct, map, slice, array)
func parseComplexField(c *gin.Context, field *reflect.Value, tag reflect.StructTag) {
	if _, ok := tag.Lookup(JSON); ok {
		ptr := field.Addr().Interface()
		if err := c.ShouldBindJSON(ptr); err == nil {
			field.Set(reflect.ValueOf(ptr).Elem())
		}
	}
	if _, ok := tag.Lookup(XML); ok {
		ptr := field.Addr().Interface()
		if err := c.ShouldBindXML(ptr); err == nil {
			field.Set(reflect.ValueOf(ptr).Elem())
		}
	}
}

// Helper function to parse pointer fields
func parsePtrField(c *gin.Context, field *reflect.Value, tag reflect.StructTag) {
	if fileTag, ok := tag.Lookup(FILE); ok {
		file, err := c.FormFile(fileTag)
		if err == nil {
			field.Set(reflect.ValueOf(file))
		}
	}
}
