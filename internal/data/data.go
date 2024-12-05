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
var ProviderSet = wire.NewSet(NewData, NewHeartbeatRepo)

// NewData .
func NewData(cData *config.Data, tools *app.Tools) (repositories.DataRepo, func(), error) {
	var ctx = context.Background()

	var repo = repositories.NewDataRepo(tools.Logger(), cData)

	var cleanup = func() {
		// 资源关闭
		for dbName, dbRepo := range repo.DbRepo().All() {
			_ = dbRepo.Close()
			tools.Logger().UseApp(ctx).Info(fmt.Sprintf("closing the data resource: `%s` db.repo successfully.", dbName))
		}

		for redisName, redisRepo := range repo.RedisRepo().All() {
			_ = redisRepo.Close()
			tools.Logger().UseApp(ctx).Info(fmt.Sprintf("closing the data resource: `%s` db.repo successfully.", redisName))
		}
	}

	return repo, cleanup, nil
}
