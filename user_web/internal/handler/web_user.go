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

type UserActions interface {
	RetrieveByID(ctx context.Context, id int) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, id int) error
	GetOrders(ctx context.Context, userID, state int) ([]entity.Order, error)
}

func (h *Handler) RetrieveByID(ctx *gin.Context) {
	id, ok := ctx.Get(userIDCtx)
	if !ok {
		h.logger.Errorf("handler.user.GetUserById couldn't get user by id: %s", errors.New("couldn't get userId from gin context"))
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
		h.logger.Errorf("handler.user.GetUserById couldn't get user with this id: %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error: %s", err),
		})
		return
	}

	user, err := h.service.RetrieveByID(ctx, idInt)
	if err != nil {
		h.logger.Errorf("handler.user.GetUserById couldn't get user by id: %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error: %s", err),
		})
		return
	}

	ctx.JSON(http.StatusOK, user)
}

func (h *Handler) UpdateUser(ctx *gin.Context) {
	id, ok := ctx.Get(userIDCtx)
	if !ok {
		h.logger.Errorf("handler.web_user.UpdateUser couldn't get user by id: %s", errors.New("couldn't get userId from gin context"))
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
		h.logger.Errorf("handler.web_user.UpdateUser couldn't get user with this id: %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error: %s", err),
		})
		return
	}

	var user entity.User
	if err = ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user.ID = idInt

	err = h.service.UpdateUser(ctx, &user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"Updated": id})
}

func (h *Handler) DeleteUser(ctx *gin.Context) {
	id, ok := ctx.Get(userIDCtx)
	if !ok {
		h.logger.Errorf("handler.web_user.DeleteUser couldn't get user by id: %s", errors.New("couldn't get userId from gin context"))
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
		h.logger.Errorf("handler.web_user.DeleteUser couldn't get user with this id: %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error: %s", err),
		})
		return
	}

	if err = h.service.DeleteUser(ctx, idInt); err != nil {
		h.logger.Errorf("handler.web_user.DeleteUser: couldn't delete user: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{"Deleted": idInt})
}

func (h *Handler) GetOrders(ctx *gin.Context) {
	id, ok := ctx.Get(userIDCtx)
	if !ok {
		h.logger.Errorf("handler.web_user.GetOrders couldn't get user by id: %s", errors.New("couldn't get userId from gin context"))
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": fmt.Sprintf("error: %s", errors.New("couldn't get userId from gin context")),
		})
		return
	}

	fmt.Println(id)

	idStr, ok := id.(string)
	if !ok {
		idStr = fmt.Sprintf("%s", id)
	}

	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		h.logger.Errorf("handler.web_user.GetOrders couldn't get user with this id: %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error: %s", err),
		})
		return
	}

	state := ctx.Param("state")
	idState, err := strconv.Atoi(state)
	if err != nil {
		h.logger.Errorf("handler.web_user.GetOrders couldn't get orders with this state: %s", err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("error: %s", err),
		})
		return
	}

	fmt.Println(idState)

	orders, err := h.service.GetOrders(ctx, idInt, idState)
	if err != nil {
		h.logger.Errorf("handler.web_user.GetOrders: couldn't get orders: %s", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, orders)
}
