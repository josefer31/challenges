package repository

import "polaris/application/domain"

type InMemoryAdRepository struct {
	domain.AdRepository
}

var ads = make([]domain.Ad, 0)

func (receiver InMemoryAdRepository) Save(ad domain.Ad) domain.Ad {
	ads = append(ads, ad)
	return ad
}
