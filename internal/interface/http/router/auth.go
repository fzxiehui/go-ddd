package router

import "github.com/gin-gonic/gin"

func registerAuth(r *gin.Engine, h *Handlers) {
	auth := r.Group("/auth")

	auth.POST("/login", h.Login.Login)
	auth.POST("/register", h.Register.Register)
}
