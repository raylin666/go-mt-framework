package data

import (
	"context"

	"mt/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type heartbeatRepo struct {
	data *Data
	log  *log.Helper
}

func NewHeartbeatRepo(data *Data, logger log.Logger) biz.HeartbeatRepo {
	return &heartbeatRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *heartbeatRepo) PONE(ctx context.Context) error {
	return nil
}
