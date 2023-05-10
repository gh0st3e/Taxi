package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"taxi/internal/entity"
)

type UserAuth interface {
	SignIn(ctx context.Context, phone, password string) (string, string, error)
	SignUp(ctx context.Context, user entity.User) (int, error)
}

func (h *Handler) SignIn(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Infof("user: %v", user)

	accessToken, refreshToken, err := h.service.SignIn(ctx, user.Phone, user.Password)
	if err != nil {
		h.logger.Errorf("handler.web_auth.SignIn couldn't sign in: %s", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, map[string]interface{}{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	})
}

func (h *Handler) SignUp(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	h.logger.Infof("user: %v", user)

	id, err := h.service.SignUp(ctx, user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"id": id})
}
