package app

import (
	"github.com/raylin666/go-utils/auth"
	"github.com/raylin666/go-utils/server/system"
	"mt/pkg/logger"
)

type Tools struct {
	logger   *logger.Logger
	datetime *system.Datetime
	jwt      auth.JWT
}

// NewTools 创建公共工具实例
func NewTools(logger *logger.Logger, datetime *system.Datetime, jwt auth.JWT) (tools *Tools) {
	tools = &Tools{logger: logger, datetime: datetime, jwt: jwt}
	return
}

// Logger 日志
func (tools *Tools) Logger() *logger.Logger { return tools.logger }

// Datetime 日期时间
func (tools *Tools) Datetime() *system.Datetime { return tools.datetime }

// JWT 鉴权认证
func (tools *Tools) JWT() auth.JWT { return tools.jwt }
