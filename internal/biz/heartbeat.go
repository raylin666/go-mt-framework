package biz

import (
	"context"
	"mt/internal/app"
	"mt/internal/data"
)

type HeartbeatUsecase struct {
	repo  data.HeartbeatRepo
	tools *app.Tools
}

func NewHeartbeatUsecase(repo data.HeartbeatRepo, tools *app.Tools) *HeartbeatUsecase {
	return &HeartbeatUsecase{repo: repo, tools: tools}
}

func (uc *HeartbeatUsecase) PONE(ctx context.Context) error {
	return uc.repo.PONE(ctx)
}
