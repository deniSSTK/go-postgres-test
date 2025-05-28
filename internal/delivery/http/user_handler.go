package http

import (
	"github.com/gin-gonic/gin"
	"go-postgres-test/internal/domain"
	"go-postgres-test/internal/usecase"
	"net/http"
)

type UserHandler struct {
	uc *usecase.UserUseCase
}

func NewUserHandler(uc *usecase.UserUseCase) *UserHandler {
	return &UserHandler{uc: uc}
}

func (h *UserHandler) RegisterRoutes(router *gin.Engine) {
	router.POST("/users/create-user", h.CreateUser)
}

func (h *UserHandler) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uc.CreateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusCreated, "success")
}
