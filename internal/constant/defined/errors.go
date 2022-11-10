package defined

import v1 "mt/api/v1"

var (
	/* 系统相关 */
	ErrorUnknownError           = v1.ErrorUnknownError("未知错误")
	ErrorServerError            = v1.ErrorServerError("服务异常")
	ErrorDataValidateError      = v1.ErrorDataValidateError("数据校验失败")
	ErrorDataSelectError        = v1.ErrorDataSelectError("数据查询失败")
	ErrorDataAlreadyExists      = v1.ErrorDataAlreadyExists("数据已存在")
	ErrorDataNotFound           = v1.ErrorDataNotFound("数据不存在")
	ErrorDataAddError           = v1.ErrorDataAddError("新增数据失败")
	ErrorDataUpdateError        = v1.ErrorDataUpdateError("更新数据失败")
	ErrorDataDeleteError        = v1.ErrorDataDeleteError("数据删除失败")
	ErrorDataResourceNotFound   = v1.ErrorDataResourceNotFound("数据资源不存在")
	ErrorDataUpdateFieldError   = v1.ErrorDataUpdateFieldError("数据属性更新失败")
	ErrorIdInvalidValueError    = v1.ErrorIdInvalidValueError("无效ID值")
	ErrorCommandInvalidNotFound = v1.ErrorCommandInvalidNotFound("无效的执行指令")
	ErrorNotLoginError          = v1.ErrorNotLoginError("请先登录后再操作")
	ErrorNotVisitAuth		    = v1.ErrorNotVisitAuth("没有访问权限")
)
