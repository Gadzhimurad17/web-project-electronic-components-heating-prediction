package http

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Handlers *Handlers
	Router   *gin.Engine
	Server   *http.Server
}

func NewServer(host string, port int, handlers *Handlers, router *gin.Engine) *Server {
	return &Server{
		Handlers: handlers,
		Router:   router,
		Server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", host, port),
		},
	}
}

func (s *Server) Start() {
	l := slog.With(slog.String("server", "http"))

	s.RegisterRouters()

	l.Info("Starting server", slog.String("addr", s.Server.Addr))
	if err := s.Router.Run(); err != nil {
		log.Fatalf("can't start server, error=%s", err)
	}
}

func (s *Server) Close(ctx context.Context) {
	slog.InfoContext(ctx, "Stopping http server on", slog.String("address", s.Server.Addr))

	s.Server.Shutdown(ctx)
}
