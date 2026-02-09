//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"

	userservice "ddd/internal/application/service/user"
	"ddd/internal/domain/user"
	"ddd/internal/infra/db"
	userinfra "ddd/internal/infra/user"
	"ddd/internal/interface/http"
	userhandler "ddd/internal/interface/http/handler/user"
	"ddd/internal/interface/http/router"
	"ddd/internal/interface/server"
	"ddd/pkg/crypto"
)

var InfraSet = wire.NewSet(
	// config
	wire.Value(db.DBConfig{Path: "test.db"}),
	wire.Value(crypto.SHA256),

	// db
	db.InitSQLite,

	// repo
	userinfra.NewSQLiteRepo,
	wire.Bind(
		new(user.Repository),
		new(*userinfra.SQLiteRepo),
	),

	// password policy
	userinfra.NewDBPasswordPolicy,
	wire.Bind(
		new(user.PasswordPolicy),
		new(*userinfra.DBPasswordPolicy),
	),
)

var ServiceSet = wire.NewSet(
	userservice.NewLoginService,
	userservice.NewRegisterService,
)

var ServerSet = wire.NewSet(
	http.NewHTTPServer,
	server.NewServer,
)

var HandlerSet = wire.NewSet(
	userhandler.NewLoginHandler,
	userhandler.NewRegisterHandler,
)

func InitServer() (*server.Server, error) {
	wire.Build(

		// infra
		InfraSet,

		// handler
		HandlerSet,

		// http handler
		wire.Struct(new(router.Handlers), "*"),

		// service
		ServiceSet,

		// server
		ServerSet,
	)
	return nil, nil
}
