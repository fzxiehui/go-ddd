package wireapp

import (
	"github.com/google/wire"

	userservice "ddd/internal/application/service/user"
)

var Set = wire.NewSet(
	userservice.NewLoginService,
	userservice.NewRegisterService,
)
