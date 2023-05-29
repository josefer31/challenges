package domain

import (
	"github.com/google/uuid"
	"time"
)

type Ad struct {
	id          uuid.UUID
	Title       string
	Description string
	Price       uint
	createdAt   time.Time
}

func (ad *Ad) GetCreatedAt() time.Time {
	return ad.createdAt
}

func (ad *Ad) GetId() uuid.UUID {
	return ad.id
}

func NewAd(id uuid.UUID, title string, description string, price uint, createdAt time.Time) *Ad {
	return &Ad{
		id,
		title,
		description,
		price,
		createdAt,
	}
}
