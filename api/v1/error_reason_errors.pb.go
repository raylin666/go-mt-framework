// Code generated by protoc-gen-go-errors. DO NOT EDIT.

package v1

import (
	fmt "fmt"
	errors "github.com/go-kratos/kratos/v2/errors"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
const _ = errors.SupportPackageIsVersion1

// 未知错误
func IsUnknown(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_UNKNOWN.String() && e.Code == 500
}

// 未知错误
func ErrorUnknown(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_UNKNOWN.String(), fmt.Sprintf(format, args...))
}

// 服务异常
func IsServer(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_SERVER.String() && e.Code == 500
}

// 服务异常
func ErrorServer(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_SERVER.String(), fmt.Sprintf(format, args...))
}

// 数据校验失败
func IsDataValidate(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_VALIDATE.String() && e.Code == 422
}

// 数据校验失败
func ErrorDataValidate(format string, args ...interface{}) *errors.Error {
	return errors.New(422, ErrorReason_DATA_VALIDATE.String(), fmt.Sprintf(format, args...))
}

// 数据查询失败
func IsDataSelect(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_SELECT.String() && e.Code == 500
}

// 数据查询失败
func ErrorDataSelect(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_DATA_SELECT.String(), fmt.Sprintf(format, args...))
}

// 数据已存在
func IsDataAlreadyExists(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_ALREADY_EXISTS.String() && e.Code == 400
}

// 数据已存在
func ErrorDataAlreadyExists(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_DATA_ALREADY_EXISTS.String(), fmt.Sprintf(format, args...))
}

// 数据不存在
func IsDataNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_NOT_FOUND.String() && e.Code == 404
}

// 数据不存在
func ErrorDataNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_DATA_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 新增数据失败
func IsDataAdd(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_ADD.String() && e.Code == 500
}

// 新增数据失败
func ErrorDataAdd(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_DATA_ADD.String(), fmt.Sprintf(format, args...))
}

// 更新数据失败
func IsDataUpdate(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_UPDATE.String() && e.Code == 500
}

// 更新数据失败
func ErrorDataUpdate(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_DATA_UPDATE.String(), fmt.Sprintf(format, args...))
}

// 数据删除失败
func IsDataDelete(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_DELETE.String() && e.Code == 500
}

// 数据删除失败
func ErrorDataDelete(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_DATA_DELETE.String(), fmt.Sprintf(format, args...))
}

// 数据资源不存在
func IsDataResourceNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_RESOURCE_NOT_FOUND.String() && e.Code == 404
}

// 数据资源不存在
func ErrorDataResourceNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_DATA_RESOURCE_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 数据属性更新失败
func IsDataUpdateField(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_DATA_UPDATE_FIELD.String() && e.Code == 500
}

// 数据属性更新失败
func ErrorDataUpdateField(format string, args ...interface{}) *errors.Error {
	return errors.New(500, ErrorReason_DATA_UPDATE_FIELD.String(), fmt.Sprintf(format, args...))
}

// 无效ID值
func IsIdInvalidValue(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_ID_INVALID_VALUE.String() && e.Code == 404
}

// 无效ID值
func ErrorIdInvalidValue(format string, args ...interface{}) *errors.Error {
	return errors.New(404, ErrorReason_ID_INVALID_VALUE.String(), fmt.Sprintf(format, args...))
}

// 无效的执行指令
func IsCommandInvalidNotFound(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_COMMAND_INVALID_NOT_FOUND.String() && e.Code == 400
}

// 无效的执行指令
func ErrorCommandInvalidNotFound(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_COMMAND_INVALID_NOT_FOUND.String(), fmt.Sprintf(format, args...))
}

// 请求参数错误
func IsRequestParams(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_REQUEST_PARAMS.String() && e.Code == 400
}

// 请求参数错误
func ErrorRequestParams(format string, args ...interface{}) *errors.Error {
	return errors.New(400, ErrorReason_REQUEST_PARAMS.String(), fmt.Sprintf(format, args...))
}

// 未登录
func IsNotLogin(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_LOGIN.String() && e.Code == 401
}

// 未登录
func ErrorNotLogin(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_NOT_LOGIN.String(), fmt.Sprintf(format, args...))
}

// 没有访问权限
func IsNotVisitAuth(err error) bool {
	if err == nil {
		return false
	}
	e := errors.FromError(err)
	return e.Reason == ErrorReason_NOT_VISIT_AUTH.String() && e.Code == 401
}

// 没有访问权限
func ErrorNotVisitAuth(format string, args ...interface{}) *errors.Error {
	return errors.New(401, ErrorReason_NOT_VISIT_AUTH.String(), fmt.Sprintf(format, args...))
}
