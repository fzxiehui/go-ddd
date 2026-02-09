package wireiface

import (
	"github.com/google/wire"

	userhandler "ddd/internal/interface/http/handler/user"
	"ddd/internal/interface/http/router"
)

var Set = wire.NewSet(
	userhandler.NewLoginHandler,
	userhandler.NewRegisterHandler,

	// router 是 interface 层的装配者
	wire.Struct(new(router.Handlers), "*"),
)
