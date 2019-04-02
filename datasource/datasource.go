package datasource

import (
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/firebase/firebase-admin-go"
	"google.golang.org/api/option"
)

type Datasource interface {
	GetFireBaseApp() (*storage.BucketHandle, error)
}

type fireBase struct {
	ctx context.Context
}

func NewDataSource(ctx context.Context) Datasource {
	return &fireBase{
		ctx: ctx,
	}
}

func (fb *fireBase) GetFireBaseApp() (*storage.BucketHandle, error) {
	opt := option.WithCredentialsFile("../annoyKey.json")
	app, err := firebase.NewApp(fb.ctx, nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error starting a new firebase app: %v", err)
	}

	client, err := app.Storage(fb.ctx)
	if err != nil {
		return nil, fmt.Errorf("error starting a new firebase client: %v", err)
	}

	bucket, err := client.DefaultBucket()
	if err != nil {
		return nil, fmt.Errorf("error getting default bucket: %v", err)
	}

	return bucket, nil
}
