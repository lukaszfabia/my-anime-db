package server

import (
	"api/internal/config"
	"api/internal/routes"
	"log"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config *config.Config
	router *gin.Engine
}

func New(cfg *config.Config) *Server {
	router := gin.Default()

	err := router.SetTrustedProxies(cfg.TrustedProxies)

	if err != nil {
		log.Fatalf("Failed to set trusted proxies: %v", err)
	}

	router.Static("/static", "./static")
	router.Static("/styles", "./styles")
	router.LoadHTMLGlob("templates/*")

	return &Server{
		config: cfg,
		router: gin.Default(),
	}
}

func (s *Server) Run() error {
	routes.DefineRoutes(s.router)
	log.Printf("Starting server on %s", s.config.Port)
	return s.router.Run(":" + s.config.Port)
}
