package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	service "ddd/internal/application/service/user"
	dto "ddd/internal/interface/http/dto/user"
)

type LoginHandler struct {
	svc *service.LoginService
}

func NewLoginHandler(svc *service.LoginService) *LoginHandler {
	return &LoginHandler{svc: svc}
}

func (h *LoginHandler) Login(c *gin.Context) {
	var req dto.LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	u, err := h.svc.Login(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "login failed"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResp{
		UserID:   u.ID,
		Username: u.Username,
	})
}
