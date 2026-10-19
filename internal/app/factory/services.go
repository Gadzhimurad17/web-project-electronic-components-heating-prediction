package factory

import (
	"heating-prediction/internal/domain/ports/repo"
	"heating-prediction/internal/domain/ports/service"
	s "heating-prediction/internal/infrastructure/service"
)

func (f *Factory) NewServices(userRepo repo.UserRepo, componentRepo repo.ComponentRepo, likeRepo repo.LikeRepo, fileRepo repo.S3Repo) (service.AuthService, service.ComponentService) {
	return f.NewAuthService(userRepo), f.NewComponentService(componentRepo, likeRepo, fileRepo)
}

func (f *Factory) NewAuthService(userRepo repo.UserRepo) service.AuthService {
	return s.NewAuthService(userRepo)
}

func (f *Factory) NewComponentService(componentRepo repo.ComponentRepo, likeRepo repo.LikeRepo, fileRepo repo.S3Repo) service.ComponentService {
	return s.NewComponentService(componentRepo, likeRepo, fileRepo)
}
