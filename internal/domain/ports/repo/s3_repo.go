package repo

import "context"

type S3Repo interface {
	Upload(ctx context.Context)
}
