package wireapp

import (
	"github.com/google/wire"

	"ddd/internal/application/service/auth"
	"ddd/internal/application/service/job"
	userservice "ddd/internal/application/service/user"
)

var Set = wire.NewSet(
	userservice.NewLoginService,
	userservice.NewRegisterService,
	auth.NewTokenService,
	job.NewScheduler,
)
