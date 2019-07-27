/**
 * @Author: huangw1
 * @Date: 2019/7/30 14:26
 */

package response

import (
	"github.com/huangw1/bbs/pkg/errno"
	"github.com/huangw1/bbs/pkg/extension"
)

type Response struct {
	ErrorCode int         `json:"errorCode"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
	Success   bool        `json:"success"`
}

func Json(code int, message string, data interface{}, success bool) *Response {
	return &Response{
		ErrorCode: code,
		Message:   message,
		Data:      data,
		Success:   success,
	}
}

func JsonData(data interface{}) *Response {
	return &Response{
		ErrorCode: 0,
		Data:      data,
		Success:   true,
	}
}

func Success() *Response {
	return &Response{
		ErrorCode: 0,
		Data:      nil,
		Success:   true,
	}
}

func Error(err *errno.CodeError) *Response {
	return ErrorCode(err.Code, err.Message)
}

func ErrorMsg(message string) *Response {
	return &Response{
		ErrorCode: 0,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}

func ErrorCode(code int, message string) *Response {
	return &Response{
		ErrorCode: code,
		Message:   message,
		Data:      nil,
		Success:   false,
	}
}

type Builder struct {
	Data map[string]interface{}
}

func NewEmptyBuilder() *Builder {
	return &Builder{Data: make(map[string]interface{})}
}

func NewBuilder(source interface{}) *Builder {
	return NewBuilderExcludes(source)
}

func NewBuilderExcludes(source interface{}, excludes ...string) *Builder {
	return &Builder{Data: extension.StructToMap(source, excludes...)}
}

func (b *Builder) Put(key string, value interface{}) *Builder {
	b.Data[key] = value
	return b
}

func (b *Builder) Build() map[string]interface{} {
	return b.Data
}

func (b *Builder) BuildResponse() *Response {
	return JsonData(b.Data)
}
