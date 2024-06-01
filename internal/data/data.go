package data

import (
	"context"
	"fmt"
	"mt/config"
	"mt/pkg/logger"
	"mt/pkg/repositories"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	repositories.NewDataRepo,
	NewHeartbeatRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	DbRepo    repositories.DbRepo
	RedisRepo repositories.RedisRepo
}

// NewData .
func NewData(c *config.Data, logger *logger.Logger, repo repositories.DataRepo) (*Data, func(), error) {
	var ctx = context.Background()
	cleanup := func() {
		// 资源关闭
		repo.DB(repositories.DbConnectionDefaultName).Close()
		logger.UseApp(ctx).Info(fmt.Sprintf("closing the data resource: %s db.repo.", repositories.DbConnectionDefaultName))
		repo.Redis(repositories.RedisConnectionDefaultName).Close()
		logger.UseApp(ctx).Info(fmt.Sprintf("closing the data resource: %s redis.repo.", repositories.RedisConnectionDefaultName))
	}

	return &Data{
		DbRepo:    repo.DbRepo(),
		RedisRepo: repo.RedisRepo(),
	}, cleanup, nil
}
