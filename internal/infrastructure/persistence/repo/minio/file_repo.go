package minio

import "github.com/minio/minio-go/v7"

type minioFileRepo struct {
	mc     *minio.Client
	bucket string
}

func NewMinioFileRepo() *minioFileRepo {
	panic("")
}

func (r *minioFileRepo) UploadVideo() {

}
