package data

import (
	"context"
	"fmt"
	"mt/config"
	"mt/internal/app"
	"mt/pkg/repositories"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData,
	NewDataRepo,
	NewHeartbeatRepo)

func NewDataRepo(tools *app.Tools, data *config.Data) repositories.DataRepo {
	return repositories.NewDataRepo(tools.Logger(), data)
}

// Data .
type Data struct {
	// TODO wrapped database client
	DbRepo    repositories.DbRepo
	RedisRepo repositories.RedisRepo
}

// NewData .
func NewData(tools *app.Tools, repo repositories.DataRepo) (*Data, func(), error) {
	var ctx = context.Background()
	cleanup := func() {
		// 资源关闭
		for dbName, dbRepo := range repo.DbRepo().All() {
			_ = dbRepo.Close()
			tools.Logger().UseApp(ctx).Info(fmt.Sprintf("closing the data resource: %s db.repo.", dbName))
		}

		for redisName, redisRepo := range repo.RedisRepo().All() {
			_ = redisRepo.Close()
			tools.Logger().UseApp(ctx).Info(fmt.Sprintf("closing the data resource: %s db.repo.", redisName))
		}
	}

	return &Data{
		DbRepo:    repo.DbRepo(),
		RedisRepo: repo.RedisRepo(),
	}, cleanup, nil
}
