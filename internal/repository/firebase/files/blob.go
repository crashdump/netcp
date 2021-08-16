package files

import (
	"context"
	"io/ioutil"
	"log"
	"time"

	"cloud.google.com/go/storage"
	firebase "firebase.google.com/go/v4"
	"github.com/crashdump/netcp/pkg/entity"
)

type Repository interface {
	Save(blobMetadata entity.BlobMetadata, blob *entity.Blob) error
	GetByID(blobMetadata entity.BlobMetadata) (*entity.Blob, error)
	Delete(blobMetadata entity.BlobMetadata) error
}

type repository struct {
	ctx           context.Context
	StorageClient *storage.Client
	BucketHandle  *storage.BucketHandle
}

//NewBlobRepo is the single instance repo that is being created.
func NewBlobRepo(fc *firebase.App, bucketName string) (Repository, error) {
	ctx := context.Background()

	s, err := fc.Storage(ctx)
	if err != nil {
		log.Printf("error initialising firebase storage: %v", err)
		return nil, err
	}

	bucketHandle, err := s.Bucket(bucketName)
	if err != nil {
		log.Printf("Unable to open Firebase storage bucket")
	}

	br := &repository{
		ctx:          ctx,
		BucketHandle: bucketHandle,
	}
	return br, err
}

func (r *repository) GetByID(blobMetadata entity.BlobMetadata) (*entity.Blob, error) {
	var blob entity.Blob

	ctx, cancel := context.WithTimeout(r.ctx, time.Second*50)
	defer cancel()

	filePath := blobMetadata.OwnerID.String() + "/" + blobMetadata.ID.String()
	br, err := r.BucketHandle.Object(filePath).NewReader(ctx)
	if err != nil {
		return &blob, err
	}

	blob.Content, err = ioutil.ReadAll(br)
	if err != nil {
		log.Printf("GetByID(): unable to read file %q: %v", blobMetadata.Filename, err)
		return &blob, err
	}

	if err := br.Close(); err != nil {
		log.Printf("GetByID(): unable to close bucket after file write %q: %v", blobMetadata.Filename, err)
		return &blob, err
	}

	log.Printf("GetByID(): downloaded file %s", filePath)

	return &blob, nil
}

func (r *repository) Save(blobMetadata entity.BlobMetadata, blob *entity.Blob) error {
	ctx, cancel := context.WithTimeout(r.ctx, time.Second*50)
	defer cancel()

	filePath := blobMetadata.OwnerID.String() + "/" + blobMetadata.ID.String()
	bw := r.BucketHandle.Object(filePath).NewWriter(ctx)
	bw.ContentType = "application/gzip" // https://datatracker.ietf.org/doc/html/rfc6713

	if _, err := bw.Write(blob.Content); err != nil {
		log.Printf("Save(): unable to write file %q: %v", blobMetadata.Filename, err)
		return err
	}
	if err := bw.Close(); err != nil {
		log.Printf("Save(): unable to close bucket after file write %q: %v", blobMetadata.Filename, err)
		return err
	}

	log.Printf("Save(): saved file %s", filePath)

	return nil
}

func (r *repository) Delete(blobMetadata entity.BlobMetadata) error {
	filePath := blobMetadata.OwnerID.String() + "/" + blobMetadata.ID.String()
	return r.BucketHandle.Object(filePath).Delete(r.ctx)
}
