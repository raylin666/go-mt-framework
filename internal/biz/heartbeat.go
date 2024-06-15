package biz

import (
	"context"
	"mt/internal/app"
)

type Heartbeat struct {
}

type HeartbeatRepo interface {
	PONE(context.Context) error
}

type HeartbeatUsecase struct {
	repo  HeartbeatRepo
	tools *app.Tools
}

func NewHeartbeatUsecase(repo HeartbeatRepo, tools *app.Tools) *HeartbeatUsecase {
	return &HeartbeatUsecase{repo: repo, tools: tools}
}

func (uc *HeartbeatUsecase) PONE(ctx context.Context) error {
	return uc.repo.PONE(ctx)
}
