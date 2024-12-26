//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"
	"mt/config"
	"mt/internal/api"
	"mt/internal/app"
	"mt/internal/biz"
	"mt/internal/data"
	"mt/internal/grpc"
	"mt/internal/repositories"
	"mt/internal/server"
	"mt/internal/service"
)

// wireApp init kratos application.
func wireApp(*config.Bootstrap, *config.Server, *config.Data, *app.Tools) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		repositories.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		api.ProviderSet,
		grpc.ProviderSet,
		newApp))
}
