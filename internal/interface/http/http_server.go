package http

import (
	"ddd/internal/config"
	"ddd/internal/interface/http/router"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(h *router.Handlers, cfg *config.Config) *http.Server {
	engine := gin.Default()
	router.Register(engine, h)

	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTP.Port),
		Handler: engine,
	}
}
