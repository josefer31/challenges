package repository

import (
	"errors"
	"github.com/google/uuid"
	. "polaris/internal/application/domain"
)

type InMemoryAds struct{}

var ads = make([]Ad, 0)

func (receiver InMemoryAds) FindAll() []Ad {
	return ads
}
func (receiver InMemoryAds) FindById(id uuid.UUID) (*Ad, error) {
	for _, ad := range ads {
		if ad.GetId() == id {
			return &ad, nil
		}
	}
	return nil, errors.New("ad not found")
}

func (receiver InMemoryAds) Save(ad Ad) Ad {
	ads = append(ads, ad)
	return ad
}
