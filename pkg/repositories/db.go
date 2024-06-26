package repositories

import "mt/pkg/db"

const (
	DbConnectionDefaultName = "default"
)

var _ DbRepo = (*dbRepo)(nil)

type DbRepo interface {
	DB(name string) db.Db
}

type dbRepo struct {
	resource map[string]db.Db
}

func (repo *dbRepo) DB(name string) db.Db {
	return repo.resource[name]
}
