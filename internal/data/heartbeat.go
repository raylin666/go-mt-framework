package data

import (
	"context"
	"mt/internal/app"
	"mt/internal/biz"
)

type heartbeatRepo struct {
	data  *Data
	tools *app.Tools
}

func NewHeartbeatRepo(data *Data, tools *app.Tools) biz.HeartbeatRepo {
	return &heartbeatRepo{
		data:  data,
		tools: tools,
	}
}

func (r *heartbeatRepo) PONE(ctx context.Context) error {
	return nil
}
