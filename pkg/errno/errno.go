/**
 * @Author: huangw1
 * @Date: 2019/7/30 14:40
 */

package errno

import (
	"fmt"
)

type CodeError struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *CodeError) Error() string {
	return fmt.Sprintf("%d: %s", e.Code, e.Message)
}

func NewError(code int, text string) *CodeError {
	return &CodeError{code, text}
}

func NewErrorMsg(message string) *CodeError {
	return &CodeError{0, message}
}
