package handlers_test

import (
	"encoding/json"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"

	"github.com/haseebh/weatherapp_auth/internal/entities/repository"
	"github.com/haseebh/weatherapp_auth/internal/handlers"
	"github.com/haseebh/weatherapp_auth/internal/usecases"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func TestUserHandler_Register(t *testing.T) {
	userUc := usecases.NewMockUserUseCase(t)

	tests := []struct {
		Name      string
		req       repository.User
		want      repository.User
		wantError bool
	}{
		{
			Name: "Valid registration",
			req: repository.User{
				Email:    "test@example.com",
				Password: "password",
				Location: "london",
				Name:     "Test User",
			},
			want: repository.User{
				Email:    "test@example.com",
				Location: "london",
				Token:    "token",
				Name:     "Test User",
			},
		},
		{
			Name:      "In-Valid registration",
			req:       repository.User{},
			wantError: true,
		},
	}
	router := setupRouter()
	userUc.On("Register", mock.Anything, mock.Anything).Return(&repository.User{}, nil)
	userUc.On("Register", mock.Anything, mock.Anything).Return(&repository.User{}, nil)
	userHandler := handlers.NewUserHandler(userUc)
	router.POST("/register", userHandler.Register)

	for _, trt := range tests {
		t.Run(trt.Name, func(t *testing.T) {

			userJSON, _ := json.Marshal(trt.req)
			req, _ := http.NewRequest("POST", "/register", strings.NewReader(string(userJSON)))
			req.Header.Set("Content-Type", "application/json")
			resp := httptest.NewRecorder()

			router.ServeHTTP(resp, req)

			result := false
			if resp.Code == http.StatusBadRequest {
				result = true
			}
			assert.Equal(t, trt.wantError, result)
		})

	}
}
