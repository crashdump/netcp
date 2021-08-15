package entity

import (
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"time"
)

type Blob struct {
	ID        uuid.UUID `json:"id"         firestore:"type:uuid"`
	Name      string    `json:"name"       firestore:"name"`
	CreatedAt time.Time `json:"created_at" firestore:"created_at"`
	Content   []byte    `json:"-"          firestore:"content"`
}

func (b Blob) String() string {
	ju, _ := json.Marshal(b)
	return string(ju)
}

type Blobs []Blob

func (b Blobs) String() string {
	ju, _ := json.Marshal(b)
	return string(ju)
}

func (b *Blob) Validate() error {
	if b.CreatedAt.IsZero() {
		return errors.New("created_at cannot be null")
	}
	return nil
}
