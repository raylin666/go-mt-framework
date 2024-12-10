package request

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/google/uuid"
	"mt/pkg/logger"
)

// Server is an server trace middleware.
func Trace() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if md, ok := metadata.FromServerContext(ctx); ok {
				// 设置请求ID
				if len(md.Get(logger.XMdKeyTraceId)) <= 0 {
					md.Set(logger.XMdKeyTraceId, uuid.New().String())
				}
			}

			reply, err = handler(ctx, req)
			return
		}
	}
}
