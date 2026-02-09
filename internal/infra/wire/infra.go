package wireinfra

import (
	"github.com/google/wire"

	"ddd/internal/domain/user"
	"ddd/internal/infra/db"
	userinfra "ddd/internal/infra/user"
	"ddd/pkg/crypto"
)

var Set = wire.NewSet(

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

	// policy
	userinfra.NewDBPasswordPolicy,
	wire.Bind(
		new(user.PasswordPolicy),
		new(*userinfra.DBPasswordPolicy),
	),
)
