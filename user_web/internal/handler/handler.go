package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"net/http"
	"strings"
	"taxi/internal/service"
)

const (
	userIDCtx  = "userID"
	tokenCtx   = "token"
	signingKey = "ds1jlg3h245wp5oih432v"
)

type UserHandler interface {
	UserAuth
	UserActions
	UserOrders
	DriverActions
}

type Handler struct {
	logger  *logrus.Logger
	service *service.Service
}

func NewHandler(logger *logrus.Logger, service *service.Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) Mount(r *gin.Engine) {
	r.POST("/user/login", h.SignIn)
	r.POST("/user/create", h.SignUp)

	api := r.Group("/api", h.UserIdentity)

	api.GET("/user", h.RetrieveByID)
	api.PUT("/user", h.UpdateUser)
	api.DELETE("/user", h.DeleteUser)
	api.GET("/orders/:state", h.GetOrders)

	api.POST("/order", h.Order)
	api.DELETE("/order/:order", h.CancelOrder)

	driver := r.Group("/driver")

	driver.POST("/login", h.DriverAuth)
	driver.POST("/orders", h.DriverOrders)
	driver.PUT("/state", h.ChangeStatus)

}

func (h *Handler) UserIdentity(ctx *gin.Context) {
	header := ctx.GetHeader("Authorization")
	if header == "" {
		h.logger.Error("handler.userIdentity.GetHeader: auth header is empty")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("auth header is empty").Error()})
		return
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		h.logger.Error("handler.userIdentity.GetHeader: invalid auth header")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("invalid auth header").Error()})
		return
	}

	stringToken := headerParts[1]

	err := h.checkTokenForExpired(stringToken)
	if err != nil {
		h.logger.Errorf("handler.userIdentity.GetHeader: token is expired %s", err)
		ctx.JSON(http.StatusForbidden, gin.H{"error": errors.New("token is expired").Error()})
		return
	}

	userID, err := service.ParseToken(stringToken, signingKey)
	if err != nil {
		h.logger.Error("handler.userIdentity.ParseToken: couldn't parse token")
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": errors.New("couldn't parse token").Error()})
		return
	}

	ctx.Set(tokenCtx, stringToken)
	ctx.Set(userIDCtx, userID)
}

func (h *Handler) checkTokenForExpired(token string) error {
	err := service.CheckTokenForExpired(token)
	if err != nil {
		return err
	}

	return nil
}
