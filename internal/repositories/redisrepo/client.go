package redisrepo

import (
	"github.com/raylin666/go-utils/cache/redis"
	"mt/pkg/repositories"
)

// NewDefaultClient 创建默认客户端
func NewDefaultClient(repo repositories.RedisRepo) redis.Client {
	return repo.Redis(repositories.RedisConnectionDefaultName).Get()
}
