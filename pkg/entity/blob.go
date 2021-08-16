package entity

import (
	"encoding/json"
	"github.com/google/uuid"
)

type Blob struct {
	ID      uuid.UUID `json:"id"            firestore:"-"` // id is stored in the doc path
	Content []byte    `json:"-"             firestore:"-"`
}

func (b Blob) String() string {
	ju, _ := json.Marshal(b.ID)
	return string(ju)
}

type Blobs []Blob
