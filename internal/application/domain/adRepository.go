package domain

import (
	"github.com/google/uuid"
)

type Ads interface {
	Save(ad Ad) Ad
	FindById(id uuid.UUID) *Ad
	FindAll() []Ad
}
