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
	IdGenerator  IdGenerator
	Clock        Clock
}

func (service *CreateAdService) Execute(request CreateAdRequest) CreateAdResponse {
	ad := NewAd(
		service.IdGenerator.Next(),
		request.Title,
		request.Description,
		request.Price,
		service.Clock.Now(),
	)

	savedAd := service.AdRepository.Save(ad)

	return CreateAdResponse{
		Id:          savedAd.GetId().String(),
		Title:       savedAd.Title,
		Description: savedAd.Description,
		CreatedAt:   savedAd.GetCreatedAt().String(),
	}
}
