package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) loginHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "login.html", gin.H{})
}

func (h *Handler) registerHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "register.html", gin.H{})
}

func (h *Handler) ordersHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "orders.html", gin.H{})
}

func (h *Handler) mainHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "main.html", gin.H{})
}

func (h *Handler) userHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "user.html", gin.H{})
}

func (h *Handler) driversHTML(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "drivers.html", gin.H{})
}
