package storage

import (
	"context"
	"log"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewMinIO() (*minio.Client, error) {
	// endpoint := getEnv("MINIO_ENDPOINT", "localhost:9000")
	// accessKey := getEnv("MINIO_ACCESS_KEY", "root")
	// secretKey := getEnv("MINIO_SECRET_KEY", "rootpassword")
	// useSSL := getEnv("MINIO_USE_SSL", "false") == "true"

	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		return nil, err
	}

	bucket := getEnv("MINIO_BUCKET", "media")
	ctx := context.Background()
	exists, err := client.BucketExists(ctx, bucket)
	if err != nil {
		return nil, err
	}
	if !exists {
		if err := client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{}); err != nil {
			return nil, err
		}
		policy := `{
          "Version": "2012-10-17",
          "Statement": [{
            "Effect": "Allow",
            "Principal": {"AWS": ["*"]},
            "Action": ["s3:GetObject"],
            "Resource": ["arn:aws:s3:::` + bucket + `/*"]
          }]
        }`
		if err := client.SetBucketPolicy(ctx, bucket, policy); err != nil {
			log.Printf("⚠️  public policy: %v", err)
		}
	}

	return client, nil
}

func getEnv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
