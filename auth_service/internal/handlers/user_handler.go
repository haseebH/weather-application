package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/haseebh/weatherapp_auth/internal/entities/repository"
	"github.com/haseebh/weatherapp_auth/internal/usecases"
)

type UserHandler struct {
	userUsecase usecases.UserUseCase
}

func NewUserHandler(userUsecase usecases.UserUseCase) *UserHandler {
	return &UserHandler{userUsecase: userUsecase}
}

func (h *UserHandler) Register(c *gin.Context) {
	var user repository.User
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := h.userUsecase.Register(c, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUsecase.Login(c, credentials.Email, credentials.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	c.JSON(http.StatusOK, user)
}
