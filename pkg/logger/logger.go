package logger

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/metadata"
	"github.com/raylin666/go-utils/logger"
	"go.uber.org/zap"
	"reflect"
	"time"
)

const (
	XMdKeyTraceId = "x-md-global-trace-id"
	LogKeyTraceId = "trace_id"
)

var _ log.Logger = (*Logger)(nil)

func (l *Logger) Log(level log.Level, keyValues ...interface{}) error {
	if len(keyValues) == 0 || len(keyValues)%2 != 0 {
		l.Logger.Warn(fmt.Sprint("keyValues must appear in pairs: ", keyValues))
		return nil
	}

	var data []zap.Field
	for i := 0; i < len(keyValues); i += 2 {
		data = append(data, zap.Any(fmt.Sprint(keyValues[i]), keyValues[i+1]))
	}

	switch level {
	case log.LevelDebug:
		l.Logger.Debug("", data...)
	case log.LevelInfo:
		l.Logger.Info("", data...)
	case log.LevelWarn:
		l.Logger.Warn("", data...)
	case log.LevelError:
		l.Logger.Error("", data...)
	case log.LevelFatal:
		l.Logger.Fatal("", data...)
	}
	return nil
}

const (
	LogApp     = "app"
	LogSQL     = "sql"
	LogRequest = "request"
	LogGrpc    = "grpc"
)

type Logger struct {
	*zap.Logger
}

func NewJSONLogger(opts ...logger.Option) (*Logger, error) {
	zapLogger, err := logger.NewJSONLogger(opts...)
	return &Logger{zapLogger}, err
}

func (l *Logger) UseApp(ctx context.Context) *zap.Logger {
	var traceId string
	if md, ok := metadata.FromServerContext(ctx); ok {
		traceId = md.Get(XMdKeyTraceId)
	}
	return l.Logger.Named(LogApp).With(zap.String(LogKeyTraceId, traceId))
}

func (l *Logger) UseSQL(ctx context.Context) *zap.Logger {
	var traceId string
	if md, ok := metadata.FromServerContext(ctx); ok {
		traceId = md.Get(XMdKeyTraceId)
	}
	return l.Logger.Named(LogSQL).With(zap.String(LogKeyTraceId, traceId))
}

func (l *Logger) UseRequest(ctx context.Context) *zap.Logger {
	var traceId string
	if md, ok := metadata.FromServerContext(ctx); ok {
		traceId = md.Get(XMdKeyTraceId)
	}
	return l.Logger.Named(LogRequest).With(zap.String(LogKeyTraceId, traceId))
}

func (l *Logger) UseGrpc(ctx context.Context) *zap.Logger {
	var traceId string
	if md, ok := metadata.FromServerContext(ctx); ok {
		traceId = md.Get(XMdKeyTraceId)
	}
	return l.Logger.Named(LogGrpc).With(zap.String(LogKeyTraceId, traceId))
}

type RequestLogFormat struct {
	ClientIp          string              `json:"client_ip"`
	Method            string              `json:"method"`
	Path              string              `json:"path"`
	RequestProto      string              `json:"request_proto"`
	RequestReferer    string              `json:"request_referer"`
	RequestUa         string              `json:"request_ua"`
	RequestPostData   string              `json:"request_post_data"`
	RequestBodyData   string              `json:"request_body_data"`
	RequestHeaderData map[string][]string `json:"request_header_data"`
	HttpCode          int                 `json:"http_code"`
	BusinessMessage   string              `json:"business_message"`
	BusinessReason    string              `json:"business_reason"`
	Args              string              `json:"args"`
	RequestTime       time.Time           `json:"request_time"`
	ResponseTime      time.Time           `json:"response_time"`
	CostSeconds       float64             `json:"cost_seconds"`
}

// RequestLog 打印请求日志
func (l *Logger) RequestLog(ctx context.Context, rlf *RequestLogFormat, err error) {
	var types = reflect.TypeOf(rlf)
	var values = reflect.ValueOf(rlf)
	var zapLogger = l.UseRequest(ctx)
	for i := 0; i < types.Elem().NumField(); i++ {
		zapLogger = zapLogger.With(zap.Any(types.Elem().Field(i).Tag.Get("json"), values.Elem().Field(i).Interface()))
	}

	zapLogger.With(zap.Error(err)).Info("请求日志")
}
