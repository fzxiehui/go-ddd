package wireinfra

import (
	"github.com/google/wire"

	"ddd/internal/infra/db"
	"ddd/internal/infra/security"
	userinfra "ddd/internal/infra/user"
)

var Set = wire.NewSet(

	// config
	// wire.Value(db.DBConfig{Path: "test.db"}),
	// wire.Value(crypto.SHA256),
	security.NewBcryptPasswordHasher,

	// db
	db.InitSQLite,

	// repo
	userinfra.NewSQLiteRepo,

	// policy
	userinfra.NewDBPasswordPolicy,
)
