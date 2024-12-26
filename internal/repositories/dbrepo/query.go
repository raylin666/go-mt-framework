package dbrepo

import (
	"gorm.io/gorm"
	"mt/internal/repositories/dbrepo/query"
	"mt/pkg/repositories"
)

// NewDefaultDb 创建默认数据库
func NewDefaultDb(repo repositories.DbRepo) *gorm.DB {
	return repo.DB(repositories.DbConnectionDefaultName).Get().DB()
}

// NewDefaultDbQuery 创建默认数据库查询
func NewDefaultDbQuery(repo repositories.DbRepo) *query.Query {
	return query.Use(NewDefaultDb(repo))
}
