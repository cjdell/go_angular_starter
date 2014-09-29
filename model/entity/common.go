package entity

import (
	"reflect"
	"strings"
)

type Entity interface {
	GetId() int64
	SetId(id int64)
}

func GetDbTable(entityType reflect.Type) string {
	field, found := entityType.FieldByName("__table")

	if !found {
		return ""
	}

	return field.Tag.Get("db")
}

func GetDbColumns(entityType reflect.Type) []string {
	fields := make([]string, entityType.NumField()-1)

	for i := 0; i < entityType.NumField(); i++ {
		var (
			field      = entityType.Field(i)
			fieldName  = field.Name
			columnName = field.Tag.Get("db")
		)

		if !strings.HasPrefix(fieldName, "__") {
			fields[i-1] = columnName
		}
	}

	return fields
}

func GetDbColumnValues(entity Entity) map[string]interface{} {
	var (
		entityType  = reflect.TypeOf(entity).Elem()
		entityValue = reflect.ValueOf(entity).Elem()
	)

	fieldValues := make(map[string]interface{})

	for i := 0; i < entityType.NumField(); i++ {
		var (
			field      = entityType.Field(i)
			fieldName  = field.Name
			columnName = field.Tag.Get("db")
			fieldValue = entityValue.Field(i).Interface()
		)

		if !strings.HasPrefix(fieldName, "__") {
			fieldValues[columnName] = fieldValue
		}
	}

	return fieldValues
}
