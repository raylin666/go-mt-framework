package request

import (
	"context"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/google/uuid"
)

const (
	// XMdTraceIdName Metadata 元数据传递保存的请求ID名称
	XMdTraceIdName = "x-md-trace-id"
)

// Server is an server trace middleware.
func Trace() middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if md, ok := metadata.FromServerContext(ctx); ok {
				// 设置请求ID
				if len(md.Get(XMdTraceIdName)) <= 0 {
					md.Set(XMdTraceIdName, uuid.New().String())
				}
			}

			reply, err = handler(ctx, req)
			return
		}
	}
}
