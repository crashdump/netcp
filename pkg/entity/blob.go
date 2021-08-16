package entity

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

type BlobMetadata struct {
	ID           uuid.UUID `json:"id"            firestore:"-"` // id is stored in the doc path
	ShortID      string    `json:"short_id"      firestore:"short_id"`
	OwnerID      uuid.UUID `json:"owner_id"      firestore:"-"` // owner is stored in the doc path
	Filename     string    `json:"filename"      firestore:"filename"`
	CreatedAt    time.Time `json:"created_at"    firestore:"created_at"`
	DownloadedAt time.Time `json:"downloaded_at" firestore:"downloaded_at"`
}

func (b BlobMetadata) String() string {
	ju, _ := json.Marshal(b)
	return string(ju)
}

type BlobMetadatas []BlobMetadata

func (b BlobMetadatas) String() string {
	ju, _ := json.Marshal(b)
	return string(ju)
}

func (b *BlobMetadata) Validate() error {
	if b.CreatedAt.IsZero() {
		return errors.New("created_at cannot be null")
	}
	return nil
}


type Blob struct {
	ID           uuid.UUID `json:"id"            firestore:"-"` // id is stored in the doc path
	Content      []byte    `json:"-"             firestore:"-"`
}

func (b Blob) String() string {
	ju, _ := json.Marshal(b.ID)
	return string(ju)
}

type Blobs []Blob