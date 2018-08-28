package convert

import (
	"reflect"

	"github.com/verystar/golib/reflectx"
)

var mapper = reflectx.NewMapper("db")

func StructToMapInterface(m interface{}) map[string]interface{} {
	v := reflect.ValueOf(m)
	fields := mapper.FieldMap(v)

	rs := make(map[string]interface{})
	for k, v := range fields {
		rs[k] = v.Interface()
	}

	return rs
}

func StructToMapString(m interface{}) map[string]string {
	maps := StructToMapInterface(m)
	rs := make(map[string]string)

	for k, v := range maps {
		rs[k] = ToStr(v)
	}

	return rs
}
