package repositories

import (
	"context"
	"fmt"
	"github.com/google/wire"
	"github.com/raylin666/go-utils/cache/redis"
	"mt/config"
	"mt/internal/app"
	"mt/internal/repositories/dbrepo"
	"mt/internal/repositories/dbrepo/query"
	"mt/internal/repositories/redisrepo"
	"mt/pkg/repositories"
)

// ProviderSet is repositories providers.
var ProviderSet = wire.NewSet(NewRepositories)

var _ DataRepo = (*Repositories)(nil)

type DataRepo interface {
	DefaultDbQuery() *query.Query
	DefaultRedisClient() redis.Client
}

type Repositories struct {
	db struct {
		defaultQuery *query.Query
	}
	redis struct {
		defaultClient redis.Client
	}
}

func NewRepositories(c *config.Bootstrap, tools *app.Tools) (DataRepo, func()) {
	var (
		ctx      = context.TODO()
		dataRepo = new(Repositories)
		repo     = repositories.NewDataRepo(tools.Logger(), c.Data)
	)

	dataRepo.db.defaultQuery = dbrepo.NewDefaultDbQuery(repo.DbRepo())
	dataRepo.redis.defaultClient = redisrepo.NewDefaultClient(repo.RedisRepo())

	cleanup := func() {
		// 资源关闭
		for dbName, dbRepo := range repo.DbRepo().All() {
			_ = dbRepo.Close()
			tools.Logger().UseApp(ctx).Info(fmt.Sprintf("closing the data repositories resource: `%s` db.repo successfully.", dbName))
		}

		for redisName, redisRepo := range repo.RedisRepo().All() {
			_ = redisRepo.Close()
			tools.Logger().UseApp(ctx).Info(fmt.Sprintf("closing the data repositories resource: `%s` db.repo successfully.", redisName))
		}
	}

	return dataRepo, cleanup
}

func (repositories *Repositories) DefaultDbQuery() *query.Query {
	//TODO implement me

	return repositories.db.defaultQuery
}

func (repositories *Repositories) DefaultRedisClient() redis.Client {
	//TODO implement me

	return repositories.redis.defaultClient
}
