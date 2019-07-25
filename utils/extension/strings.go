/**
 * @Author: huangw1
 * @Date: 2019/7/25 17:45
 */

package extension

import (
	"reflect"
	"strconv"
	"strings"
)

func MustInt(s string) int {

	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func IntDefault(s string, defaultVal ...int) int {
	getDefault := func() int {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		return getDefault()
	}
	return i
}

func MustInt64(s string) int64 {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		panic(err)
	}
	return i
}

func Int64Default(s string, defaultVal ...int64) int64 {
	getDefault := func() int64 {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return getDefault()
	}
	return i
}

func Substr(s string, start, length int) string {
	runes := []rune(s)
	if start < 0 {
		start = 0
	}
	if start > len(runes) {
		start = start % len(runes)
	}
	end := start + length
	if end > len(runes) {
		end = len(runes)
	} else {
		end = start + length
	}
	return string(runes[start:end])
}

func RuneLen(s string) int {
	bt := []rune(s)
	return len(bt)
}

func GetSummary(s string, length int) string {
	summary := Substr(s, 0, length)
	if RuneLen(summary) > length {
		summary += "..."
	}
	return summary
}

func Contains(slice interface{}, target interface{}) (bool, int) {
	t := reflect.TypeOf(slice).Kind()
	if t == reflect.Slice || t == reflect.Array {
		v := reflect.ValueOf(slice)
		for i := 0; i < v.Len(); i++ {
			if target == v.Index(i).Interface() {
				return true, i
			}
		}
	}
	if t == reflect.String && reflect.TypeOf(target).Kind() == reflect.String {
		return strings.Contains(slice.(string), target.(string)), strings.Index(slice.(string), target.(string))
	}
	return false, -1
}
