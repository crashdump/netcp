package storage

import (
	"context"
	"log"
	"strings"
	"time"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/storage"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/google/uuid"
)

type BlobRepository interface {
	Save(blob *entity.Blob) error
	GetByID(ID uuid.UUID) (*entity.Blob, error)
	Delete(ID uuid.UUID) error
}

type repository struct {
	ctx           context.Context
	StorageClient *storage.Client
	BucketName    string
}

//NewRepo is the single instance repo that is being created.
func NewRepo(fc *firebase.App, bucketName string) (BlobRepository, error) {
	ctx := context.Background()

	s, err := fc.Storage(ctx)
	if err != nil {
		log.Printf("error initialising firebase storage: %v", err)
		return nil, err
	}

	return &repository{
		ctx:           ctx,
		BucketName:    bucketName,
		StorageClient: s,
	}, nil
}

func (r *repository) GetByID(id uuid.UUID) (*entity.Blob, error) {
	//

	return &entity.Blob{}, nil
}

func (r *repository) Save(blob *entity.Blob) error {
	b, err := r.StorageClient.Bucket(r.BucketName)
	if err != nil {
		log.Printf("Unable to open Firebase storage bucket")
	}
	bw := b.Object(blob.ID.String()).NewWriter(r.ctx)
	bw.ContentType = "application/gzip" // https://datatracker.ietf.org/doc/html/rfc6713
	bw.Metadata = map[string]string{
		// TODO: "x-goog-meta-owner-id": TBD,
		"x-goog-meta-filename": blob.Name,
		"x-goog-meta-created-at": time.Now().String(),
	}

	if _, err := bw.Write([]byte(strings.Repeat("f", 1024*4) + "\n")); err != nil {
		log.Printf("createFile: unable to write data to bucket %q, file %q: %v", r.BucketName, blob.Name, err)
		return err
	}
	if err := bw.Close(); err != nil {
		log.Printf("createFile: unable to close bucket %q, file %q: %v", r.BucketName, blob.Name, err)
		return err
	}

	return nil
}

func (r *repository) Delete(uuid uuid.UUID) error {
	//

	return nil
}