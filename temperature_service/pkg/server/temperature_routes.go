package server

import (
	"github.com/haseebh/weatherapp_temperature/internal/handlers"
	"github.com/haseebh/weatherapp_temperature/internal/middleware"
)

func (s *GinServer) TemperatureHandlers(uc *handlers.TemperatureHandler) {

	g := s.server.Group("/temperature/api/v1")
	g.Use(middleware.AuthMiddleware())
	{
		g.POST("/fetch/:location", uc.FetchTemperature)
		g.GET("/temperature/:location/:period", uc.GetTemperatureByPeriod)
	}
}
