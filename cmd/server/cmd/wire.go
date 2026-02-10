//go:build wireinject
// +build wireinject

package cmd

import (
	"github.com/google/wire"

	wireapp "ddd/internal/application/wire"
	"ddd/internal/config"
	wireinfra "ddd/internal/infra/wire"
	wireiface "ddd/internal/interface/wire"
	server "ddd/internal/server/server"
	wireserver "ddd/internal/server/wire"
)

func InitServer(cfg *config.Config) (*server.Server, error) {
	wire.Build(
		wireinfra.Set,
		wireapp.Set,
		wireiface.Set,
		wireserver.Set,
	)
	return nil, nil
}
