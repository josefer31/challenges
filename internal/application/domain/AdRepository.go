package domain

import (
	"github.com/google/uuid"
)

type AdRepository interface {
	Save(ad Ad) Ad
	FindById(id uuid.UUID) (Ad, error)
}
