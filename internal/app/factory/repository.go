package factory

import (
	"heating-prediction/internal/domain/ports/repo"
	"heating-prediction/internal/infrastructure/persistence/repo/minio"
	"heating-prediction/internal/infrastructure/persistence/repo/postgres"

	m "github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"gorm.io/gorm"
)

func (f *Factory) NewRepos() (repo.ComponentRepo, repo.LikeRepo, repo.UserRepo, repo.S3Repo) {
	stg, err := minio.Connect(f.cfg)
	if err != nil {
		panic("minio isn`t connected")
	}
	db, err := postgres.Connect(f.cfg)
	if err != nil {
		panic("postgres isn`t connected")
	}
	return f.NewComponetRepo(db), f.NewLikeRepo(db), f.NewUserRepo(db), f.NewFileRepo(stg)
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

func (f *Factory) NewFileRepo(minioClient *m.Client) repo.S3Repo {
	minioClient, err := m.New(f.cfg.MinioEndpoint, &m.Options{
		Creds: credentials.NewStaticV4(
			f.cfg.MinioAccesKey,
			f.cfg.MinioSecretKey,
			"",
		),
		Secure: f.cfg.MinioUseSSL,
	})

	if err != nil {
		return nil
	}
	return minio.NewMinioFileRepo(minioClient, f.cfg.MinioBucket)
}
