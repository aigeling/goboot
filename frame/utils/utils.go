package utils

import (
	"reflect"
	"strings"
)

func IsNil(o any) bool {
	v := reflect.ValueOf(o)
	if v.Kind() == reflect.Ptr {
		return v.IsNil()
	}
	return false
}

func IsSlice(o any) bool {

	v := reflect.ValueOf(o)
	return v.Kind().String() == "slice"
}

// 获得结构体字段tag中json的name
func TagJsonName(o any, fieldName string) (jsonName string) {
	v := reflect.TypeOf(o)

	f, ok := v.FieldByName("Name")
	if ok {
		jsonTag := f.Tag.Get("json")
		if jsonTag != "" {
			jsonName = strings.Split(jsonTag, ",")[0]
		} else {
			jsonName = fieldName
		}
	}

	return
}
