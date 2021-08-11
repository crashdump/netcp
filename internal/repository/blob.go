package repository

import (
	"github.com/crashdump/netcp/internal/model"

	"github.com/google/uuid"
)

type BlobRepository struct {
}

func (u *BlobRepository) FindAll() model.Blobs {
	var blobs model.Blobs
	GetGorm().Find(&blobs)

	return blobs
}

func (u *BlobRepository) FindByID(id uuid.UUID) model.Blob {
	var blob model.Blob
	GetGorm().First(&blob, id)

	return blob
}

func (u *BlobRepository) Save(blob model.Blob) model.Blob {
	GetGorm().Save(&blob)

	return blob
}

func (u *BlobRepository) Delete(blob model.Blob) {
	GetGorm().Delete(&blob)
}