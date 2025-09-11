package errors

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	errorsPb "mt/api/errors"
)

type Option func(opt *option)

func WithMessage(format string) Option {
	return func(opt *option) {
		opt.format = format
	}
}

func IsAdd(is bool) Option {
	return func(opt *option) {
		opt.isAdd = is
	}
}

type option struct {
	format string
	isAdd  bool
}

func Is(err, target error) bool { return errors.Is(err, target) }

type Error struct{ *option }

func New(opts ...Option) *Error {
	var err = new(Error)
	var o = new(option)
	for _, opt := range opts {
		opt(o)
	}
	err.option = o
	return err
}

func getFormat(err *Error, format string) string {
	if len(err.format) > 0 {
		if err.isAdd {
			format = fmt.Sprintf("%s: %s", format, err.format)
		} else {
			format = err.format
		}
	}

	return format
}

func (err *Error) Unknown(args ...interface{}) *errors.Error {
	return errorsPb.ErrorUnknown(getFormat(err, "未知错误"), args...)
}

func (err *Error) Server(args ...interface{}) *errors.Error {
	return errorsPb.ErrorServer(getFormat(err, "服务异常"), args...)
}

func (err *Error) DataValidate(args ...interface{}) *errors.Error {
	return errorsPb.ErrorDataValidate(getFormat(err, "数据校验失败"), args...)
}

func (err *Error) DataSelect(args ...interface{}) *errors.Error {
	return errorsPb.ErrorDataSelect(getFormat(err, "数据查询失败"), args...)
}

func (err *Error) DataAlreadyExists(args ...interface{}) *errors.Error {
	return errorsPb.ErrorDataAlreadyExists(getFormat(err, "数据已存在"), args...)
}

func (err *Error) DataNotFound(args ...interface{}) *errors.Error {
	return errorsPb.ErrorDataNotFound(getFormat(err, "数据不存在"), args...)
}

func (err *Error) DataAdd(args ...interface{}) *errors.Error {
	return errorsPb.ErrorDataAdd(getFormat(err, "新增数据失败"), args...)
}

func (err *Error) DataUpdate(args ...interface{}) *errors.Error {
	return errorsPb.ErrorDataUpdate(getFormat(err, "更新数据失败"), args...)
}

func (err *Error) DataDelete(args ...interface{}) *errors.Error {
	return errorsPb.ErrorDataDelete(getFormat(err, "数据删除失败"), args...)
}

func (err *Error) DataResourceNotFound(args ...interface{}) *errors.Error {
	return errorsPb.ErrorDataResourceNotFound(getFormat(err, "数据资源不存在"), args...)
}

func (err *Error) DataUpdateField(args ...interface{}) *errors.Error {
	return errorsPb.ErrorDataUpdateField(getFormat(err, "数据属性更新失败"), args...)
}

func (err *Error) IdInvalidValue(args ...interface{}) *errors.Error {
	return errorsPb.ErrorIdInvalidValue(getFormat(err, "无效ID值"), args...)
}

func (err *Error) CommandInvalidNotFound(args ...interface{}) *errors.Error {
	return errorsPb.ErrorCommandInvalidNotFound(getFormat(err, "无效的执行指令"), args...)
}

func (err *Error) RequestParams(args ...interface{}) *errors.Error {
	return errorsPb.ErrorRequestParams(getFormat(err, "请求参数错误"), args...)
}

func (err *Error) NotLogin(args ...interface{}) *errors.Error {
	return errorsPb.ErrorNotLogin(getFormat(err, "未登录帐号"), args...)
}

func (err *Error) NotVisitAuth(args ...interface{}) *errors.Error {
	return errorsPb.ErrorNotVisitAuth(getFormat(err, "没有访问权限"), args...)
}
