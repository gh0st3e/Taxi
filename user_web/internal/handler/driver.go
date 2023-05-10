package handler

import (
	"context"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"taxi/internal/entity"
)

type DriverActions interface {
	DriverAuth(ctx context.Context, phone, password string) (entity.Driver, error)
	DriverOrders(ctx context.Context, driverID int, date string) ([]entity.Order, error)
	ChangeStatus(ctx context.Context, id, state int) error
	UpdateOrdersStatus(ctx context.Context, orders []entity.Order) error
}

type OrdersRequest struct {
	ID   int
	Date string
}

func (h *Handler) DriverAuth(ctx *gin.Context) {
	driver := entity.Driver{}

	err := ctx.ShouldBindJSON(&driver)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	driver, err = h.service.DriverAuth(ctx, driver.Phone, driver.Password)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, driver)
}

func (h *Handler) DriverOrders(ctx *gin.Context) {
	ordersReq := OrdersRequest{}

	err := ctx.ShouldBindJSON(&ordersReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	orders, err := h.service.DriverOrders(ctx, ordersReq.ID, ordersReq.Date)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (h *Handler) ChangeStatus(ctx *gin.Context) {
	changeStateReq := entity.ChangeStateRequest{}

	err := ctx.ShouldBindJSON(&changeStateReq)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.ChangeStatus(ctx, changeStateReq.ID, changeStateReq.State)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "OK")
}

func (h *Handler) UpdateOrdersStatus(ctx *gin.Context) {
	request := entity.ChangeManyStateRequest{}

	err := ctx.ShouldBindJSON(&request)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = h.service.UpdateOrdersStatus(ctx, request.Orders)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, "OK")
}
