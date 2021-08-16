package blob

import (
	"fmt"
	"github.com/crashdump/netcp/internal/repository/firebase/files"
	"github.com/crashdump/netcp/internal/repository/firebase/metadata"
	"github.com/crashdump/netcp/pkg/entity"
	"github.com/google/uuid"
	"github.com/teris-io/shortid"
	"time"
)

type Service interface {
	DownloadByID(ID uuid.UUID) (*entity.Blob, error)
	DownloadByShortID(ID string) (*entity.Blob, error)
	Upload(filename string, blob *entity.Blob) error
	Remove(ID uuid.UUID) error
}

type service struct {
	blobRepository     files.Repository
	metadataRepository metadata.Repository
}

//NewService is used to create a single instance of the service
func NewService(r files.Repository, m metadata.Repository) Service {
	return &service{
		blobRepository:     r,
		metadataRepository: m,
	}
}

func (s *service) DownloadByShortID(id string) (*entity.Blob, error) {
	userid := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	meta, err := s.metadataRepository.GetByShortID(id, userid)
	if err != nil {
		return &entity.Blob{}, err
	}

	return s.blobRepository.GetByID(meta)
}

func (s *service) DownloadByID(id uuid.UUID) (*entity.Blob, error) {
	return s.blobRepository.GetByID(entity.BlobMetadata{
		ID: id,
	})
}

func (s *service) Upload(filename string, blob *entity.Blob) error {
	userid := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	meta := entity.BlobMetadata{
		ID:        uuid.New(),
		ShortID:   shortid.MustGenerate(),
		OwnerID:   userid,
		Filename:  filename,
		CreatedAt: time.Now(),
	}

	err := s.blobRepository.Save(meta, blob)
	if err != nil {
		return fmt.Errorf("upload(): unable to save metadata for file %s", filename)
	}
	return s.blobRepository.Save(meta, blob)

}

func (s *service) Remove(id uuid.UUID) error {
	return s.blobRepository.Delete(entity.BlobMetadata{
		ID: id,
	})
}
