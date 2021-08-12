package model

import (
	"encoding/json"
	"errors"
	"time"
)

type Blob struct {
	Model

	UploadedAt   time.Time `json:"uploaded_at" firestore:"uploaded_at"`
	DownloadedAt time.Time `json:"downloaded_at" firestore:"downloaded_at"`
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
	if b.UploadedAt.IsZero() {
		return errors.New("uploaded_at cannot be null")
	}
	return nil
}