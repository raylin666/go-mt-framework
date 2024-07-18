package repositories

import "mt/pkg/db"

const (
	DbConnectionDefaultName = "default"
)

var _ DbRepo = (*dbRepo)(nil)

type DbRepo interface {
	Count() int
	Has(name string) bool
	DB(name string) db.Db
}

type dbRepo struct {
	resource map[string]db.Db
}

func (repo *dbRepo) Count() int {
	return len(repo.resource)
}

func (repo *dbRepo) Has(name string) bool {
	if _, ok := repo.resource[name]; ok {
		return true
	}

	return false
}

func (repo *dbRepo) DB(name string) db.Db {
	return repo.resource[name]
}
