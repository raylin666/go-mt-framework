package data

import (
	"context"
	"mt/internal/app"
	"mt/internal/biz"
	"mt/pkg/repositories"
)

type heartbeatRepo struct {
	data  repositories.DataRepo
	tools *app.Tools
}

func NewHeartbeatRepo(repo repositories.DataRepo, tools *app.Tools) biz.HeartbeatRepo {
	return &heartbeatRepo{
		data:  repo,
		tools: tools,
	}
}

func (r *heartbeatRepo) PONE(ctx context.Context) error {
	return nil
}
