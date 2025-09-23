package logger

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	pkgLogger "mt/pkg/logger"
	"time"
)

// Server is an server logging middleware.
func Server(log *pkgLogger.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				message   string
				reason    string
				operation string
				headers   = make(map[string][]string)
			)
			startTime := time.Now()
			if info, ok := transport.FromServerContext(ctx); ok {
				operation = info.Operation()
				for _, hVal := range info.RequestHeader().Keys() {
					headers[hVal] = []string{info.RequestHeader().Get(hVal)}
				}
			}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				message = se.Message
				reason = se.Reason
			}

			if operation == "/v1.Heartbeat/PONE" {
				return
			}

			log.RequestLog(ctx, &pkgLogger.RequestLogFormat{
				Path:              operation,
				Args:              extractArgs(req),
				RequestHeaderData: headers,
				RequestTime:       startTime,
				ResponseTime:      time.Now(),
				HttpCode:          int(code),
				BusinessMessage:   message,
				BusinessReason:    reason,
				CostSeconds:       time.Since(startTime).Seconds(),
			}, err)
			return
		}
	}
}

// Client is an client logging middleware.
func Client(log *pkgLogger.Logger) middleware.Middleware {
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			var (
				code      int32
				message   string
				reason    string
				operation string
				headers   = make(map[string][]string)
			)
			startTime := time.Now()
			if info, ok := transport.FromClientContext(ctx); ok {
				operation = info.Operation()
				for _, hVal := range info.RequestHeader().Keys() {
					headers[hVal] = []string{info.RequestHeader().Get(hVal)}
				}
			}
			reply, err = handler(ctx, req)
			if se := errors.FromError(err); se != nil {
				code = se.Code
				message = se.Message
				reason = se.Reason
			}

			if operation == "/v1.Heartbeat/PONE" {
				return
			}

			log.RequestLog(ctx, &pkgLogger.RequestLogFormat{
				Path:              operation,
				Args:              extractArgs(req),
				RequestHeaderData: headers,
				RequestTime:       startTime,
				ResponseTime:      time.Now(),
				HttpCode:          int(code),
				BusinessMessage:   message,
				BusinessReason:    reason,
				CostSeconds:       time.Since(startTime).Seconds(),
			}, err)
			return
		}
	}
}

// extractArgs returns the string of the req
func extractArgs(req interface{}) string {
	if stringer, ok := req.(fmt.Stringer); ok {
		return stringer.String()
	}
	return fmt.Sprintf("%+v", req)
}

// extractError returns the string of the error
func extractError(err error) (log.Level, string) {
	if err != nil {
		return log.LevelError, fmt.Sprintf("%+v", err)
	}
	return log.LevelInfo, ""
}
