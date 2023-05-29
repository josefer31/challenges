package domain

import (
	"github.com/google/uuid"
)

type Ads interface {
	Save(ad *Ad) (*Ad, error)
	FindById(id uuid.UUID) (*Ad, error)
	FindAll() []Ad
}
