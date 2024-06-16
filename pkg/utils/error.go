package utils

import "github.com/go-kratos/kratos/v2/errors"

// ErrorMessage 获取错误信息
func ErrorMessage(err *errors.Error) (code uint32, message string) {
	code = uint32(err.Code)
	message = err.Message
	return
}
