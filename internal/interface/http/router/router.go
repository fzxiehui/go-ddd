package router

import (
	handler "ddd/internal/interface/http/handler/user"

	"github.com/gin-gonic/gin"
)

type Handlers struct {
	Login    *handler.LoginHandler
	Register *handler.RegisterHandler
}

func Register(r *gin.Engine, h *Handlers) {

	registerAuth(r, h)
}
