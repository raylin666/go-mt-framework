package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Heartbeat struct {
}

type HeartbeatRepo interface {
	PONE(context.Context) error
}

type HeartbeatUsecase struct {
	repo HeartbeatRepo
	log  *log.Helper
}

func NewHeartbeatUsecase(repo HeartbeatRepo, logger log.Logger) *HeartbeatUsecase {
	return &HeartbeatUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *HeartbeatUsecase) PONE(ctx context.Context) error {
	return uc.repo.PONE(ctx)
}
