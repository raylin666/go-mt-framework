package api

import (
	"github.com/google/wire"
	"github.com/gorilla/mux"
	"mt/config"
	"mt/internal/app"
	"mt/internal/grpc"
)

// ProviderSet is api.handler providers.
var ProviderSet = wire.NewSet(NewHandler)

type Handler struct {
	r          *mux.Router
	grpcClient grpc.GrpcClient
	tools      *app.Tools
	config     *config.Bootstrap
	Prefix     string
}

func NewHandler(config *config.Bootstrap, tools *app.Tools, grpcClient grpc.GrpcClient) *Handler {
	return &Handler{
		r:          mux.NewRouter(),
		grpcClient: grpcClient,
		tools:      tools,
		config:     config,
		Prefix:     "/app/",
	}
}

func (h *Handler) Router() *mux.Router {

	return h.r
}
