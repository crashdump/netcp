package model

import (
	"github.com/google/uuid"
	"time"
)

type Model struct {
	ID        uuid.UUID `form:"id" gorm:"type:uuid;primary_key"`
	CreatedAt time.Time `binding:"-" form:"-"`
	UpdatedAt time.Time `binding:"-" form:"-"`
}