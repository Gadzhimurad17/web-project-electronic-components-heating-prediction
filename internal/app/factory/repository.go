package factory

import (
	"heating-prediction/internal/domain/ports/repo"
	"heating-prediction/internal/infrastructure/persistence/repo/postgres"

	"gorm.io/gorm"
)

func (f *Factory) NewRepos() (repo.ComponentRepo, repo.LikeRepo, repo.UserRepo) {

	db, err := postgres.Connect(f.cfg)
	if err != nil {
		panic("postgres isn`t connected")
	}
	return f.NewComponetRepo(db), f.NewLikeRepo(db), f.NewUserRepo(db)
}

func (f *Factory) NewUserRepo(db *gorm.DB) repo.UserRepo {
	return postgres.NewUserRepo(db)
}

func (f *Factory) NewComponetRepo(db *gorm.DB) repo.ComponentRepo {
	return postgres.NewComponentRepo(db)
}

func (f *Factory) NewLikeRepo(db *gorm.DB) repo.LikeRepo {
	return postgres.NewLikeRepo(db)
}
