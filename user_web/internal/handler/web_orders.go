package handler

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"net/http"
	"strconv"
	"taxi/internal/entity"
)

type UserOrders interface {
	Order(ctx context.Context, order *entity.Order) error
	CancelOrder(ctx context.Context, id int) error
}

func (h *Handler) Order(ctx *gin.Context) {
	id, ok := ctx.Get(userIDCtx)
	if !ok {
		h.logger.Errorf("handler.web_orders.Order couldn't get user by id: %s", errors.New("couldn't get userId from gin context"))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error: %s", errors.New("couldn't get userId from gin context")),
		})
		return
	}

	idStr, ok := id.(string)
	if !ok {
		idStr = fmt.Sprintf("%s", id)
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Errorf("handler.web_orders.Order couldn't get user with this id: %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error: %s", err),
		})
		return
	}

	var order entity.Order

	err = ctx.ShouldBindJSON(&order)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	order.UserID = idInt

	err = h.service.Order(ctx, &order)
	if err != nil {
		h.logger.Errorf("handler.web_orders.Order couldn't set order %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, "OK")
}

func (h *Handler) CancelOrder(ctx *gin.Context) {
	orderID := ctx.Param("id")

	idInt, err := strconv.Atoi(orderID)
	if err != nil {
		h.logger.Errorf("handler.web_orders.CancelOrder couldn't get order with this id: %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error: %s", err),
		})
		return
	}

	err = h.service.CancelOrder(ctx, idInt)
	if err != nil {
		h.logger.Errorf("handler.web_orders.CancelOrder couln't cancel order: %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, "OK")
}
