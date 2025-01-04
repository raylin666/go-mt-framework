package data

import (
	"context"
	"mt/internal/app"
	"mt/internal/repositories"
)

type HeartbeatRepo interface {
	PONE(context.Context) error
}

type heartbeatRepo struct {
	data  repositories.DataRepo
	tools *app.Tools
}

func NewHeartbeatRepo(repo repositories.DataRepo, tools *app.Tools) HeartbeatRepo {
	return &heartbeatRepo{
		data:  repo,
		tools: tools,
	}
}

func (r *heartbeatRepo) PONE(ctx context.Context) error {
	return nil
}
