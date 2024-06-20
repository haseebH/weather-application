package server

import (
	"github.com/haseebh/weatherapp_auth/internal/handlers"
)

func (s *GinServer) UserHandlers(uc *handlers.UserHandler) {
	g := s.server.Group("/rbac/api/v1")
	{
		g.POST("/register", uc.Register)
		g.POST("/login", uc.Login)
		g.GET("/verify", uc.Login)
	}
}
