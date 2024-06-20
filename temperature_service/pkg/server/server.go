package server

import (
	"github.com/gin-gonic/gin"
	"github.com/haseebh/weatherapp_temperature/pkg/config"
)

type Server interface {
	Run() error
}

var ()

type GinServer struct {
	server *gin.Engine
}

func GetHTTPServer() *GinServer {
	r := gin.Default()

	return &GinServer{
		server: r,
	}
}
func (s *GinServer) Run() error {
	return s.server.Run(":" + config.LoadConfig().ServerPort)
}
