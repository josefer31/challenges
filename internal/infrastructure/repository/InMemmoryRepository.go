package repository

import (
	"errors"
	"github.com/google/uuid"
	. "polaris/internal/application/domain"
)

type InMemoryAdRepository struct {
	AdRepository
}

func (receiver InMemoryAdRepository) FindById(id uuid.UUID) (Ad, error) {
	for _, ad := range ads {
		if ad.GetId() == id {
			return ad, nil
		}
	}
	return Ad{}, errors.New("ad not found")
}

var ads = make([]Ad, 0)

func (receiver InMemoryAdRepository) Save(ad Ad) Ad {
	ads = append(ads, ad)
	return ad
}
