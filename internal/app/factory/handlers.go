package factory

import (
	"heating-prediction/internal/interfaces/http"

	"github.com/gin-gonic/gin"
)

func (f *Factory) NewHandlers() *http.Handlers {
	componentRepo, likeRepo, userRepo, fileRepo := f.NewRepos()
	authService, componentService := f.NewServices(userRepo, componentRepo, likeRepo, fileRepo)
	return &http.Handlers{
		Service: componentService,
		Auth:    authService,
	}
}

func (f *Factory) NewHttpServer(host string, port int, handlers *http.Handlers, router *gin.Engine) *http.Server {
	return http.NewServer(host, port, handlers, router)
}

func (f *Factory) NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	gin.ForceConsoleColor()

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	return router
}
