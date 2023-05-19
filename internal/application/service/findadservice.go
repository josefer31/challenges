package service

import (
	"github.com/google/uuid"
	. "polaris/internal/application/domain"
)

type FindAdRequest struct {
	Id string
}

type FindAdResponse struct {
	Id          string
	Title       string
	Description string
	CreatedAt   string
}

type FindAdService struct {
	AdRepository Ads
}

func (service FindAdService) Execute(request FindAdRequest) (FindAdResponse, error) {
	id, _ := uuid.Parse(request.Id)
	savedAd, notFoundError := service.AdRepository.FindById(id)

	if notFoundError != nil {
		return FindAdResponse{}, notFoundError
	}

	return FindAdResponse{
		Id:          savedAd.GetId().String(),
		Title:       savedAd.Title,
		Description: savedAd.Description,
		CreatedAt:   savedAd.GetCreatedAt().String(),
	}, nil
}
