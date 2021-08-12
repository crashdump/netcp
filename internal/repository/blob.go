package repository

import (
	"github.com/crashdump/netcp/internal/model"

	"github.com/google/uuid"
)

type BlobRepository struct {
}

func (u *BlobRepository) FindAll() model.Blobs {
	var blobs model.Blobs
	//blobs, err := Firebase().firestore.GetAll(context.Background(), []*firestore.DocumentRef{
	//	Firebase().firestore.Doc("blob/username"),
	//})
	//if err != nil {
	//	return nil, err
	//}

	//.Find(&blobs)

	return blobs
}

func (u *BlobRepository) FindByID(id uuid.UUID) model.Blob {
	var blob model.Blob
	Firebase() //.First(&blob, id)

	return blob
}

func (u *BlobRepository) Save(blob model.Blob) model.Blob {
	Firebase() //.Save(&blob)

	return blob
}

func (u *BlobRepository) Delete(blob model.Blob) {
	Firebase() //.Delete(&blob)
}