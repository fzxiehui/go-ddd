package wireiface

import (
	"github.com/google/wire"

	"ddd/internal/interface/grpc/handler/user"
	userhandler "ddd/internal/interface/http/handler/user"
	"ddd/internal/interface/http/router"
)

var Set = wire.NewSet(
	// HTTP
	userhandler.NewLoginHandler,
	userhandler.NewRegisterHandler,

	// GRPC
	user.NewAuthHandler,

	// router 是 interface 层的装配者
	wire.Struct(new(router.Handlers), "*"),
)
