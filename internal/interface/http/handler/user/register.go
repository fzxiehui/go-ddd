package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	service "ddd/internal/application/service/user"
	dto "ddd/internal/interface/http/dto/user"
)

type RegisterHandler struct {
	svc *service.RegisterService
}

func NewRegisterHandler(svc *service.RegisterService) *RegisterHandler {
	return &RegisterHandler{svc: svc}
}

func (h *RegisterHandler) Register(c *gin.Context) {
	var req dto.RegisterReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}

	u, err := h.svc.Register(req.Username, req.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.RegisterResp{
		UserID:   u.ID,
		Username: u.Username,
	})
}
