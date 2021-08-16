package blob

import (
	"github.com/crashdump/netcp/internal/repository/firebase/storage"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/google/uuid"
)

type Service interface {
	Download(ID uuid.UUID) (*entity.BlobMetadata, error)
	Upload(Blob *entity.BlobMetadata) error
	Remove(ID uuid.UUID) error
}

type service struct {
	repository storage.BlobRepository
}

//NewService is used to create a single instance of the service
func NewService(r storage.BlobRepository) Service {
	return &service{
		repository: r,
	}
}

func (s *service) Download(ID uuid.UUID) (*entity.BlobMetadata, error) {
	return s.repository.GetByID(ID)
}

func (s *service) Upload(Blob *entity.BlobMetadata) error {
	return s.repository.Save(Blob)
}

func (s *service) Remove(ID uuid.UUID) error {
	return s.repository.Delete(ID)
}