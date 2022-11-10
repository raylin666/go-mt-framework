package data

import (
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
	cleanup := func() {
		logger.UseApp().Info("closing the data resources")
	}

	return &Data{
		DbRepo:    repo.DbRepo(),
		RedisRepo: repo.RedisRepo(),
	}, cleanup, nil
}
