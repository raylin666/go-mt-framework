package errors

import (
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	pb "mt/api/v1"
)

type Option func(opt *option)

func WithMessage(format string) Option {
	return func(opt *option) {
		opt.format = format
	}
}

type option struct{ format string }

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

func (err *Error) Unknown(args ...interface{}) *errors.Error {
	var format = "未知错误"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorUnknown(err.format, args...)
}

func (err *Error) Server(args ...interface{}) *errors.Error {
	var format = "服务异常"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorServer(err.format, args...)
}

func (err *Error) DataValidate(args ...interface{}) *errors.Error {
	var format = "数据校验失败"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorDataValidate(err.format, args...)
}

func (err *Error) DataSelect(args ...interface{}) *errors.Error {
	var format = "数据查询失败"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorDataSelect(err.format, args...)
}

func (err *Error) DataAlreadyExists(args ...interface{}) *errors.Error {
	var format = "数据已存在"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorDataAlreadyExists(err.format, args...)
}

func (err *Error) DataNotFound(args ...interface{}) *errors.Error {
	var format = "数据不存在"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorDataNotFound(err.format, args...)
}

func (err *Error) DataAdd(args ...interface{}) *errors.Error {
	var format = "新增数据失败"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorDataAdd(err.format, args...)
}

func (err *Error) DataUpdate(args ...interface{}) *errors.Error {
	var format = "更新数据失败"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorDataUpdate(err.format, args...)
}

func (err *Error) DataDelete(args ...interface{}) *errors.Error {
	var format = "数据删除失败"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorDataDelete(err.format, args...)
}

func (err *Error) DataResourceNotFound(args ...interface{}) *errors.Error {
	var format = "数据资源不存在"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorDataResourceNotFound(err.format, args...)
}

func (err *Error) DataUpdateField(args ...interface{}) *errors.Error {
	var format = "数据属性更新失败"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorDataUpdateField(err.format, args...)
}

func (err *Error) IdInvalidValue(args ...interface{}) *errors.Error {
	var format = "无效ID值"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorIdInvalidValue(err.format, args...)
}

func (err *Error) CommandInvalidNotFound(args ...interface{}) *errors.Error {
	var format = "无效的执行指令"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorCommandInvalidNotFound(err.format, args...)
}

func (err *Error) RequestParams(args ...interface{}) *errors.Error {
	var format = "请求参数错误"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorRequestParams(err.format, args...)
}

func (err *Error) NotLogin(args ...interface{}) *errors.Error {
	var format = "未登录帐号"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorNotLogin(err.format, args...)
}

func (err *Error) NotVisitAuth(args ...interface{}) *errors.Error {
	var format = "没有访问权限"
	if len(err.format) > 0 {
		err.format = fmt.Sprintf("%s: %s", format, err.format)
	}

	return pb.ErrorNotVisitAuth(err.format, args...)
}