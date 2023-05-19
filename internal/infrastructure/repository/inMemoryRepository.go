package repository

import (
	"github.com/google/uuid"
	. "polaris/internal/application/domain"
)

type InMemoryAds struct{}

func (receiver *InMemoryAds) DeleteAll() {
	ads = make([]Ad, 0)
}

var ads = make([]Ad, 0)

func (receiver *InMemoryAds) FindAll() []Ad {
	return ads
}
func (receiver *InMemoryAds) FindById(id uuid.UUID) *Ad {
	for _, ad := range ads {
		if ad.GetId() == id {
			return &ad
		}
	}
	return nil
}

func (receiver *InMemoryAds) Save(ad Ad) Ad {
	ads = append(ads, ad)
	return ad
}

func NewInMemoryAds() Ads { return &InMemoryAds{} }
