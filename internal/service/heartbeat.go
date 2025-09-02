package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	heartbeatPb "mt/api/heartbeat"

	"mt/internal/biz"
)

type HeartbeatService struct {
	heartbeatPb.UnimplementedHeartbeatServer

	uc *biz.HeartbeatUsecase
}

func NewHeartbeatService(uc *biz.HeartbeatUsecase) *HeartbeatService {
	return &HeartbeatService{uc: uc}
}

func (s *HeartbeatService) PONE(ctx context.Context, req *emptypb.Empty) (*heartbeatPb.PONEResponse, error) {
	err := s.uc.PONE(ctx)
	if err != nil {
		return nil, err
	}

	return &heartbeatPb.PONEResponse{Message: "PONE"}, nil
}
