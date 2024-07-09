package repositories

import "mt/pkg/cache"

const (
	RedisConnectionDefaultName = "default"
)

var _ RedisRepo = (*redisRepo)(nil)

type RedisRepo interface {
	Count() int
	Redis(name string) cache.Redis
}

type redisRepo struct {
	resource map[string]cache.Redis
}

func (repo *redisRepo) Count() int {
	return len(repo.resource)
}

func (repo *redisRepo) Redis(name string) cache.Redis {
	return repo.resource[name]
}
