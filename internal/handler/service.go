package handler

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
	DownloadByID(ID uuid.UUID) (*entity.Blob, entity.BlobMetadata, error)
	DownloadByShortID(ID string) (*entity.Blob, entity.BlobMetadata, error)
	Upload(filename string, blob *entity.Blob) (string, error)
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

func (s *service) DownloadByShortID(id string) (*entity.Blob, entity.BlobMetadata, error) {
	userid := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	meta, err := s.metadataRepository.GetByShortID(id, userid)
	if err != nil {
		return &entity.Blob{}, entity.BlobMetadata{}, err
	}

	blob, err := s.blobRepository.GetByID(meta)
	return blob, meta, err
}

func (s *service) DownloadByID(id uuid.UUID) (*entity.Blob, entity.BlobMetadata, error) {
	userid := uuid.MustParse("00000000-0000-0000-0000-000000000000")

	meta, err := s.metadataRepository.GetByID(id, userid)
	if err != nil {
		return &entity.Blob{}, entity.BlobMetadata{}, err
	}

	blob, err := s.blobRepository.GetByID(meta)
	return blob, meta, err
}

func (s *service) Upload(filename string, blob *entity.Blob) (string, error) {
	userid := uuid.MustParse("00000000-0000-0000-0000-000000000000")
	meta := entity.BlobMetadata{
		ID:        uuid.New(),
		ShortID:   shortid.MustGenerate(),
		OwnerID:   userid,
		Filename:  filename,
		CreatedAt: time.Now(),
	}

	err := s.metadataRepository.Save(meta)
	if err != nil {
		return "", fmt.Errorf("upload(): unable to save metadata for file %s", filename)
	}

	err = s.blobRepository.Save(meta, blob)
	if err != nil {
		return "", fmt.Errorf("upload(): unable to save blob for file %s", filename)
	}

	return meta.ShortID, nil
}

func (s *service) Remove(id uuid.UUID) error {
	return s.blobRepository.Delete(entity.BlobMetadata{
		ID: id,
	})
}
