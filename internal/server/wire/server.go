package wireserver

import (
	"github.com/google/wire"

	"ddd/internal/interface/http"
	"ddd/internal/server/server"
)

var Set = wire.NewSet(
	http.NewHTTPServer,
	server.NewServer,
)
