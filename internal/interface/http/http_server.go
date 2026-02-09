package http

import (
	"ddd/internal/interface/http/router"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(h *router.Handlers) *gin.Engine {
	r := gin.Default()
	router.Register(r, h)
	return r
}
