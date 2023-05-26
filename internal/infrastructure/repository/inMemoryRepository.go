package repository

import (
	"github.com/google/uuid"
	"polaris/internal/application/domain"
)

type InMemoryAds struct{}

func (receiver *InMemoryAds) DeleteAll() {
	ads = make([]domain.Ad, 0)
}

var ads = make([]domain.Ad, 0)

func (receiver *InMemoryAds) FindAll() []domain.Ad {
	return ads
}
func (receiver *InMemoryAds) FindById(id uuid.UUID) (*domain.Ad, error) {
	for _, ad := range ads {
		if ad.GetId() == id {
			return &ad, nil
		}
	}
	return nil, domain.NewAdNotFoundError(id)
}

func (receiver *InMemoryAds) Save(ad *domain.Ad) (*domain.Ad, error) {
	ads = append(ads, *ad)
	return ad, nil
}

func NewInMemoryAds() domain.Ads { return &InMemoryAds{} }
