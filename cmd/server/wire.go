//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	wireapp "ddd/internal/application/wire"
	wireinfra "ddd/internal/infra/wire"
	wireiface "ddd/internal/interface/wire"
	server "ddd/internal/server/server"
	wireserver "ddd/internal/server/wire"
)

func InitServer() (*server.Server, error) {
	wire.Build(
		wireinfra.Set,
		wireapp.Set,
		wireiface.Set,
		wireserver.Set,
	)
	return nil, nil
}
