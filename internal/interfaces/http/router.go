package http

func (s *Server) RegisterRouters() {
	s.RegisterAuthRouters()
	s.RegisterUtilitRouters()
}

func (s *Server) RegisterAuthRouters() {
	s.Router.POST("/api/register", s.Handlers.Register)
	s.Router.POST("/api/login", s.Handlers.Login)
	s.Router.POST("/api/logout", s.Handlers.Logout)
}

func (s *Server) RegisterUtilitRouters() {
	s.Router.PUT("/api/component/:component_id", s.Handlers.ComponentStatusChange)
	s.Router.GET("/api/component/:component_id", s.Handlers.GetFeedByID)
	s.Router.GET("/api/components/", s.Handlers.CatalogPage)
	s.Router.GET("/api/component/draft", s.Handlers.GetCurrentUserDraft)
	s.Router.POST("/api/component/:component_id/like", s.Handlers.LikeComponent)
	s.Router.POST("/api/component/", s.Handlers.CreateComponent)
	s.Router.DELETE("/api/component/:component_id", s.Handlers.DeleteComponent)
}
