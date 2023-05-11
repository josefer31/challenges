package service

import . "polaris/internal/application/domain"

type CreateAdRequest struct {
	Title       string
	Description string
	Price       uint
}

type CreateAdResponse struct {
	Id          string
	Title       string
	Description string
	CreatedAt   string
}

type CreateAdService struct {
	AdRepository AdRepository
}

func (service CreateAdService) Execute(request CreateAdRequest) CreateAdResponse {
	ad := CreateAd(request.Title, request.Description, request.Price)
	savedAd := service.AdRepository.Save(ad)
	return CreateAdResponse{
		Id:          savedAd.GetId().String(),
		Title:       savedAd.GetTitle(),
		Description: savedAd.GetDescription(),
		CreatedAt:   savedAd.GetCreatedAt().String(),
	}
}
