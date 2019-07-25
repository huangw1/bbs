/**
 * @Author: huangw1
 * @Date: 2019/7/25 15:34
 */

package util

import (
	"fmt"
	"github.com/huangw1/bbs/utils/extension"
	"regexp"
	"strings"
)

const (
	ValidateRequire = "required"
	ValidateLength  = "length"
	ValidateEmail   = "email"
	ValidateCompare = "compare"

	RangeField   = "range"
	CompareField = "field"
	ErrorField   = "error"
	OPERATOR     = "operator"
)

var EmailRegexp = regexp.MustCompile(`^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+((\.[a-zA-Z0-9_-]{2,3}){1,2})$`)

func Validate(params map[string]interface{}, rules map[string]map[string]map[string]string) (message string) {
	for field, validateRule := range rules {
		value := params[field]
		if info, ok := validateRule[ValidateRequire]; ok {
			if value == "" {
				return info[ErrorField]
			}
		}
		if info, ok := validateRule[ValidateLength]; ok {
			var length int
			if v, ok := value.(int); ok {
				length = v
			}
			if v, ok := value.(string); ok {
				length = len(v)
			}
			if rangeInfo, ok := info[RangeField]; ok {
				if message = checkRange(length, rangeInfo, info[ErrorField]); message != "" {
					return message
				}

			}
		}
		if info, ok := validateRule[ValidateEmail]; ok {
			var str string
			if v, ok := value.(string); ok {
				str = v
			} else {
				return info[ErrorField]
			}
			if !EmailRegexp.MatchString(str) {
				return info[ErrorField]
			}
		}
		if compareInfo, ok := validateRule[ValidateCompare]; ok {
			compared := compareInfo[CompareField]
			operator := compareInfo[OPERATOR]
			switch operator {
			case "=":
				if params[compared] != value {
					return compareInfo[ErrorField]
				}
			case ">":
			case "<":
			default:
			}
		}
	}
	return
}

func checkRange(val int, lenRange string, msg string) (message string) {
	pair := strings.SplitN(lenRange, ",", 2)
	pair[0] = strings.TrimSpace(pair[0])
	pair[1] = strings.TrimSpace(pair[1])
	var min, max = 0, 0
	if pair[0] == "" {
		max = extension.MustInt(pair[1])
		if val > max {
			return fmt.Sprintf(msg, min, max)
		}
	}
	if pair[1] == "" {
		min = extension.MustInt(pair[0])
		if val < min {
			return fmt.Sprintf(msg, min, max)
		}
	}
	min = extension.MustInt(pair[0])
	max = extension.MustInt(pair[1])
	if val < min || val > max {
		return fmt.Sprintf(msg, min, max)
	}
	return ""
}
