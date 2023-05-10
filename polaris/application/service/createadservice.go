package service

import (
	"polaris/application/domain"
)

type CreateAdRequest struct {
	Title       string
	Description string
	Price       uint
}

type CreateAdService struct {
	AdRepository domain.AdRepository
}

func (service CreateAdService) Execute(request CreateAdRequest) domain.Ad {
	ad := domain.CreateAd(request.Title, request.Description, request.Price)
	return service.AdRepository.Save(ad)
}
