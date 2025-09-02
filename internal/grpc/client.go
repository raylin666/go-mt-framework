package grpc

import (
	"context"
	"fmt"
	"github.com/raylin666/go-utils/server/system"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	heartbeatPb "mt/api/heartbeat"
	"mt/internal/app"
	"mt/pkg/logger"
)

var _ ClientInterface = (*Client)(nil)

type ClientInterface interface {
	Heartbeat() heartbeatPb.HeartbeatClient
}

type Client struct {
	ctx         context.Context
	environment system.Environment
	logger      *logger.Logger

	connects []*grpc.ClientConn

	heartbeatClient heartbeatPb.HeartbeatClient
}

func NewGrpcClient(tools *app.Tools) (grpcClient ClientInterface, cleanup func(), err error) {
	var client = &Client{
		ctx:         context.TODO(),
		environment: tools.Environment(),
		logger:      tools.Logger(),
	}

	cleanup = func() {
		client.close()
		tools.Logger().UseGrpc(client.ctx).Info("closing the grpc clients successfully.")
	}

	err = client.connect()

	return client, cleanup, err
}

func (client *Client) connect() error {
	// 心跳服务客户端
	heartbeatEndpoint := client.getHeartbeatEndpoint()
	heartbeatClientConn, err := dial(client.ctx, heartbeatEndpoint, client.logger)
	if err != nil {
		client.logger.UseGrpc(client.ctx).Error(fmt.Sprintf("The heartbeat service client `%s` connected error.", heartbeatEndpoint), zap.Error(err))
		return err
	}
	client.connects = append(client.connects, heartbeatClientConn)
	client.heartbeatClient = heartbeatPb.NewHeartbeatClient(heartbeatClientConn)
	client.logger.UseGrpc(client.ctx).Info(fmt.Sprintf("The heartbeat service client `%s` connected successfully.", heartbeatEndpoint))

	return nil
}

func (client *Client) close() {
	for _, conn := range client.connects {
		conn.Close()
	}
}

// getHeartbeatEndpoint 获取心跳服务地址
func (client *Client) getHeartbeatEndpoint() string {
	return HeartbeatGrpcClientEndpoint
}

func (client *Client) Heartbeat() heartbeatPb.HeartbeatClient {
	//TODO implement me

	return client.heartbeatClient
}
