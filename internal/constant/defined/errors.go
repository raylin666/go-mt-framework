package defined

import api "mt/api/v1"

var (
	/* 系统相关 */
	ErrorUnknown                = api.ErrorUnknown("未知错误")
	ErrorServer                 = api.ErrorServer("服务异常")
	ErrorDataValidate           = api.ErrorDataValidate("数据校验失败")
	ErrorDataSelect             = api.ErrorDataSelect("数据查询失败")
	ErrorDataAlreadyExists      = api.ErrorDataAlreadyExists("数据已存在")
	ErrorDataNotFound           = api.ErrorDataNotFound("数据不存在")
	ErrorDataAdd                = api.ErrorDataAdd("新增数据失败")
	ErrorDataUpdate             = api.ErrorDataUpdate("更新数据失败")
	ErrorDataDelete             = api.ErrorDataDelete("数据删除失败")
	ErrorDataResourceNotFound   = api.ErrorDataResourceNotFound("数据资源不存在")
	ErrorDataUpdateField        = api.ErrorDataUpdateField("数据属性更新失败")
	ErrorIdInvalidValue         = api.ErrorIdInvalidValue("无效ID值")
	ErrorCommandInvalidNotFound = api.ErrorCommandInvalidNotFound("无效的执行指令")
	ErrorNotLogin               = api.ErrorNotLogin("未登录")
	ErrorNotVisitAuth           = api.ErrorNotVisitAuth("没有访问权限")
)
