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

func NewAd(title string, description string, price uint) Ad {
	return Ad{
		uuid.New(), title, description, price, time.Now(),
	}
}
