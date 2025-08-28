package server

import (
	"github.com/go-kratos/kratos/v2/middleware/metadata"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"mt/api/v1"
	"mt/config"
	"mt/internal/app"
	"mt/internal/middleware/auth"
	logging "mt/internal/middleware/logger"
	"mt/internal/middleware/validate"
	"mt/internal/service"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(
	c *config.Bootstrap,
	heartbeatService *service.HeartbeatService,
	tools *app.Tools) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			metadata.Server(),
			logging.Server(tools.Logger()),
			auth.NewJWTAuthServer(tools.JWT()),
		),
	}
	if c.Server.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Server.Grpc.Network))
	}
	if c.Server.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Server.Grpc.Addr))
	}
	if c.Server.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Server.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	v1.RegisterHeartbeatServer(srv, heartbeatService)
	return srv
}
