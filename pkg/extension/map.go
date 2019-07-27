/**
 * @Author: huangw1
 * @Date: 2019/7/30 20:14
 */

package extension

import (
	"reflect"
	"strings"
	"unicode"
	"unicode/utf8"
)

func StructToMap(source interface{}, excludes ...string) map[string]interface{} {
	data := make(map[string]interface{})
	fields := reflect.TypeOf(source)
	values := reflect.ValueOf(source)
	fillMap(data, fields, values, excludes...)
	return data
}

func fillMap(data map[string]interface{}, fields reflect.Type, values reflect.Value, excludes ...string) {
	if fields.Kind() == reflect.Ptr {
		fields = fields.Elem()
	}
	if values.Kind() == reflect.Ptr {
		values = values.Elem()
	}
	for i := 0; i < fields.NumField(); i++ {
		field := fields.Field(i)
		value := values.Field(i)
		if field.Anonymous {
			fillMap(data, field.Type, value, excludes...)
		} else if !ContainsIgnoreCase(field.Name, excludes...) {
			if IsExported(field.Name) {
				name := field.Tag.Get("json")
				if len(name) > 0 {
					data[name] = value.Interface()
				} else {
					data[field.Name] = value.Interface()
				}
			}
		}
	}
}

func ContainsIgnoreCase(search string, excludes ...string) bool {
	if len(excludes) == 0 {
		return false
	}
	if len(search) == 0 {
		return false
	}
	search = strings.ToLower(search)
	for _, str := range excludes {
		if strings.ToLower(str) == search {
			return true
		}
	}
	return false
}

func IsExported(field string) bool {
	f, _ := utf8.DecodeRuneInString(field)
	return unicode.IsUpper(f)
}
