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
	Price       uint
	CreatedAt   string
}

type CreateAdService interface {
	Execute(request CreateAdRequest) CreateAdResponse
}

type CreateAdServiceImpl struct {
	AdRepository Ads
	IdGenerator  IdGenerator
	Clock        Clock
}

func (service *CreateAdServiceImpl) Execute(request CreateAdRequest) CreateAdResponse {
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
		Price:       savedAd.Price,
		CreatedAt:   savedAd.GetCreatedAt().String(),
	}
}

func NewCreateAdService(ads Ads, idGenerator IdGenerator, clock Clock) CreateAdService {
	return &CreateAdServiceImpl{
		AdRepository: ads,
		IdGenerator:  idGenerator,
		Clock:        clock,
	}
}
